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

func GetLogicalSwitchAttributes() (*string, *string, *models.LogicalSwitchAttributes) {
	u, _ := uuid.NewV4()
	fmt.Printf("Logical Switch ID Generated: %s\n", u.String())
	Id := agilec.String(u.String())
	Name := agilec.String("logicSwitchName")
	LogicalSwitch := models.LogicalSwitchAttributes{}
	LogicalSwitch.Description = agilec.String("Created By GO")
	LogicalSwitch.LogicNetworkId = agilec.String("5308df55-1709-404f-b4f8-4d8947d8f0c4")
	LogicalSwitch.MacAddress = agilec.String("00:00:5E:00:01:02")
	LogicalSwitch.TenantId = agilec.String("cd27d9cf-9be0-4852-a560-2d6e05fd3c1e")
	LogicalSwitch.StormSuppress = &models.LogicalSwitchStormSuppress{
		BroadcastEnable: agilec.Bool(false),
		MulticastEnable: agilec.Bool(false),
		UnicastEnable:   agilec.Bool(false),
	}
	LogicalSwitch.Additional = &models.LogicalSwitchAdditional{
		Producer: agilec.String("GOLANG"),
	}
	return Id, Name, &LogicalSwitch
}

func TestCreateLogicalSwitch(t *testing.T) {
	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes()
	defer DeleteLogicalSwitch(*id)
	client := helper.GetClient()
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.Nil(t, err)
	logicalSwitch, err := client.GetLogicalSwitch(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalSwitch.Id)
	assert.Equal(t, name, logicalSwitch.Name)
	assert.Equal(t, *logicalSwitchAttr.Description, *logicalSwitch.Description)
	assert.Equal(t, *logicalSwitchAttr.LogicNetworkId, *logicalSwitch.LogicNetworkId)
	assert.Equal(t, *logicalSwitchAttr.TenantId, *logicalSwitch.TenantId)
	assert.Equal(t, *logicalSwitchAttr.MacAddress, *logicalSwitch.MacAddress)
}

func TestCreateLogicalSwitchDuplicate(t *testing.T) {
	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes()
	defer DeleteLogicalSwitch(*id)
	client := helper.GetClient()
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.Nil(t, err)
	err = client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.NotNil(t, err)
}

func TestCreateLogicalSwitchInvalidID(t *testing.T) {
	_, name, logicalSwitchAttr := GetLogicalSwitchAttributes()
	id := agilec.String("dummy")
	client := helper.GetClient()
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.NotNil(t, err)
}

func TestUpdateLogicalSwitch(t *testing.T) {
	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes()
	defer DeleteLogicalSwitch(*id)
	client := helper.GetClient()
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	logicalSwitchAttr.Description = agilec.String("Updated From GO")
	logicalSwitch, err := client.UpdateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.Nil(t, err)
	assert.Equal(t, logicalSwitchAttr.Description, logicalSwitch.Description)
	getLogicalSwitch := GetLogicalSwitch(*id)
	assert.Equal(t, getLogicalSwitch.Description, logicalSwitch.Description)
}

func TestUpdateNonExistingLogicalSwitch(t *testing.T) {
	client := helper.GetClient()
	u, _ := uuid.NewV4()
	_, err := client.UpdateLogicalSwitch(agilec.String(u.String()), agilec.String("dummy"), &models.LogicalSwitchAttributes{})
	assert.NotNil(t, err)
}

func TestGetLogicalSwitch(t *testing.T) {
	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes()
	defer DeleteLogicalSwitch(*id)
	client := helper.GetClient()
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	logicalSwitch := GetLogicalSwitch(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalSwitch.Id, id)
	assert.Equal(t, name, logicalSwitch.Name, name)
	assert.Equal(t, logicalSwitchAttr.Description, logicalSwitch.Description)
	assert.Equal(t, logicalSwitchAttr.TenantId, logicalSwitch.TenantId)
	assert.Equal(t, logicalSwitchAttr.LogicNetworkId, logicalSwitch.LogicNetworkId)
	assert.Equal(t, logicalSwitchAttr.MacAddress, logicalSwitch.MacAddress)
}

func TestGetNonExistLogicalSwitch(t *testing.T) {
	client := helper.GetClient()
	u, _ := uuid.NewV4()
	_, err := client.GetLogicalSwitch(u.String())
	assert.NotNil(t, err)
}

func TestListLogicalSwitches(t *testing.T) {
	id, name, logicalSwitchAttr := GetLogicalSwitchAttributes()
	defer DeleteLogicalSwitch(*id)
	client := helper.GetClient()
	err := client.CreateLogicalSwitch(id, name, logicalSwitchAttr)
	assert.Nil(t, err)
	queryParameters := &models.LogicalSwitchRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListLogicalSwitches(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}

func TestDeleteLogicalSwitch(t *testing.T) {
	id, _, _ := GetLogicalSwitchAttributes()
	client := helper.GetClient()
	err := client.DeleteLogicalSwitch(*id)
	assert.Nil(t, err)
}

func GetLogicalSwitch(id string) *models.LogicalSwitch {
	client := helper.GetClient()
	logicalSwitch, _ := client.GetLogicalSwitch(id)
	return logicalSwitch
}

func DeleteLogicalSwitch(id string) {
	client := helper.GetClient()
	_ = client.DeleteLogicalSwitch(id)
}
