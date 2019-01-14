package api

import (
	"log"
	"time"
)

const lpUpdatePlayersCat string = "update-players"

// publishUpdatePlayers publishes the latest players update UNIX nano timestamp
// to all the connected clients.
func (a *API) publishUpdatePlayers() {
	a.lpUpdatedTS = time.Now().UTC().UnixNano()

	err := a.lpManager.Publish(lpUpdatePlayersCat, map[string]interface{}{
		"updated": a.lpUpdatedTS,
	})

	if err != nil {
		log.Println("Failed to publish update-players long polling message:", err.Error())
	}
}
