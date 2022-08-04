package client

import (
	"github.com/claranet/agilec-go-client/models"
	log "github.com/sirupsen/logrus"
	"strings"
)

const EndPortModuleURL = "/controller/dc/v3/logicnetwork/endports"

func (sm *ServiceManager) CreateEndPort(id, name *string, endPortAttr *models.EndPortAttributes) error {
	log.Debug("Begin Create End Port")
	endPort := models.NewEndPort(id, name, *endPortAttr)
	_, err := sm.Post(EndPortModuleURL,
		&RequestOpts{
			Body: models.EndPortResponse{
				EndPort: *endPort,
			},
		})
	return err
}

func (sm *ServiceManager) GetEndPort(id string) (*models.EndPort, error) {
	log.Debug("Begin Get End Port")
	var response models.EndPortResponse

	if _, err := sm.Get(strings.ToLower(models.EndPortModuleName), EndPortModuleURL, id, &RequestOpts{
		Response: &response,
	}); err != nil {
		return nil, err
	}

	return &response.EndPort, nil
}

func (sm *ServiceManager) DeleteEndPort(id string) error {
	log.Debug("Begin Delete End Port")
	_, err := sm.Del(strings.ToLower(models.EndPortModuleName), EndPortModuleURL, id, nil)
	return err
}

func (sm *ServiceManager) UpdateEndPort(id, name *string, endPortAttr *models.EndPortAttributes) (*models.EndPort, error) {
	log.Debug("Begin Update End Port")
	var response models.EndPortResponse
	endPort := models.NewEndPort(id, name, *endPortAttr)

	_, err := sm.Put(strings.ToLower(models.EndPortModuleName), EndPortModuleURL, *endPort.Id, &RequestOpts{
		Body: models.EndPortResponse{
			EndPort: *endPort,
		},
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.EndPort, nil
}

func (sm *ServiceManager) ListEndPorts(queryParameters *models.EndPortRequestOpts) ([]*models.EndPort, error) {
	log.Debug("Begin Get End Ports")
	var response models.EndPortListResponse

	_, err := sm.List(EndPortModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response:    &response,
	})

	return response.EndPorts, err
}
