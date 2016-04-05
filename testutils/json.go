package testutils

import (
	"errors"
	"github.com/bcokert/render-cloud/model"
	"reflect"
	"testing"
)

func ExpectJsonEncoding(t *testing.T, object interface{}, expectedJson string) {
	output, err := model.ToJson(model.DefaultMarshaler, object)
	if err != nil {
		t.Errorf("An error occurred while trying to json encode %s, error: %s", reflect.TypeOf(object), err.Error())
	}
	if output != expectedJson {
		t.Errorf("%s json encoded to %q, expected %q", reflect.TypeOf(object), output, expectedJson)
	}
}

type BadWriter struct{}

func (*BadWriter) Write(bytes []byte) (int, error) {
	return 0, errors.New("Oh No!")
}
