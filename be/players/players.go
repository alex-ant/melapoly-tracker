package players

import (
	"errors"
	"sort"
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

// InitialAmount returns the set initial amount.
func (p *Players) InitialAmount() int {
	return p.initialAmount
}

// Salary returns the set salary.
func (p *Players) Salary() int {
	return p.salary
}

// GetAllIDs returns ID of all the players.
func (p *Players) GetAllIDs() []string {
	var res []string

	p.mu.Lock()
	defer p.mu.Unlock()

	for _, player := range p.players {
		res = append(res, player.ID)
	}

	sort.Strings(res)

	return res
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

// RemovePlayer deletes a user identified by the provided token and, if admin,
// assigns a new random admin.
func (p *Players) RemovePlayer(token string) error {
	player, playerErr := p.GetPlayer(token)
	if playerErr != nil {
		return playerErr
	}

	isAdmin, isAdminErr := p.IsAdmin(player.ID)
	if isAdminErr != nil {
		return isAdminErr
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	delete(p.players, token)

	if isAdmin {
		p.adminPlayer = ""

		for _, player := range p.players {
			p.adminPlayer = player.ID
			break
		}
	}

	return nil
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

// GetPlayerByID returns player data identified by the passed ID.
func (p *Players) GetPlayerByID(id string) (Player, error) {
	token, tokenErr := p.getTokenByID(id)
	if tokenErr != nil {
		return Player{}, tokenErr
	}

	return p.GetPlayer(token)
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
