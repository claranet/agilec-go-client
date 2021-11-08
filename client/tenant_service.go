package client

import (
	"github.com/outscope-solutions/acdn-go-client/models"
)

const TenantModuleURL = "/controller/dc/v3/tenants"

// TODO Tenant Name Already Exists Error
func (sm *ServiceManager) CreateTenant(tenant *models.Tenant) (*models.Tenant, error) {
	_, err := sm.Post(models.TenantModuleName, TenantModuleURL, tenant, nil)
	return tenant, err
}

func (sm *ServiceManager) GetTenant(id string) (*models.Tenant, error) {
	var response models.TenantResponseBody
	opts := &RequestOpts{}
	opts.JSONResponse = &response

	_, err := sm.Get(models.TenantModuleName, TenantModuleURL, id, opts)
	if err != nil {
		return nil, err
	}
	tenant := models.TenantFromResponse(opts.JSONResponse.(*models.TenantResponseBody))
	return tenant, nil
}

func (sm *ServiceManager) DeleteTenant(id string) error {
	_, err := sm.Del(models.TenantModuleName, TenantModuleURL, id, nil)
	return err
}

func (sm *ServiceManager) UpdateTenant(id string, tenant *models.Tenant) (*models.Tenant, error) {
	_, err := sm.Put(models.TenantModuleName, TenantModuleURL, id, tenant, nil)
	return tenant, err
}

func (sm *ServiceManager) GetTenants(queryParmeters *models.TenantRequestParameters) (*models.TenantResponseBody, error) {
	var response models.TenantResponseBody
	opts := &RequestOpts{}
	opts.JSONResponse = &response

	_, err := sm.List(TenantModuleURL, opts, queryParmeters)
	tenantResponse := opts.JSONResponse.(*models.TenantResponseBody)

	return tenantResponse, err
}
