package models

const ExternalGatewayModuleName = "external-gateway"

type ExternalGatewayRequestOpts struct {
	BaseRequestOpts
}

type ExternalGatewayResponse struct {
	ExternalGatewayList
	TotalNum int32 `json:"totalNum,omitempty"`
}

type ExternalGatewayList struct {
	ExternalGateways []*ExternalGateway `json:"externalGateway"`
}

type ExternalGateway struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	ExternalGatewayAttributes
}

type ExternalGatewayAttributes struct {
	Description    string   `json:"description,omitempty"`
	GatewayType    string   `json:"gatewayType,omitempty"`
	PublicIpPool   []string `json:"publicIppool,omitempty"`
	ServiceIpPool  []string `json:"serviceIppool,omitempty"`
	IsTelcoGateway bool     `json:"isTelcoGateway,omitempty"`
	VrfName        string   `json:"vrfName,omitempty"`
	IsAllShared    bool     `json:"isAllShared,omitempty"`
}

func (resp *ExternalGatewayResponse) Count() int {
	return len(resp.ExternalGateways)
}
