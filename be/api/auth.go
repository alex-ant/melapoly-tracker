package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type authRequest struct {
	Token string `json:"token"`
}

type authResponse struct {
	ValidToken string `json:"validToken"`
}

func (a *API) authHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve request data.
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		respond("", nil, "failed to read request body", http.StatusBadRequest, w)
		return
	}

	var data authRequest
	dataErr := json.Unmarshal(body, &data)
	if dataErr != nil {
		respond("", nil, "failed to unmarshal request body", http.StatusBadRequest, w)
		return
	}
	r.Body.Close()

	fmt.Println("request data:", data)

	// Assemble response.
	resp := authResponse{
		ValidToken: "abcd",
	}

	respond("results", resp, "ok", http.StatusOK, w)
}
