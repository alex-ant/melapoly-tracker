package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

func TestGetTokenByID(t *testing.T) {
	p := New(10, 5)
	p.players["token1"] = Player{
		ID: "id1",
	}

	res, resErr := p.getTokenByID("id1")
	require.NoError(t, resErr, "Error getting player token by ID")
	require.Equal(t, "token1", res, "Invalid player token received")

	_, res2Err := p.getTokenByID("id2")
	require.Error(t, res2Err, "Error determinig non-existent player")
}

func TestValidID(t *testing.T) {
	p := New(10, 5)
	p.players["token1"] = Player{
		ID: "id1",
	}

	require.True(t, p.validID("id1"), "Failed to confirm valid player ID")
	require.False(t, p.validID("invalidid"), "Failed to confirm invalid player ID")
}

func TestIsAdmin(t *testing.T) {
	// test valid admin user
	p := New(10, 5)
	p.players["token1"] = Player{
		ID: "id1",
	}

	p.adminPlayer = "id1"

	res, resErr := p.IsAdmin("id1")
	require.NoError(t, resErr, "Error determinig whether a user is admin")
	require.True(t, res, "Failed to confirm admin player")

	// test user which is not the admin
	p.players["token2"] = Player{
		ID: "id2",
	}

	res2, res2Err := p.IsAdmin("id2")
	require.NoError(t, res2Err, "Error determinig whether a user is admin")
	require.False(t, res2, "Failed to confirm non-admin player")

	// test non-existent user
	_, res3Err := p.IsAdmin("id3")
	require.Error(t, res3Err, "Error determinig non-existing user")
}

func TestPlayerExists(t *testing.T) {
	p := New(10, 5)
	p.players["token1"] = Player{}

	require.True(t, p.PlayerExists("token1"), "Invalid player exists determination")
	require.False(t, p.PlayerExists("token2"), "Invalid player exists determination")

	require.Equal(t, 10, p.initialAmount, "Invalid inital amount has been set to players processor")
	require.Equal(t, 5, p.salary, "Invalid salary has been set to players processor")
}

func TestAddPlayer(t *testing.T) {
	p := New(10, 5)

	token, err := p.AddPlayer("John")
	require.NoError(t, err, "Failed to add a new player")
	require.True(t, p.PlayerExists(token), "A player hasn't been added")
	require.Equal(t, 10, p.players[token].CashAmount, "Invalid initial amount has been assigned to a new player")
	require.Equal(t, p.adminPlayer, p.players[token].ID, "Failed to assign admin player")
}

func TestGetPlayer(t *testing.T) {
	p := New(10, 5)

	token, err := p.AddPlayer("John")
	require.NoError(t, err, "Failed to add a new player")

	player, playerErr := p.GetPlayer(token)
	require.NoError(t, playerErr, "Failed to get player")

	require.Equal(t, Player{
		Name:       "John",
		CashAmount: 10,
		ID:         p.players[token].ID,
	}, player, "Invalid player data received")

	_, expectedErr := p.GetPlayer("invalid-token")
	require.Error(t, expectedErr, "No error on non-existent player data request")
}
