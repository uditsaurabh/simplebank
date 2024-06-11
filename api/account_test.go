package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/uditsaurabh/go-simple-bank/orm"
	mock_orm "github.com/uditsaurabh/go-simple-bank/orm/mock"
	"github.com/uditsaurabh/go-simple-bank/util"
)

func TestGetAccountApi(t *testing.T) {
	account := randomAccount()
	ctrl:= gomock.NewController(t)
	defer ctrl.Finish()
	store:=mock_orm.NewMockStore(ctrl)
	store.EXPECT().
		GetAccount(gomock.Any(),gomock.Eq(account.ID)).
		Times(1).
		Return(account,nil)
	
	server:=NewServer(store)
	recorder:=httptest.NewRecorder()
	url:=fmt.Sprintf("/accounts/%d",account.ID)
	request,err:=http.NewRequest(http.MethodGet,url,nil)
	require.NoError(t,err)
	server.router.ServeHTTP(recorder,request)
	require.Equal(t,http.StatusOK,recorder.Code)

}

func randomAccount() orm.Account {
	return orm.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrencies(),
	}
}
