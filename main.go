package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GersonTf/fire-backend/api"
	"github.com/GersonTf/fire-backend/config"
	"github.com/GersonTf/fire-backend/storage"
	"github.com/GersonTf/fire-backend/utils"
)

func main() {
	printBanner()

	//creates the listenAddr flag, so we can start the app with a custom addr
	listenAddr := flag.String("listenaddr", ":8080", "the server address")
	flag.Parse()

	slog.Info("loading app configuration...")
	cfg := config.LoadConfig()

	slog.Info("Starting a DB connection...")
	store, err := storage.NewMongoStorage(cfg)
	if err != nil {
		log.Fatalf("failed to create a storage connection: %v", err)
	}

	slog.Info("connected to the DB!")
	slog.Info("Starting server ...")
	//Start server with graceful shutdown
	server := api.NewServer(*listenAddr, store)
	slog.Info("Server running on", "port", *listenAddr)

	// Create a channel to listen for termination signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a new goroutine
	go func() {
		if err := server.Start(); err != nil {
			utils.LogError("Failure starting the server", err)
		}
	}()

	// Block until a signal is received
	<-stopChan

	slog.Info("Exit signal received, starting cleanup")

	// Perform cleanup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 	if err := server.Shutdown(shutdownCtx); err != nil { TODO handle
	// 		log.Fatalf("failed to gracefully shutdown the server: %v", err)
	// }

	if err = store.Disconnect(ctx); err != nil {
		utils.LogError("Failed to disconnect from store", err)
	}

	slog.Info("graceful shutdown finished correctly")
}

func printBanner() {
	slog.Info(`
	/    // \/  __\/  __/  / \ /\/ \   /__ __\/ \/ \__/|/  _ \/__ __\/  __/
	|  __\| ||  \/||  \    | | ||| |     / \  | || |\/||| / \|  / \  |  \  
	| |   | ||    /|  /_   | \_/|| |_/\  | |  | || |  ||| |-||  | |  |  /_ 
	\_/   \_/\_/\_\\____\  \____/\____/  \_/  \_/\_/  \|\_/ \|  \_/  \____\
	`)
}
