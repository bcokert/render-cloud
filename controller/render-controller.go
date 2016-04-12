package controller

import (
	"encoding/json"
	"fmt"
	"github.com/bcokert/render-cloud/dto/response"
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/utils"
	"net/http"
	"github.com/bcokert/render-cloud/raytracer"
	"github.com/bcokert/render-cloud/raytracer/illumination/phong"
	"github.com/bcokert/render-cloud/image"
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

	var scene model.Scene
	err := model.FromJson(request.Body, &scene)
	if err != nil {
		badRequest(responseWriter, response.ErrorResponse{Message: utils.StringPointer("Failed to decode post data. Expected a Scene object: " + err.Error())})
		return
	}

	illuminator := phong.PhongIlluminator{}
	colors, err := raytracer.TraceScene(scene, illuminator, 300, 300)
	if  err != nil {
		badRequest(responseWriter, response.ErrorResponse{Message: utils.StringPointer("Failed to raytrace scene: " + err.Error())})
		return
	}

	pngWriter := image.PNGImageWriter{}
	err = pngWriter.WriteImage(image.DefaultPNGEncoder, colors, 300, 300, "testout.png")
	if err != nil {
		badRequest(responseWriter, response.ErrorResponse{Message: utils.StringPointer("Failed to write file with raytracer output: " + err.Error())})
		return
	}

	okRequest(responseWriter, scene)
}
