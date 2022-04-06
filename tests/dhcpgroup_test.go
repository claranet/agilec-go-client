package tests

import (
	agilec "agilec-go-client/client"
	"agilec-go-client/models"
	helper "agilec-go-client/tests/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDHCPGroup(t *testing.T) {
	client := helper.GetClient()
	_, err := client.GetDHCPGroup("0ad60070-9bf9-4def-bf6e-3a86157eea1c")
	if assert.NotNil(t, err) {
		response, ok := err.(*agilec.ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}
		assert.Equal(t, "The Resource don't exists.", response.ErrorMessage)
		assert.Equal(t, "/controller/dc/v3/publicservice/dhcpgroups/dhcpgroup/0ad60070-9bf9-4def-bf6e-3a86157eea1c", response.URL)
		assert.Equal(t, 0, response.HttpStatusCode)
		assert.Equal(t, "Get", response.Method)
	}
}

func TestListDHCPGroups(t *testing.T) {
	queryParameters := &models.DHCPGroupRequestOpts{}
	queryParameters.PageSize = 3
	queryParameters.PageIndex = 1
	client := helper.GetClient()
	response, err := client.ListDHCPGroups(queryParameters)
	assert.Equal(t, 0, len(response))
	assert.Nil(t, err)
}
