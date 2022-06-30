package models

const LogicalSwitchModuleName = "switch"

type LogicalSwitchRequestOpts struct {
	Producer        string `url:"producer,omitempty"`
	TenantId        string `url:"tenantId,omitempty"`
	LogicNetworkId  string `url:"logicNetworkId,omitempty"`
	LogicSwitchName string `url:"logicSwitchName,omitempty"`
	BaseRequestOpts
}

type LogicalSwitchResponse struct {
	LogicalSwitchList
}

type LogicalSwitchList struct {
	LogicalSwitches []*LogicalSwitch `json:"switch"`
}

type LogicalSwitch struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
	LogicalSwitchAttributes
}

type LogicalSwitchAttributes struct {
	Description    *string                     `json:"description,omitempty"`
	LogicNetworkId *string                     `json:"logicNetworkId,omitempty"`
	Vni            *int32                      `json:"vni,omitempty"`
	Bd             *int32                      `json:"bd,omitempty"`
	MacAddress     *string                     `json:"macAddress,omitempty"`
	TenantId       *string                     `json:"tenantId,omitempty"`
	StormSuppress  *LogicalSwitchStormSuppress `json:"stormSuppress,omitempty"`
	Additional     *LogicalSwitchAdditional    `json:"additional,omitempty"`
}

type LogicalSwitchStormSuppress struct {
	BroadcastEnable  *bool   `json:"broadcastEnable,omitempty"`
	MulticastEnable  *bool   `json:"multicastEnable,omitempty"`
	UnicastEnable    *bool   `json:"unicastEnable,omitempty"`
	BroadcastCbs     *string `json:"broadcastCbs,omitempty"`
	BroadcastCbsUnit *string `json:"broadcastCbsUnit,omitempty"`
	BroadcastCir     *int64  `json:"broadcastCir,omitempty"`
	BroadcastCirUnit *string `json:"broadcastCirUnit,omitempty"`
	UnicastCbs       *string `json:"unicastCbs,omitempty"`
	UnicastCbsUnit   *string `json:"unicastCbsUnit,omitempty"`
	UnicastCir       *int64  `json:"unicastCir,omitempty"`
	UnicastCirUnit   *string `json:"unicastCirUnit,omitempty"`
	MulticastCbs     *string `json:"multicastCbs,omitempty"`
	MulticastCbsUnit *string `json:"multicastCbsUnit,omitempty"`
	MulticastCir     *int64  `json:"multicastCir,omitempty"`
	MulticastCirUnit *string `json:"multicastCirUnit,omitempty"`
}

type LogicalSwitchAdditional struct {
	Producer *string `json:"producer,omitempty"`
	CreateAt *string `json:"createAt,omitempty"`
	UpdateAt *string `json:"updateAt,omitempty"`
}

func NewLogicalSwitch(id, name *string, logicalSwitchAttr LogicalSwitchAttributes) *LogicalSwitch {
	return &LogicalSwitch{
		Id:                      id,
		Name:                    name,
		LogicalSwitchAttributes: logicalSwitchAttr,
	}
}

func (resp *LogicalSwitchResponse) Count() int {
	return len(resp.LogicalSwitches)
}
