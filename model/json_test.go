package model_test

import (
	"errors"
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"testing"
	"bytes"
	"fmt"
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

func TestFromJsonSuccess(t *testing.T) {
	var testModel TestModel2
	jsonString := "{\"Int\": 66, \"Str\": \"Hello\"}"
	data := bytes.NewReader([]byte(jsonString))

	err := model.FromJson(data, &testModel)

	if err != nil {
		t.Errorf("FromJson failed to decode a json string %s", jsonString)
	}

	expectedModel := TestModel2{66, "Hello"}
	if testModel != expectedModel {
		t.Errorf("FromJson decoded a model improperly. Received %s, but expected %s", testModel, expectedModel)
	}
}

func TestFromJsonErrorDecoding(t *testing.T) {
	var testModel TestModel2
	jsonString := "\\fd89ha"
	data := bytes.NewReader([]byte(jsonString))

	err := model.FromJson(data, &testModel)

	if err == nil {
		t.Errorf("FromJson did not fail when given an illegal string")
	}

	expectedError := "Failed to json decode input: invalid character '\\\\' looking for beginning of value"
	if err.Error() != expectedError {
		t.Errorf("FromJson did not return the correct error message when decoding an illegal string. Expected %s, received %s", expectedError, err.Error())
	}
}
