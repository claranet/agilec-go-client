package client

import (
	"fmt"
	"github.com/outscope-solutions/acdn-go-client/container"
	"github.com/outscope-solutions/acdn-go-client/models"
	"net/http"
)

type ServiceManager struct {
	//MOURL  string
	client *Client
}

// RequestOpts customizes the behavior of the provider.Request() method.
type RequestOpts struct {

	QueryStringParameters map[string]interface{}

	// JSONBody, if provided, will be encoded as JSON and used as the body of the HTTP request. The
	// content type of the request will default to "application/json" unless overridden by MoreHeaders.
	// It's an error to specify both a JSONBody and a RawBody.
	JSONBody []byte

	// JSONResponse, if provided, will be populated with the contents of the response body parsed as
	// JSON.
	JSONResponse interface{}

	// OkCodes contains a list of numeric HTTP status codes that should be interpreted as success. If
	// the response has a different code, an error will be returned.
	OkCodes []int

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

func (err *UnexpectedResponseCodeError) Error() string {
	return fmt.Sprintf(
		"Expected HTTP response code %v when accessing [%s %s], but got %d instead\n%s",
		err.Expected, err.Method, err.URL, err.Actual, err.Body,
	)
}

func createJsonPayload(modulename string, payload []byte) ([]byte, error) {
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": [%s]
	}`, modulename, payload))

	body, err := container.ParseJSON(containerJSON)

	return body.Bytes(), err
}

func (opts *RequestOpts) PrepareBody(modulename string, obj models.Model) ([]byte, error) {
	cont, err := obj.ToJson()
	if err != nil {
		return nil, err
	}
	body, err := createJsonPayload(modulename, cont)
	return body, err
}

func (opts *RequestOpts) setBody(modulename string, obj models.Model) error {
	body, err := opts.PrepareBody(modulename, obj)

	if err != nil {
		return err
	}

	opts.JSONBody = body
	return nil
}

func NewServiceManager(client *Client) *ServiceManager {

	sm := &ServiceManager{
		//MOURL:  moURL,
		client: client,
	}
	return sm
}

func (sm *ServiceManager) Post(modulename, url string, obj models.Model, opts *RequestOpts) (*http.Response, error) {

	if opts == nil {
		opts = &RequestOpts{}
	}

	err := opts.setBody(modulename, obj)

	if err != nil {
		return nil, err
	}

	return sm.client.MakeRestRequest("POST", url, *opts, true)
}

func (sm *ServiceManager) Get(modulename, url string, id string, opts *RequestOpts) (*http.Response, error) {
	fURL := fmt.Sprintf("%s/%s/%s", url, modulename, id)
	return sm.client.MakeRestRequest("GET", fURL, *opts, true)
}

func (sm *ServiceManager) Del(modulename, url string, id string, opts *RequestOpts) (*http.Response, error) {
	fURL := fmt.Sprintf("%s/%s/%s", url, modulename, id)
	if opts == nil {
		opts = &RequestOpts{}
	}
	return sm.client.MakeRestRequest("DELETE", fURL, *opts, true)
}

func (sm *ServiceManager) Put(modulename, url, id string, obj models.Model, opts *RequestOpts) (*http.Response, error) {

	fURL := fmt.Sprintf("%s/%s/%s", url, modulename, id)

	if opts == nil {
		opts = &RequestOpts{}
	}

	err := opts.setBody(modulename, obj)

	if err != nil {
		return nil, err
	}

	return sm.client.MakeRestRequest("PUT", fURL, *opts, true)
}

func (sm *ServiceManager) List(url string, opts *RequestOpts, queryParameters interface{}) (*http.Response, error) {

	if opts == nil {
		opts = &RequestOpts{}
	}

	if queryParameters != nil && opts.QueryStringParameters == nil {
		queryParametersMap , err := StructToMap(queryParameters)

		if err != nil {
			return nil, err
		}

		opts.QueryStringParameters = queryParametersMap
	}

	return sm.client.MakeRestRequest("GET", url, *opts, true)
}
