package client

import (
	"fmt"
	"github.com/claranet/agilec-go-client/models"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateLogicalNetworkCompleteSuccess(t *testing.T) {
	client := GetClientTest()

	id, name, logicalNetworkAttr := GetLogicalNetworkAttributes(t, client)
	err := client.CreateLogicalNetwork(id, name, logicalNetworkAttr)
	defer DeleteLogicalNetwork(t, *id)
	assert.Nil(t, err)

	logicalNetwork, err := client.GetLogicalNetwork(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalNetwork.Id)
	assert.Equal(t, name, logicalNetwork.Name)
	assert.Equal(t, logicalNetworkAttr.Description, logicalNetwork.Description)
	assert.Equal(t, logicalNetworkAttr.MulticastCapability, logicalNetwork.MulticastCapability)
	assert.Equal(t, logicalNetworkAttr.IsVpcDeployed, logicalNetwork.IsVpcDeployed)
	assert.Equal(t, logicalNetworkAttr.Type, logicalNetwork.Type)
	assert.Equal(t, logicalNetworkAttr.Additional.Producer, logicalNetwork.Additional.Producer)
	assert.Equal(t, logicalNetworkAttr.FabricId, logicalNetwork.FabricId)
	assert.Equal(t, logicalNetworkAttr.TenantId, logicalNetwork.TenantId)
}

func TestCreateLogicalNetworkDuplicate(t *testing.T) {
	client := GetClientTest()
	id, name, logicalNetworkAttr := GetLogicalNetworkAttributes(t, client)
	defer DeleteLogicalNetwork(t, *id)
	err := client.CreateLogicalNetwork(id, name, logicalNetworkAttr)
	assert.Nil(t, err)
	err = client.CreateLogicalNetwork(id, name, logicalNetworkAttr)
	assert.NotNil(t, err)
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}

		assert.Contains(t, response.ErrorMessage, "The VPC "+fmt.Sprint(*id)+" already exists.")
		assert.Equal(t, "/controller/dc/v3/logicnetwork/networks", response.URL)
		assert.Equal(t, 400, response.HttpStatusCode)
		assert.Equal(t, "Post", response.Method)
	}
}

func TestCreateLogicalNetworkInvalidID(t *testing.T) {
	client := GetClientTest()
	_, name, logicalNetworkAttr := GetLogicalNetworkAttributes(t, client)
	id := String("dummy")
	defer DeleteLogicalNetwork(t, *id)
	err := client.CreateLogicalNetwork(id, name, logicalNetworkAttr)
	assert.NotNil(t, err)
}

func TestUpdateLogicalNetwork(t *testing.T) {
	client := GetClientTest()
	id, name, logicalNetworkAttr := GetLogicalNetworkAttributes(t, client)
	defer DeleteLogicalNetwork(t, *id)
	err := client.CreateLogicalNetwork(id, name, logicalNetworkAttr)
	assert.Nil(t, err)
	logicalNetworkAttr.Description = String("Updated From Go")
	logicalNetwork, err := client.UpdateLogicalNetwork(id, name, logicalNetworkAttr)
	assert.Nil(t, err)
	assert.Equal(t, logicalNetworkAttr.Description, logicalNetwork.Description)
	getlogicalNetwork := GetLogicalNetwork(t, *id)
	assert.Equal(t, getlogicalNetwork.Description, logicalNetwork.Description)
}

func TestUpdateNonExistingLogicalNetwork(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.UpdateLogicalNetwork(String(u.String()), String("dummy"), &models.LogicalNetworkAttributes{})
	if assert.NotNil(t, err) {
		if assert.NotNil(t, err) {
			response, ok := err.(*ErrorResponse)
			if !ok {
				t.Error("Wrong Error Response")
			}
			assert.Contains(t, response.ErrorMessage, "The parameter LogicNetworkEntity is null.")
			assert.Contains(t, response.URL, "/controller/dc/v3/logicnetwork/networks/network/")
			assert.Equal(t, 400, response.HttpStatusCode)
			assert.Equal(t, "Put", response.Method)
		}
	}
}

