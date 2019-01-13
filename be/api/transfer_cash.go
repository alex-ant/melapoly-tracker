package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type transferCashRequest struct {
	Token  string `json:"token"`
	IDFrom string `json:"idFrom"`
	IDTo   string `json:"idTo"`
	Amount int    `json:"amount"`
}

func (a *API) transferCashHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve request data.
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		respond("", nil, "failed to read request body", http.StatusBadRequest, w)
		return
	}

	var data transferCashRequest
	dataErr := json.Unmarshal(body, &data)
	if dataErr != nil {
		respond("", nil, "failed to unmarshal request body", http.StatusBadRequest, w)
		return
	}
	r.Body.Close()

	// Check if the requester is admin.
	reqPlayer, reqPlayerErr := a.playersProc.GetPlayer(data.Token)
	if reqPlayerErr != nil {
		respond("", nil, "failed to get requester data: "+reqPlayerErr.Error(), http.StatusBadRequest, w)
		return
	}

	isAdmin, isAdminErr := a.playersProc.IsAdmin(reqPlayer.ID)
	if isAdminErr != nil {
		respond("", nil, "failed to determine if the requester is admin: "+isAdminErr.Error(), http.StatusInternalServerError, w)
		return
	}

	if !isAdmin {
		respond("", nil, "only admin can perform this action", http.StatusBadRequest, w)
		return
	}

	transferCashErr := a.playersProc.TransferCash(data.IDFrom, data.IDTo, data.Amount)
	if transferCashErr != nil {
		respond("", nil, "failed to transfer cash: "+transferCashErr.Error(), http.StatusBadRequest, w)
		return
	}

	respond("", nil, "ok", http.StatusOK, w)
}
