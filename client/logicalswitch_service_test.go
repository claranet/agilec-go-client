package client

import (
	"fmt"
	"github.com/claranet/agilec-go-client/models"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateLogicalSwitchCompleteSuccess(t *testing.T) {
	client := GetClientTest()

	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes(t, client)
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	defer DeleteLogicalSwitch(t, *id)
	assert.Nil(t, err)

	logicalSwitch, err := client.GetLogicalSwitch(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalSwitch.Id)
	assert.Equal(t, name, logicalSwitch.Name)
	assert.Equal(t, logicalSwitchAttr.Description, logicalSwitch.Description)
	assert.Equal(t, logicalSwitchAttr.Vni, logicalSwitch.Vni)
	assert.Equal(t, logicalSwitchAttr.Bd, logicalSwitch.Bd)
	assert.Equal(t, logicalSwitchAttr.MacAddress, logicalSwitch.MacAddress)
	assert.Equal(t, logicalSwitchAttr.Additional.Producer, logicalSwitch.Additional.Producer)
	assert.Equal(t, logicalSwitchAttr.LogicNetworkId, logicalSwitch.LogicNetworkId)
	assert.Equal(t, logicalSwitchAttr.StormSuppress.MulticastEnable, logicalSwitch.StormSuppress.MulticastEnable)
	assert.Equal(t, logicalSwitchAttr.StormSuppress.UnicastEnable, logicalSwitch.StormSuppress.UnicastEnable)
	assert.Equal(t, logicalSwitchAttr.StormSuppress.BroadcastEnable, logicalSwitch.StormSuppress.BroadcastEnable)
}

func TestCreateLogicalSwitchDuplicate(t *testing.T) {
	client := GetClientTest()
	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes(t, client)
	defer DeleteLogicalSwitch(t, *id)
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.Nil(t, err)
	err = client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.NotNil(t, err)
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}

		assert.Contains(t, response.ErrorMessage, "The object "+fmt.Sprint(*id)+" is being used.")
		assert.Equal(t, "/controller/dc/v3/logicnetwork/switchs", response.URL)
		assert.Equal(t, 409, response.HttpStatusCode)
		assert.Equal(t, "Post", response.Method)
	}
}

func TestCreateLogicalSwitchInvalidID(t *testing.T) {
	client := GetClientTest()
	_, name, logicalSwitchAttr := GetLogicalSwitchAttributes(t, client)
	id := String("dummy")
	defer DeleteLogicalSwitch(t, *id)
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.NotNil(t, err)
}

func TestUpdateLogicalSwitch(t *testing.T) {
	client := GetClientTest()
	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes(t, client)
	defer DeleteLogicalSwitch(t, *id)
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.Nil(t, err)
	logicalSwitchAttr.Description = String("Updated From Go")
	logicalSwitch, err := client.UpdateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.Nil(t, err)
	assert.Equal(t, logicalSwitchAttr.Description, logicalSwitch.Description)
	getlogicalSwitch := GetLogicalSwitch(t, *id)
	assert.Equal(t, getlogicalSwitch.Description, logicalSwitch.Description)
}

func TestUpdateNonExistingLogicalSwitch(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.UpdateLogicalSwitch(String(u.String()), String("dummy"), &models.LogicalSwitchAttributes{LogicNetworkId: String("Dummy")})
	if assert.NotNil(t, err) {
		if assert.NotNil(t, err) {
			response, ok := err.(*ErrorResponse)
			if !ok {
				t.Error("Wrong Error Response")
			}
			assert.Contains(t, response.ErrorMessage, "The parameter logicswitch is null.")
			assert.Contains(t, response.URL, "/controller/dc/v3/logicnetwork/switchs/switch/")
			assert.Equal(t, 400, response.HttpStatusCode)
			assert.Equal(t, "Put", response.Method)
		}
	}
}

