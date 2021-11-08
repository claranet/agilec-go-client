package models

import (
	"encoding/json"
)

const TenantModuleName = "tenant"

type TenantRequestParameters struct {
	Producer string `json:"producer,omitempty"`
	BaseRequestParameters
}

type TenantResponseBody struct {
	Tenant []Tenant `json:"tenant"`
	BaseResponseAttributes
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

func (tenant *Tenant) ToJson() ([]byte, error) {
	return json.Marshal(&tenant)
}

func TenantFromContainerList(body *TenantResponseBody, index int) *Tenant {
	return &body.Tenant[index]
}

func TenantFromResponse(body *TenantResponseBody) *Tenant {
	return TenantFromContainerList(body, 0)
}
