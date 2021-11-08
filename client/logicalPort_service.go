package client

import (
	"github.com/outscope-solutions/acdn-go-client/models"
)

const LogicalPortModuleURL = "/controller/dc/v3/logicnetwork/ports"

//func (sm *ServiceManager) CreateTenant(tenant *models.Tenant) (*models.Tenant, error) {
//	_, err := sm.Post(models.TenantModuleName, LogicalPortModuleURL, tenant, nil)
//	return tenant, err
//}
//
//func (sm *ServiceManager) GetTenant(id string) (*models.Tenant, error) {
//	var response models.TenantResponseBody
//	opts := &RequestOpts{}
//	opts.JSONResponse = &response
//
//	_, err := sm.Get(models.TenantModuleName, LogicalPortModuleURL, id, opts)
//	if err != nil {
//		return nil, err
//	}
//	tenant := models.TenantFromResponse(opts.JSONResponse.(*models.TenantResponseBody))
//	return tenant, nil
//}
//
//func (sm *ServiceManager) DeleteTenant(id string) error {
//	_, err := sm.Del(models.TenantModuleName, LogicalPortModuleURL, id, nil)
//	return err
//}
//
//func (sm *ServiceManager) UpdateTenant(id string, tenant *models.Tenant) (*models.Tenant, error) {
//	_, err := sm.Put(models.TenantModuleName, LogicalPortModuleURL, id, tenant, nil)
//	return tenant, err
//}

func (sm *ServiceManager) GetLogicalPorts(queryParmeters *models.LogicalPortListOpts) (*models.LogicalPortResponseBody, error) {
	var response models.LogicalPortResponseBody
	opts := &RequestOpts{}
	opts.JSONResponse = &response

	_, err := sm.List(LogicalPortModuleURL, opts, queryParmeters)
	logicalPortResponse := opts.JSONResponse.(*models.LogicalPortResponseBody)

	return logicalPortResponse, err
}

