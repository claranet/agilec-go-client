package client

import (
	"github.com/outscope-solutions/acdn-go-client/models"
)

const TenantModuleURL = "/controller/dc/v3/tenants"

func (sm *ServiceManager) CreateTenant(tenant *models.Tenant) error {
	_, err := sm.Post(TenantModuleURL,
		&RequestOpts{
			Body:     models.TenantList{
				Tenant: []models.Tenant{*tenant},
			},
		})
	return err
}

func (sm *ServiceManager) GetTenant(id string) (*models.Tenant, error) {
	var response models.TenantResponseBody

	_, err := sm.Get(models.TenantModuleName, TenantModuleURL, id, &RequestOpts{
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.Tenant[0], nil
}

func (sm *ServiceManager) DeleteTenant(id string) error {
	_, err := sm.Del(models.TenantModuleName, TenantModuleURL, id, nil)
	return err
}

func (sm *ServiceManager) UpdateTenant(id string, tenant *models.Tenant) error {
	_, err := sm.Put(models.TenantModuleName, TenantModuleURL, id, &RequestOpts{
		Body:  models.TenantList{
			Tenant: []models.Tenant{*tenant},
		}})
	return err
}

func (sm *ServiceManager) GetTenants(queryParameters *models.TenantRequestOpts) (*models.TenantResponseBody, error) {
	var response models.TenantResponseBody

	_, err := sm.List(TenantModuleURL, &RequestOpts{
		QueryString: queryParameters,
		Response: &response,
	})

	return &response, err
}
