package model

import (
	"encoding/json"
	"errors"
)

// The default marshaler for ToJson. If you don't know what to pass it, just pass the default
var DefaultMarshaler = json.Marshal

// Coverts a struct (really any interface pointer) into a json string
// Requires a function that creates a new encoder for the object. Use model.DefaultEncoder in general
func ToJson(marshal func(interface{}) ([]byte, error), object interface{}) (string, error) {
	result, err := marshal(object)
	if err != nil {
		return "", errors.New("Failed to json encode object: " + err.Error())
	}
	return string(result), nil
}
