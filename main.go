package main

import (
	"database/sql"
	"log"

	"github.com/fredele20/Golang-backend-master/api"
	db "github.com/fredele20/Golang-backend-master/db/sqlc"
	"github.com/fredele20/Golang-backend-master/util"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable"
)

func main() {
	config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	defer conn.Close()

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
