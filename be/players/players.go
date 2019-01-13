package players

import (
	"errors"
	"sync"
)

// Player contains single player's data.
type Player struct {
	Name       string
	CashAmount int
	ID         string
}

// Players contains players data.
type Players struct {
	mu            sync.Mutex
	initialAmount int
	salary        int
	players       map[string]Player
	adminPlayer   string
}

// New returns new Players instance.
func New(initialAmount, salary int) *Players {
	return &Players{
		initialAmount: initialAmount,
		players:       make(map[string]Player),
		salary:        salary,
	}
}

// PlayerExists determines whether a player identified by the passed token
// exists.
func (p *Players) PlayerExists(token string) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, e := p.players[token]
	return e
}

// AddPlayer adds a new player with the passed parameters and returns the
// corresponding identification token.
func (p *Players) AddPlayer(name string) (string, error) {
	token, tokenErr := randomHex(16)
	if tokenErr != nil {
		return "", tokenErr
	}

	id, idErr := randomHex(16)
	if idErr != nil {
		return "", idErr
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.players[token] = Player{
		ID:         id,
		Name:       name,
		CashAmount: p.initialAmount,
	}

	if p.adminPlayer == "" {
		p.adminPlayer = id
	}

	return token, nil
}

// GetPlayer returns player data identified by the passed token.
func (p *Players) GetPlayer(token string) (Player, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	player, e := p.players[token]
	if !e {
		return Player{}, errors.New("Invalid token provided")
	}

	return player, nil
}

// IsAdmin tells whether a player identified by the passed ID is admin.
func (p *Players) IsAdmin(id string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.validID(id) {
		return false, errors.New("Invalid player ID provided")
	}

	return p.adminPlayer == id, nil
}
