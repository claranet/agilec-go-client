package client

import (
	"fmt"
	"github.com/claranet/agilec-go-client/models"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEndPortCompleteSuccess(t *testing.T) {
	client := GetClientTest()

	id, name, endPortAttr := GetEndPortAttributes(t, client)
	err := client.CreateEndPort(id, name, endPortAttr)
	defer DeleteEndPort(t, *id)
	assert.Nil(t, err)

	endPort, err := client.GetEndPort(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, endPort.Id)
	assert.Equal(t, name, endPort.Name)
	assert.Equal(t, endPortAttr.Description, endPort.Description)
	assert.Equal(t, endPortAttr.VmName, endPort.VmName)
	assert.Equal(t, endPortAttr.Location, endPort.Location)
	assert.Equal(t, endPortAttr.Ipv4, endPort.Ipv4)
	assert.Equal(t, endPortAttr.Ipv6, endPort.Ipv6)
	assert.Equal(t, endPortAttr.LogicNetworkId, endPort.LogicNetworkId)
}

func TestCreateEndPortDuplicate(t *testing.T) {
	client := GetClientTest()
	id, name, endPortAttr := GetEndPortAttributes(t, client)
	defer DeleteEndPort(t, *id)
	err := client.CreateEndPort(id, name, endPortAttr)
	assert.Nil(t, err)
	err = client.CreateEndPort(id, name, endPortAttr)
	assert.NotNil(t, err)
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}

		assert.Contains(t, response.ErrorMessage, "The End Port ID already exists.")
		assert.Equal(t, "/controller/dc/v3/logicnetwork/endports", response.URL)
		assert.Equal(t, 409, response.HttpStatusCode)
		assert.Equal(t, "Post", response.Method)
	}
}

func TestCreateEndPortInvalidID(t *testing.T) {
	client := GetClientTest()
	_, name, endPortAttr := GetEndPortAttributes(t, client)
	id := String("dummy")
	defer DeleteEndPort(t, *id)
	err := client.CreateEndPort(id, name, endPortAttr)
	assert.NotNil(t, err)
}

func TestUpdateEndPort(t *testing.T) {
	client := GetClientTest()
	id, name, endPortAttr := GetEndPortAttributes(t, client)
	defer DeleteEndPort(t, *id)
	err := client.CreateEndPort(id, name, endPortAttr)
	assert.Nil(t, err)
	endPortAttr.Description = String("Updated From Go")
	endPort, err := client.UpdateEndPort(id, name, endPortAttr)
	assert.Nil(t, err)
	assert.Equal(t, endPortAttr.Description, endPort.Description)
	getEndPort := GetEndPort(t, *id)
	assert.Equal(t, getEndPort.Description, endPort.Description)
}

func TestUpdateNonExistingEndPort(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.UpdateEndPort(String(u.String()), String("dummy"), &models.EndPortAttributes{})
	if assert.NotNil(t, err) {
		if assert.NotNil(t, err) {
			response, ok := err.(*ErrorResponse)
			if !ok {
				t.Error("Wrong Error Response")
			}
			assert.Contains(t, response.ErrorMessage, "The end port "+fmt.Sprint(u)+" does not exist.")
			assert.Contains(t, response.URL, "/controller/dc/v3/logicnetwork/endports/endport/")
			assert.Equal(t, 400, response.HttpStatusCode)
			assert.Equal(t, "Put", response.Method)
		}
	}
}

func TestGetEndPort(t *testing.T) {
	client := GetClientTest()
	id, name, endPortAttr := GetEndPortAttributes(t, client)
	defer DeleteEndPort(t, *id)
	err := client.CreateEndPort(id, name, endPortAttr)
	endPort, err := client.GetEndPort(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, endPort.Id, id)
	assert.Equal(t, name, endPort.Name, name)
	assert.Equal(t, endPortAttr.Description, endPort.Description)
	assert.Equal(t, endPortAttr.Location, endPort.Location)
	assert.Equal(t, endPortAttr.VmName, endPort.VmName)
	assert.Equal(t, endPortAttr.Ipv4, endPort.Ipv4)
	assert.Equal(t, endPortAttr.Ipv6, endPort.Ipv6)
	assert.Equal(t, endPortAttr.LogicNetworkId, endPort.LogicNetworkId)
	assert.Equal(t, endPortAttr.LogicPortId, endPort.LogicPortId)
}

// TODO check why it is getting response nil
// func TestGetNonExistEndPort(t *testing.T) {
// 	client := GetClientTest()
// 	u, _ := uuid.NewV4()
// 	_, err := client.GetEndPort(u.String())
// 	if assert.NotNil(t, err) {
// 		response, ok := err.(*ErrorResponse)

// 		if !ok {
// 			t.Error("Wrong Error Response")
// 		}
// 		assert.Equal(t, "The Resource don't exists.", response.ErrorMessage)
// 		assert.Equal(t, "/controller/dc/v3/logicnetwork/endports/endport/"+u.String(), response.URL)
// 		assert.Equal(t, 0, response.HttpStatusCode)
// 		assert.Equal(t, "Get", response.Method)
// 	}
// }

func TestListEndPorts(t *testing.T) {
	client := GetClientTest()
	id, name, endPortAttr := GetEndPortAttributes(t, client)
	defer DeleteEndPort(t, *id)
	err := client.CreateEndPort(id, name, endPortAttr)
	assert.Nil(t, err)
	queryParameters := &models.EndPortRequestOpts{}
	queryParameters.PageSize = 1
	response, err := client.ListEndPorts(queryParameters)
	assert.Equal(t, 1, len(response))
	assert.Nil(t, err)
}

func TestDeleteEndPort(t *testing.T) {
	client := GetClientTest()
	id, _, _ := GetEndPortAttributes(t, client)
	err := client.DeleteEndPort(*id)
	assert.Nil(t, err)
}

func GetEndPortAttributes(t *testing.T, client *Client) (*string, *string, *models.EndPortAttributes) {
	t.Helper()
	u, _ := uuid.NewV4()
	fmt.Printf("End port ID Generated: %s\n", u.String())
	Id := String(u.String())
	Name := String("CLARANET-GO-TESTS-001")

	EndPort := models.EndPortAttributes{}
	EndPort.Description = String("Created By GO")
	EndPort.Location = String("10")
	EndPort.VmName = String("10")
	ipv4 := "192.168.1.1"
	EndPort.Ipv4 = []*string{&ipv4}
	ipv6 := "FE80::A1"
	EndPort.Ipv6 = []*string{&ipv6}

	// Get Logical network
	responseLogicalNetwork, err := client.ListLogicalNetworks(&models.LogicalNetworkRequestOpts{
		BaseRequestOpts: models.BaseRequestOpts{
			PageSize:  1,
			PageIndex: 1,
		},
	})
	if err != nil {
		t.Logf("GetEndPortAttributes %s", err)
	}
	assert.Equal(t, 1, len(responseLogicalNetwork))

	EndPort.LogicNetworkId = responseLogicalNetwork[0].Id

	return Id, Name, &EndPort
}

func GetEndPort(t *testing.T, id string) *models.EndPort {
	t.Helper()
	client := GetClientTest()
	endPort, _ := client.GetEndPort(id)
	return endPort
}

func DeleteEndPort(t *testing.T, id string) {
	t.Helper()
	client := GetClientTest()
	_ = client.DeleteEndPort(id)
}