func TestGetLogicalSwitch(t *testing.T) {
	client := GetClientTest()
	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes(t, client)
	defer DeleteLogicalSwitch(t, *id)
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	logicalSwitch, err := client.GetLogicalSwitch(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalSwitch.Id, id)
	assert.Equal(t, name, logicalSwitch.Name, name)
	assert.Equal(t, logicalSwitchAttr.Description, logicalSwitch.Description)
	assert.Equal(t, logicalSwitchAttr.Vni, logicalSwitch.Vni)
	assert.Equal(t, logicalSwitchAttr.Bd, logicalSwitch.Bd)
	assert.Equal(t, logicalSwitchAttr.MacAddress, logicalSwitch.MacAddress)
	assert.Equal(t, logicalSwitchAttr.LogicNetworkId, logicalSwitch.LogicNetworkId)
	assert.Equal(t, logicalSwitchAttr.Additional.Producer, logicalSwitch.Additional.Producer)
	assert.Equal(t, logicalSwitchAttr.StormSuppress.BroadcastEnable, logicalSwitch.StormSuppress.BroadcastEnable)
	assert.Equal(t, logicalSwitchAttr.StormSuppress.MulticastEnable, logicalSwitch.StormSuppress.MulticastEnable)
	assert.Equal(t, logicalSwitchAttr.StormSuppress.UnicastEnable, logicalSwitch.StormSuppress.UnicastEnable)
}

func TestGetNonExistLogicalSwitch(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.GetLogicalSwitch(u.String())
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}
		assert.Equal(t, "The Resource don't exists.", response.ErrorMessage)
		assert.Equal(t, "/controller/dc/v3/logicnetwork/switchs/switch/"+u.String(), response.URL)
		assert.Equal(t, 0, response.HttpStatusCode)
		assert.Equal(t, "Get", response.Method)
	}
}

func TestListLogicalSwitches(t *testing.T) {
	client := GetClientTest()
	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes(t, client)
	defer DeleteLogicalSwitch(t, *id)
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.Nil(t, err)
	queryParameters := &models.LogicalSwitchRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListLogicalSwitches(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}

func TestDeleteLogicalSwitch(t *testing.T) {
	client := GetClientTest()
	id, _, _ := GetLogicalSwitchAttributes(t, client)
	err := client.DeleteLogicalSwitch(*id)
	assert.Nil(t, err)
}

func GetLogicalSwitchAttributes(t *testing.T, client *Client) (*string, *string, *models.LogicalSwitchAttributes) {
	t.Helper()
	u, _ := uuid.NewV4()
	fmt.Printf("Logical Switch ID Generated: %s\n", u.String())
	Id := String(u.String())
	Name := String("CLARANET-GO-TESTS-001")

	LogicalSwitch := models.LogicalSwitchAttributes{}
	LogicalSwitch.Description = String("Created By GO")
	LogicalSwitch.Vni = Int32(6000)
	LogicalSwitch.Bd = Int32(10000)
	LogicalSwitch.MacAddress = String("00:00:00:00:00:01")
	LogicalSwitch.Additional = &models.LogicalSwitchAdditional{
		Producer:    String("Test Producer"),
	}
	LogicalSwitch.StormSuppress = &models.LogicalSwitchStormSuppress{
		BroadcastEnable: Bool(false),  
		MulticastEnable: Bool(false),
		UnicastEnable: Bool(false),
	}

	// Get Logical network
	responseLogicalNetwork, err := client.ListLogicalNetworks(&models.LogicalNetworkRequestOpts{
		BaseRequestOpts: models.BaseRequestOpts{
			PageSize:  1,
			PageIndex: 1,
		},
	})
	if err != nil {
		t.Logf("GetLogicalSwitchAttributes %s", err)
	}
	assert.Equal(t, 1, len(responseLogicalNetwork))

	LogicalSwitch.LogicNetworkId = responseLogicalNetwork[0].Id

	return Id, Name, &LogicalSwitch
}

func GetLogicalSwitch(t *testing.T, id string) *models.LogicalSwitch {
	t.Helper()
	client := GetClientTest()
	logicalSwitch, _ := client.GetLogicalSwitch(id)
	return logicalSwitch
}

func DeleteLogicalSwitch(t *testing.T, id string) {
	t.Helper()
	client := GetClientTest()
	_ = client.DeleteLogicalSwitch(id)
}