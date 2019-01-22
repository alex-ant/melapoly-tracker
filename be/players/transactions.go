package players

import (
	"errors"
	"strings"
	"time"
)

// BankID specifies an ID reserved by the bank.
const BankID string = "BANK"

// Transaction contains single transaction data.
type Transaction struct {
	FromID string
	ToID   string
	Amount int
	TS     time.Time
}

func (p *Players) addTransaction(fromID, toID string, amount int) error {
	p.tMu.Lock()
	defer p.tMu.Unlock()

	if strings.TrimSpace(fromID) == "" {
		return errors.New("fromID is not provided")
	}

	if strings.TrimSpace(toID) == "" {
		return errors.New("toID is not provided")
	}

	p.transactions = append(p.transactions, Transaction{
		FromID: fromID,
		ToID:   toID,
		Amount: amount,
		TS:     time.Now().UTC(),
	})

	return nil
}

// GetTransactions returns all recorded transactions.
func (p *Players) GetTransactions() []Transaction {
	p.tMu.Lock()
	defer p.tMu.Unlock()

	return p.transactions
}
