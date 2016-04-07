package response_test

import (
	"github.com/bcokert/render-cloud/testutils"
	"testing"
	"github.com/bcokert/render-cloud/dto/response"
	"github.com/bcokert/render-cloud/utils"
)

func TestErrorResponseJsonEncodes(t *testing.T) {
	errorResponse := response.ErrorResponse{
		Message: utils.StringPointer("Failed to do the thing"),
		Code: utils.IntPointer(293),
		Reason: utils.StringPointer("Because Reasons"),
	}

	expectedJson := "{\"message\":\"Failed to do the thing\",\"code\":293,\"reason\":\"Because Reasons\"}"

	testutils.ExpectJsonEncoding(t, &errorResponse, expectedJson)
}
