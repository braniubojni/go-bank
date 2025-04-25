package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mockdb "github.com/simplebank/db/mock"
	db "github.com/simplebank/db/sqlc"
	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
)

func genRandomAccount() db.Accounts {
	return db.Accounts{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchToAcc(t *testing.T, body *bytes.Buffer, account db.Accounts) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Accounts
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}

func TestGetAccountAPI(t *testing.T) {
	account := genRandomAccount()

	mockStore := &mockdb.MockStore{
		GetAccountFunc: func(ctx context.Context, id int64) (db.Accounts, error) {
			require.NotZero(t, id)
			return account, nil
		},
	}

	server := NewServer(mockStore)

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/accounts/%d", account.ID)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchToAcc(t, recorder.Body, account)
}

func TestCreateAccountAPI(t *testing.T) {
	account := genRandomAccount()

	mockStore := &mockdb.MockStore{
		CreateAccountFunc: func(ctx context.Context, arg db.CreateAccountParams) (db.Accounts, error) {
			require.Equal(t, account.Owner, arg.Owner)
			require.Equal(t, account.Currency, arg.Currency)
			require.Equal(t, int64(0), arg.Balance)
			return account, nil
		},
	}

	server := NewServer(mockStore)

	recorder := httptest.NewRecorder()

	requestBody := fmt.Sprintf(`{"owner":"%s", "currency":"%s"}`, account.Owner, account.Currency)
	request, err := http.NewRequest(http.MethodPost, "/accounts", strings.NewReader(requestBody))
	require.NoError(t, err)
	request.Header.Set("Content-Type", "application/json")

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusCreated, recorder.Code)
	requireBodyMatchToAcc(t, recorder.Body, account)
}
