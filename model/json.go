package model

import (
	"encoding/json"
	"errors"
	"io"
)

// The default marshaler for ToJson. If you don't know what to pass it, just pass the default
var DefaultMarshaler = json.Marshal

// Coverts a struct into a json string
// Requires a function that creates a new encoder for the object. Use model.DefaultEncoder in general
func ToJson(marshal func(interface{}) ([]byte, error), object interface{}) (string, error) {
	result, err := marshal(object)
	if err != nil {
		return "", errors.New("Failed to json encode object: " + err.Error())
	}
	return string(result), nil
}

func FromJson (content io.Reader, model interface{}) error {
	decoder := json.NewDecoder(content)
	err := decoder.Decode(model)
	if err != nil {
		return errors.New("Failed to json decode input: " + err.Error())
	}
	return nil
}
