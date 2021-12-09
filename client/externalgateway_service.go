package client

import (
	"github.com/outscope-solutions/agilec-go-client/models"
	log "github.com/sirupsen/logrus"
)

const ExternalGatewayModuleURL = "/controller/dc/v3/publicservice/external-gateways"

func (sm *ServiceManager) GetExternalGateway(id string) (*models.ExternalGateway, error) {
	log.Debug("Begin Get External Gateway")
	var response models.ExternalGatewayResponse

	_, err := sm.Get(models.ExternalGatewayModuleName, ExternalGatewayModuleURL, id, &RequestOpts{
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.ExternalGateways[0], nil
}

func (sm *ServiceManager) ListExternalGateways(queryParameters *models.ExternalGatewayRequestOpts) ([]*models.ExternalGateway, error) {
	log.Debug("Begin Get External Gateway")
	var response models.ExternalGatewayResponse

	_, err := sm.List(ExternalGatewayModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response:    &response,
	})

	return response.ExternalGateways, err
}
