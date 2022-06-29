package client

import (
	"github.com/claranet/agilec-go-client/models"
	log "github.com/sirupsen/logrus"
)

const LogicalRouterModuleURL = "/controller/dc/v3/logicnetwork/routers"

func (sm *ServiceManager) CreateLogicalRouter(id, name *string, logicalRouterAttr *models.LogicalRouterAttributes) error {
	log.Debug("Begin Create Logical Router")
	logicalRouter := models.NewLogicalRouter(id, name, *logicalRouterAttr)
	_, err := sm.Post(LogicalRouterModuleURL,
		&RequestOpts{
			Body: models.LogicalRouterResponse{
				LogicalRouter: *logicalRouter,
			},
		})
	return err
}

func (sm *ServiceManager) GetLogicalRouter(id string) (*models.LogicalRouter, error) {
	log.Debug("Begin Get Logical Router")
	var response models.LogicalRouterListResponse

	_, err := sm.Get(models.LogicalRouterModuleName, LogicalRouterModuleURL, id, &RequestOpts{
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.LogicalRouters[0], nil
}

func (sm *ServiceManager) DeleteLogicalRouter(id string) error {
	log.Debug("Begin Delete Logical Router")
	_, err := sm.Del(models.LogicalRouterModuleName, LogicalRouterModuleURL, id, nil)
	return err
}

func (sm *ServiceManager) UpdateLogicalRouter(id, name *string, logicalRouterAttr *models.LogicalRouterAttributes) (*models.LogicalRouter, error) {
	log.Debug("Begin Update Logical Router")
	var response models.LogicalRouterResponse
	logicalRouter := models.NewLogicalRouter(id, name, *logicalRouterAttr)

	_, err := sm.Put(models.LogicalRouterModuleName, LogicalRouterModuleURL, *logicalRouter.Id, &RequestOpts{
		Body: models.LogicalRouterResponse{
			LogicalRouter: *logicalRouter,
		},
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.LogicalRouter, nil
}

func (sm *ServiceManager) ListLogicalRouters(queryParameters *models.LogicalRouterRequestOpts) ([]*models.LogicalRouter, error) {
	log.Debug("Begin Get Logical Routers")
	var response models.LogicalRouterListResponse

	_, err := sm.List(LogicalRouterModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response:    &response,
	})

	return response.LogicalRouters, err
}
