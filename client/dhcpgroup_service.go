package client

import (
	"github.com/outscope-solutions/agilec-go-client/models"
	log "github.com/sirupsen/logrus"
)

const DHCPGroupModuleURL = "/controller/dc/v3/publicservice/dhcpgroups"

func (sm *ServiceManager) GetDHCPGroup(id string) (*models.DHCPGroup, error) {
	log.Debug("Begin Get DHCP Group")
	var response models.DHCPGroupResponse

	_, err := sm.Get(models.DHCPGroupModuleName, DHCPGroupModuleURL, id, &RequestOpts{
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.DHCPGroups[0], nil
}

func (sm *ServiceManager) ListDHCPGroups(queryParameters *models.DHCPGroupRequestOpts) ([]*models.DHCPGroup, error) {
	log.Debug("Begin Get DHCP Group")
	var response models.DHCPGroupResponse

	_, err := sm.List(DHCPGroupModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response:    &response,
	})

	return response.DHCPGroups, err
}
