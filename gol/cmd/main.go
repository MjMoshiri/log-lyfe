package main

import (
	"flag"
	"github.com/mjmoshiri/log-lyfe/gol/internal/api"
	"github.com/mjmoshiri/log-lyfe/gol/internal/api/router"
	"log"
)

func main() {
	serverConfigPath := flag.String("server", "config/server.yaml", "Path to server config file")
	dbConfigPath := flag.String("db", "config/db.yaml", "Path to db config file")
	flag.Parse()

	serverConfig, err := initServerConfig(*serverConfigPath)
	if err != nil {
		log.Fatalf("Error initializing server config: %v", err)
	}

	dbConfig, err := initDBConfig(*dbConfigPath)
	if err != nil {
		log.Fatalf("Error initializing DB config: %v", err)
	}

	db, err := initDB(dbConfig)
	if err != nil {
		log.Fatalf("Error initializing DB: %v", err)
	}
	defer db.Close()
	app := api.New(serverConfig)
	router.SetupRoutes(&app, db, serverConfig)
	app.ListenAndServe()
}
