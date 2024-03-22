package main

import (
	"github.com/nattapat-w/chatapp/config"
	"github.com/nattapat-w/chatapp/database"
	"github.com/nattapat-w/chatapp/server"
)

func main() {
	cfg := config.GetConfig()

	// db := database.NewMySQLDatabase()
	db := database.NewPostgresDatabase(&cfg)

	server.NewFiberServer(&cfg, db.GetDb()).Start()
	// server.NewEchoServer(&cfg, db.GetDb()).Start()
}
