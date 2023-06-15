package client

import (
	"fmt"
	"github.com/claranet/agilec-go-client/models"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateLogicalRouterCompleteSuccess(t *testing.T) {
	client := GetClientTest()

	id, name, tenantAttr := GetTenantAttributes(t, client)
	err := client.CreateTenant(id, name, tenantAttr)
	defer DeleteTenant(t, *id)
	assert.Nil(t, err)

	id, name, logicalNetworkAttr := GetLogicalNetworkAttributes(t, client)
	err = client.CreateLogicalNetwork(id, name, logicalNetworkAttr)
	defer DeleteLogicalNetwork(t, *id)
	assert.Nil(t, err)

	id, name, logicalRouterAttr := GetLogicalRouterAttributes(t, client)
	err = client.CreateLogicalRouter(id, name, logicalRouterAttr)
	defer DeleteLogicalRouter(t, *id)
	assert.Nil(t, err)
	logicalRouter, err := client.GetLogicalRouter(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalRouter.Id)
	assert.Equal(t, name, logicalRouter.Name)
	assert.Equal(t, logicalRouterAttr.LogicNetworkId, logicalRouter.LogicNetworkId)
	assert.Equal(t, logicalRouterAttr.Description, logicalRouter.Description)
	assert.Equal(t, logicalRouterAttr.Type, logicalRouter.Type)
	assert.Equal(t, logicalRouterAttr.Vni, logicalRouter.Vni)
	assert.Equal(t, logicalRouterAttr.VrfName, logicalRouter.VrfName)
	assert.Equal(t, logicalRouterAttr.RouterLocations[0].FabricId, logicalRouter.RouterLocations[0].FabricId)
	assert.Equal(t, logicalRouterAttr.RouterLocations[0].FabricRole, logicalRouter.RouterLocations[0].FabricRole)
}

func TestCreateLogicalRouterDuplicate(t *testing.T) {
	client := GetClientTest()
	id, name, logicalRouterAttr := GetLogicalRouterAttributes(t, client)
	defer DeleteLogicalRouter(t, *id)
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.Nil(t, err)
	err = client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.NotNil(t, err)
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}

		assert.Contains(t, response.ErrorMessage, "The object "+fmt.Sprint(*id)+" is being used.")
		assert.Equal(t, "/controller/dc/v3/logicnetwork/routers", response.URL)
		assert.Equal(t, 409, response.HttpStatusCode)
		assert.Equal(t, "Post", response.Method)
	}
}

func TestCreateLogicalRouterInvalidID(t *testing.T) {
	client := GetClientTest()
	_, name, logicalRouterAttr := GetLogicalRouterAttributes(t, client)
	id := String("dummy")
	defer DeleteLogicalRouter(t, *id)
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.NotNil(t, err)
}

func TestUpdateLogicalRouter(t *testing.T) {
	client := GetClientTest()
	id, name, logicalRouterAttr := GetLogicalRouterAttributes(t, client)
	defer DeleteLogicalRouter(t, *id)
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.Nil(t, err)
	logicalRouterAttr.Description = String("Updated From Go")
	logicalRouter, err := client.UpdateLogicalRouter(id, name, logicalRouterAttr)
	assert.Nil(t, err)
	getLogicalRouter := GetLogicalRouter(t, *id)
	assert.Equal(t, logicalRouterAttr.Description, logicalRouter.Description)
	assert.Equal(t, getLogicalRouter.Description, logicalRouter.Description)

}

func TestUpdateNonExistingLogicalRouter(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.UpdateLogicalRouter(String(u.String()), String("dummy"), &models.LogicalRouterAttributes{})
	if assert.NotNil(t, err) {
		if assert.NotNil(t, err) {
			response, ok := err.(*ErrorResponse)
			if !ok {
				t.Error("Wrong Error Response")
			}
			assert.Contains(t, response.ErrorMessage, "The logicRouter does not exist.")
			assert.Contains(t, response.URL, "/controller/dc/v3/logicnetwork/routers")
			assert.Equal(t, 409, response.HttpStatusCode)
			assert.Equal(t, "Put", response.Method)
		}
	}
}

