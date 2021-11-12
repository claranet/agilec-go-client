package client

import (
	"github.com/outscope-solutions/acdcn-go-client/models"
	log "github.com/sirupsen/logrus"
)

const TenantModuleURL = "/controller/dc/v3/tenants"

func (sm *ServiceManager) CreateTenant(id, name string, tenantAttr *models.TenantAttributes) (*models.Tenant, error) {
	log.Debug("Begin Create Tenant")
	tenant := models.NewTenant(id, name, *tenantAttr)
	_, err := sm.Post(TenantModuleURL,
		&RequestOpts{
			Body: models.TenantList{
				Tenants: []models.Tenant{*tenant},
			},
		})
	return tenant, err
}

func (sm *ServiceManager) GetTenant(id string) (*models.Tenant, error) {
	log.Debug("Begin Get Tenant")
	var response models.TenantResponse

	_, err := sm.Get(models.TenantModuleName, TenantModuleURL, id, &RequestOpts{
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.Tenants[0], nil
}

func (sm *ServiceManager) DeleteTenant(id string) error {
	log.Debug("Begin Delete Tenant")
	_, err := sm.Del(models.TenantModuleName, TenantModuleURL, id, nil)
	return err
}

func (sm *ServiceManager) UpdateTenant(id, name string, tenantAttr *models.TenantAttributes) (*models.Tenant, error) {
	log.Debug("Begin Update Tenant")
	tenant := models.NewTenant(id, name, *tenantAttr)
	_, err := sm.Put(models.TenantModuleName, TenantModuleURL, tenant.Id, &RequestOpts{
		Body: models.TenantList{
			Tenants: []models.Tenant{*tenant},
		}})
	return tenant, err
}

// TODO Return []Tenants
func (sm *ServiceManager) GetTenants(queryParameters *models.TenantRequestOpts) (*[]models.Tenant, error) {
	log.Debug("Begin Get Tenants")
	var response models.TenantResponse

	_, err := sm.List(TenantModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response:    &response,
	})

	return &response.Tenants, err
}
