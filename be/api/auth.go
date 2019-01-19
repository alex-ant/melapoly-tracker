package api

import (
	"net/http"

	"github.com/go-zoo/bone"
)

type authResponse struct {
	Authenticated bool `json:"authenticated"`
	PlayerData    struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		CashAmount int    `json:"cashAmount"`
		IsAdmin    bool   `json:"isAdmin"`
		Color      string `json:"color"`
	} `json:"playerData"`
}

func (a *API) authHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve request data.
	token := bone.GetValue(r, "token")

	// Validate received token.
	if !a.playersProc.PlayerExists(token) {
		respond("auth", authResponse{
			Authenticated: false,
		}, "ok", http.StatusOK, w)
		return
	}

	playerData, playerDataErr := a.playersProc.GetPlayer(token)
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
	resp.PlayerData.Color = a.colorSet.GetColor(playerData.ID).Hex()

	respond("auth", resp, "ok", http.StatusOK, w)
}