func TestGetLogicalNetwork(t *testing.T) {
	client := GetClientTest()
	id, name, logicalNetworkAttr := GetLogicalNetworkAttributes(t, client)
	defer DeleteLogicalNetwork(t, *id)
	err := client.CreateLogicalNetwork(id, name, logicalNetworkAttr)
	logicalNetwork, err := client.GetLogicalNetwork(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalNetwork.Id, id)
	assert.Equal(t, name, logicalNetwork.Name, name)
	assert.Equal(t, logicalNetworkAttr.Description, logicalNetwork.Description)
	assert.Equal(t, logicalNetworkAttr.TenantId, logicalNetwork.TenantId)
	assert.Equal(t, logicalNetworkAttr.MulticastCapability, logicalNetwork.MulticastCapability)
	assert.Equal(t, logicalNetworkAttr.IsVpcDeployed, logicalNetwork.IsVpcDeployed)
	assert.Equal(t, logicalNetworkAttr.FabricId, logicalNetwork.FabricId)
	assert.Equal(t, logicalNetworkAttr.Type, logicalNetwork.Type)
	assert.Equal(t, logicalNetworkAttr.Additional.Producer, logicalNetwork.Additional.Producer)
}

func TestGetNonExistLogicalNetwork(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.GetLogicalNetwork(u.String())
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}
		assert.Equal(t, "The Resource don't exists.", response.ErrorMessage)
		assert.Equal(t, "/controller/dc/v3/logicnetwork/networks/network/"+u.String(), response.URL)
		assert.Equal(t, 0, response.HttpStatusCode)
		assert.Equal(t, "Get", response.Method)
	}
}

func TestListLogicalNetworks(t *testing.T) {
	client := GetClientTest()
	id, name, logicalNetworkAttr := GetLogicalNetworkAttributes(t, client)
	defer DeleteLogicalNetwork(t, *id)
	err := client.CreateLogicalNetwork(id, name, logicalNetworkAttr)
	assert.Nil(t, err)
	queryParameters := &models.LogicalNetworkRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListLogicalNetworks(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}

func TestDeleteLogicalNetwork(t *testing.T) {
	client := GetClientTest()
	id, _, _ := GetLogicalNetworkAttributes(t, client)
	err := client.DeleteLogicalNetwork(*id)
	assert.Nil(t, err)
}

func GetLogicalNetworkAttributes(t *testing.T, client *Client) (*string, *string, *models.LogicalNetworkAttributes) {
	t.Helper()
	u, _ := uuid.NewV4()
	fmt.Printf("Logical network ID Generated: %s\n", u.String())
	Id := String(u.String())
	Name := String("CLARANET-GO-TESTS-001")

	LogicalNetwork := models.LogicalNetworkAttributes{}
	LogicalNetwork.Description = String("Created By GO")
	LogicalNetwork.MulticastCapability = Bool(false)
	LogicalNetwork.IsVpcDeployed = Bool(true)
	LogicalNetwork.Type = String("Transit")
	LogicalNetwork.Additional = &models.LogicalNetworkAdditional{
		Producer:    String("Test Producer"),
	}
	
	// Get Fabric
	responseFabric, err := client.ListFabrics(&models.FabricRequestOpts{
		BaseRequestOpts: models.BaseRequestOpts{
			PageSize:  1,
			PageIndex: 1,
		},
	})
	if err != nil {
		t.Logf("GetLogicalNetworkAttributes %s", err)
	}
	assert.Equal(t, 1, len(responseFabric))

	LogicalNetwork.FabricId = []*string{responseFabric[0].Id}

	// Get Tenant
	responseTenants, err := client.ListTenants(&models.TenantRequestOpts{
		BaseRequestOpts: models.BaseRequestOpts{
			PageSize:  1,
			PageIndex: 1,
		},
	})
	if err != nil {
		t.Logf("GetLogicalNetworkAttributes %s", err)
	}
	assert.Equal(t, 1, len(responseTenants))

	LogicalNetwork.TenantId = responseTenants[0].Id

	return Id, Name, &LogicalNetwork
}

func GetLogicalNetwork(t *testing.T, id string) *models.LogicalNetwork {
	t.Helper()
	client := GetClientTest()
	logicalNetwork, _ := client.GetLogicalNetwork(id)
	return logicalNetwork
}

func DeleteLogicalNetwork(t *testing.T, id string) {
	t.Helper()
	client := GetClientTest()
	_ = client.DeleteLogicalNetwork(id)
}