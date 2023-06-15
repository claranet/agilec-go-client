package client

// import (
// 	"testing"
// 	uuid "github.com/nu7hatch/gouuid"
// 	"github.com/claranet/agilec-go-client/models"
// 	"github.com/stretchr/testify/assert"
// )

// TODO Cannot enable due to missing dhcp groups on agile controller (nor ability to create it)

// var existingDHCPGroup models.DHCPGroup

// func TestListDHCPGroups(t *testing.T) {
// 	client := GetClientTest()
// 	queryParameters := &models.DHCPGroupRequestOpts{}
// 	queryParameters.PageSize = 3
// 	response, err := client.ListDHCPGroups(queryParameters)
// 	assert.Equal(t, 1, len(response))
// 	assert.Nil(t, err)
// 	existingDHCPGroup = *response[0]
// }

// func TestGetDHCPGroups(t *testing.T) {
// 	client := GetClientTest()
// 	dhcpGroup, err := client.GetDHCPGroup(*existingDHCPGroup.Id)
// 	assert.Nil(t, err)
// 	assert.Equal(t, existingDHCPGroup.Id, dhcpGroup.Id, existingDHCPGroup.Id)
// 	assert.Equal(t, existingDHCPGroup.Name, dhcpGroup.Name, existingDHCPGroup.Name)
// 	assert.Equal(t, existingDHCPGroup.ServerIp[0], dhcpGroup.ServerIp[0])
// 	assert.Equal(t, existingDHCPGroup.DHCPGroupAttributes.LogicRouterId, dhcpGroup.DHCPGroupAttributes.LogicRouterId)
// 	assert.Equal(t, existingDHCPGroup.DHCPGroupAttributes.VrfName, dhcpGroup.DHCPGroupAttributes.VrfName)
// 	assert.Equal(t, existingDHCPGroup.Description, dhcpGroup.Description)
// 	assert.Equal(t, existingDHCPGroup.DHCPGroupAttributes.Producer, dhcpGroup.DHCPGroupAttributes.Producer)
// 	assert.Equal(t, existingDHCPGroup.Dhcpgroupl2vni.Ipv4Cidr, dhcpGroup.Dhcpgroupl2vni.Ipv4Cidr)
// 	assert.Equal(t, existingDHCPGroup.Dhcpgroupl2vni.Ipv6Cidr, dhcpGroup.Dhcpgroupl2vni.Ipv6Cidr)
// 	assert.Equal(t, existingDHCPGroup.Dhcpgroupl2vni.L2Vni, dhcpGroup.Dhcpgroupl2vni.L2Vni)
// }

// func TestGetNonDHCPGroup(t *testing.T) {
// 	client := GetClientTest()
// 	u, _ := uuid.NewV4()
// 	_, err := client.GetDHCPGroup(u.String())
// 	if assert.NotNil(t, err) {
// 		response, ok := err.(*ErrorResponse)

// 		if !ok {
// 			t.Error("Wrong Error Response")
// 		}
// 		assert.Equal(t, "The Resource don't exists.", response.ErrorMessage)
// 		assert.Equal(t, "/controller/dc/v3//publicservice/dhcpgroups"+u.String(), response.URL)
// 		assert.Equal(t, 0, response.HttpStatusCode)
// 		assert.Equal(t, "Get", response.Method)
// 	}
// }