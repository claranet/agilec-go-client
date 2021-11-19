package models

type ModelResponse interface {
	Count() int
}

type BaseRequestOpts struct {
	PageIndex int32 `url:"pageIndex,omitempty"`
	PageSize  int32 `url:"pageSize,omitempty"`
}
