package tests

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/outscope-solutions/acdn-go-client/models"
	helper "github.com/outscope-solutions/acdn-go-client/tests/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO Make Request without Required Parameters and also parameters with wrong type. The Response is Diferent
func GetTenant() *models.Tenant {
	u, _ := uuid.NewV4()
	fmt.Printf("Tenant ID Generated: %s\n", u.String())
	Tenant := models.Tenant{}
	Tenant.Id = u.String()
	Tenant.Name = "OUTSCOPE-GO-TESTS-001"
	Tenant.Description = "Created By GO"
	Tenant.Producer = "Default"
	Tenant.MulticastCapability = true
	Tenant.Quota = &models.TenantQuota{
		LogicVasNum:    10,
		LogicRouterNum: 5,
		LogicSwitchNum: 6,
	}
	Tenant.MulticastQuota = &models.TenantMulticastQuota{
		AclNum:     10,
		AclRuleNum: 10,
	}
	Tenant.ResPool = &models.TenantResPool{
		ExternalGatewayIds: []string{"1", "2"},
		FabricIds:          []string{"1", "2"},
		VmmIds:             []string{"1", "2"},
		DhcpGroupIds:       []string{"1", "2"},
	}

	return &Tenant
}

func TestCreateTenant (t *testing.T) {
	tenant := GetTenant()
	defer DeleteTenant(tenant.Id)
	client := helper.GetClient()
	err := client.CreateTenant(tenant)
	assert.Nil(t, err)
}

func TestCreateTenantDuplicate (t *testing.T) {
	tenant := GetTenant()
	defer DeleteTenant(tenant.Id)
	client := helper.GetClient()
	err := client.CreateTenant(tenant)
	assert.Nil(t, err)
	err = client.CreateTenant(tenant)
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "The tenant id already exist.", err)
	}
}

func TestCreateTenantInvalidID (t *testing.T) {
	tenant := GetTenant()
	tenant.Id = "dummy"
	client := helper.GetClient()
	err := client.CreateTenant(tenant)
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "Invalid UUID format.", err)
	}
}


//func TestCreateTenantWithoutID (t *testing.T) {  // Make request error for type conversion
//	tenant := GetTenant()
//	tenant.Id = ""
//	client := helper.GetClient()
//	err := client.CreateTenant(tenant)
//	if assert.NotNil(t, err) {
//		assert.EqualError(t, err, "The tenant id already exist.", err)
//	}
//}

func TestUpdateTenant (t *testing.T) {
	tenant := GetTenant()
	defer DeleteTenant(tenant.Id)
	client := helper.GetClient()
	err := client.CreateTenant(tenant)
	tenant.Description = "Updated From GO"
	err = client.UpdateTenant(tenant)
	assert.Nil(t, err)
}

func TestGetTenants (t *testing.T) {
	client := helper.GetClient()
	queryParameters := &models.TenantRequestOpts{}
	queryParameters.Producer = "default"
	queryParameters.PageSize = 2
	response, err := client.GetTenants(queryParameters)
	assert.Greater(t, response.TotalNum, int64(1))
	assert.Nil(t, err)
}

func TestDeleteTenant (t *testing.T) {
	tenant := GetTenant()
	client := helper.GetClient()
	err := client.DeleteTenant(tenant.Id)
	assert.Nil(t, err)
}

func DeleteTenant(id string) {
	client := helper.GetClient()
	_ = client.DeleteTenant(id)
}