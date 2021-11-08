package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
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
	Body interface {}

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

// UnexpectedResponseCodeError is returned by the Request method when a response code other than
// those listed in OkCodes is encountered.
type UnexpectedResponseCodeError struct {
	URL      string
	Method   string
	Expected []int
	Actual   int
	Body     []byte
}

type ErrorResponse struct {
	Errors *Errors                  `json:"errors"`
	//URL      string
	//Method   string
	//Expected []int
	//Actual   int
}

type Errors struct {
	 Error *[]Error `json:"error"`
}

type Error struct {
	ErrorType string `json:"error-type,omitempty"`
	ErrorLevel string `json:"error-level,omitempty"`
	ErrorTag string `json:"error-tag,omitempty"`
	ErrorAppTag string `json:"error-app-tag,omitempty"`
	ErrorPath string `json:"error-path,omitempty"`
	ErrorMessage string `json:"error-message,omitempty"`
	ErrorInfo *ErrorInfo `json:"error-info,omitempty"`
}

type ErrorInfo struct {
	ErrorCode string `json:"error-code,omitempty"`
	ErrorParas []ErrorPara `json:"error-paras,omitempty"`
}

type ErrorPara struct {
	ErrorPara string `json:"error-para,omitempty"`
}

func (err *ErrorResponse) Error() string {
	error := *err.Errors.Error
	return fmt.Sprintf(error[0].ErrorMessage)
	//return fmt.Sprintf(
	//	"Expected HTTP response code %v when accessing [%s %s], but got %d instead\n%s",
	//	err.Expected, err.Method, err.URL, err.Actual, error[0].ErrorMessage,
	//)
}

func NewServiceManager(client *Client) *ServiceManager {

	sm := &ServiceManager{
		client: client,
	}
	return sm
}

func (sm *ServiceManager) Post(url string, opts *RequestOpts) (*resty.Response, error) {
	return sm.client.NewRequest("Post", url, *opts, true)
}

func (sm *ServiceManager) Get(modulename, url, id string, opts *RequestOpts) (*resty.Response, error) {
	fURL := fmt.Sprintf("%s/%s/%s", url, modulename, id)
	return sm.client.NewRequest("Get", fURL, *opts,true)
}

func (sm *ServiceManager) Del(modulename, url, id string, opts *RequestOpts) (*resty.Response, error) {
	fURL := fmt.Sprintf("%s/%s/%s", url, modulename, id)
	reqOpts := &RequestOpts{}
	if opts != nil {
		reqOpts = opts
	}
	return sm.client.NewRequest("Delete", fURL, *reqOpts, true)
}

func (sm *ServiceManager) Put(modulename, url, id string, opts *RequestOpts) (*resty.Response, error) {
	fURL := fmt.Sprintf("%s/%s/%s", url, modulename, id)
	return sm.client.NewRequest("Put", fURL, *opts,true)
}

func (sm *ServiceManager) List(url string, opts *RequestOpts) (*resty.Response, error) {
	return sm.client.NewRequest("Get", url, *opts,true)
}