package players

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (p *Players) validID(id string) bool {
	for _, player := range p.players {
		if player.ID == id {
			return true
		}
	}

	return false
}

func (p *Players) getTokenByID(id string) (string, error) {
	for token, player := range p.players {
		if player.ID == id {
			return token, nil
		}
	}

	return "", errors.New("Invalid player ID")
}
