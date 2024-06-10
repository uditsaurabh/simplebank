package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uditsaurabh/go-simple-bank/util"
)

func TestCreateAccount(t *testing.T) {
	account, err, args := createRandomAccount()
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, args.Owner, account.Owner)
}

func createRandomAccount() (Account, error, CreateAccountParams) {
	args := CreateAccountParams{
		Owner:       util.RandomOwner(),
		Balance:     util.RandomMoney(),
		Currency:    util.RandomCurrencies(),
		CountryCode: util.RandomCountryCode(),
	}
	account, error := testQueries.CreateAccount(context.Background(), args)
	return account, error, args
}

func TestGetAccount(t *testing.T) {
	account1, err, _ := createRandomAccount()
	require.NoError(t, err)
	savedAccountFromDb, err := testQueries.GetAccount(context.Background(), account1.ID)
	fmt.Println(savedAccountFromDb, "<---------->", account1)
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	require.Equal(t, account1.Currency, savedAccountFromDb.Currency)
	require.Equal(t, account1.Owner, savedAccountFromDb.Owner)
	require.Equal(t, account1.Balance, savedAccountFromDb.Balance)
	require.Equal(t, account1.CountryCode, savedAccountFromDb.CountryCode)
}

func TestUpdateAccount(t *testing.T) {
	account1, err, _ := createRandomAccount()
	require.NoError(t, err)
	testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		ID:          account1.ID,
		Balance:     int64(1000),
		Currency:    account1.Currency,
		CountryCode: account1.CountryCode,
		Owner:       account1.Owner,
	})
	savedAccountFromDb, err := testQueries.GetAccount(context.Background(), account1.ID)
	fmt.Println(savedAccountFromDb, "<---------->", account1)
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	require.Equal(t, account1.Currency, savedAccountFromDb.Currency)
	require.Equal(t, account1.Owner, savedAccountFromDb.Owner)
	require.Equal(t, int64(1000), int64(savedAccountFromDb.Balance))
	require.Equal(t, account1.CountryCode, savedAccountFromDb.CountryCode)
}

func TestDeleteAccount(t *testing.T) {
	account1, err, _ := createRandomAccount()
	fmt.Println("the account created with details -->", account1)
	require.NoError(t, err)
	testQueries.DeleteAccount(context.Background(), account1.ID)
	_, err = testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
}