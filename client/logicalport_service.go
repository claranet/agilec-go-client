package client

import (
	"github.com/claranet/agilec-go-client/models"
	log "github.com/sirupsen/logrus"
)

const LogicalPortModuleURL = "/controller/dc/v3/logicnetwork/ports"

func (sm *ServiceManager) CreateLogicalPort(id, name *string, logicalPortAttr *models.LogicalPortAttributes) error {
	log.Debug("Begin Create Logical Port")
	logicalPort := models.NewLogicalPort(id, name, *logicalPortAttr)
	_, err := sm.Post(LogicalPortModuleURL,
		&RequestOpts{
			Body: models.LogicalPortList{
				LogicalPorts: []*models.LogicalPort{logicalPort},
			},
		})
	return err
}

func (sm *ServiceManager) GetLogicalPort(id string) (*models.LogicalPort, error) {
	log.Debug("Begin Get Logical Port")
	var response models.LogicalPortResponse

	if _, err := sm.Get(models.LogicalPortModuleName, LogicalPortModuleURL, id, &RequestOpts{
		Response: &response,
	}); err != nil {
		return nil, err
	}

	return response.LogicalPorts[0], nil
}

func (sm *ServiceManager) DeleteLogicalPort(id string) error {
	log.Debug("Begin Delete Logical Port")
	_, err := sm.Del(models.LogicalPortModuleName, LogicalPortModuleURL, id, nil)
	return err
}

func (sm *ServiceManager) UpdateLogicalPort(id, name *string, logicalPortAttr *models.LogicalPortAttributes) (*models.LogicalPort, error) {
	log.Debug("Begin Update Logical Port")
	var response models.LogicalPortResponse
	logicalPort := models.NewLogicalPort(id, name, *logicalPortAttr)

	_, err := sm.Put(models.LogicalPortModuleName, LogicalPortModuleURL, *logicalPort.Id, &RequestOpts{
		Body: models.LogicalPortList{
			LogicalPorts: []*models.LogicalPort{logicalPort},
		},
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.LogicalPorts[0], nil
}

func (sm *ServiceManager) ListLogicalPorts(queryParameters *models.LogicalPortRequestOpts) ([]*models.LogicalPort, error) {
	log.Debug("Begin Get Logical Ports")
	var response models.LogicalPortResponse

	_, err := sm.List(LogicalPortModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response:    &response,
	})

	return response.LogicalPorts, err
}
