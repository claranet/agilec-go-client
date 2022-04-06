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

func GetLogicalNetworkAttributes() (*string, *string, *models.LogicalNetworkAttributes) {
	u, _ := uuid.NewV4()
	fmt.Printf("Logical Network ID Generated: %s\n", u.String())
	Id := agilec.String(u.String())
	Name := agilec.String("CLARANET-GO-TESTS-001")

	LogicalNetwork := models.LogicalNetworkAttributes{}
	LogicalNetwork.Description = agilec.String("Created By GO")
	LogicalNetwork.MulticastCapability = agilec.Bool(false)
	LogicalNetwork.Type = agilec.String("Transit")
	LogicalNetwork.TenantId = agilec.String("7e0ba3e8-280d-420c-951a-b2fe79b4b68a")
	LogicalNetwork.FabricId = []*string{agilec.String("f1429224-1860-4bdb-8cc8-98ccc0f5563a")}

	LogicalNetwork.Additional = &models.LogicalNetworkAdditional{
		Producer: agilec.String("GOLANG"),
	}
	return Id, Name, &LogicalNetwork
}

func TestCreateLogicalNetwork(t *testing.T) {
	id, name, logicalNetworktAttr := GetLogicalNetworkAttributes()
	defer DeleteLogicalNetwork(*id)
	client := helper.GetClient()
	err := client.CreateLogicalNetwork(id, name, logicalNetworktAttr)
	assert.Nil(t, err)
	logicalNetwork, err := client.GetLogicalNetwork(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalNetwork.Id)
	assert.Equal(t, name, logicalNetwork.Name)
}

func TestCreateLogicalNetworkDuplicate(t *testing.T) {
	id, name, logicalNetworktAttr := GetLogicalNetworkAttributes()
	defer DeleteLogicalNetwork(*id)
	client := helper.GetClient()
	err := client.CreateLogicalNetwork(id, name, logicalNetworktAttr)
	assert.Nil(t, err)
	err = client.CreateLogicalNetwork(id, name, logicalNetworktAttr)
	assert.NotNil(t, err)
}

func TestCreateLogicalNetworkInvalidID(t *testing.T) {
	_, name, logicalNetworktAttr := GetLogicalNetworkAttributes()
	id := agilec.String("dummy")
	client := helper.GetClient()
	err := client.CreateLogicalNetwork(id, name, logicalNetworktAttr)
	assert.NotNil(t, err)
}

func TestUpdateLogicalNetwork(t *testing.T) {
	id, name, logicalNetworktAttr := GetLogicalNetworkAttributes()
	defer DeleteLogicalNetwork(*id)
	client := helper.GetClient()
	err := client.CreateLogicalNetwork(id, name, logicalNetworktAttr)
	logicalNetworktAttr.Description = agilec.String("Updated From GO")
	logicalNetwork, err := client.UpdateLogicalNetwork(id, name, logicalNetworktAttr)
	assert.Nil(t, err)
	assert.Equal(t, logicalNetworktAttr.Description, logicalNetwork.Description)
	getLogicalNetwork := GetLogicalNetwork(*id)
	assert.Equal(t, getLogicalNetwork.Description, logicalNetwork.Description)
}

func TestUpdateNonExistingLogicalNetwork(t *testing.T) {
	client := helper.GetClient()
	u, _ := uuid.NewV4()
	_, err := client.UpdateLogicalNetwork(agilec.String(u.String()), agilec.String("dummy"), &models.LogicalNetworkAttributes{})
	assert.NotNil(t, err)
}

func TestGetLogicalNetwork(t *testing.T) {
	id, name, logicalNetworktAttr := GetLogicalNetworkAttributes()
	defer DeleteLogicalNetwork(*id)
	client := helper.GetClient()
	err := client.CreateLogicalNetwork(id, name, logicalNetworktAttr)
	logicalNetwork := GetLogicalNetwork(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalNetwork.Id, id)
	assert.Equal(t, name, logicalNetwork.Name, name)
	assert.Equal(t, logicalNetworktAttr.Description, logicalNetwork.Description)
	assert.Equal(t, logicalNetworktAttr.TenantId, logicalNetwork.TenantId)
	assert.Equal(t, logicalNetworktAttr.MulticastCapability, logicalNetwork.MulticastCapability)
	assert.Equal(t, logicalNetworktAttr.FabricId, logicalNetwork.FabricId)
}

func TestGetNonExistLogicalNetwork(t *testing.T) {
	client := helper.GetClient()
	u, _ := uuid.NewV4()
	_, err := client.GetLogicalNetwork(u.String())
	assert.NotNil(t, err)
}

func TestListLogicalNetworks(t *testing.T) {
	id, name, logicalNetworktAttr := GetLogicalNetworkAttributes()
	defer DeleteLogicalNetwork(*id)
	client := helper.GetClient()
	err := client.CreateLogicalNetwork(id, name, logicalNetworktAttr)
	assert.Nil(t, err)
	queryParameters := &models.LogicalNetworkRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListLogicalNetworks(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}

func TestDeleteLogicalNetwork(t *testing.T) {
	id, _, _ := GetLogicalNetworkAttributes()
	client := helper.GetClient()
	err := client.DeleteLogicalNetwork(*id)
	assert.Nil(t, err)
}

func GetLogicalNetwork(id string) *models.LogicalNetwork {
	client := helper.GetClient()
	logicalNetwork, _ := client.GetLogicalNetwork(id)
	return logicalNetwork
}

func DeleteLogicalNetwork(id string) {
	client := helper.GetClient()
	_ = client.DeleteLogicalNetwork(id)
}
