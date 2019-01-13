package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type addPlayerRequest struct {
	Name string `json:"name"`
}

type addPlayerResponse struct {
	Token string `json:"token"`
}

func (a *API) addPlayerHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve request data.
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		respond("", nil, "failed to read request body", http.StatusBadRequest, w)
		return
	}

	var data addPlayerRequest
	dataErr := json.Unmarshal(body, &data)
	if dataErr != nil {
		respond("", nil, "failed to unmarshal request body", http.StatusBadRequest, w)
		return
	}
	r.Body.Close()

	// Validate received data.
	if strings.TrimSpace(data.Name) == "" {
		respond("auth", nil, "name field is required", http.StatusBadRequest, w)
		return
	}

	token, addErr := a.playersProc.AddPlayer(data.Name)
	if addErr != nil {
		respond("auth", nil, "failed to add new player: "+addErr.Error(), http.StatusBadRequest, w)
		return
	}

	resp := addPlayerResponse{
		Token: token,
	}

	respond("player", resp, "ok", http.StatusOK, w)
}
