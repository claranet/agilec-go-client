package client

import (
	"errors"
	"log"

	//"errors"
	"fmt"
	"github.com/outscope-solutions/acdn-go-client/container"
	"github.com/outscope-solutions/acdn-go-client/models"
	//"log"
	//"strconv"
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

//func createJsonPayload(payload map[string]interface{}) (*container.Container, error) {
func createJsonPayload(classname string, payload []byte) (*container.Container, error) {
	//containerJSON := []byte(fmt.Sprintf(`{
	//	"%s": [{
	//
	//	}]
	//}`, payload["classname"]))
	//
	//return container.ParseJSON(containerJSON)

	containerJSON := []byte(fmt.Sprintf(`{
		"%s": [%s]
	}`, classname, payload))

	return container.ParseJSON(containerJSON)
}

func (sm *ServiceManager) Save(classname string, url string, obj models.Model) error {

	jsonPayload, err := sm.PrepareModel(classname, obj)

	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", url, jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	return CheckForErrors(cont, "POST", sm.client.skipLoggingPayload)
}

func (sm *ServiceManager) Get(url string, id int) (*container.Container, error) {
	finalURL := fmt.Sprintf("%s/%d", url, id)
	req, err := sm.client.MakeRestRequest("GET", finalURL, nil, true)

	if err != nil {
		return nil, err
	}

	obj, _, err := sm.client.Do(req)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, errors.New("Empty response body")
	}
	log.Printf("[DEBUG] Exit from GET %s", finalURL)
	return obj, CheckForErrors(obj, "GET", sm.client.skipLoggingPayload)
}

func (sm *ServiceManager) PrepareModel(classname string, obj models.Model) (*container.Container, error) {
	cont, err := obj.ToMap()
	if err != nil {
		return nil, err
	}
	jsonPayload, err := createJsonPayload(classname, cont)
	return jsonPayload, err
}

// TODO
func CheckForErrors(cont *container.Container, method string, skipLoggingPayload bool) error {
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
	return nil
}
