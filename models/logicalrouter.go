package models

const LogicalRouterModuleName = "router"

type LogicalRouterRequestOpts struct {
	Producer       string `url:"producer,omitempty"`
	Type           string `url:"type,omitempty"`
	TenantId       string `url:"tenantId,omitempty"`
	LogicNetworkId string `url:"logicNetworkId,omitempty"`
	FabricIds      string `url:"fabricIds,omitempty"`
	TransitConnect string `url:"transitConnect,omitempty"`
	BaseRequestOpts
}

type LogicalRouterListResponse struct {
	LogicalRouterList
}

type LogicalRouterList struct {
	LogicalRouters []*LogicalRouter `json:"router"`
}

type LogicalRouterResponse struct {
	LogicalRouter LogicalRouter `json:"router"`
}

type LogicalRouter struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
	LogicalRouterAttributes
}

type LogicalRouterAttributes struct {
	Description     *string                   `json:"description,omitempty"`
	LogicNetworkId  *string                   `json:"logicNetworkId,omitempty"`
	RouterLocations []*LogicalRouterLocations `json:"routerLocations,omitempty"`
	Type            *string                   `json:"type,omitempty"`
	Vni             *int32                    `json:"vni,omitempty"`
	VrfName         *string                   `json:"vrfName,omitempty"`
	Additional      *LogicalRouterAdditional  `json:"additional,omitempty"`
}

type LogicalRouterLocations struct {
	FabricId    *string                              `json:"fabricId,omitempty"`
	FabricRole  *string                              `json:"fabricRole,omitempty"`
	FabricName  *string                              `json:"fabricName,omitempty"`
	DeviceGroup []*LogicalRouterLocationsDeviceGroup `json:"deviceGroup,omitempty"`
}

type LogicalRouterLocationsDeviceGroup struct {
	DeviceId *string `json:"deviceId,omitempty"`
	DeviceIp *string `json:"deviceIp,omitempty"`
}

type LogicalRouterAdditional struct {
	Producer *string `json:"producer,omitempty"`
	CreateAt *string `json:"createAt,omitempty"`
	UpdateAt *string `json:"updateAt,omitempty"`
}

func NewLogicalRouter(id, name *string, logicalRouterAttr LogicalRouterAttributes) *LogicalRouter {
	return &LogicalRouter{
		Id:                      id,
		Name:                    name,
		LogicalRouterAttributes: logicalRouterAttr,
	}
}

func (resp *LogicalRouterListResponse) Count() int {
	return len(resp.LogicalRouters)
}
