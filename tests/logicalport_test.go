package tests

//import (
//	"fmt"
//	agilec "github.com/claranet/agilec-go-client/client"
//	"github.com/claranet/agilec-go-client/models"
//	helper "github.com/claranet/agilec-go-client/tests/helpers"
//	uuid "github.com/nu7hatch/gouuid"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func GetLogicalPortAttributes() (*string, *string, *models.LogicalPortAttributes) {
//	u, _ := uuid.NewV4()
//	fmt.Printf("Logical Port ID Generated: %s\n", u.String())
//	Id := agilec.String(u.String())
//	Name := agilec.String("CLARANET-GO-TESTS-001")
//	LogicalPort := models.LogicalPortAttributes{}
//	LogicalPort.Description = agilec.String("Created By GO")
//	LogicalPort.LogicSwitchId = agilec.String("6c0a96d3-0789-47e6-9dbc-66ac5ba2e519")
//	LogicalPort.MetaData = agilec.String("Metadata")
//	LogicalPort.TenantId = agilec.String("11ade37a-79d0-482f-a7a0-6ad070e1d05d")
//	LogicalPort.FabricId = agilec.String("f1429224-1860-4bdb-8cc8-98ccc0f5563a")
//	LogicalPort.AccessInfo = &models.LogicalPortAccessInfo{
//		Mode: agilec.String("Uni"),
//		Type: agilec.String("Dot1q"),
//		Vlan: agilec.Int32(1215),
//		Qinq: &models.LogicalPortAccessInfoQinq{
//			InnerVidBegin: agilec.Int32(10),
//			InnerVidEnd:   agilec.Int32(10),
//			OuterVidBegin: agilec.Int32(10),
//			OuterVidEnd:   agilec.Int32(10),
//			RewriteAction: agilec.String("PopDouble"),
//		},
//		Location: []*models.LogicalPortAccessInfoLocation{
//			{
//				DeviceGroupId: agilec.String("e13784fb-499f-4c30-8f9c-e49e6c98fdbb"),
//				DeviceId:      agilec.String("9e3a5bee-3d95-3bf7-90f5-09bd2177324b"),
//				PortId:        agilec.String("589c87dd-7222-3c09-87b7-d09a236af285"),
//				//PortName:      agilec.String("10GE1/0/1"),
//				//DeviceIp:      agilec.String("192.168.1.1"),
//			},
//		},
//		SubinterfaceNumber: agilec.Int32(10),
//	}
//	LogicalPort.Additional = &models.LogicalPortAdditional{
//		Producer: agilec.String("GOLANG"),
//	}
//	return Id, Name, &LogicalPort
//}
//
//func TestCreateLogicalPort(t *testing.T) {
//	id, name, logicalPortAttr := GetLogicalPortAttributes()
//	defer DeleteLogicalPort(*id)
//	client := helper.GetClient()
//	err := client.CreateLogicalPort(id, name, logicalPortAttr)
//	assert.Nil(t, err)
//	logicalPort, err := client.GetLogicalPort(*id)
//	assert.Nil(t, err)
//	assert.Equal(t, id, logicalPort.Id)
//	assert.Equal(t, name, logicalPort.Name)
//}
//
//func TestCreateLogicalPortDuplicate(t *testing.T) {
//	id, name, logicalPortAttr := GetLogicalPortAttributes()
//	defer DeleteLogicalPort(*id)
//	client := helper.GetClient()
//	err := client.CreateLogicalPort(id, name, logicalPortAttr)
//	assert.Nil(t, err)
//	err = client.CreateLogicalPort(id, name, logicalPortAttr)
//	assert.NotNil(t, err)
//}
//
//func TestCreateLogicalPortInvalidID(t *testing.T) {
//	_, name, logicalPortAttr := GetLogicalPortAttributes()
//	id := agilec.String("dummy")
//	client := helper.GetClient()
//	err := client.CreateLogicalPort(id, name, logicalPortAttr)
//	assert.NotNil(t, err)
//}
//
//func TestUpdateLogicalPort(t *testing.T) {
//	id, name, logicalPortAttr := GetLogicalPortAttributes()
//	defer DeleteLogicalPort(*id)
//	client := helper.GetClient()
//	err := client.CreateLogicalPort(id, name, logicalPortAttr)
//	logicalPortAttr.Description = agilec.String("Updated From GO")
//	logicalPort, err := client.UpdateLogicalPort(id, name, logicalPortAttr)
//	assert.Nil(t, err)
//	assert.Equal(t, logicalPortAttr.Description, logicalPort.Description)
//	getLogicalPort := GetLogicalPort(*id)
//	assert.Equal(t, getLogicalPort.Description, logicalPort.Description)
//}
//
//func TestUpdateNonExistingLogicalPort(t *testing.T) {
//	client := helper.GetClient()
//	u, _ := uuid.NewV4()
//	_, err := client.UpdateLogicalPort(agilec.String(u.String()), agilec.String("dummy"), &models.LogicalPortAttributes{})
//	assert.NotNil(t, err)
//}
//
//func TestGetLogicalPort(t *testing.T) {
//	id, name, logicalPortAttr := GetLogicalPortAttributes()
//	defer DeleteLogicalPort(*id)
//	client := helper.GetClient()
//	err := client.CreateLogicalPort(id, name, logicalPortAttr)
//	logicalPort := GetLogicalPort(*id)
//	assert.Nil(t, err)
//	assert.Equal(t, id, logicalPort.Id, id)
//	assert.Equal(t, name, logicalPort.Name, name)
//	assert.Equal(t, logicalPortAttr.Description, logicalPort.Description)
//	assert.Equal(t, logicalPortAttr.TenantId, logicalPort.TenantId)
//	//assert.Equal(t, logicalPortAttr.MetaData, logicalPort.MetaData)
//	assert.Equal(t, logicalPortAttr.FabricId, logicalPort.FabricId)
//	assert.Equal(t, logicalPortAttr.LogicSwitchId, logicalPort.LogicSwitchId)
//	assert.Equal(t, logicalPortAttr.AccessInfo.Mode, logicalPort.AccessInfo.Mode)
//	assert.Equal(t, logicalPortAttr.AccessInfo.Type, logicalPort.AccessInfo.Type)
//	assert.Equal(t, logicalPortAttr.AccessInfo.Vlan, logicalPort.AccessInfo.Vlan)
//	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.InnerVidBegin, logicalPort.AccessInfo.Qinq.InnerVidBegin)
//	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.InnerVidEnd, logicalPort.AccessInfo.Qinq.InnerVidEnd)
//	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.OuterVidBegin, logicalPort.AccessInfo.Qinq.OuterVidBegin)
//	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.OuterVidEnd, logicalPort.AccessInfo.Qinq.OuterVidEnd)
//	assert.Equal(t, logicalPortAttr.AccessInfo.Qinq.RewriteAction, logicalPort.AccessInfo.Qinq.RewriteAction)
//	assert.Equal(t, logicalPortAttr.Additional.Producer, logicalPort.Additional.Producer)
//}
//
//func TestGetNonExistLogicalPort(t *testing.T) {
//	client := helper.GetClient()
//	u, _ := uuid.NewV4()
//	_, err := client.GetLogicalPort(u.String())
//	assert.NotNil(t, err)
//}
//
//func TestListLogicalPorts(t *testing.T) {
//	id, name, logicalPortAttr := GetLogicalPortAttributes()
//	defer DeleteLogicalPort(*id)
//	client := helper.GetClient()
//	err := client.CreateLogicalPort(id, name, logicalPortAttr)
//	assert.Nil(t, err)
//	queryParameters := &models.LogicalPortRequestOpts{}
//	queryParameters.PageSize = 3
//	response, err := client.ListLogicalPorts(queryParameters)
//	assert.Equal(t, 3, len(response))
//	assert.Nil(t, err)
//}
//
//func TestDeleteLogicalPort(t *testing.T) {
//	id, _, _ := GetLogicalPortAttributes()
//	client := helper.GetClient()
//	err := client.DeleteLogicalPort(*id)
//	assert.Nil(t, err)
//}
//
//func GetLogicalPort(id string) *models.LogicalPort {
//	client := helper.GetClient()
//	logicalPort, _ := client.GetLogicalPort(id)
//	return logicalPort
//}
//
//func DeleteLogicalPort(id string) {
//	client := helper.GetClient()
//	_ = client.DeleteLogicalPort(id)
//}
