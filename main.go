package main

import (
	"SimpleBank/api"
	db "SimpleBank/db/sqlc"
	"SimpleBank/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
