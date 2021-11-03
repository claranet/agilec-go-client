package client

import (
	"github.com/outscope-solutions/acdn-go-client/client"
	"os"
	"testing"
)

func TestClientAuthenticateOK(t *testing.T) {
	client := GetTestClientOK()
	//err := client.Authenticate()
	//if err != nil {
	//	t.Error(err)
	//}

	//Tenant := models.TenantAttributes{}
	//Tenant.Name = "Outscope"
	//Tenant.Description = "Created By GO"
	//Tenant.Producer = "Default"
	//Tenant.MulticastCapability = true
	//Tenant.Quota = &models.TenantQuota{
	//	LogicVasNum: 10,
	//	LogicRouterNum: 5,
	//	LogicSwitchNum: 6,
	//}
	//Tenant.MulticastQuota = &models.TenantMulticastQuota{
	//	AclNum:     10,
	//	AclRuleNum: 10,
	//}
	//Tenant.ResPool = &models.TenantResPool{
	//	ExternalGatewayIds: []string{"1", "2"},
	//	FabricIds:          []string{"1", "2"},
	//	VmmIds:             []string{"1", "2"},
	//	DhcpGroupIds:       []string{"1", "2"},
	//}
	//client.CreateTenant("outscope_001", Tenant)
	//tenant, _ := client.ReadTenant(12)
	//fmt.Println(tenant.Producer)

	//if client.AuthToken.Token == "" {
	//	t.Error("Token is empty")
	//}
}

//func TestClientAuthenticateNOK(t *testing.T) {
//	client := GetTestClientNOK()
//	err := client.Authenticate()
//
//	if _, ok := err.(*UnexpectedResponseCodeError); ok {
//		fmt.Println("This is an int")
//	}
//
//	//switch t := err.(type) {
//	//default:
//	//	t.Error("Unexpected Error")
//	//case *client.UnexpectedResponseCodeError:
//	//
//	//}
//	//
//	//fmt.Println(reflect.TypeOf(err))
//	//if err != nil {
//	//	t.Error(err)
//	//}
//}

func GetTestClientOK() *client.Client {
	return client.GetClient(
		os.Getenv("AC_HOST"),
		os.Getenv("AC_USERNAME"),
		client.Password(os.Getenv("AC_PASSWORD")),
		client.Insecure(true),
		client.SkipLoggingPayload(true))
}

func GetTestClientNOK() *client.Client {
	return client.GetClient(
		os.Getenv("AC_HOST"),
		os.Getenv("AC_USERNAME"),
		client.Password("NOK"),
		client.Insecure(true),
		client.SkipLoggingPayload(true))
}
