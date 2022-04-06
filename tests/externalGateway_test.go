package tests

import (
	"github.com/claranet/agilec-go-client/models"
	helper "github.com/claranet/agilec-go-client/tests/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetExternalGateway(t *testing.T) {
	client := helper.GetClient()
	externalGateway, err := client.GetExternalGateway("0ad60070-9bf9-4def-bf6e-3a86157eea1c")
	assert.Nil(t, err)
	assert.Equal(t, "DUAL-EXT-GW", *externalGateway.Name)
	assert.Empty(t, externalGateway.Description)
	assert.Equal(t, "Public", *externalGateway.GatewayType)
	assert.Len(t, externalGateway.PublicIpPool, 0)
	assert.Len(t, externalGateway.ServiceIpPool, 2)
	assert.Equal(t, "0.0.0.0/0", *externalGateway.ServiceIpPool[0])
	assert.Equal(t, "::/0", *externalGateway.ServiceIpPool[1])
	assert.Nil(t, externalGateway.IsTelcoGateway)
	assert.Equal(t, "DUAL-DUMMY", *externalGateway.VrfName)
	assert.True(t, *externalGateway.IsAllShared)
}

func TestListExternalGateway(t *testing.T) {
	queryParameters := &models.ExternalGatewayRequestOpts{}
	queryParameters.PageSize = 3
	queryParameters.PageIndex = 1
	client := helper.GetClient()
	response, err := client.ListExternalGateways(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}
