package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer() (CreateTransferParams, Transfer, error) {
	_, account_1, _ := createRandomAccount()
	_, account_2, _ := createRandomAccount()
	arg := CreateTransferParams{account_1.ID, account_2.ID, RandomAmountTransfer()}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	return arg, transfer, err
}
func TestCreateTransfer(t *testing.T) {
	arg, transfer, err := createRandomTransfer()

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.Amount, arg.Amount)
	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

func getTransferData(id int64) (Transfer, error) {
	return testQueries.GetTransfer(context.Background(), id)
}
func TestGetTransfer(t *testing.T) {
	_, transfer, _ := createRandomTransfer()
	transferGet, err := getTransferData(transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transferGet)
	require.Equal(t, transferGet.Amount, transfer.Amount)
	require.Equal(t, transferGet.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transferGet.ToAccountID, transfer.ToAccountID)
	require.NotZero(t, transferGet.ID)
	require.NotZero(t, transferGet.CreatedAt)
}

func TestUpdateTransfer(t *testing.T) {
	_, transfer, _ := createRandomTransfer()
	arg := UpdateTransferParams{transfer.ID, RandomAmountTransfer()}
	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTransfer)
	require.Equal(t, updatedTransfer.Amount, arg.Amount)
	require.NotZero(t, updatedTransfer.FromAccountID)
	require.NotZero(t, updatedTransfer.ToAccountID)
	require.NotZero(t, updatedTransfer.ID)
	require.NotZero(t, updatedTransfer.CreatedAt)
}

func TestDeleteTransfer(t *testing.T) {
	_, transfer, _ := createRandomTransfer()
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	deletedTransfer, err := getTransferData(transfer.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedTransfer)
}

func TestGetListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer()
	}
	arg := ListTransferParams{5, 5}
	listTransfer, err := testQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listTransfer, int(arg.Limit))
}
