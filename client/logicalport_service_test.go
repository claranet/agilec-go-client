package client

import (
	"fmt"
	"github.com/claranet/agilec-go-client/models"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateLogicalPortCompleteSuccess(t *testing.T) {
	client := GetClientTest()

	id, name, logicalPortAttr := GetLogicalPortAttributes(t, client)
	err := client.CreateLogicalPort(id, name, logicalPortAttr)
	defer DeleteLogicalPort(t, *id)
	assert.Nil(t, err)

	logicalPort, err := client.GetLogicalPort(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalPort.Id, id)
	assert.Equal(t, name, logicalPort.Name, name)
	assert.Equal(t, logicalPortAttr.Description, logicalPort.Description)
	// assert.Equal(t, logicalPortAttr.TenantId, logicalPort.TenantId)
	assert.Equal(t, logicalPortAttr.FabricId, logicalPort.FabricId)
	assert.Equal(t, logicalPortAttr.LogicSwitchId, logicalPort.LogicSwitchId)
	// assert.Equal(t, logicalPortAttr.MetaData, logicalPort.MetaData)
	assert.Equal(t, logicalPortAttr.Additional.Producer, logicalPort.Additional.Producer)
	assert.Equal(t, logicalPortAttr.AccessInfo.Mode, logicalPort.AccessInfo.Mode)
	assert.Equal(t, logicalPortAttr.AccessInfo.Type, logicalPort.AccessInfo.Type)
	assert.Equal(t, logicalPortAttr.AccessInfo.Vlan, logicalPort.AccessInfo.Vlan)
	assert.Equal(t, logicalPortAttr.AccessInfo.Location[0].PortName, logicalPort.AccessInfo.Location[0].PortName)
	assert.Equal(t, logicalPortAttr.AccessInfo.Location[0].DeviceIp, logicalPort.AccessInfo.Location[0].DeviceIp)
	assert.Equal(t, logicalPortAttr.AccessInfo.Location[1].PortName, logicalPort.AccessInfo.Location[1].PortName)
	assert.Equal(t, logicalPortAttr.AccessInfo.Location[1].DeviceIp, logicalPort.AccessInfo.Location[1].DeviceIp)
	assert.Equal(t, logicalPortAttr.AccessInfo.SubinterfaceNumber, logicalPort.AccessInfo.SubinterfaceNumber)
	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.InnerVidBegin, logicalPort.AccessInfo.Qinq.InnerVidBegin)
	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.InnerVidEnd, logicalPort.AccessInfo.Qinq.InnerVidEnd)
	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.OuterVidBegin, logicalPort.AccessInfo.Qinq.OuterVidBegin)
	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.OuterVidEnd, logicalPort.AccessInfo.Qinq.OuterVidEnd)
}

func TestCreateLogicalPortDuplicate(t *testing.T) {
	client := GetClientTest()
	id, name, logicalPortAttr := GetLogicalPortAttributes(t, client)
	defer DeleteLogicalPort(t, *id)
	err := client.CreateLogicalPort(id, name, logicalPortAttr)
	assert.Nil(t, err)
	err = client.CreateLogicalPort(id, name, logicalPortAttr)
	assert.NotNil(t, err)
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}

		assert.Contains(t, response.ErrorMessage, "Parameter logic port id already exists.")
		assert.Equal(t, "/controller/dc/v3/logicnetwork/ports", response.URL)
		assert.Equal(t, 400, response.HttpStatusCode)
		assert.Equal(t, "Post", response.Method)
	}
}

func TestCreateLogicalPortInvalidID(t *testing.T) {
	client := GetClientTest()
	_, name, logicalPortAttr := GetLogicalPortAttributes(t, client)
	id := String("dummy")
	defer DeleteLogicalPort(t, *id)
	err := client.CreateLogicalPort(id, name, logicalPortAttr)
	assert.NotNil(t, err)
}

func TestUpdateLogicalPort(t *testing.T) {
	client := GetClientTest()
	id, name, logicalPortAttr := GetLogicalPortAttributes(t, client)
	defer DeleteLogicalPort(t, *id)
	err := client.CreateLogicalPort(id, name, logicalPortAttr)
	assert.Nil(t, err)
	logicalPortAttr.Description = String("Updated From Go")
	logicalPort, err := client.UpdateLogicalPort(id, name, logicalPortAttr)
	assert.Nil(t, err)
	assert.Equal(t, logicalPortAttr.Description, logicalPort.Description)
	getlogicalPort := GetLogicalPort(t, *id)
	assert.Equal(t, getlogicalPort.Description, logicalPort.Description)
}

func TestGetLogicalPort(t *testing.T) {
	client := GetClientTest()
	id, name, logicalPortAttr := GetLogicalPortAttributes(t, client)
	defer DeleteLogicalPort(t, *id)
	err := client.CreateLogicalPort(id, name, logicalPortAttr)
	logicalPort, err := client.GetLogicalPort(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalPort.Id, id)
	assert.Equal(t, name, logicalPort.Name, name)
	assert.Equal(t, logicalPortAttr.Description, logicalPort.Description)
	// assert.Equal(t, logicalPortAttr.TenantId, logicalPort.TenantId)
	assert.Equal(t, logicalPortAttr.FabricId, logicalPort.FabricId)
	assert.Equal(t, logicalPortAttr.LogicSwitchId, logicalPort.LogicSwitchId)
	// assert.Equal(t, logicalPortAttr.MetaData, logicalPort.MetaData)
	assert.Equal(t, logicalPortAttr.Additional.Producer, logicalPort.Additional.Producer)
	assert.Equal(t, logicalPortAttr.AccessInfo.Mode, logicalPort.AccessInfo.Mode)
	assert.Equal(t, logicalPortAttr.AccessInfo.Type, logicalPort.AccessInfo.Type)
	assert.Equal(t, logicalPortAttr.AccessInfo.Vlan, logicalPort.AccessInfo.Vlan)
	assert.Equal(t, logicalPortAttr.AccessInfo.SubinterfaceNumber, logicalPort.AccessInfo.SubinterfaceNumber)
	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.InnerVidBegin, logicalPort.AccessInfo.Qinq.InnerVidBegin)
	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.InnerVidEnd, logicalPort.AccessInfo.Qinq.InnerVidEnd)
	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.OuterVidBegin, logicalPort.AccessInfo.Qinq.OuterVidBegin)
	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.OuterVidEnd, logicalPort.AccessInfo.Qinq.OuterVidEnd)
}

