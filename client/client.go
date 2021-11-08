// TODO: NewClient()
// TODO: Implement Multi Version

package client

import (
	"crypto/tls"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"log"
	"net/url"
	"reflect"
	"sync"
	"time"
	"github.com/google/go-querystring/query"
)

// Client is the main entry point
type Client struct {
	BaseURL string
	insecure           bool
	httpClient         *resty.Client
	AuthToken          *Auth
	l                  sync.Mutex
	username           string
	password           string
	//skipLoggingPayload bool
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

func Insecure(insecure bool) Option {
	return func(client *Client) {
		client.insecure = insecure
	}
}

//func SkipLoggingPayload(skipLoggingPayload bool) Option {
//	return func(client *Client) {
//		client.skipLoggingPayload = skipLoggingPayload
//	}
//}

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

	if client.httpClient == nil {
		client.httpClient = resty.New()
		client.httpClient.SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: client.insecure,
		})
		client.httpClient.SetBaseURL(clientUrl)
	}

	//var timeout time.Duration
	//if client.reqTimeoutSet {
	//	timeout = time.Second * time.Duration(client.reqTimeoutVal)
	//} else {
	//	timeout = time.Second * time.Duration(DefaultReqTimeoutVal)
	//}

	//client.httpClient.Timeout = timeout
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

func (c *Client) NewRequest(method, url string, opts RequestOpts, authenticated bool) (*resty.Response, error) {
	request := c.httpClient.R().SetHeader("Content-Type", "application/json")
	request.SetError(&ErrorResponse{})
	if authenticated {
		err := c.InjectAuthenticationHeader(request)
		if err != nil {
			return nil, err
		}
	}

	if opts.QueryString != nil {
		q, err := query.Values(opts.QueryString)
		if err != nil {
			return nil, err
		}
		request.SetQueryString(q.Encode())
	}

	if opts.Body != nil {
		request.SetBody(opts.Body)
	}

	if opts.Response != nil {
		request.SetResult(&opts.Response)
	}

	doRequest := reflect.ValueOf(request).MethodByName(method).Call([]reflect.Value{reflect.ValueOf(url)})
	response := doRequest[0].Interface().(*resty.Response)

	if response.IsError() {
		return nil, response.Error().(*ErrorResponse)
	}
	return response, nil
}

func (c *Client) Authenticate() error {
	var result AuthResponse

	resp, err := c.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(fmt.Sprintf(authPayload, c.username, c.password))).
		SetResult(&result).

		Post("/controller/v2/tokens")

	if err != nil {
		return err
	}

	if resp.IsSuccess() {
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