package api

import (
	"net/http"
	"strings"
)

type playerData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CashAmount int    `json:"cashAmount"`
	IsAdmin    bool   `json:"isAdmin"`
	You        bool   `json:"you"`
	Color      string `json:"color"`
}

func (a *API) getAllPlayersHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve token.
	token := r.Header.Get("X-Token")
	if strings.TrimSpace(token) == "" {
		respond("auth", nil, "X-Token header is required", http.StatusBadRequest, w)
		return
	}

	// Get requester's info.
	currentPlayer, currentPlayerErr := a.playersProc.GetPlayer(token)
	if currentPlayerErr != nil {
		respond("auth", nil, "failed to get player data: "+currentPlayerErr.Error(), http.StatusBadRequest, w)
		return
	}

	var res []playerData

	for _, id := range a.playersProc.GetAllIDs() {
		player, playerErr := a.playersProc.GetPlayerByID(id)
		if playerErr != nil {
			respond("auth", nil, "failed to get player data: "+playerErr.Error(), http.StatusInternalServerError, w)
			return
		}

		isAdmin, isAdminErr := a.playersProc.IsAdmin(id)
		if isAdminErr != nil {
			respond("", nil, "failed to determine if the requester is admin: "+isAdminErr.Error(), http.StatusInternalServerError, w)
			return
		}

		res = append(res, playerData{
			ID:         player.ID,
			Name:       player.Name,
			CashAmount: player.CashAmount,
			IsAdmin:    isAdmin,
			You:        player.ID == currentPlayer.ID,
			Color:      a.colorSet.GetColor(player.ID).Hex(),
		})
	}

	respond("players", res, "ok", http.StatusOK, w)
}
