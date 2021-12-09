package client

import (
	"github.com/outscope-solutions/agilec-go-client/models"
	log "github.com/sirupsen/logrus"
)

const LogicalNetworkModuleURL = "/controller/dc/v3/logicnetwork/networks"

func (sm *ServiceManager) CreateLogicalNetwork(id, name *string, logicalNetworkAttr *models.LogicalNetworkAttributes) error {
	log.Debug("Begin Create Logical Network")
	logicalNetwork := models.NewLogicalNetwork(id, name, *logicalNetworkAttr)
	_, err := sm.Post(LogicalNetworkModuleURL,
		&RequestOpts{
			Body: models.LogicalNetworkResponse{
				LogicalNetwork: *logicalNetwork,
			},
		})
	return err
}

func (sm *ServiceManager) GetLogicalNetwork(id string) (*models.LogicalNetwork, error) {
	log.Debug("Begin Get Logical Network")
	var response models.LogicalNetworkListResponse

	_, err := sm.Get(models.LogicalNetworkModuleName, LogicalNetworkModuleURL, id, &RequestOpts{
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.LogicalNetworks[0], nil
}

func (sm *ServiceManager) DeleteLogicalNetwork(id string) error {
	log.Debug("Begin Delete Logical Network")
	_, err := sm.Del(models.LogicalNetworkModuleName, LogicalNetworkModuleURL, id, nil)
	return err
}

func (sm *ServiceManager) UpdateLogicalNetwork(id, name *string, logicalNetworkAttr *models.LogicalNetworkAttributes) (*models.LogicalNetwork, error) {
	log.Debug("Begin Update Logical Network")
	var response models.LogicalNetworkResponse
	logicalNetwork := models.NewLogicalNetwork(id, name, *logicalNetworkAttr)

	_, err := sm.Put(models.LogicalNetworkModuleName, LogicalNetworkModuleURL, *logicalNetwork.Id, &RequestOpts{
		Body: models.LogicalNetworkResponse{
			LogicalNetwork: *logicalNetwork,
		},
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.LogicalNetwork, nil
}

func (sm *ServiceManager) ListLogicalNetworks(queryParameters *models.LogicalNetworkRequestOpts) ([]*models.LogicalNetwork, error) {
	log.Debug("Begin Get Logical Networks")
	var response models.LogicalNetworkListResponse

	_, err := sm.List(LogicalNetworkModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response:    &response,
	})

	return response.LogicalNetworks, err
}
