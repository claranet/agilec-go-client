package models

const TenantModuleName = "tenant"

type TenantRequestOpts struct {
	Producer string `json:"producer,omitempty"`
	BaseRequestOpts
}

type TenantResponseBody struct {
	TenantList
	BaseResponseAttributes
}

type TenantList struct {
	Tenant []Tenant `json:"tenant"`
}

type Tenant struct {
	Id                  string                `json:"id"`
	Name                string                `json:"name,omitempty"`
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