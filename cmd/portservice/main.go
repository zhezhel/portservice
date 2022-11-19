package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"

	"portservice/pkg/state"
	"portservice/pkg/store"
	"portservice/usecase"
)

func main() {

	ctx := context.Background()

	// context sigterm, sighup

	filepath := os.Getenv("FILE_SOURCE")
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Unable to open file %q: %v\n", filepath, err)
	}
	defer file.Close()

	connStr := os.Getenv("PG_DSN")
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	if err := store.RunMigrations(ctx, conn); err != nil {
		log.Fatalln(err)
	}

	portStore, err := store.NewPorts(conn)
	if err != nil {
		log.Fatalln(err)
	}

	stateManager, err := state.NewStateManager(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fileIngestor, err := usecase.NewFileIngestor(portStore, stateManager, 20)
	if err != nil {
		log.Fatalln(err)
	}

	err = fileIngestor.Start(ctx, file)
	log.Print(err)
}
