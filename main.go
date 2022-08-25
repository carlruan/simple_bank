package main

import (
	"database/sql"
	"github.com/carlruan/simple_bank/api"
	db "github.com/carlruan/simple_bank/db/sqlc"
	"github.com/carlruan/simple_bank/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration from viper")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server on", config.ServerAddress, err)
	}
}
