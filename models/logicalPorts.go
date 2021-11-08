package models

import (
	"encoding/json"
)

const LogicalPortsModuleName = "port"

type LogicalPortListOpts struct {
	Producer      string `json:"producer,omitempty"`
	LogicSwitchId string `json:"logicSwitchId,omitempty"`
	//BaseRequestParameters
}

type LogicalPortResponseBody struct {
	Port []LogicalPort `json:"port"`
	BaseResponseAttributes
}

type LogicalPort struct {
	Id            string                 `json:"id"`
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	LogicSwitchId string                 `json:"logicSwitchId,omitempty"`
	MetaData      string                 `json:"metaData,omitempty"`
	AccessInfo    *LogicalPortAccessInfo `json:"accessInfo,omitempty"`
	TenantId      string                 `json:"tenantId,omitempty"`
	Additional    *LogicalPortAddition   `json:"additional,omitempty"`
	FabricId      string                 `json:"fabricId,omitempty"`
}

type LogicalPortAccessInfo struct {
	Mode               string                     `json:"mode,omitempty"`
	Type               string                     `json:"type,omitempty"`
	Vlan               int32                      `json:"vlan,omitempty"`
	Qinq               *LogicalPortQinq           `json:"qinq,omitempty"`
	Location           []LogicalPortDevicePortDto `json:"location,omitempty"`
	SubinterfaceNumber int32                      `json:"subinterfaceNumber,omitempty"`
}

type LogicalPortAddition struct {
	Producer string `json:"producer,omitempty"`
	CreateAt string `json:"CreateAt,omitempty"`
	UpdateAt string `json:"UpdateAt,omitempty"`
}

type LogicalPortQinq struct {
	InnerVidBegin int32  `json:"innerVidBegin,omitempty"`
	InnerVidEnd   int32  `json:"innerVidEnd,omitempty"`
	OuterVidBegin int32  `json:"outerVidBegin,omitempty"`
	OuterVidEnd   int32  `json:"outerVidEnd,omitempty"`
	RewriteAction string `json:"rewriteAction,omitempty"`
}

type LogicalPortDevicePortDto struct {
	DeviceGroupId string `json:"deviceGroupId,omitempty"`
	DeviceId      string `json:"deviceId,omitempty"`
	PortId        string `json:"portId,omitempty"`
	PortName      string `json:"portName,omitempty"`
	DeviceIp      string `json:"deviceIp,omitempty"`
}

func (port *LogicalPort) ToJson() ([]byte, error) {
	return json.Marshal(&port)
}