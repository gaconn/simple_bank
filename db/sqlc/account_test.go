package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount() (CreateAccountParams, Account, error) {
	arg := CreateAccountParams{
		RandomOwner(),
		RandomBalance(),
		RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	return arg, account, err
}
func TestCreateAccount(t *testing.T) {

	arg, account, err := createRandomAccount()
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	_, newAccount, _ := createRandomAccount()
	account, err := getAccountData(newAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Balance, account.Balance)
	require.Equal(t, newAccount.Currency, account.Currency)
	require.Equal(t, newAccount.ID, account.ID)
	require.Equal(t, newAccount.CreatedAt, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	_, account, _ := createRandomAccount()
	arg := UpdateAccountParams{
		account.ID,
		RandomBalance(),
	}
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, updatedAccount.Balance, arg.Balance)
	require.Equal(t, updatedAccount.ID, arg.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, account.CreatedAt, updatedAccount.CreatedAt)
}

func getAccountData(id int64) (Account, error) {
	return testQueries.GetAccount(context.Background(), id)
}

func TestDeleteAccount(t *testing.T) {
	_, account, _ := createRandomAccount()
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	deletedAccount, _ := getAccountData(account.ID)
	require.NoError(t, err)
	require.Empty(t, deletedAccount)
}

func TestGetListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount()
	}

	arg := ListAccountsParams{
		5, 5,
	}
	listAccount, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listAccount, int(arg.Limit))
}
