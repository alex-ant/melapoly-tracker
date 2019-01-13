package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type authRequest struct {
	Token string `json:"token"`
}

type authResponse struct {
	Authenticated bool `json:"authenticated"`
	PlayerData    struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		CashAmount int    `json:"cashAmount"`
		IsAdmin    bool   `json:"isAdmin"`
	} `json:"playerData"`
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

	// Validate received token.
	if !a.playersProc.PlayerExists(data.Token) {
		respond("auth", authResponse{
			Authenticated: false,
		}, "ok", http.StatusOK, w)
		return
	}

	playerData, playerDataErr := a.playersProc.GetPlayer(data.Token)
	if playerDataErr != nil {
		respond("auth", nil, "failed to get player data: "+playerDataErr.Error(), http.StatusBadRequest, w)
		return
	}

	isAdmin, isAdminErr := a.playersProc.IsAdmin(playerData.ID)
	if isAdminErr != nil {
		respond("auth", nil, "failed to determine if the player is admin: "+isAdminErr.Error(), http.StatusInternalServerError, w)
		return
	}

	resp := authResponse{
		Authenticated: true,
	}

	resp.PlayerData.ID = playerData.ID
	resp.PlayerData.Name = playerData.Name
	resp.PlayerData.CashAmount = playerData.CashAmount
	resp.PlayerData.IsAdmin = isAdmin

	respond("auth", resp, "ok", http.StatusOK, w)
}
