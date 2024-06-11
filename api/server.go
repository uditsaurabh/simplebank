package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	store "github.com/uditsaurabh/go-simple-bank/orm"
)

type Server struct {
	store  store.Store
	router *gin.Engine
}

func NewServer(store store.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	if v,ok:=binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("currency",validCurrency)
	}
	router.GET("/accounts/:id", server.getAccount)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccount)
	router.GET("/transfer/:id", server.getTransfer)
	router.GET("/entry/:id", server.getEntry)
	router.POST("/transfers", server.createTransfer)
	router.POST("/users",server.createUser)
	router.POST("/login",server.Login)


	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
