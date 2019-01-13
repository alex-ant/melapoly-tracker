package api

import (
	"net/http"
)

type playerData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CashAmount int    `json:"cashAmount"`
	IsAdmin    bool   `json:"isAdmin"`
}

func (a *API) getAllPlayersHandler(w http.ResponseWriter, r *http.Request) {
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
		})
	}

	respond("players", res, "ok", http.StatusOK, w)
}
