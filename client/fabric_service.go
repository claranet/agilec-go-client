package client

import (
	"github.com/outscope-solutions/acdcn-go-client/models"
	log "github.com/sirupsen/logrus"
)

const FabricModuleURL = "/controller/dc/v3/physicalnetwork/fabricresource/fabrics"

func (sm *ServiceManager) GetFabric(id string) (*models.Fabric, error) {
	log.Debug("Begin Get Fabric")
	var response models.FabricResponse

	_, err := sm.Get(models.FabricModuleName, FabricModuleURL, id, &RequestOpts{
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.Fabrics[0], nil
}

func (sm *ServiceManager) ListFabrics(queryParameters *models.FabricRequestOpts) ([]*models.Fabric, error) {
	log.Debug("Begin Get Fabrics")
	var response models.FabricResponse

	_, err := sm.List(FabricModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response:    &response,
	})

	return response.Fabrics, err
}