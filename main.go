package main

import (
	"database/sql"
	_ "fmt"
	"log"
	_ "os"
	_ "testing"

	_ "github.com/lib/pq"
	api "github.com/uditsaurabh/go-simple-bank/api"
	db "github.com/uditsaurabh/go-simple-bank/orm"
	"github.com/uditsaurabh/go-simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	log.Println("setting up test database...")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
