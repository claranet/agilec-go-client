package tests

import (
	"fmt"
	agilec "github.com/claranet/agilec-go-client/client"
	"github.com/claranet/agilec-go-client/models"
	helper "github.com/claranet/agilec-go-client/tests/helpers"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetTenantAttributes() (*string, *string, *models.TenantAttributes) {
	u, _ := uuid.NewV4()
	fmt.Printf("Tenant ID Generated: %s\n", u.String())
	Id := agilec.String(u.String())
	Name := agilec.String("CLARANET-GO-TESTS-001")

	Tenant := models.TenantAttributes{}
	Tenant.Description = agilec.String("Created By GO")
	Tenant.Producer = agilec.String("GOLANG")
	Tenant.MulticastCapability = agilec.Bool(true)
	Tenant.Quota = &models.TenantQuota{
		LogicVasNum:    agilec.Int32(10),
		LogicRouterNum: agilec.Int32(5),
		LogicSwitchNum: agilec.Int32(6),
	}
	Tenant.MulticastQuota = &models.TenantMulticastQuota{
		AclNum:     agilec.Int32(10),
		AclRuleNum: agilec.Int32(10),
	}
	Tenant.ResPool = &models.TenantResPool{
		//ExternalGatewayIds: []*string{"15e608a6-8d60-4c5a-89c0-a99e3cd967ff"},
		FabricIds: []*string{agilec.String("804c7c74-5586-48bf-9cea-96a6d4d3f3a5")},
		//VmmIds:             []*string{"1", "2"},
		//DhcpGroupIds:       []*string{"1", "2"},
	}

	return Id, Name, &Tenant
}

func TestCreateTenantComplete(t *testing.T) {
	id, name, tenantAttr := GetTenantAttributes()
	defer DeleteTenant(*id)
	client := helper.GetClient()
	err := client.CreateTenant(id, name, tenantAttr)
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
	id, name, tenantAttr := GetTenantAttributes()
	defer DeleteTenant(*id)
	client := helper.GetClient()
	err := client.CreateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	err = client.CreateTenant(id, name, tenantAttr)
	assert.NotNil(t, err)
	if assert.NotNil(t, err) {
		response, ok := err.(*agilec.ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}

		assert.Contains(t, response.ErrorMessage, "The tenant id already exist.")
		assert.Equal(t, "/controller/dc/v3/tenants", response.URL)
		assert.Equal(t, 400, response.HttpStatusCode)
		assert.Equal(t, "Post", response.Method)
	}
}

func TestCreateTenantInvalidID(t *testing.T) {
	_, name, tenantAttr := GetTenantAttributes()
	id := agilec.String("dummy")
	client := helper.GetClient()
	err := client.CreateTenant(id, name, tenantAttr)
	if assert.NotNil(t, err) {
		if assert.NotNil(t, err) {
			response, ok := err.(*agilec.ErrorResponse)
			if !ok {
				t.Error("Wrong Error Response")
			}
			assert.Contains(t, response.ErrorMessage, "Invalid UUID format.")
			assert.Equal(t, "/controller/dc/v3/tenants", response.URL)
			assert.Equal(t, 400, response.HttpStatusCode)
			assert.Equal(t, "Post", response.Method)
		}
	}
}

func TestUpdateTenant(t *testing.T) {
	id, name, tenantAttr := GetTenantAttributes()
	defer DeleteTenant(*id)
	client := helper.GetClient()
	err := client.CreateTenant(id, name, tenantAttr)
	tenantAttr.Description = agilec.String("Updated From Go")
	tenantAttr.MulticastCapability = agilec.Bool(false)
	tenantAttr.MulticastQuota = nil
	tenant, err := client.UpdateTenant(id, name, tenantAttr)
	assert.Nil(t, err)
	getTenant := GetTenant(*id)
	assert.Equal(t, tenantAttr.Description, getTenant.Description)
	assert.Equal(t, tenant.Producer, getTenant.Producer)
	assert.Equal(t, tenantAttr.Description, getTenant.Description)
	assert.Equal(t, tenantAttr.MulticastCapability, getTenant.MulticastCapability)
	assert.Equal(t, tenant.ResPool.FabricIds[0], getTenant.ResPool.FabricIds[0])
	assert.Equal(t, tenant.Quota, getTenant.Quota)

}

func TestUpdateNonExistingTenant(t *testing.T) {
	client := helper.GetClient()
	u, _ := uuid.NewV4()
	_, err := client.UpdateTenant(agilec.String(u.String()), agilec.String("dummy"), &models.TenantAttributes{})
	if assert.NotNil(t, err) {
		if assert.NotNil(t, err) {
			response, ok := err.(*agilec.ErrorResponse)
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
	id, name, tenantAttr := GetTenantAttributes()
	defer DeleteTenant(*id)
	client := helper.GetClient()
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
	client := helper.GetClient()
	u, _ := uuid.NewV4()
	_, err := client.GetTenant(u.String())
	if assert.NotNil(t, err) {
		response, ok := err.(*agilec.ErrorResponse)

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
	id, name, tenantAttr := GetTenantAttributes()
	client := helper.GetClient()
	defer DeleteTenant(*id)
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
	id, _, _ := GetTenantAttributes()
	client := helper.GetClient()
	err := client.DeleteTenant(*id)
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
