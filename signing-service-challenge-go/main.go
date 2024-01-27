package main

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/api"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	"log"
)

const (
	ListenAddress = ":8080"
	storageType   = "inMemory"
	// TODO: add further configuration parameters here ...
)

type repo persistence.StorageInterface

func main() {
	repo := persistence.NewInMemoryRepository()

	server := api.NewServer(ListenAddress, repo)

	if err := server.Run(); err != nil {
		log.Fatal("Could not start server on ", ListenAddress)
	}
}
