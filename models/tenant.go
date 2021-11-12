package models

const TenantModuleName = "tenant"

type TenantRequestOpts struct {
	Producer string `url:"producer,omitempty"`
	BaseRequestOpts
}

type TenantResponse struct {
	TenantList
	//BaseResponseAttributes
}

type TenantList struct {
	Tenants []Tenant `json:"tenant"`
}

type Tenant struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	TenantAttributes
}

type TenantAttributes struct {
	Description         string                `json:"description,omitempty"`
	Producer            string                `json:"producer,omitempty"`
	CreateAt            string                `json:"createAt,omitempty"`
	UpdateAt            string                `json:"updateAt,omitempty"`
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

func NewTenant(id, name string, tenantAttr TenantAttributes) *Tenant {
	return &Tenant{
		Id:               id,
		Name:             name,
		TenantAttributes: tenantAttr,
	}
}