func TestGetLogicalRouter(t *testing.T) {
	client := GetClientTest()
	id, name, logicalRouterAttr := GetLogicalRouterAttributes(t, client)
	defer DeleteLogicalRouter(t, *id)
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	logicalRouter, err := client.GetLogicalRouter(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalRouter.Id)
	assert.Equal(t, name, logicalRouter.Name)
	assert.Equal(t, logicalRouterAttr.LogicNetworkId, logicalRouter.LogicNetworkId)
	assert.Equal(t, logicalRouterAttr.Description, logicalRouter.Description)
	assert.Equal(t, logicalRouterAttr.Type, logicalRouter.Type)
	assert.Equal(t, logicalRouterAttr.Vni, logicalRouter.Vni)
	assert.Equal(t, logicalRouterAttr.VrfName, logicalRouter.VrfName)
	assert.Equal(t, logicalRouterAttr.RouterLocations[0].FabricId, logicalRouter.RouterLocations[0].FabricId)
	assert.Equal(t, logicalRouterAttr.RouterLocations[0].FabricRole, logicalRouter.RouterLocations[0].FabricRole)
}

func TestGetNonExistLogicalRouter(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.GetLogicalRouter(u.String())
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}
		assert.Equal(t, "The Resource don't exists.", response.ErrorMessage)
		assert.Equal(t, "/controller/dc/v3/logicnetwork/routers/router/"+u.String(), response.URL)
		assert.Equal(t, 0, response.HttpStatusCode)
		assert.Equal(t, "Get", response.Method)
	}
}

func TestListLogicalRouters(t *testing.T) {
	client := GetClientTest()
	id, name, logicalRouterAttr := GetLogicalRouterAttributes(t, client)
	defer DeleteLogicalRouter(t, *id)
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.Nil(t, err)
	queryParameters := &models.LogicalRouterRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListLogicalRouters(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}

func TestDeleteLogicalRouter(t *testing.T) {
	client := GetClientTest()
	id, _, _ := GetLogicalRouterAttributes(t, client)
	err := client.DeleteLogicalRouter(*id)
	assert.Nil(t, err)
}

func GetLogicalRouterAttributes(t *testing.T, client *Client) (*string, *string, *models.LogicalRouterAttributes) {
	t.Helper()
	u, _ := uuid.NewV4()
	fmt.Printf("Logical router ID Generated: %s\n", u.String())
	Id := String(u.String())
	Name := String("CLARANET-GO-TESTS-001")

	LogicalRouter := models.LogicalRouterAttributes{}
	LogicalRouter.Description = String("Created By GO")

	// Get LogicalNetwork
	responseLogicalNetwork, err := client.ListLogicalNetworks(&models.LogicalNetworkRequestOpts{
		BaseRequestOpts: models.BaseRequestOpts{
			PageSize:  1,
			PageIndex: 1,
		},
	})
	if err != nil {
		t.Logf("GetLogicalNetworkAttributes %s", err)
	}
	assert.Equal(t, 1, len(responseLogicalNetwork))

	LogicalRouter.LogicNetworkId = responseLogicalNetwork[0].Id
	LogicalRouter.Type = String("Normal")
	LogicalRouter.Vni = Int32(6054)
	LogicalRouter.VrfName = String("Test_VRF")
	LogicalRouter.Additional = &models.LogicalRouterAdditional{
		Producer:    String("Test Producer"),
	}
	
	// Get Fabric
	response, err := client.ListFabrics(&models.FabricRequestOpts{
		BaseRequestOpts: models.BaseRequestOpts{
			PageSize:  1,
			PageIndex: 1,
		},
	})
	if err != nil {
		t.Logf("GetLogicalRouterAttributes %s", err)
	}
	assert.Equal(t, 1, len(response))

	LogicalRouter.RouterLocations = []*models.LogicalRouterLocations{
		{
			FabricId:    String(*response[0].Id),
		},
	}

	return Id, Name, &LogicalRouter
}

func GetLogicalRouter(t *testing.T, id string) *models.LogicalRouter {
	t.Helper()
	client := GetClientTest()
	logicalRouter, _ := client.GetLogicalRouter(id)
	return logicalRouter
}

func DeleteLogicalRouter(t *testing.T, id string) {
	t.Helper()
	client := GetClientTest()
	_ = client.DeleteLogicalRouter(id)
}