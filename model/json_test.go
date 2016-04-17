package model_test

import (
	"errors"
	"fmt"
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"testing"
)

type TestModel struct {
	Arr []int
	Str string
	Mod *TestModel
}

type TestModel2 struct {
	Int int
	Str string
}

func (model TestModel2) String() string {
	return fmt.Sprintf("{Int: %d, Str: %s}", model.Int, model.Str)
}

func TestToJsonSuccess(t *testing.T) {
	data := TestModel{[]int{1, 2, 3}, "asdfjk\"dsaf&f", &TestModel{[]int{}, "", nil}}
	testutils.ExpectJsonEncoding(t, data, "{\"Arr\":[1,2,3],\"Str\":\"asdfjk\\\"dsaf\\u0026f\",\"Mod\":{\"Arr\":[],\"Str\":\"\",\"Mod\":null}}")
}

func TestToJsonError(t *testing.T) {
	data := TestModel{[]int{1, 2, 3}, "asdfjk\"dsaf&f", &TestModel{[]int{}, "", nil}}

	var badMarshaler = func(object interface{}) ([]byte, error) {
		return []byte{}, errors.New("Oh teh Noes!")
	}

	output, err := model.ToJson(badMarshaler, data)

	if output != "" {
		t.Errorf("ToJson returned a non-empty json response when it failed: %q", output)
	}

	if err == nil {
		t.Errorf("ToJson did not return an error when the underlying Encode returned an error")
	}
}
