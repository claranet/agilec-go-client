package models

const EndPortModuleName = "endPort"

type EndPortRequestOpts struct {
	LogicNetworkId string `url:"logicNetworkId,omitempty"`
	tenantId       string `url:"tenantId,omitempty"`
	BaseRequestOpts
}

type EndPortListResponse struct {
	EndPortList
}

type EndPortResponse struct {
	EndPort EndPort `json:"endPort"`
}

type EndPortList struct {
	EndPorts []*EndPort `json:"endPorts"`
}

type EndPort struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
	EndPortAttributes
}

type EndPortAttributes struct {
	LogicPortId    *string   `json:"logicPortId,omitempty"`
	Description    *string   `json:"description,omitempty"`
	LogicNetworkId *string   `json:"logicNetworkId,omitempty"`
	Location       *string   `json:"location,omitempty"`
	VmName         *string   `json:"vmName,omitempty"`
	Ipv4           []*string `json:"ipv4,omitempty"`
	Ipv6           []*string `json:"ipv6,omitempty"`
}

func NewEndPort(id, name *string, endPortAttr EndPortAttributes) *EndPort {
	return &EndPort{
		Id:                id,
		Name:              name,
		EndPortAttributes: endPortAttr,
	}
}

func (resp *EndPortListResponse) Count() int {
	return len(resp.EndPorts)
}

func (resp *EndPortResponse) Count() int {
	return 1
}
