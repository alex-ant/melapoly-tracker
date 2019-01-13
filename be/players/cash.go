package players

import "errors"

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

// AddSalary adds salary to the specified player.
func (p *Players) AddSalary(id string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	token, tokenErr := p.getTokenByID(id)
	if tokenErr != nil {
		return tokenErr
	}

	player := p.players[token]
	player.CashAmount += p.salary
	p.players[token] = player

	return nil
}

// DeductCash deducts the specified amount of cash from the player.
func (p *Players) DeductCash(id string, amount int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	token, tokenErr := p.getTokenByID(id)
	if tokenErr != nil {
		return tokenErr
	}

	player := p.players[token]

	// Return an error if the player has insufficient amount of cash available.
	if player.CashAmount < amount {
		return errors.New("insufficient amount of cash available")
	}

	player.CashAmount -= amount
	p.players[token] = player

	return nil
}

// TransferCash sends cash from one player to another. An error is thrown if
// the sender has insufficient amount of cash available.
func (p *Players) TransferCash(fromID, toID string, amount int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	tokenFrom, tokenFromErr := p.getTokenByID(fromID)
	if tokenFromErr != nil {
		return tokenFromErr
	}

	tokenTo, tokenToErr := p.getTokenByID(toID)
	if tokenToErr != nil {
		return tokenToErr
	}

	playerFrom := p.players[tokenFrom]
	playerTo := p.players[tokenTo]

	// Return an error if the player has insufficient amount of cash available.
	if playerFrom.CashAmount < amount {
		return errors.New("insufficient amount of cash available")
	}

	// Deduct cash from sender.
	playerFrom.CashAmount -= amount
	p.players[tokenFrom] = playerFrom

	// Add cash to receiver.
	playerTo.CashAmount += amount
	p.players[tokenTo] = playerTo

	return nil
}
