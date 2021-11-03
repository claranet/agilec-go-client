package client

import (
	"github.com/outscope-solutions/acdn-go-client/models"
)

const URL = "/controller/dc/v3/tenants"

// TODO Tenant Name Already Exists Error
func (sm *ServiceManager) CreateTenant(id string, tenantAttr models.TenantAttributes) (*models.Tenant, error) {
	tenant := models.NewTenant(id, tenantAttr)
	_, err := sm.Post(tenant.BaseAttributes.ClassName, URL, tenant, nil)
	return tenant, err
}

func (sm *ServiceManager) GetTenant(id string) (*models.Tenant, error) {
	var response models.TenantBody
	opts := &RequestOpts{}
	opts.JSONResponse = &response

	_, err := sm.Get("tenant", URL , id, opts )
	if err != nil {
		return nil, err
	}
	tenant := models.TenantFromResponse(opts.JSONResponse.(*models.TenantBody))
	return tenant, nil
}

func (sm *ServiceManager) DeleteTenant(id string) error {
	_, err := sm.Del("tenant", URL , id, nil )
	return err
}

func (sm *ServiceManager) UpdateTenant(id string, tenantAttr models.TenantAttributes) (*models.Tenant, error) {
	tenant := models.NewTenant(id, tenantAttr)
	_, err := sm.Put(tenant.BaseAttributes.ClassName, URL , id, tenant, nil )
	return tenant, err
}