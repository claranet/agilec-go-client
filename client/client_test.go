package client

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/outscope-solutions/acdn-go-client/client"
	"github.com/outscope-solutions/acdn-go-client/models"
	"os"
	"testing"
)

func TestClientAuthenticateOK(t *testing.T) {

	u, _ := uuid.NewV4()
	fmt.Println(u)
	client := GetTestClientOK()
	//err := client.Authenticate()
	//if err != nil {
	//	t.Error(err)
	//}

	//Tenant := models.Tenant{}
	//Tenant.Id = u.String()
	//Tenant.Name = "Outscope-demo-001"
	//Tenant.Description = "Created By GO"
	//Tenant.Producer = "Default"
	//Tenant.MulticastCapability = true
	//Tenant.Quota = &models.TenantQuota{
	//	LogicVasNum:    10,
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
	////_, err := client.CreateTenant(&Tenant)
	//
	//Tenant.Description = "Updated FROM GOLANG"

	//_, err := client.UpdateTenant(u.String(), &Tenant)

	//err := client.DeleteTenant("51c5b574-a4fa-44a9-617b-ac058674f4b7")

	ListOpts := &models.LogicalPortListOpts{}
	ListOpts.PageSize = 10
	ListOpts.PageIndex = 2
	ListOpts.LogicSwitchId = "bf1a554f-7cfa-4816-8aee-8353b12e844b"


	logicalPorts, err := client.GetLogicalPorts(ListOpts)

	logicalPorts.Next()

	if err != nil {
		t.Error(err)
	}

	fmt.Println("-----------")

	fmt.Println(logicalPorts.PageIndex)



	fmt.Println("-----------")
	//.previous()

	for _, element := range logicalPorts.Port {
		fmt.Println(element.Name)
	}


	//tenant, _ := client.GetTenant("ed0cb9aa-df73-4bd3-a71b-893c248e2c5c")
	//fmt.Println(tenant.MulticastCapability)

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
