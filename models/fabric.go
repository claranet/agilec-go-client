package models

const FabricModuleName = "fabric"

type FabricRequestOpts struct {
	BaseRequestOpts
}

type FabricResponse struct {
	FabricList
	TotalNum string `json:"totalNum,omitempty"`
}

type FabricList struct {
	Fabrics []*Fabric `json:"fabric"`
}

type Fabric struct {
	Id                     string `json:"id"`
	Name                   string `json:"name"`
	NetworkType            string `json:"networkType"`
	PhysicalNetworkMode    string `json:"physicalNetworkMode"`
	MicroSegmentCapability bool   `json:"microSegmentCapability"`
	FabricAttributes
}

type FabricAttributes struct {
	Description             string   `json:"description,omitempty"`
	BgpEvpnEnable           bool     `json:"bgpEvpnEnable,omitempty"`
	ExtInterfaceType        string   `json:"ExtInterfaceType,omitempty"`
	ArpBroadcastSuppression bool     `json:"arpBroadcastSuppression,omitempty"`
	DciSplitGroup           string   `json:"dciSplitGroup,omitempty"`
	MulticastCapability     bool     `json:"multicastCapability,omitempty"`
	SegmentMasks            []string `json:"segmentMasks,omitempty"`
	AclModel                string   `json:"aclModel,omitempty"`
	SfcCapability           string   `json:"sfcCapability,omitempty"`
	FaultService            bool     `json:"faultService,omitempty"`
	UnderLay                string   `json:"underLay,omitempty"`
	Management              string   `json:"management,omitempty"`
}

func (resp *FabricResponse) Count() int {
	return len(resp.Fabrics)
}
