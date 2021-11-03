package client

import (
	"github.com/outscope-solutions/acdn-go-client/client"
	"github.com/outscope-solutions/acdn-go-client/models"
	"testing"
)

func TestClientAuthenticate(t *testing.T) {

	client := GetTestClient()

	Tenant := models.TenantAttributes{}
	Tenant.Name = "Outscope"
	Tenant.Description = "Created By GO"
	Tenant.Producer = "Default"
	Tenant.MulticastCapability = true
	Tenant.Quota = &models.TenantQuota{
		LogicVasNum: 10,
		LogicRouterNum: 5,
		LogicSwitchNum: 6,
	}
	Tenant.MulticastQuota = &models.TenantMulticastQuota{
		AclNum:     10,
		AclRuleNum: 10,
	}
	Tenant.ResPool = &models.TenantResPool{
		ExternalGatewayIds: []string{"1", "2"},
		FabricIds:          []string{"1", "2"},
		VmmIds:             []string{"1", "2"},
		DhcpGroupIds:       []string{"1", "2"},
	}
	client.CreateTenant("outscope_001", Tenant)
	//tenant, _ := client.ReadTenant(12)
	//fmt.Println(tenant.Producer)
	//err := client.Authenticate()
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//if client.AuthToken.Token == "" {
	//	t.Error("Token is empty")
	//}
}

func GetTestClient() *client.Client {

	return client.GetClient("https://u0bw69ss5e.execute-api.eu-west-1.amazonaws.com", "admin", client.Password("cisco123"), client.Insecure(true), client.SkipLoggingPayload(false))

	//return client.GetClient("https://192.168.10.102", "admin", client.Password("cisco123"), client.Insecure(true), client.SkipLoggingPayload(false))
}
