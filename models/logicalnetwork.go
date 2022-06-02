package models

const LogicalNetworkModuleName = "network"

type LogicalNetworkRequestOpts struct {
	Producer     string `url:"producer,omitempty"`
	Type         string `url:"type,omitempty"`
	TenantId     string `url:"tenantId,omitempty"`
	FilterTenant bool   `url:"filterTenant,omitempty"`
	BaseRequestOpts
}

type LogicalNetworkListResponse struct {
	LogicalNetworkList
}

type LogicalNetworkList struct {
	LogicalNetworks []*LogicalNetwork `json:"network"`
}

type LogicalNetworkResponse struct {
	LogicalNetwork LogicalNetwork `json:"network"`
}

type LogicalNetwork struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
	LogicalNetworkAttributes
}

type LogicalNetworkAttributes struct {
	Description         *string                   `json:"description,omitempty"`
	TenantId            *string                   `json:"tenantId,omitempty"`
	FabricId            []*string                 `json:"fabricId,omitempty"`
	MulticastCapability *bool                     `json:"multicastCapability,omitempty"`
	Type                *string                   `json:"type,omitempty"`
	IsVpcDeployed       *bool                     `json:"isVpcDeployed,omitempty"`
	Additional          *LogicalNetworkAdditional `json:"additional,omitempty"`
}

type LogicalNetworkAdditional struct {
	Producer *string `json:"producer,omitempty"`
	CreateAt *string `json:"createAt,omitempty"`
	UpdateAt *string `json:"updateAt,omitempty"`
}

func NewLogicalNetwork(id, name *string, logicalNetworkAttr LogicalNetworkAttributes) *LogicalNetwork {
	return &LogicalNetwork{
		Id:                       id,
		Name:                     name,
		LogicalNetworkAttributes: logicalNetworkAttr,
	}
}

func (resp *LogicalNetworkListResponse) Count() int {
	return len(resp.LogicalNetworks)
}
