package api

import (
	"net/http"
	"strings"
)

func (a *API) removePlayerHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve token.
	token := r.Header.Get("X-Token")
	if strings.TrimSpace(token) == "" {
		respond("auth", nil, "X-Token header is required", http.StatusBadRequest, w)
		return
	}

	removeErr := a.playersProc.RemovePlayer(token)
	if removeErr != nil {
		respond("auth", nil, "failed to remove player: "+removeErr.Error(), http.StatusBadRequest, w)
		return
	}

	a.publishUpdatePlayers()

	respond("", nil, "ok", http.StatusOK, w)
}
