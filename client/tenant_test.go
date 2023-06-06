package client

import (
	"fmt"
	"github.com/claranet/agilec-go-client/models"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTenantCompleteSuccess(t *testing.T) {
	client := GetClientTest()
	id, name, tenantAttr := GetTenantAttributes(t, client)
	err := client.CreateTenant(id, name, tenantAttr)
	defer DeleteTenant(t, *id)
	assert.Nil(t, err)
	tenant, err := client.GetTenant(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, tenant.Id)
	assert.Equal(t, name, tenant.Name)
	assert.Equal(t, tenantAttr.Producer, tenant.Producer)
	assert.Equal(t, tenantAttr.Description, tenant.Description)
	assert.Equal(t, tenantAttr.MulticastCapability, tenant.MulticastCapability)
	assert.Equal(t, tenantAttr.ResPool.FabricIds[0], tenant.ResPool.FabricIds[0])
	assert.Equal(t, tenantAttr.Quota, tenant.Quota)
	assert.Equal(t, tenantAttr.MulticastQuota, tenant.MulticastQuota)
}

func TestCreateTenantDuplicate(t *testing.T) {
	client := GetClientTest()
	id, name, tenantAttr := GetTenantAttributes(t, client)
	defer DeleteTenant(t, *id)
	err := client.CreateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	err = client.CreateTenant(id, name, tenantAttr)
	assert.NotNil(t, err)
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}

		assert.Contains(t, response.ErrorMessage, "The tenant ID "+fmt.Sprint(*id)+" already exists.")
		assert.Equal(t, "/controller/dc/v3/tenants", response.URL)
		assert.Equal(t, 400, response.HttpStatusCode)
		assert.Equal(t, "Post", response.Method)
	}
}

func TestUpdateTenant(t *testing.T) {
	client := GetClientTest()
	id, name, tenantAttr := GetTenantAttributes(t, client)
	defer DeleteTenant(t, *id)
	err := client.CreateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	tenantAttr.Description = String("Updated From Go")
	tenantAttr.MulticastCapability = Bool(false)
	tenantAttr.MulticastQuota = nil
	tenant, err := client.UpdateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	getTenant := GetTenant(t, *id)
	assert.Equal(t, tenantAttr.Description, getTenant.Description)
	assert.Equal(t, tenant.Producer, getTenant.Producer)
	assert.Equal(t, tenantAttr.Description, getTenant.Description)
	assert.Equal(t, tenantAttr.MulticastCapability, getTenant.MulticastCapability)
	assert.Equal(t, tenant.ResPool.FabricIds[0], getTenant.ResPool.FabricIds[0])
	assert.Equal(t, tenant.Quota, getTenant.Quota)
}

func TestUpdateNonExistingTenant(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.UpdateTenant(String(u.String()), String("dummy"), &models.TenantAttributes{})
	if assert.NotNil(t, err) {
		if assert.NotNil(t, err) {
			response, ok := err.(*ErrorResponse)
			if !ok {
				t.Error("Wrong Error Response")
			}
			assert.Contains(t, response.ErrorMessage, "tenant not exist.")
			assert.Contains(t, response.URL, "/controller/dc/v3/tenants")
			assert.Equal(t, 400, response.HttpStatusCode)
			assert.Equal(t, "Put", response.Method)
		}
	}
}

func TestGetTenant(t *testing.T) {
	client := GetClientTest()
	id, name, tenantAttr := GetTenantAttributes(t, client)
	defer DeleteTenant(t, *id)
	err := client.CreateTenant(id, name, tenantAttr)
	tenant, err := client.GetTenant(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, tenant.Id, id)
	assert.Equal(t, name, tenant.Name, name)
	assert.Equal(t, tenantAttr.Description, tenant.Description)
	assert.Equal(t, tenantAttr.Producer, tenant.Producer)
	assert.Equal(t, tenantAttr.MulticastCapability, tenant.MulticastCapability)
	assert.Equal(t, tenantAttr.Quota, tenant.Quota)
	assert.Equal(t, tenantAttr.ResPool.FabricIds[0], tenant.ResPool.FabricIds[0])
}

func TestGetNonExistTenant(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.GetTenant(u.String())
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}
		assert.Equal(t, "The Resource don't exists.", response.ErrorMessage)
		assert.Equal(t, "/controller/dc/v3/tenants/tenant/"+u.String(), response.URL)
		assert.Equal(t, 0, response.HttpStatusCode)
		assert.Equal(t, "Get", response.Method)
	}
}

func TestListTenants(t *testing.T) {
	client := GetClientTest()
	id, name, tenantAttr := GetTenantAttributes(t, client)
	defer DeleteTenant(t, *id)
	err := client.CreateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	queryParameters := &models.TenantRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListTenants(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
	queryParameters.Producer = *tenantAttr.Producer
	queryParameters.PageSize = 3
	response, err = client.ListTenants(queryParameters)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, tenantAttr.Producer, (response)[0].Producer)
	assert.Nil(t, err)
}

func TestDeleteTenant(t *testing.T) {
	client := GetClientTest()
	id, _, _ := GetTenantAttributes(t, client)
	err := client.DeleteTenant(*id)
	assert.Nil(t, err)
}

func GetTenantAttributes(t *testing.T, client *Client) (*string, *string, *models.TenantAttributes) {
	t.Helper()
	u, _ := uuid.NewV4()
	fmt.Printf("Tenant ID Generated: %s\n", u.String())
	Id := String(u.String())
	Name := String("CLARANET-GO-TESTS-001")

	Tenant := models.TenantAttributes{}
	Tenant.Description = String("Created By GO")
	Tenant.Producer = String("GOLANG")
	Tenant.MulticastCapability = Bool(true)
	Tenant.Quota = &models.TenantQuota{
		LogicVasNum:    Int32(10),
		LogicRouterNum: Int32(5),
		LogicSwitchNum: Int32(6),
	}
	Tenant.MulticastQuota = &models.TenantMulticastQuota{
		AclNum:     Int32(10),
		AclRuleNum: Int32(10),
	}

	// Get Fabric
	response, err := client.ListFabrics(&models.FabricRequestOpts{
		BaseRequestOpts: models.BaseRequestOpts{
			PageSize:  1,
			PageIndex: 1,
		},
	})
	if err != nil {
		t.Logf("GetTenantAttributes %s", err)
	}
	assert.Equal(t, 1, len(response))
	Tenant.ResPool = &models.TenantResPool{
		//ExternalGatewayIds: []*string{"15e608a6-8d60-4c5a-89c0-a99e3cd967ff"},
		FabricIds: []*string{response[0].Id},
		//VmmIds:             []*string{"1", "2"},
		//DhcpGroupIds:       []*string{"1", "2"},
	}

	return Id, Name, &Tenant
}

func GetTenant(t *testing.T, id string) *models.Tenant {
	t.Helper()
	client := GetClientTest()
	tenant, _ := client.GetTenant(id)
	return tenant
}

func DeleteTenant(t *testing.T, id string) {
	t.Helper()
	client := GetClientTest()
	_ = client.DeleteTenant(id)
}
