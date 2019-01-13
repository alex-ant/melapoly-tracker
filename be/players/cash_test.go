package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferCash(t *testing.T) {
	p := New(10, 5)
	p.players["token1"] = Player{
		ID:         "id1",
		CashAmount: 100,
	}
	p.players["token2"] = Player{
		ID:         "id2",
		CashAmount: 100,
	}

	err := p.TransferCash("id1", "id2", 20)
	require.NoError(t, err, "Failed to transfer cash")
	require.Equal(t, 80, p.players["token1"].CashAmount, "Invalid cash amount has been transfered")
	require.Equal(t, 120, p.players["token2"].CashAmount, "Invalid cash amount has been transfered")

	err2 := p.TransferCash("id1", "id2", 101)
	require.Error(t, err2, "Transfered cash while the amount was insufficient")

	err3 := p.TransferCash("id3", "id2", 101)
	require.Error(t, err3, "Error determinig non-existent player")

	err4 := p.TransferCash("id1", "id3", 101)
	require.Error(t, err4, "Error determinig non-existent player")
}

func TestDeductCash(t *testing.T) {
	p := New(10, 5)
	p.players["token1"] = Player{
		ID:         "id1",
		CashAmount: 100,
	}

	err := p.DeductCash("id1", 25)
	require.NoError(t, err, "Failed to deduct cash")
	require.Equal(t, 75, p.players["token1"].CashAmount, "Invalid cash amount has been deducted")

	err2 := p.DeductCash("id1", 76)
	require.Error(t, err2, "Deducted too much cash")

	err3 := p.DeductCash("id2", 76)
	require.Error(t, err3, "Error determinig non-existent player")
}

func TestAddSalary(t *testing.T) {
	p := New(10, 5)
	p.players["token1"] = Player{
		ID:         "id1",
		CashAmount: 100,
	}

	err := p.AddSalary("id1")
	require.NoError(t, err, "Failed to add salary")
	require.Equal(t, 105, p.players["token1"].CashAmount, "Invalid salary amount has been added")

	err2 := p.AddSalary("id2")
	require.Error(t, err2, "Error determinig non-existent player")
}

func TestAddCash(t *testing.T) {
	p := New(10, 5)
	p.players["token1"] = Player{
		ID:         "id1",
		CashAmount: 100,
	}

	err := p.AddCash("id1", 30)
	require.NoError(t, err, "Failed to add cash")
	require.Equal(t, 130, p.players["token1"].CashAmount, "Invalid cash amount has been added")

	err2 := p.AddCash("id2", 30)
	require.Error(t, err2, "Error determinig non-existent player")
}
