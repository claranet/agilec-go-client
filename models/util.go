package models

import (
	"github.com/outscope-solutions/acdn-go-client/container"
	"strings"
)

func StripQuotes(word string) string {
	if strings.HasPrefix(word, "\"") && strings.HasSuffix(word, "\"") {
		return strings.TrimSuffix(strings.TrimPrefix(word, "\""), "\"")
	}
	return word
}

func G(cont *container.Container, key string) interface{} {
	return StripQuotes(cont.S(key).String())
}

func A(data map[string]interface{}, key, value string) {

	if value != "" {
		data[key] = value
	}

	if value == "{}" {
		data[key] = ""
	}
}

func toStringMap(intf interface{}) map[string]interface{} {

	result := make(map[string]interface{})
	temp := intf.(map[string]interface{})

	for key, value := range temp {
		A(result, key, value.(string))
	}

	return result
}
