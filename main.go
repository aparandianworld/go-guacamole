package main

import (
	"log"

	"github.com/aparandianworld/go-guacamole/internal/database"
	"github.com/aparandianworld/go-guacamole/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to initialize database client: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
