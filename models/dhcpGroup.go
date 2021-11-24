package models

const DHCPGroupModuleName = "dhcpgroup"

type DHCPGroupRequestOpts struct {
	BaseRequestOpts
}

type DHCPGroupResponse struct {
	DHCPGroupList
}

type DHCPGroupList struct {
	DHCPGroups []DHCPGroup `json:"dhcpgroup"`
}

type DHCPGroup struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	ServerIp    []string `json:"serverIp"`
	DHCPGroupAttributes
}

type DHCPGroupAttributes struct {
	LogicRouterId  string          `json:"logicRouterId,omitempty"`
	VrfName        string          `json:"vrfName,omitempty"`
	Producer       string          `json:"producer,omitempty"`
	Dhcpgroupl2vni *DHCPGroupL2Vni `json:"dhcpgroupl2vni,omitempty"`
}

type DHCPGroupL2Vni struct {
	L2Vni    int32  `json:"l2Vni,omitempty"`
	Ipv4Cidr string `json:"ipv4Cidr,omitempty"`
	Ipv6Cidr string `json:"ipv6Cidr,omitempty"`
}

func (resp *DHCPGroupResponse) Count() int {
	return len(resp.DHCPGroups)
}
