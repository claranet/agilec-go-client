// TODO: NewClient()
// TODO: Implement Multi Version

package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/outscope-solutions/acdn-go-client/container"
	"io/ioutil"
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

func (c *Client) MakeRestRequest(method string, path string, body *container.Container, authenticated bool) (*http.Request, error) {

	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	fURL, err := url.Parse(c.BaseURL.String())
	if err != nil {
		return nil, err
	}

	fURL = fURL.ResolveReference(pathURL)

	var req *http.Request
	log.Printf("[DEBUG] BaseURL: %s, pathURL: %s, fURL: %s", c.BaseURL.String(), pathURL.String(), fURL.String())

	if method == "GET" {
		req, err = http.NewRequest(method, fURL.String(), nil)
	} else {
		req, err = http.NewRequest(method, fURL.String(), bytes.NewBuffer(body.Bytes()))
	}

	fmt.Println(req.Body)

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

	if !c.skipLoggingPayload {
		log.Printf("HTTP request after injection %s %s %v", method, path, req)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request) (*container.Container, *http.Response, error) {
	log.Printf("[DEBUG] Begining DO method %s", req.URL.String())

	if !c.skipLoggingPayload {
		log.Printf("\n\n\n HTTP request: %v", req.Body)
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, nil, err
	}

	log.Printf("\nHTTP Request: %s %s", req.Method, req.URL.String())
	if !c.skipLoggingPayload {
		log.Printf("\nHTTP Response: %d %s %v", resp.StatusCode, resp.Status, resp)
	} else {
		log.Printf("\nHTTP Response: %d %s", resp.StatusCode, resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyStr := string(bodyBytes)
	resp.Body.Close()

	if !c.skipLoggingPayload {
		log.Printf("\n HTTP response unique string %s %s %s", req.Method, req.URL.String(), bodyStr)
	}
	obj, err := container.ParseJSON(bodyBytes)

	if err != nil {

		log.Printf("Error occured while json parsing %+v", err)
		return nil, resp, err
	}
	log.Printf("[DEBUG] Exit from do method")

	return obj, resp, err
}
