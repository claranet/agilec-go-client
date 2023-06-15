package client

import (
	"testing"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/claranet/agilec-go-client/models"
	"github.com/stretchr/testify/assert"
)

var existingFabric models.Fabric

func TestListFabrics(t *testing.T) {
	client := GetClientTest()
	queryParameters := &models.FabricRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListFabrics(queryParameters)
	assert.Equal(t, 1, len(response))
	assert.Nil(t, err)
	existingFabric = *response[0]
}

func TestGetFabric(t *testing.T) {
	client := GetClientTest()
	fabric, err := client.GetFabric(*existingFabric.Id)
	assert.Nil(t, err)
	assert.Equal(t, existingFabric.Id, fabric.Id, existingFabric.Id)
	assert.Equal(t, existingFabric.Name, fabric.Name, existingFabric.Name)
	assert.Equal(t, existingFabric.NetworkType, fabric.NetworkType)
	assert.Equal(t, existingFabric.PhysicalNetworkMode, fabric.PhysicalNetworkMode)
	assert.Equal(t, existingFabric.MicroSegmentCapability, fabric.MicroSegmentCapability)
	assert.Equal(t, existingFabric.Description, fabric.Description)
	assert.Equal(t, existingFabric.BgpEvpnEnable, fabric.BgpEvpnEnable)
	assert.Equal(t, existingFabric.ExtInterfaceType, fabric.ExtInterfaceType)
	assert.Equal(t, existingFabric.ArpBroadcastSuppression, fabric.ArpBroadcastSuppression)
	assert.Equal(t, existingFabric.DciSplitGroup, fabric.DciSplitGroup)
	assert.Equal(t, existingFabric.MulticastCapability, fabric.MulticastCapability)
	assert.Equal(t, existingFabric.SegmentMasks, fabric.SegmentMasks)
	assert.Equal(t, existingFabric.AclModel, fabric.AclModel)
	assert.Equal(t, existingFabric.SfcCapability, fabric.SfcCapability)
	assert.Equal(t, existingFabric.FaultService, fabric.FaultService)
	assert.Equal(t, existingFabric.UnderLay, fabric.UnderLay)
	assert.Equal(t, existingFabric.Management, fabric.Management)
}

func TestGetNonExistFabric(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.GetFabric(u.String())
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}
		assert.Equal(t, "The Resource don't exists.", response.ErrorMessage)
		assert.Equal(t, "/controller/dc/v3/physicalnetwork/fabricresource/fabrics/fabric/"+u.String(), response.URL)
		assert.Equal(t, 0, response.HttpStatusCode)
		assert.Equal(t, "Get", response.Method)
	}
}