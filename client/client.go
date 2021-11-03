// TODO: NewClient()
// TODO: Implement Multi Version

package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

const DefaultMOURL = "/controller/dc/v3"

// Client is the main entry point
type Client struct {
	BaseURL            *url.URL
	MOURL              string
	insecure           bool
	httpClient         *http.Client
	AuthToken          *Auth
	username           string
	password           string
	skipLoggingPayload bool
	*ServiceManager
}

// RequestOpts customizes the behavior of the provider.Request() method.
type RequestOpts struct {
	// JSONBody, if provided, will be encoded as JSON and used as the body of the HTTP request. The
	// content type of the request will default to "application/json" unless overridden by MoreHeaders.
	// It's an error to specify both a JSONBody and a RawBody.
	JSONBody []byte
	//// RawBody contains an io.ReadSeeker that will be consumed by the request directly. No content-type
	//// will be set unless one is provided explicitly by MoreHeaders.
	//RawBody io.ReadSeeker

	// JSONResponse, if provided, will be populated with the contents of the response body parsed as
	// JSON.
	JSONResponse interface{}
	//// OkCodes contains a list of numeric HTTP status codes that should be interpreted as success. If
	//// the response has a different code, an error will be returned.
	//OkCodes []int

	//// MoreHeaders specifies additional HTTP headers to be provide on the request. If a header is
	//// provided with a blank value (""), that header will be *omitted* instead: use this to suppress
	//// the default Accept header or an inferred Content-Type, for example.
	//MoreHeaders map[string]string
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
		MOURL:    DefaultMOURL,
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
	client.ServiceManager = NewServiceManager(client.MOURL, client)
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

func (c *Client) MakeRestRequest(method string, path string, options RequestOpts/*, authenticated bool*/) (*http.Response, error) {
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

	// TODO Authentication
	//if authenticated {
	//	req, err = c.InjectAuthenticationHeader(req, rpath)
	//	if err != nil {
	//		return req, err
	//	}
	//}

	//if !c.skipLoggingPayload {
	//	log.Printf("HTTP request after injection %s %s %v", method, path, req)
	//}

	// Populate the request headers. Apply options.MoreHeaders last, to give the caller the chance to
	// modify or omit any header.
	if contentType != nil {
		req.Header.Set("Content-Type", *contentType)
	}

	req.Header.Set("Accept", applicationJSON)

	//for k, v := range client.AuthenticatedHeaders() {
	//	req.Header.Add(k, v)
	//}
	//
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

	//if resp.StatusCode == http.StatusUnauthorized {
	//	if client.ReauthFunc != nil {
	//		err = client.ReauthFunc()
	//		if err != nil {
	//			return nil, fmt.Errorf("Error trying to re-authenticate: %s", err)
	//		}
	//		if options.RawBody != nil {
	//			options.RawBody.Seek(0, 0)
	//		}
	//		resp.Body.Close()
	//		resp, err = client.Request(method, url, options)
	//		if err != nil {
	//			return nil, fmt.Errorf("Successfully re-authenticated, but got error executing request: %s", err)
	//		}
	//
	//		return resp, nil
	//	}
	//}
	//
	//// Allow default OkCodes if none explicitly set
	//if options.OkCodes == nil {
	//	options.OkCodes = defaultOkCodes(method)
	//}
	//
	//// Validate the HTTP response status.
	//var ok bool
	//for _, code := range options.OkCodes {
	//	if resp.StatusCode == code {
	//		ok = true
	//		break
	//	}
	//}
	//if !ok {
	//	body, _ := ioutil.ReadAll(resp.Body)
	//	resp.Body.Close()
	//	return resp, &UnexpectedResponseCodeError{
	//		URL:      url,
	//		Method:   method,
	//		Expected: options.OkCodes,
	//		Actual:   resp.StatusCode,
	//		Body:     body,
	//	}
	//}

	// Parse the response body as JSON, if requested to do so.
	if options.JSONResponse != nil {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(options.JSONResponse); err != nil {
			return nil, err
		}
	}

	return resp, nil
}

//func (c *Client) Do(req *http.Request) (*container.Container, *http.Response, error) {
//	log.Printf("[DEBUG] Begining DO method %s", req.URL.String())
//
//	if !c.skipLoggingPayload {
//		log.Printf("\n\n\n HTTP request: %v", req.Body)
//	}
//
//	resp, err := c.httpClient.Do(req)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	log.Printf("\nHTTP Request: %s %s", req.Method, req.URL.String())
//	if !c.skipLoggingPayload {
//		log.Printf("\nHTTP Response: %d %s %v", resp.StatusCode, resp.Status, resp)
//	} else {
//		log.Printf("\nHTTP Response: %d %s", resp.StatusCode, resp.Status)
//	}
//
//	bodyBytes, err := ioutil.ReadAll(resp.Body)
//	bodyStr := string(bodyBytes)
//	resp.Body.Close()
//
//	if !c.skipLoggingPayload {
//		log.Printf("\n HTTP response unique string %s %s %s", req.Method, req.URL.String(), bodyStr)
//	}
//	obj, err := container.ParseJSON(bodyBytes)
//
//	if err != nil {
//
//		log.Printf("Error occured while json parsing %+v", err)
//		return nil, resp, err
//	}
//	log.Printf("[DEBUG] Exit from do method")
//
//	return obj, resp, err
//}
