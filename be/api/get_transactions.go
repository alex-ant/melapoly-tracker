package api

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/alex-ant/melapoly-tracker/be/players"
)

type transactionData struct {
	FromID   string    `json:"fromID"`
	FromName string    `json:"fromName"`
	ToID     string    `json:"toID"`
	ToName   string    `json:"toName"`
	Amount   int       `json:"amount"`
	TS       time.Time `json:"ts"`
}

func (a *API) getTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var res []transactionData

	// nameCache is used to cache names per ID in the format of map[<id>]<name>
	nameCache := map[string]string{
		players.BankID: players.BankID,
	}

	getName := func(id string) (string, error) {
		// Attempt to fetch the name from the cache first.
		fromName, fromNameOK := nameCache[id]
		if !fromNameOK {
			if id != players.BankID {
				// Fetch player name by ID and add to cache.
				player, playerErr := a.playersProc.GetPlayerByID(id)
				if playerErr != nil {
					return "", fmt.Errorf("failed to get player data: %v", playerErr)
				}

				nameCache[id] = player.Name

				fromName = player.Name
			} else {
				fromName = players.BankID
			}
		}

		return fromName, nil
	}

	for _, trans := range a.playersProc.GetTransactions() {
		fromName, fromNameErr := getName(trans.FromID)
		if fromNameErr != nil {
			respond("", nil, fromNameErr.Error(), http.StatusInternalServerError, w)
			return
		}

		toName, toNameErr := getName(trans.ToID)
		if toNameErr != nil {
			respond("", nil, toNameErr.Error(), http.StatusInternalServerError, w)
			return
		}

		res = append(res, transactionData{
			FromID:   trans.FromID,
			FromName: fromName,
			ToID:     trans.ToID,
			ToName:   toName,
			Amount:   trans.Amount,
			TS:       trans.TS,
		})
	}

	// Show the recent transactions first.
	sort.Slice(res, func(i, j int) bool {
		return res[i].TS.After(res[j].TS)
	})

	respond("transactions", res, "ok", http.StatusOK, w)
}
