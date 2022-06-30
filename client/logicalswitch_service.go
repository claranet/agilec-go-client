package client

import (
	"github.com/claranet/agilec-go-client/models"
	log "github.com/sirupsen/logrus"
)

const LogicalSwitchModuleURL = "/controller/dc/v3/logicnetwork/switchs"

func (sm *ServiceManager) CreateLogicalSwitch(id, name *string, logicalSwitchAttr *models.LogicalSwitchAttributes) error {
	log.Debug("Begin Create Logical Switch")
	logicalSwitch := models.NewLogicalSwitch(id, name, *logicalSwitchAttr)
	_, err := sm.Post(LogicalSwitchModuleURL,
		&RequestOpts{
			Body: models.LogicalSwitchList{
				LogicalSwitches: []*models.LogicalSwitch{logicalSwitch},
			},
		})
	return err
}

func (sm *ServiceManager) GetLogicalSwitch(id string) (*models.LogicalSwitch, error) {
	log.Debug("Begin Get Logical Switch")
	var response models.LogicalSwitchResponse

	_, err := sm.Get(models.LogicalSwitchModuleName, LogicalSwitchModuleURL, id, &RequestOpts{
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.LogicalSwitches[0], nil
}

func (sm *ServiceManager) DeleteLogicalSwitch(id string) error {
	log.Debug("Begin Delete Logical Switch")
	_, err := sm.Del(models.LogicalSwitchModuleName, LogicalSwitchModuleURL, id, nil)
	return err
}

func (sm *ServiceManager) UpdateLogicalSwitch(id, name *string, logicalSwitchAttr *models.LogicalSwitchAttributes) (*models.LogicalSwitch, error) {
	log.Debug("Begin Update Logical Switch")
	var response models.LogicalSwitchResponse
	logicalSwitch := models.NewLogicalSwitch(id, name, *logicalSwitchAttr)

	_, err := sm.Put(models.LogicalSwitchModuleName, LogicalSwitchModuleURL, *logicalSwitch.Id, &RequestOpts{
		Body: models.LogicalSwitchList{
			LogicalSwitches: []*models.LogicalSwitch{logicalSwitch},
		},
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.LogicalSwitches[0], nil
}

func (sm *ServiceManager) ListLogicalSwitches(queryParameters *models.LogicalSwitchRequestOpts) ([]*models.LogicalSwitch, error) {
	log.Debug("Begin Get Logical Switches")
	var response models.LogicalSwitchResponse

	_, err := sm.List(LogicalSwitchModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response:    &response,
	})

	return response.LogicalSwitches, err
}
