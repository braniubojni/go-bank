package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) (Entries, Accounts) {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return entry, account
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1, _ := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
}

func TestUpdateEntry(t *testing.T) {
	entry1, _ := createRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.NotEqual(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
}

func TestDeleteEntry(t *testing.T) {
	entry1, _ := createRandomEntry(t)

	deletedEntry, err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedEntry)

	require.Equal(t, entry1.Amount, deletedEntry.Amount)
	require.Equal(t, entry1.ID, deletedEntry.ID)
	require.Equal(t, entry1.AccountID, deletedEntry.AccountID)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {
	for range 10 {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
