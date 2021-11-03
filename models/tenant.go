package models

import (
	"encoding/json"
	"github.com/outscope-solutions/acdn-go-client/container"
)

const TenantClassName = "tenant"

type Tenant struct {
	BaseAttributes
	TenantAttributes
}

type TenantAttributes struct {
	Name                string                `json:"name,omitempty"`
	Description         string                `json:"description,omitempty"`
	Producer            string                `json:"producer,omitempty"`
	MulticastCapability bool                  `json:"multicastCapability,omitempty"`
	Quota               *TenantQuota          `json:"quota,omitempty"`
	MulticastQuota      *TenantMulticastQuota `json:"multicastQuota,omitempty"`
	ResPool             *TenantResPool        `json:"ResPool,omitempty"`
}

type TenantQuota struct {
	LogicVasNum    int32 `json:"logicVasNum,omitempty"`
	LogicRouterNum int32 `json:"logicRouterNum,omitempty"`
	LogicSwitchNum int32 `json:"logicSwitchNum,omitempty"`
}

type TenantMulticastQuota struct {
	AclNum     int32 `json:"aclNum,omitempty"`
	AclRuleNum int32 `json:"aclRuleNum,omitempty"`
}

type TenantResPool struct {
	ExternalGatewayIds []string `json:"externalGatewayIds,omitempty"`
	FabricIds          []string `json:"fabricIds,omitempty"`
	VmmIds             []string `json:"vmmIds,omitempty"`
	DhcpGroupIds       []string `json:"dhcpGroupIds,omitempty"`
}

func NewTenant(id string, tenantAttrs TenantAttributes) *Tenant {
	return &Tenant{
		BaseAttributes: BaseAttributes{
			Id:        id,
			ClassName: TenantClassName,
		},
		TenantAttributes: tenantAttrs,
	}
}

//func (tenant *Tenant) ToMap() (map[string]interface{}, error) {
func (tenant *Tenant) ToMap() ([]byte, error) {

	//e, err := json.Marshal(&tenant)
	//dynamic := make(map[string]interface{})
	//json.Unmarshal([]byte(e), &dynamic)

	//tenantMap, err := tenant.BaseAttributes.ToMap()
	//if err != nil {
	//	return nil, err
	//}
	//A(tenantMap, "name", tenant.Name)
	//A(tenantMap, "description", tenant.Description)
	//A(tenantMap, "producer", tenant.Producer)
	//B(tenantMap, "multicastCapability", tenant.MulticastCapability)
	//C(tenantMap, "quota", tenant.Quota)
	return json.Marshal(&tenant)
}

func TenantFromContainerList(cont *container.Container, index int) *Tenant {

	TenantCont := cont.Index(index)
	return &Tenant{
		BaseAttributes{
			Id:        G(TenantCont, "id").(string),
			ClassName: TenantClassName,
		},

		TenantAttributes{
			Name:                G(TenantCont, "name").(string),
			Description:         G(TenantCont, "description").(string),
			Producer:            G(TenantCont, "producer").(string),
			MulticastCapability: G(TenantCont, "multicastCapability").(bool),
		},
	}
}

func TenantFromContainer(cont *container.Container) *Tenant {
	return TenantFromContainerList(cont.S(TenantClassName), 0)
}
