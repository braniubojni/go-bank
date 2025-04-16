package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/simplebank/api"
	"github.com/simplebank/config"
	db "github.com/simplebank/db/sqlc"
)

func main() {
	cfg, err := config.LoadConfig(".")
	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
