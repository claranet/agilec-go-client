package models

const LogicalPortModuleName = "port"

type LogicalPortRequestOpts struct {
	Producer      string `url:"producer,omitempty"`
	LogicSwitchId string `url:"logicSwitchId,omitempty"`
	BaseRequestOpts
}

type LogicalPortResponse struct {
	LogicalPortList
}

type LogicalPortList struct {
	LogicalPorts []*LogicalPort `json:"port"`
}

type LogicalPort struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
	LogicalPortAttributes
}

type LogicalPortAttributes struct {
	Description   *string                `json:"description,omitempty"`
	TenantId      *string                `json:"tenantId,omitempty"`
	FabricId      *string                `json:"fabricId,omitempty"`
	LogicSwitchId *string                `json:"logicSwitchId,omitempty"`
	MetaData      *string                `json:"metaData,omitempty"`
	AccessInfo    *LogicalPortAccessInfo `json:"accessInfo,omitempty"`
	Additional    *LogicalPortAdditional `json:"additional,omitempty"`
}

type LogicalPortAccessInfo struct {
	Mode               *string                          `json:"mode,omitempty"`
	Type               *string                          `json:"type,omitempty"`
	Vlan               *int32                           `json:"vlan,omitempty"`
	Qinq               *LogicalPortAccessInfoQinq       `json:"qinq,omitempty"`
	Location           []*LogicalPortAccessInfoLocation `json:"location"`
	SubinterfaceNumber *int32                           `json:"subinterfaceNumber,omitempty"`
}

type LogicalPortAccessInfoQinq struct {
	InnerVidBegin *int32  `json:"innerVidBegin,omitempty"`
	InnerVidEnd   *int32  `json:"innerVidEnd,omitempty"`
	OuterVidBegin *int32  `json:"outerVidBegin,omitempty"`
	OuterVidEnd   *int32  `json:"outerVidEnd,omitempty"`
	RewriteAction *string `json:"rewriteAction,omitempty"`
}

type LogicalPortAccessInfoLocation struct {
	DeviceGroupId *string `json:"deviceGroupId,omitempty"`
	DeviceId      *string `json:"deviceId,omitempty"`
	PortId        *string `json:"portId,omitempty"`
	PortName      *string `json:"portName,omitempty"`
	DeviceIp      *string `json:"deviceIp,omitempty"`
}

type LogicalPortAdditional struct {
	Producer *string `json:"producer,omitempty"`
	CreateAt *string `json:"createAt,omitempty"`
	UpdateAt *string `json:"updateAt,omitempty"`
}

func NewLogicalPort(id, name *string, logicalPortAttr LogicalPortAttributes) *LogicalPort {
	return &LogicalPort{
		Id:                    id,
		Name:                  name,
		LogicalPortAttributes: logicalPortAttr,
	}
}

func (resp *LogicalPortResponse) Count() int {
	return len(resp.LogicalPorts)
}