func TestGetNonExistLogicalPort(t *testing.T) {
	client := GetClientTest()
	u, _ := uuid.NewV4()
	_, err := client.GetLogicalPort(u.String())
	if assert.NotNil(t, err) {
		response, ok := err.(*ErrorResponse)

		if !ok {
			t.Error("Wrong Error Response")
		}
		assert.Equal(t, "The Resource don't exists.", response.ErrorMessage)
		assert.Equal(t, "/controller/dc/v3/logicnetwork/ports/port/"+u.String(), response.URL)
		assert.Equal(t, 0, response.HttpStatusCode)
		assert.Equal(t, "Get", response.Method)
	}
}

func TestListLogicalPort(t *testing.T) {
	client := GetClientTest()
	id, name, logicalPortAttr := GetLogicalPortAttributes(t, client)
	defer DeleteLogicalPort(t, *id)
	err := client.CreateLogicalPort(id, name, logicalPortAttr)
	assert.Nil(t, err)
	queryParameters := &models.LogicalPortRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListLogicalPorts(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}

func TestDeleteLogicalPort(t *testing.T) {
	client := GetClientTest()
	id, _, _ := GetLogicalPortAttributes(t, client)
	err := client.DeleteLogicalPort(*id)
	assert.Nil(t, err)
}

func GetLogicalPortAttributes(t *testing.T, client *Client) (*string, *string, *models.LogicalPortAttributes) {
	t.Helper()
	u, _ := uuid.NewV4()
	fmt.Printf("Logical Port ID Generated: %s\n", u.String())
	Id := String(u.String())
	Name := String("CLARANET-GO-TESTS-001")

	LogicalPort := models.LogicalPortAttributes{}
	LogicalPort.Description = String("Created By GO")
	LogicalPort.AccessInfo = &models.LogicalPortAccessInfo{
		Mode: String("Uni"),
		Type: String("Dot1q"),
		Vlan: Int32(1218),
		Qinq: &models.LogicalPortAccessInfoQinq{
			InnerVidBegin: Int32(10),
			InnerVidEnd: Int32(10),
			OuterVidBegin: Int32(10),
			OuterVidEnd: Int32(10),
			RewriteAction: String("POPDOUBLE"),
		},
		Location: []*models.LogicalPortAccessInfoLocation{
			{
				PortName: String("Eth-Trunk11"),
				DeviceIp: String("10.32.32.131"),
			},
			{
				PortName: String("Eth-Trunk11"),
				DeviceIp: String("10.32.32.130"),
			},
		},
		SubinterfaceNumber: Int32(16),
	}
	LogicalPort.MetaData = String("METADATA test")
	LogicalPort.Additional = &models.LogicalPortAdditional{
		Producer:    String("Test Producer"),
	}
	
	// Get Fabric
	responseFabric, err := client.ListFabrics(&models.FabricRequestOpts{
		BaseRequestOpts: models.BaseRequestOpts{
			PageSize:  1,
			PageIndex: 1,
		},
	})
	if err != nil {
		t.Logf("GetLogicalPortAttributes %s", err)
	}
	assert.Equal(t, 1, len(responseFabric))

	LogicalPort.FabricId = responseFabric[0].Id

	// Get Tenant
	// responseTenants, err := client.ListTenants(&models.TenantRequestOpts{
	// 	BaseRequestOpts: models.BaseRequestOpts{
	// 		PageSize:  1,
	// 		PageIndex: 1,
	// 	},
	// })
	// if err != nil {
	// 	t.Logf("GetLogicalPortAttributes %s", err)
	// }
	// assert.Equal(t, 1, len(responseTenants))

	// LogicalPort.TenantId = responseTenants[0].Id

	// Get Logical Switch
	responseLogicalSwitch, err := client.ListLogicalSwitches(&models.LogicalSwitchRequestOpts{
		BaseRequestOpts: models.BaseRequestOpts{
			PageSize:  1,
			PageIndex: 1,
		},
	})
	if err != nil {
		t.Logf("GetLogicalPortAttributes %s", err)
	}
	assert.Equal(t, 1, len(responseLogicalSwitch))

	LogicalPort.LogicSwitchId = responseLogicalSwitch[0].Id

	return Id, Name, &LogicalPort
}

func GetLogicalPort(t *testing.T, id string) *models.LogicalPort {
	t.Helper()
	client := GetClientTest()
	logicalPort, _ := client.GetLogicalPort(id)
	return logicalPort
}

func DeleteLogicalPort(t *testing.T, id string) {
	t.Helper()
	client := GetClientTest()
	_ = client.DeleteLogicalPort(id)
}