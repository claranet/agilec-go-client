package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type ServiceManager struct {
	client *Client
}

// RequestOpts customizes the behavior of the provider.Request() method.
type RequestOpts struct {
	QueryString interface{}

	// JSONBody, if provided, will be encoded as JSON and used as the body of the HTTP request. The
	// content type of the request will default to "application/json" unless overridden by MoreHeaders.
	// It's an error to specify both a JSONBody and a RawBody.
	Body interface{}

	// JSONResponse, if provided, will be populated with the contents of the response body parsed as
	// JSON.
	Response interface{}

	// OkCodes contains a list of numeric HTTP status codes that should be interpreted as success. If
	// the response has a different code, an error will be returned.
	//OkCodes []int

	//// MoreHeaders specifies additional HTTP headers to be provide on the request. If a header is
	//// provided with a blank value (""), that header will be *omitted* instead: use this to suppress
	//// the default Accept header or an inferred Content-Type, for example.
	//MoreHeaders map[string]string
}

type ErrorResponse struct {
	ErrorMessage   string
	ErrorCode      string
	URL            string
	Method         string
	HttpStatusCode int
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf(
		"HTTP Error response status code %d when accessing [%s %s]. Error Message: %s - Error Code: %s",
		err.HttpStatusCode, err.Method, err.URL, err.ErrorMessage, err.ErrorCode,
	)
}

func NewServiceManager(client *Client) *ServiceManager {
	sm := &ServiceManager{
		client: client,
	}
	return sm
}

func (sm *ServiceManager) Post(url string, opts *RequestOpts) (*resty.Response, error) {
	log.Debug("Creating Post Request")
	return sm.client.Request("Post", url, *opts)
}

func (sm *ServiceManager) Get(modulename, url, id string, opts *RequestOpts) (*resty.Response, error) {
	log.Debug("Creating Get Request")
	fURL := fmt.Sprintf("%s/%s/%s", url, modulename, id)
	return sm.client.Request("Get", fURL, *opts)
}

func (sm *ServiceManager) Del(modulename, url, id string, opts *RequestOpts) (*resty.Response, error) {
	log.Debug("Creating Delete Request")
	fURL := fmt.Sprintf("%s/%s/%s", url, modulename, id)
	reqOpts := &RequestOpts{}
	if opts != nil {
		reqOpts = opts
	}
	return sm.client.Request("Delete", fURL, *reqOpts)
}

func (sm *ServiceManager) Put(modulename, url, id string, opts *RequestOpts) (*resty.Response, error) {
	log.Debug("Creating Put Request")
	fURL := fmt.Sprintf("%s/%s/%s", url, modulename, id)
	return sm.client.Request("Put", fURL, *opts)
}

func (sm *ServiceManager) List(url string, opts *RequestOpts) (*resty.Response, error) {
	log.Debug("Creating Get Request")
	return sm.client.Request("Get", url, *opts)
}
