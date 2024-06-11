package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uditsaurabh/go-simple-bank/orm"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency" binding:"required currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("the account ids are -->", req.FromAccountID, req.ToAccountID, req)
	arg := orm.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	if !server.validAccount(ctx, arg.FromAccountID, req.Currency) {
		ctx.JSON(http.StatusBadRequest, "Currency mismatch")
		return
	}
	if !server.validAccount(ctx, arg.ToAccountID, req.Currency) {
		ctx.JSON(http.StatusBadRequest, "Currency mismatch")
		return
	}
	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	fmt.Println("account ID------>", accountID)
	account, err := server.store.GetAccount(ctx, accountID)
	fmt.Println("account------>", account)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false

	}
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch", accountID)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}
