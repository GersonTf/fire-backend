package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/GersonTf/fire-backend/api"
	"github.com/GersonTf/fire-backend/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("server running on port: ", *listenAddr)
	log.Fatal(server.Start())

}
