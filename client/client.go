// TODO: NewClient()
// TODO: Implement Multi Version

package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//const DefaultMOURL = "/controller/dc/v3"

// Client is the main entry point
type Client struct {
	BaseURL *url.URL
	//MOURL              string
	insecure           bool
	httpClient         *http.Client
	AuthToken          *Auth
	l                  sync.Mutex
	username           string
	password           string
	skipLoggingPayload bool
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

func SkipLoggingPayload(skipLoggingPayload bool) Option {
	return func(client *Client) {
		client.skipLoggingPayload = skipLoggingPayload
	}
}

func initClient(clientUrl, username string, options ...Option) *Client {
	var transport *http.Transport
	bUrl, err := url.Parse(clientUrl)
	if err != nil {
		// cannot move forward if url is undefined
		log.Fatal(err)
	}
	client := &Client{
		BaseURL:  bUrl,
		username: username,
		//MOURL:    DefaultMOURL,
	}

	for _, option := range options {
		option(client)
	}

	if client.httpClient == nil {
		transport = client.useInsecureHTTPClient(client.insecure)
		client.httpClient = &http.Client{
			Transport: transport,
		}
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
		bUrl, err := url.Parse(clientUrl)
		if err != nil {
			// cannot move forward if url is undefined
			log.Fatal(err)
		}
		if bUrl != clientImpl.BaseURL {
			clientImpl = initClient(clientUrl, username, options...)
		}
	}
	return clientImpl
}

func (c *Client) useInsecureHTTPClient(insecure bool) *http.Transport {
	transport := http.DefaultTransport.(*http.Transport)
	transport.TLSClientConfig = &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
		PreferServerCipherSuites: true,
		InsecureSkipVerify:       insecure,
		MinVersion:               tls.VersionTLS11,
		MaxVersion:               tls.VersionTLS12,
	}
	return transport

}

var applicationJSON = "application/json"

func (c *Client) MakeRestRequest(method string, path string, options RequestOpts, authenticated bool) (*http.Response, error) {
	var body io.ReadSeeker
	var contentType *string

	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	fURL, err := url.Parse(c.BaseURL.String())
	if err != nil {
		return nil, err
	}

	fURL = fURL.ResolveReference(pathURL)

	// Derive the content body by either encoding an arbitrary object as JSON, or by taking a provided
	// io.ReadSeeker as-is. Default the content-type to application/json.
	if options.JSONBody != nil {
		body = bytes.NewReader(options.JSONBody)
		contentType = &applicationJSON
	}

	var req *http.Request
	log.Printf("[DEBUG] BaseURL: %s, pathURL: %s, fURL: %s", c.BaseURL.String(), pathURL.String(), fURL.String())

	req, err = http.NewRequest(method, fURL.String(), body)

	if err != nil {
		return nil, err
	}

	if c.skipLoggingPayload {
		log.Printf("HTTP request %s %s", method, path)
	} else {
		log.Printf("HTTP request %s %s %v", method, path, req)
	}

	if authenticated {
		req, err = c.InjectAuthenticationHeader(req, path)
		if err != nil {
			return nil, err
		}
	}

	if !c.skipLoggingPayload {
		log.Printf("HTTP request after injection %s %s %v", method, path, req)
	}

	// Populate the request headers. Apply options.MoreHeaders last, to give the caller the chance to
	// modify or omit any header.
	if contentType != nil {
		req.Header.Set("Content-Type", *contentType)
	}

	//// Set the User-Agent header
	//req.Header.Set("User-Agent", client.UserAgent.Join())
	//
	//if options.MoreHeaders != nil {
	//	for k, v := range options.MoreHeaders {
	//		if v != "" {
	//			req.Header.Set(k, v)
	//		} else {
	//			req.Header.Del(k)
	//		}
	//	}
	//}

	// Set connection parameter to close the connection immediately when we've got the response
	req.Close = true

	// Issue the request.
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if !c.skipLoggingPayload {
		log.Printf("\nHTTP Response: %d %s %v", resp.StatusCode, resp.Status, resp)
	} else {
		log.Printf("\nHTTP Response: %d %s", resp.StatusCode, resp.Status)
	}

	//// Allow default OkCodes if none explicitly set
	if options.OkCodes == nil {
		options.OkCodes = defaultOkCodes(method)
	}

	// Validate the HTTP response status.
	var ok bool
	for _, code := range options.OkCodes {
		if resp.StatusCode == code {
			ok = true
			break
		}
	}
	if !ok {
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return resp, &UnexpectedResponseCodeError{
			URL:      fURL.String(),
			Method:   method,
			Expected: options.OkCodes,
			Actual:   resp.StatusCode,
			Body:     body,
		}
	}

	// Parse the response body as JSON, if requested to do so.
	if options.JSONResponse != nil {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(options.JSONResponse); err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *Client) Authenticate() error {
	method := "POST"
	path := "/controller/v2/tokens"
	var response AuthResponse

	opts := &RequestOpts{}
	opts.JSONBody = []byte(fmt.Sprintf(authPayload, c.username, c.password))
	opts.JSONResponse = &response
	opts.OkCodes = []int{200}

	c.skipLoggingPayload = true

	_, err := c.MakeRestRequest(method, path, *opts, false)
	if err != nil {
		return err
	}

	c.skipLoggingPayload = false
	if err != nil {
		log.Printf("[DEBUG]: ERROR %s", err)
		return err
	}

	token := opts.JSONResponse.(*AuthResponse).Data.(map[string]interface{})["token_id"].(string)
	expiredDateStr := opts.JSONResponse.(*AuthResponse).Data.(map[string]interface{})["expiredDate"].(string)
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

func defaultOkCodes(method string) []int {
	switch {
	case method == "GET":
		return []int{200}
	case method == "POST":
		return []int{202, 204}
	case method == "PUT":
		return []int{200}
	//case method == "PATCH":
	//	return []int{200, 204}
	case method == "DELETE":
		return []int{204}
	}

	return []int{}
}
