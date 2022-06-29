package tests

import (
	"fmt"
	agilec "github.com/claranet/agilec-go-client/client"
	"github.com/claranet/agilec-go-client/models"
	helper "github.com/claranet/agilec-go-client/tests/helpers"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetLogicalRouterAttributes() (*string, *string, *models.LogicalRouterAttributes) {
	u, _ := uuid.NewV4()
	fmt.Printf("Logical Router ID Generated: %s\n", u.String())
	Id := agilec.String(u.String())
	Name := agilec.String("logicRouterName")
	LogicalRouter := models.LogicalRouterAttributes{}
	LogicalRouter.Description = agilec.String("Created By GO")
	LogicalRouter.Type = agilec.String("Normal")
	LogicalRouter.LogicNetworkId = agilec.String("acfd8aaf-c6dc-499d-8020-bebd85b1f0e6")
	LogicalRouter.RouterLocations = []*models.LogicalRouterLocations{{
		FabricId: agilec.String("f1429224-1860-4bdb-8cc8-98ccc0f5563a"),
	},
	}
	//LogicalRouter.Vni = agilec.Int32(10)
	LogicalRouter.VrfName = agilec.String("ManagementRT_67766776")
	LogicalRouter.Additional = &models.LogicalRouterAdditional{
		Producer: agilec.String("GOLANG"),
	}
	return Id, Name, &LogicalRouter
}

func TestCreateLogicalRouter(t *testing.T) {
	id, name, logicalRouterAttr := GetLogicalRouterAttributes()
	defer DeleteLogicalRouter(*id)
	client := helper.GetClient()
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.Nil(t, err)
	logicalRouter, err := client.GetLogicalRouter(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalRouter.Id)
	assert.Equal(t, name, logicalRouter.Name)
}

func TestCreateLogicalRouterDuplicate(t *testing.T) {
	id, name, logicalRouterAttr := GetLogicalRouterAttributes()
	defer DeleteLogicalRouter(*id)
	client := helper.GetClient()
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.Nil(t, err)
	err = client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.NotNil(t, err)
}

func TestCreateLogicalRouterInvalidID(t *testing.T) {
	_, name, logicalRouterAttr := GetLogicalRouterAttributes()
	id := agilec.String("dummy")
	client := helper.GetClient()
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.NotNil(t, err)
}

func TestUpdateLogicalRouter(t *testing.T) {
	id, name, logicalRouterAttr := GetLogicalRouterAttributes()
	defer DeleteLogicalRouter(*id)
	client := helper.GetClient()
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	logicalRouterAttr.Description = agilec.String("Updated From GO")
	logicalRouter, err := client.UpdateLogicalRouter(id, name, logicalRouterAttr)
	assert.Nil(t, err)
	assert.Equal(t, logicalRouterAttr.Description, logicalRouter.Description)
	getLogicalRouter := GetLogicalRouter(*id)
	assert.Equal(t, getLogicalRouter.Description, logicalRouter.Description)
}

func TestUpdateNonExistingLogicalRouter(t *testing.T) {
	client := helper.GetClient()
	u, _ := uuid.NewV4()
	_, err := client.UpdateLogicalRouter(agilec.String(u.String()), agilec.String("dummy"), &models.LogicalRouterAttributes{})
	assert.NotNil(t, err)
}

func TestGetLogicaRouter(t *testing.T) {
	id, name, logicalRouterAttr := GetLogicalRouterAttributes()
	defer DeleteLogicalRouter(*id)
	client := helper.GetClient()
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	logicalRouter := GetLogicalRouter(*id)
	assert.Nil(t, err)
	assert.Equal(t, id, logicalRouter.Id, id)
	assert.Equal(t, name, logicalRouter.Name, name)
	assert.Equal(t, logicalRouterAttr.Description, logicalRouter.Description)
	assert.Equal(t, logicalRouterAttr.Type, logicalRouter.Type)
	//assert.Equal(t, logicalRouterAttr.Vni, logicalRouter.Vni)
	assert.Equal(t, logicalRouterAttr.VrfName, logicalRouter.VrfName)
	assert.Equal(t, logicalRouterAttr.LogicNetworkId, logicalRouter.LogicNetworkId)
	//assert.Equal(t, logicalRouterAttr.RouterLocations, logicalRouter.RouterLocations)
}

func TestGetNonExistLogicalRouter(t *testing.T) {
	client := helper.GetClient()
	u, _ := uuid.NewV4()
	_, err := client.GetLogicalRouter(u.String())
	assert.NotNil(t, err)
}

func TestListLogicalRouters(t *testing.T) {
	id, name, logicalRouterAttr := GetLogicalRouterAttributes()
	defer DeleteLogicalRouter(*id)
	client := helper.GetClient()
	err := client.CreateLogicalRouter(id, name, logicalRouterAttr)
	assert.Nil(t, err)
	queryParameters := &models.LogicalRouterRequestOpts{}
	queryParameters.PageSize = 3
	response, err := client.ListLogicalRouters(queryParameters)
	assert.Equal(t, 3, len(response))
	assert.Nil(t, err)
}

func TestDeleteLogicalRouter(t *testing.T) {
	id, _, _ := GetLogicalRouterAttributes()
	client := helper.GetClient()
	err := client.DeleteLogicalRouter(*id)
	assert.Nil(t, err)
}

func GetLogicalRouter(id string) *models.LogicalRouter {
	client := helper.GetClient()
	logicalRouter, _ := client.GetLogicalRouter(id)
	return logicalRouter
}

func DeleteLogicalRouter(id string) {
	client := helper.GetClient()
	_ = client.DeleteLogicalRouter(id)
}
