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

func NewServiceManager(moURL string, client *Client) *ServiceManager {

	sm := &ServiceManager{
		//MOURL:  moURL,
		client: client,
	}
	return sm
}

func (opts *RequestOpts) PrepareBody(classname string, obj models.Model) ([]byte, error) {
	cont, err := obj.ToJson()
	if err != nil {
		return nil, err
	}
	body, err := createJsonPayload(classname, cont)
	return body, err
}

func (opts *RequestOpts) setBody(classname string, obj models.Model) error {
	body, err := opts.PrepareBody(classname, obj)

	if err != nil {
		return err
	}

	opts.JSONBody = body
	return nil
}

func createJsonPayload(classname string, payload []byte) ([]byte, error) {
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": [%s]
	}`, classname, payload))

	body, err := container.ParseJSON(containerJSON)

	return body.Bytes(), err
}

func (sm *ServiceManager) Post(classname string, url string, obj models.Model, JSONResponse *interface{}, opts *RequestOpts) (*http.Response, error) {

	if opts == nil {
		opts = &RequestOpts{}
	}

	err := opts.setBody(classname, obj)

	if err != nil {
		return nil, err
	}

	//req, err := sm.client.MakeRestRequest("POST", url, *opts /*, true*/)
	//if err != nil {
	//	return err
	//}

	return sm.client.MakeRestRequest("POST", url, *opts /*, true*/)
	//return CheckForErrors(cont, "POST", sm.client.skipLoggingPayload)
}

//func (sm *ServiceManager) Get(url string, id int) (*container.Container, error) {
//	finalURL := fmt.Sprintf("%s/%d", url, id)
//	req, err := sm.client.MakeRestRequest("GET", finalURL, nil/*, true*/)
//
//	if err != nil {
//		return nil, err
//	}
//
//	obj, _, err := sm.client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	if obj == nil {
//		return nil, errors.New("Empty response body")
//	}
//	log.Printf("[DEBUG] Exit from GET %s", finalURL)
//	return obj, CheckForErrors(obj, "GET", sm.client.skipLoggingPayload)
//}

// TODO
//func CheckForErrors(cont *container.Container, method string, skipLoggingPayload bool) error {
	//number, err := strconv.Atoi(models.G(cont, "totalCount"))
	//if err != nil {
	//	if !skipLoggingPayload {
	//		log.Printf("[DEBUG] Exit from errors, Unable to parse error count from response %v", cont)
	//	} else {
	//		log.Printf("[DEBUG] Exit from errors %s", err.Error())
	//	}
	//	return err
	//}
	//imdata := cont.S("imdata").Index(0)
	//if number > 0 {
	//
	//	if imdata.Exists("error") {
	//
	//		if models.StripQuotes(imdata.Path("error.attributes.code").String()) == "103" {
	//			if !skipLoggingPayload {
	//				log.Printf("[DEBUG] Exit from error 103 %v", cont)
	//			}
	//			return nil
	//		} else {
	//			if models.StripQuotes(imdata.Path("error.attributes.text").String()) == "" && models.StripQuotes(imdata.Path("error.attributes.code").String()) == "403" {
	//				if !skipLoggingPayload {
	//					log.Printf("[DEBUG] Exit from authentication error 403 %v", cont)
	//				}
	//				return errors.New("Unable to authenticate. Please check your credentials")
	//			}
	//			if !skipLoggingPayload {
	//				log.Printf("[DEBUG] Exit from errors %v", cont)
	//			}
	//
	//			return errors.New(models.StripQuotes(imdata.Path("error.attributes.text").String()))
	//		}
	//	}
	//
	//}
	//
	//if imdata.String() == "{}" && method == "GET" {
	//	if !skipLoggingPayload {
	//		log.Printf("[DEBUG] Exit from error (Empty response) %v", cont)
	//	}
	//
	//	return errors.New("Error retrieving Object: Object may not exists")
	//}
	//if !skipLoggingPayload {
	//	log.Printf("[DEBUG] Exit from errors %v", cont)
	//}
//	return nil
//}
