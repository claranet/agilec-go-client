package client

import (
	"github.com/outscope-solutions/acdn-go-client/models"
)

const URL = "/controller/dc/v3/tenants"

// TODO Tenant Name Already Exists Error
func (sm *ServiceManager) CreateTenant(tenant *models.Tenant) (*models.Tenant, error) {
	_, err := sm.Post(models.TenantModuleName, URL, tenant, nil)
	return tenant, err
}

func (sm *ServiceManager) GetTenant(id string) (*models.Tenant, error) {
	var response models.TenantResponseBody
	opts := &RequestOpts{}
	opts.JSONResponse = &response

	_, err := sm.Get(models.TenantModuleName, URL, id, opts)
	if err != nil {
		return nil, err
	}
	tenant := models.TenantFromResponse(opts.JSONResponse.(*models.TenantResponseBody))
	return tenant, nil
}

func (sm *ServiceManager) DeleteTenant(id string) error {
	_, err := sm.Del(models.TenantModuleName, URL, id, nil)
	return err
}

func (sm *ServiceManager) UpdateTenant(id string, tenant *models.Tenant) (*models.Tenant, error) {
	_, err := sm.Put(models.TenantModuleName, URL, id, tenant, nil)
	return tenant, err
}

func (sm *ServiceManager) GetTenants(queryParmeters *models.TenantRequestParameters) (*models.TenantResponseBody, error) {
	var response models.TenantResponseBody
	opts := &RequestOpts{}
	opts.JSONResponse = &response

	_, err := sm.List(URL, opts, queryParmeters)
	tenantResponse := opts.JSONResponse.(*models.TenantResponseBody)

	return tenantResponse, err
}
