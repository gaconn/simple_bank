package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/quan12xz/simple_bank/api"
	db "github.com/quan12xz/simple_bank/db/sqlc"
	"github.com/quan12xz/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to load config file: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to db: ", err)
		return
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.APIAddress)

	if err != nil {
		log.Fatal("Cannot connect to server: ", err)
		return
	}
}
