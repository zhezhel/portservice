package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v4"

	"portservice/pkg/state"
	"portservice/pkg/store"
	"portservice/usecase"
)

func main() {
	log.Printf("Start: port service")
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	config, err := LoadConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = run(ctx, config)
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("Shutdown: port service")
}

func run(ctx context.Context, config Config) error {
	file, err := os.Open(config.FileSource)
	if err != nil {
		log.Printf("Unable to open file %q: %v\n", config.FileSource, err)
		return err
	}
	defer file.Close()

	conn, err := pgx.Connect(ctx, config.PostgresDSN)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return err
	}
	defer conn.Close(ctx)

	if err := store.RunMigrations(ctx, conn); err != nil {
		return err
	}

	portStore, err := store.NewPorts(conn)
	if err != nil {
		return err
	}

	stateManager, err := state.NewManager(conn)
	if err != nil {
		return err
	}

	fileIngestor, err := usecase.NewFileIngestor(
		portStore, stateManager, config.ObjectsCountLimit,
	)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err = fileIngestor.Start(ctx, file)
		wg.Done()
	}()

	wg.Wait()

	return err
}
