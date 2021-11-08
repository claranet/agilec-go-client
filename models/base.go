package models

type BaseResponseAttributes struct {
	TotalNum  int64 `json:"totalNum,omitempty"`
	PageIndex int32 `json:"pageIndex,omitempty"`
	PageSize  int32 `json:"pageSize,omitempty"`
}

type BaseRequestOpts struct {
	PageIndex int32 `json:"pageIndex,omitempty"`
	PageSize  int32 `json:"pageSize,omitempty"`
}