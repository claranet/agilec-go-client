// TODO: NewClient()
// TODO: Implement Multi Version

package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
	log "github.com/sirupsen/logrus"
)

// Client is the main entry point
type Client struct {
	BaseURL          string
	insecure         bool
	httpClient       *resty.Client
	AuthToken        *Auth
	l                sync.Mutex
	username         string
	password         string
	timeout          int
	retryCount       int
	retryWaitTime    int
	retryMaxWaitTime int
	skipLoggingPayload bool
	logLevel int
	*ServiceManager
}

// singleton implementation of a client
var clientImpl *Client

type Option func(*Client)

func Password(password string) Option {
	return func(client *Client) {
		client.password = password
	}
}

func TimeoutInSeconds(timeout int) Option {
	return func(client *Client) {
		client.timeout = timeout
	}
}

func RetryCount(retryCount int) Option {
	return func(client *Client) {
		client.retryCount = retryCount
	}
}

func RetryWaitTimeInSeconds(retryWaitTime int) Option {
	return func(client *Client) {
		client.retryWaitTime = retryWaitTime
	}
}

func RetryMaxWaitTimeSeconds(retryMaxWaitTime int) Option {
	return func(client *Client) {
		client.retryMaxWaitTime = retryMaxWaitTime
	}
}

func Insecure(insecure bool) Option {
	return func(client *Client) {
		client.insecure = insecure
	}
}

func SkipLoggingPayload(skipLoggingPayload bool) Option {
	return func(client *Client) {
		client.skipLoggingPayload = skipLoggingPayload
	}
}

func LogLevel(level int) Option {
	return func(client *Client) {
		client.logLevel = level
	}
}

func initClient(clientUrl, username string, options ...Option) *Client {
	//bUrl, err := url.Parse(clientUrl)
	//if err != nil {
	//	// cannot move forward if url is undefined
	//	log.Fatal(err)
	//}

	client := &Client{
		BaseURL:  clientUrl,
		username: username,
	}

	for _, option := range options {
		option(client)
	}

	log.SetLevel(log.Level(client.logLevel))

	log.Debug("Begin Client Initialization")

	if client.httpClient == nil {
		client.httpClient = resty.New()
		client.httpClient.SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: client.insecure,
		})
		client.httpClient.SetBaseURL(clientUrl)

		client.httpClient.SetTimeout(time.Duration(client.timeout) * time.Second)

		// Configure Retry Policy
		client.httpClient.
			SetRetryCount(client.retryCount).
			SetRetryWaitTime(time.Duration(client.timeout) * time.Second).
			SetRetryMaxWaitTime(time.Duration(client.timeout) * time.Second)

		client.httpClient.OnError(func(req *resty.Request, err error) {
			log.Fatal("Connection Timeout to Huawei Agile Controller")
		})
	}


	client.ServiceManager = NewServiceManager(client)
	return client
}

// GetClient returns a singleton
func GetClient(clientUrl, username string, options ...Option) *Client {

	if clientImpl == nil {
		clientImpl = initClient(clientUrl, username, options...)
	} else {
		// making sure it is the same client
		//bUrl, err := url.Parse(clientUrl)
		//if err != nil {
		//	// cannot move forward if url is undefined
		//	log.Fatal(err)
		//}
		if clientUrl != clientImpl.BaseURL {
			clientImpl = initClient(clientUrl, username, options...)
		}
	}
	return clientImpl
}

// NewClient returns a new Instance of the client - allowing for simultaneous connections to the same APIC
func NewClient(clientUrl, username string, options ...Option) *Client {
	// making sure it is the same client
	_, err := url.Parse(clientUrl)
	if err != nil {
		// cannot move forward if url is undefined
		log.Fatal(err)
	}

	// initClient always returns a new struct, so always create a new pointer to allow for
	// multiple object instances
	newClientImpl := initClient(clientUrl, username, options...)

	return newClientImpl
}

func (c *Client) Request(method, url string, opts RequestOpts, authenticated bool) (*resty.Response, error) {
	log.Debug("Begin New Request")
	request := c.httpClient.R().ForceContentType("application/json")
	request.SetError(&ErrorResponse{})
	if authenticated {
		log.Debug("Request Needs authentication")
		err := c.InjectAuthenticationHeader(request)
		if err != nil {
			return nil, err
		}
	}

	if opts.QueryString != nil {
		log.Debug("Injection Query String Parameters")
		q, err := query.Values(opts.QueryString)
		if err != nil {
			return nil, err
		}
		request.SetQueryString(q.Encode())
		log.Debugf("Query String: %s", q.Encode())
	}

	if opts.Body != nil {
		log.Debug("Setting Request Body")
		request.SetBody(opts.Body)
		if !c.skipLoggingPayload {
			log.Debugf("Request Body %+v", opts.Body)
		}
	}

	if c.skipLoggingPayload {
		log.Infof("HTTP request %s %s", method, url)
	} else {
		log.Infof("HTTP request %s %s %v", method, url, opts.Body)
	}

	if opts.Response != nil {
		log.Debug("Configure Response format")
		request.SetResult(&opts.Response)
	}

	log.Info("Send Request")
	doRequest := reflect.ValueOf(request).MethodByName(method).Call([]reflect.Value{reflect.ValueOf(url)})
	response := doRequest[0].Interface().(*resty.Response)

	if response.IsError() {
		log.Error("Error occurred")
		return nil, CheckForErrors(response, method, url)
	}
	log.Info("Request sent with success")
	return response, nil
}

func (c *Client) Authenticate() error {
	log.Debug("Begin authentication")
	var result AuthResponse

	resp, err := c.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(fmt.Sprintf(authPayload, c.username, c.password))).
		SetResult(&result).
		Post("/controller/v2/tokens")

	if err != nil {
		log.Error("Authentication Error occurred")
		return err
	}

	if resp.IsSuccess() {
		log.Info("Authentication with Success")
		token := result.Data.(map[string]interface{})["token_id"].(string)
		expiredDateStr := result.Data.(map[string]interface{})["expiredDate"].(string)
		expiredDate, err := time.Parse(expiredDateLayout, expiredDateStr)

		if err != nil {
			return err
		}

		if c.AuthToken == nil {
			c.AuthToken = &Auth{}
		}
		c.AuthToken.Token = token
		c.AuthToken.Expiry = expiredDate
		c.AuthToken.CaclulateOffset()
		return nil
	}

	return fmt.Errorf("authentication Failed")
}

func CheckForErrors(response *resty.Response, method, url string) *ErrorResponse {
	error := &ErrorResponse{
		Method:         method,
		URL:            url,
		HttpStatusCode: response.StatusCode(),
	}
	if response.Header().Get("Content-Type") == "text/plain" {
		error.ErrorMessage = response.String()
	} else {
		var msgErrorTemplate interface{}
		err := json.Unmarshal(response.Body(), &msgErrorTemplate)
		if err != nil {
			error.ErrorMessage = "Error Message Not supported. Please open an Issue"
		}
		if strings.Contains(response.String(), "errmsg") {
			error.ErrorMessage = msgErrorTemplate.(map[string]string)["errmsg"]
		} else {
			error.ErrorMessage = msgErrorTemplate.(map[string]map[string][1]map[string]string)["errors"]["error"][0]["error-message"]
		}
	}
	return error
}
