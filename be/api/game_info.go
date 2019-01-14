package api

import (
	"net/http"
)

type getGameInfoResponse struct {
	InitialAmount int `json:"initialAmount"`
	Salary        int `json:"salary"`
}

func (a *API) getGameInfoHandler(w http.ResponseWriter, r *http.Request) {
	resp := getGameInfoResponse{
		InitialAmount: a.playersProc.InitialAmount(),
		Salary:        a.playersProc.Salary(),
	}

	respond("gameInfo", resp, "ok", http.StatusOK, w)
}
