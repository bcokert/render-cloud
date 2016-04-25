package controller

import (
	"encoding/json"
	"fmt"
	"github.com/bcokert/render-cloud/dto/response"
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/utils"
	"net/http"
	"github.com/bcokert/render-cloud/raytracer"
	"github.com/bcokert/render-cloud/raytracer/illumination"
	"github.com/bcokert/render-cloud/image"
	"github.com/bcokert/render-cloud/validation"
)

func badRequest(responseWriter http.ResponseWriter, output response.ErrorResponse) {
	responseWriter.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(responseWriter).Encode(output); err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(responseWriter, "{\"message\": \"Failed to encode error response: %q %d\"}", *output.Message, *output.Code)
	}
}

func okRequest(responseWriter http.ResponseWriter, output interface{}) {
	responseWriter.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(responseWriter).Encode(output); err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(responseWriter, "\"message\": \"Failed to encode response %q\"}", output)
	}
}

func PostRender(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		badRequest(responseWriter, response.ErrorResponse{Message: utils.StringPointer("POST /render requires a json body containing a scene, but nothing was sent.")})
		return
	}

	var postRequest model.ScenePostRequest
	if err := json.NewDecoder(request.Body).Decode(&postRequest); err != nil {
		badRequest(responseWriter, response.ErrorResponse{Message: utils.StringPointer("POST /render received an invalid PostSceneRequest"), Reason: utils.StringPointer(err.Error())})
		return
	}

	validator := validation.NewValidator()
	var scene model.Scene
	if err := scene.FromPostRequest(validator, postRequest); err != nil {
		badRequest(responseWriter, response.ErrorResponse{Message: utils.StringPointer("Failed to decode scene from request"), Reason: utils.StringPointer(err.Error())})
		return
	}

	illuminator := illumination.PhongIlluminator{}
	colors, err := raytracer.DefaultRaytracer{}.TraceScene(scene, illuminator, 300, 300)
	if  err != nil {
		badRequest(responseWriter, response.ErrorResponse{Message: utils.StringPointer("Failed to raytrace scene"), Reason: utils.StringPointer(err.Error())})
		return
	}

	pngWriter := image.PNGImageWriter{}
	err = pngWriter.WriteImage(image.DefaultPNGEncoder, colors, 300, 300, "testout.png")
	if err != nil {
		badRequest(responseWriter, response.ErrorResponse{Message: utils.StringPointer("Failed to write file with raytracer output"), Reason: utils.StringPointer(err.Error())})
		return
	}

	okRequest(responseWriter, scene)
}
