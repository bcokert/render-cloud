package testutils

import (
	"reflect"
	"testing"
	"encoding/json"
)

func ExpectJsonEncoding(t *testing.T, object interface{}, expectedJson string) {
	output, err := json.Marshal(object)
	if err != nil {
		t.Errorf("An error occurred while trying to json encode %s, error: %s", reflect.TypeOf(object), err.Error())
	}
	if string(output) != expectedJson {
		t.Errorf("%s json encoded to %q, expected %q", reflect.TypeOf(object), output, expectedJson)
	}
}

type BadWriter struct{}
