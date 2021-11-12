package tests

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/outscope-solutions/acdn-go-client/models"
	helper "github.com/outscope-solutions/acdn-go-client/tests/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetTenantAttributes() (string, string, *models.TenantAttributes) {
	u, _ := uuid.NewV4()
	fmt.Printf("Tenant ID Generated: %s\n", u.String())
	Id := u.String()
	Name := "OUTSCOPE-GO-TESTS-001"

	Tenant := models.TenantAttributes{}
	Tenant.Description = "Created By GO"
	Tenant.Producer = "GOLANG"
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

	return Id, Name, &Tenant
}

func TestCreateTenant(t *testing.T) {
	id, name, tenantAttr := GetTenantAttributes()
	defer DeleteTenant(id)
	client := helper.GetClient()
	tenant, err := client.CreateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	assert.Equal(t, id, tenant.Id)
	assert.Equal(t, name, tenant.Name)
}

func TestCreateTenantDuplicate(t *testing.T) {
	id, name, tenantAttr := GetTenantAttributes()
	defer DeleteTenant(id)
	client := helper.GetClient()
	_, err := client.CreateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	_ , err = client.CreateTenant(id, name, tenantAttr)
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "HTTP Error response status code 400 when accessing [Post /controller/dc/v3/tenants]. Error Message: The tenant id already exist. - Error Code: ", err)
	}
}

func TestCreateTenantInvalidID(t *testing.T) {
	_, name, tenantAttr := GetTenantAttributes()
	id := "dummy"
	client := helper.GetClient()
	_, err := client.CreateTenant(id, name, tenantAttr)
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "HTTP Error response status code 400 when accessing [Post /controller/dc/v3/tenants]. Error Message: Invalid UUID format. - Error Code: ", err)
	}
}

func TestUpdateTenant(t *testing.T) {
	id, name, tenantAttr := GetTenantAttributes()
	description :=  "Updated From GO"
	defer DeleteTenant(id)
	client := helper.GetClient()
	_, err := client.CreateTenant(id, name, tenantAttr)
	tenantAttr.Description = description
	tenant, err := client.UpdateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	assert.Equal(t, description, tenant.Description)
	getTenant := GetTenant(id)
	assert.Equal(t, getTenant.Description, tenant.Description)
}

func TestGetTenant(t *testing.T) {
	id, name, tenantAttr := GetTenantAttributes()
	defer DeleteTenant(id)
	client := helper.GetClient()
	_, err := client.CreateTenant(id, name, tenantAttr)
	tenant, err := client.GetTenant(id)
	assert.Nil(t, err)
	assert.Equal(t, id, tenant.Id, id)
	assert.Equal(t, name, tenant.Name, name)
	assert.Equal(t, tenantAttr.Description, tenant.Description )
	assert.Equal(t, tenantAttr.Producer, tenant.Producer)
	assert.Equal(t, tenantAttr.MulticastCapability, tenant.MulticastCapability)
	assert.Equal(t, tenantAttr.Quota, tenant.Quota)
	//assert.Equal(t, tenantAttr.ResPool, tenant.ResPool)
}

func TestListTenants(t *testing.T) {
	id, name, tenantAttr := GetTenantAttributes()
	client := helper.GetClient()
	defer DeleteTenant(id)
	_, err := client.CreateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	queryParameters := &models.TenantRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.GetTenants(queryParameters)
	assert.Equal(t, 3, len(*response))
	assert.Nil(t, err)
	queryParameters.Producer = tenantAttr.Producer
	queryParameters.PageSize = 3
	response, err = client.GetTenants(queryParameters)
	assert.Equal(t, 1, len(*response))
	assert.Equal(t, tenantAttr.Producer, (*response)[0].Producer)
	assert.Nil(t, err)
}

func TestDeleteTenant(t *testing.T) {
	id, _, _ := GetTenantAttributes()
	client := helper.GetClient()
	err := client.DeleteTenant(id)
	assert.Nil(t, err)
}

func GetTenant(id string) *models.Tenant {
	client := helper.GetClient()
	tenant, _ := client.GetTenant(id)
	return tenant
}

func DeleteTenant(id string) {
	client := helper.GetClient()
	_ = client.DeleteTenant(id)
}
