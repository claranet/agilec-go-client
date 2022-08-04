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

func GetEndPortAttributes() (*string, *string, *models.EndPortAttributes) {
	u, _ := uuid.NewV4()
	fmt.Printf("End Port ID Generated: %s\n", u.String())
	Id := agilec.String(u.String())
	Name := agilec.String("CLARANET-GO-TESTS-001")
	EndPort := models.EndPortAttributes{}
	EndPort.Description = agilec.String("Created By GO")
	EndPort.LogicNetworkId = agilec.String("5308df55-1709-404f-b4f8-4d8947d8f0c4")
	EndPort.LogicPortId = agilec.String("8682b032-6bb3-46f8-b7b5-2c8bcfceefff")
	EndPort.Location = agilec.String("10")
	EndPort.VmName = agilec.String("10")
	EndPort.Ipv4 = []*string{
		agilec.String("192.168.1.1"),
	}
	EndPort.Ipv6 = []*string{
		agilec.String("FE80::A1"),
	}

	return Id, Name, &EndPort
}

func TestCreateEndPort(t *testing.T) {
	id, name, endPortAttr := GetEndPortAttributes()
	defer DeleteEndPort(*id)
	client := helper.GetClient()
	err := client.CreateEndPort(id, name, endPortAttr)
	assert.Nil(t, err)
	endPort, err := client.GetEndPort(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, endPort.Id)
	assert.Equal(t, name, endPort.Name)
	assert.Equal(t, *endPortAttr.LogicPortId, *endPort.LogicPortId)
	assert.Equal(t, *endPortAttr.Description, *endPort.Description)
	assert.Equal(t, *endPortAttr.Location, *endPort.Location)
	assert.Equal(t, *endPortAttr.LogicNetworkId, *endPort.LogicNetworkId)
	assert.Equal(t, endPortAttr.Ipv4, endPort.Ipv4)
	assert.Equal(t, endPortAttr.Ipv6, endPort.Ipv6)
	assert.Equal(t, *endPortAttr.VmName, *endPort.VmName)
}

func TestCreateEndPortDuplicate(t *testing.T) {
	id, name, endPortAttr := GetEndPortAttributes()
	defer DeleteEndPort(*id)
	client := helper.GetClient()
	err := client.CreateEndPort(id, name, endPortAttr)
	assert.Nil(t, err)
	err = client.CreateEndPort(id, name, endPortAttr)
	assert.NotNil(t, err)
}

func TestCreateEndPortInvalidID(t *testing.T) {
	_, name, endPortAttr := GetEndPortAttributes()
	id := agilec.String("dummy")
	client := helper.GetClient()
	err := client.CreateEndPort(id, name, endPortAttr)
	assert.NotNil(t, err)
}

func TestUpdateEndPort(t *testing.T) {
	id, name, endPortAttr := GetEndPortAttributes()
	defer DeleteEndPort(*id)
	client := helper.GetClient()
	err := client.CreateEndPort(id, name, endPortAttr)
	endPortAttr.Description = agilec.String("Updated From GO")
	endPort, err := client.UpdateEndPort(id, name, endPortAttr)
	assert.Nil(t, err)
	assert.Equal(t, endPortAttr.Description, endPort.Description)
	getEndPort := GetEndPort(*id)
	assert.Equal(t, getEndPort.Description, endPort.Description)
}

func TestUpdateNonExistingEndPort(t *testing.T) {
	client := helper.GetClient()
	u, _ := uuid.NewV4()
	_, err := client.UpdateEndPort(agilec.String(u.String()), agilec.String("dummy"), &models.EndPortAttributes{})
	assert.NotNil(t, err)
}

func TestGetEndPort(t *testing.T) {
	id, name, endPortAttr := GetEndPortAttributes()
	defer DeleteEndPort(*id)
	client := helper.GetClient()
	err := client.CreateEndPort(id, name, endPortAttr)
	endPort := GetEndPort(*id)
	assert.Nil(t, err)
	assert.Nil(t, err)
	assert.Equal(t, id, endPort.Id)
	assert.Equal(t, name, endPort.Name)
	assert.Equal(t, *endPortAttr.LogicPortId, *endPort.LogicPortId)
	assert.Equal(t, *endPortAttr.Description, *endPort.Description)
	assert.Equal(t, *endPortAttr.Location, *endPort.Location)
	assert.Equal(t, *endPortAttr.LogicNetworkId, *endPort.LogicNetworkId)
	assert.Equal(t, endPortAttr.Ipv4, endPort.Ipv4)
	assert.Equal(t, endPortAttr.Ipv6, endPort.Ipv6)
	assert.Equal(t, *endPortAttr.VmName, *endPort.VmName)
}

func TestListEndPorts(t *testing.T) {
	id, name, endPortAttr := GetEndPortAttributes()
	defer DeleteEndPort(*id)
	client := helper.GetClient()
	err := client.CreateEndPort(id, name, endPortAttr)
	assert.Nil(t, err)
	queryParameters := &models.EndPortRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListEndPorts(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}

func TestDeleteEndPort(t *testing.T) {
	id, name, endPortAttr:= GetEndPortAttributes()
	defer DeleteEndPort(*id)
	client := helper.GetClient()
	err := client.CreateEndPort(id, name, endPortAttr)
	assert.Nil(t, err)
	err = client.DeleteLogicalPort(*id)
	assert.Nil(t, err)
}

func GetEndPort(id string) *models.EndPort {
	client := helper.GetClient()
	endPort, _ := client.GetEndPort(id)
	return endPort
}

func DeleteEndPort(id string) {
	client := helper.GetClient()
	_ = client.DeleteEndPort(id)
}
