package db

import (
	"context"
	"golang-backend-structure/util"
	"testing"

	"github.com/stretchr/testify/require"
)


func createRandomTransfer(t *testing.T) Transfer {
	args := CreateTransferParams{
		FromAccountID: util.RandomInt(1, 1000),
		ToAccountID: util.RandomInt(1, 1000),
		Amount: util.RandomMoney(),
	}

	transfer, _ := testQueries.CreateTransfer(context.Background(), args)
	// require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}
