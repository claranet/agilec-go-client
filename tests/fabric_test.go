package tests

import (
	"agilec-go-client/models"
	helper "agilec-go-client/tests/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFabric(t *testing.T) {
	client := helper.GetClient()
	fabric, err := client.GetFabric("aa1a1e72-e2af-44b8-adb3-6f1bf82291f6")
	assert.Nil(t, err)
	assert.Equal(t, "SDN-teste", *fabric.Name)
	assert.Empty(t, *fabric.Description)
	assert.Equal(t, "Distributed", *fabric.NetworkType)
	assert.Equal(t, "Vxlan", *fabric.PhysicalNetworkMode)
	assert.True(t, *fabric.BgpEvpnEnable)
	assert.Equal(t, "Vbdif", *fabric.ExtInterfaceType)
	assert.True(t, *fabric.ArpBroadcastSuppression)
	assert.Empty(t, fabric.DciSplitGroup)
	assert.True(t, *fabric.MicroSegmentCapability)
	assert.False(t, *fabric.MulticastCapability)
	assert.Nil(t, fabric.SegmentMasks)
	assert.Empty(t, fabric.AclModel)
	assert.Equal(t, "null", *fabric.SfcCapability)
}

func TestListFabrics(t *testing.T) {
	queryParameters := &models.FabricRequestOpts{}
	queryParameters.PageSize = 3
	queryParameters.PageIndex = 1
	client := helper.GetClient()
	response, err := client.ListFabrics(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}
