package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry() (CreateEntryParams, Entry, error) {
	_, account, _ := createRandomAccount()
	arg := CreateEntryParams{
		account.ID,
		RandomAmount(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	return arg, entry, err
}
func TestCreateEntry(t *testing.T) {
	arg, entry, err := createRandomEntry()
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

func getEntryData(id int64) (Entry, error) {

	return testQueries.GetEntry(context.Background(), id)
}
func TestGetEntry(t *testing.T) {
	_, newEntry, _ := createRandomEntry()
	entry, err := getEntryData(newEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

}

func TestUpdateEntry(t *testing.T) {
	_, entry, _ := createRandomEntry()
	arg := UpdateEntryParams{
		entry.ID,
		RandomAmount(),
	}
	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)
	require.Equal(t, updatedEntry.Amount, arg.Amount)
}

func TestDeleteEntry(t *testing.T) {
	_, entry, _ := createRandomEntry()
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	deletedEntry, err := getEntryData(entry.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedEntry)
}

func TestGetListEntry(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry()
	}

	arg := ListEntriesParams{5, 5}
	listEntry, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, listEntry)
	require.Len(t, listEntry, int(arg.Limit))
}
