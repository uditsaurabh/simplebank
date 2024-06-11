package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/uditsaurabh/go-simple-bank/orm"
	util "github.com/uditsaurabh/go-simple-bank/util"
)

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserParams
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	hashed_password, err := util.HashPassword(req.Password)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("the hashed password as follows ---->",hashed_password)
	
	arg := db.CreateUserParams{
		Username:     req.Username,
		HashPassword: hashed_password,
		FullName:     req.FullName,
		Email:        req.Email,
	}

	result, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (server *Server) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := util.CheckPasswordHash(req.Password, user.HashPassword)
	if !res {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Password Matched")

}
