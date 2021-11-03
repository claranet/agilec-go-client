package models

import (
	"encoding/json"
	"github.com/outscope-solutions/acdn-go-client/container"
)

type BaseAttributes struct {
	Id        string `json:"id"`
	ClassName string `json:"-"`
}

func (ba *BaseAttributes) ToJson() (string, error) {
	data, err := json.Marshal(ba)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (ba *BaseAttributes) ToMap() (map[string]interface{}, error) {

	jsonData, err := ba.ToJson()
	if err != nil {
		return nil, err
	}
	cont, err := container.ParseJSON([]byte(jsonData))
	if err != nil {
		return nil, err
	}
	cont.Set(ba.ClassName, "classname")
	if err != nil {
		return nil, err
	}
	return toStringMap(cont.Data()), nil
}
