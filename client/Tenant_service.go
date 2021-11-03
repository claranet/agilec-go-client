package client

import (
	"github.com/outscope-solutions/acdn-go-client/models"
)

const URL = "/controller/dc/v3/tenants"

func (sm *ServiceManager) CreateTenant(id string, tenantAttr models.TenantAttributes) (*models.Tenant, error) {
	tenant := models.NewTenant(id, tenantAttr)
	_, err := sm.Post(tenant.BaseAttributes.ClassName, URL, tenant, nil, nil)
	return tenant, err
}

//func (sm *ServiceManager) ReadTenant(id int) (*models.Tenant, error) {
//	cont, err := sm.Get(URL, id)
//	if err != nil {
//		return nil, err
//	}
//	tenant := models.TenantFromContainer(cont)
//	return tenant, nil
//}
