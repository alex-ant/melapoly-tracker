package players

// AddCash adds passed cash amount to the specified player.
func (p *Players) AddCash(id string, amount int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	token, tokenErr := p.getTokenByID(id)
	if tokenErr != nil {
		return tokenErr
	}

	player := p.players[token]
	player.CashAmount += amount
	p.players[token] = player

	return nil
}
