package controller

import (
	"encoding/json"
	"github.com/bcokert/render-cloud/model"
	"net/http"
	"fmt"
)

func PostRender(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var scene model.Scene
	err := decoder.Decode(&scene)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(response, "Failed to parse post data: %s", err.Error())
	}

	fmt.Fprint(response, "You sent in: ")
	if err := json.NewEncoder(response).Encode(scene); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(response, "Failed to encode response")
	} else {
		response.WriteHeader(http.StatusOK)
	}
}
