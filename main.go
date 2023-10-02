package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GersonTf/fire-backend/api"
	"github.com/GersonTf/fire-backend/config"
	"github.com/GersonTf/fire-backend/storage"
)

func main() {
	printBanner()

	//creates the listenAddr flag, so we can start the app with a custom addr
	listenAddr := flag.String("listenaddr", ":8080", "the server address")
	flag.Parse()

	fmt.Println("loading app configuration...")
	cfg := config.LoadConfig()

	fmt.Println("Starting a DB connection...")
	store, err := storage.NewMongoStorage(cfg)
	if err != nil {
		log.Fatalf("failed to create a storage connection: %v", err)
	}

	fmt.Println("connected to the DB!")
	fmt.Println("Starting server ...")
	//Start server with graceful shutdown
	server := api.NewServer(*listenAddr, store)
	fmt.Println("server running on port: ", *listenAddr)

	// Create a channel to listen for termination signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a new goroutine
	go func() {
		log.Fatal(server.Start())
	}()

	// Block until a signal is received
	<-stopChan

	fmt.Println("Exit signal received, starting cleanup")

	// Perform cleanup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 	if err := server.Shutdown(shutdownCtx); err != nil { TODO handle
	// 		log.Fatalf("failed to gracefully shutdown the server: %v", err)
	// }

	if err = store.Disconnect(ctx); err != nil {
		log.Fatalf("failed to disconnect from MongoDB222: %v", err)
	}

	fmt.Println("graceful shutdown finished correctly")
}

func printBanner() {
	fmt.Println(`
	/    // \/  __\/  __/  / \ /\/ \   /__ __\/ \/ \__/|/  _ \/__ __\/  __/
	|  __\| ||  \/||  \    | | ||| |     / \  | || |\/||| / \|  / \  |  \  
	| |   | ||    /|  /_   | \_/|| |_/\  | |  | || |  ||| |-||  | |  |  /_ 
	\_/   \_/\_/\_\\____\  \____/\____/  \_/  \_/\_/  \|\_/ \|  \_/  \____\
	`)
}
