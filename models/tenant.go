package models

const TenantModuleName = "tenant"

type TenantRequestOpts struct {
	Producer string `url:"producer,omitempty"`
	BaseRequestOpts
}

type TenantResponse struct {
	TenantList
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
	Quota               *TenantQuota          `json:"quota"`
	MulticastQuota      *TenantMulticastQuota `json:"multicastQuota"`
	ResPool             *TenantResPool        `json:"ResPool,omitempty"`
}

type TenantQuota struct {
	LogicVasNum    int32 `json:"logicVasNum"`
	LogicRouterNum int32 `json:"logicRouterNum"`
	LogicSwitchNum int32 `json:"logicSwitchNum"`
}

type TenantMulticastQuota struct {
	AclNum     int32 `json:"aclNum"`
	AclRuleNum int32 `json:"aclRuleNum"`
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

func (resp *TenantResponse) Count() int {
	return len(resp.Tenants)
}