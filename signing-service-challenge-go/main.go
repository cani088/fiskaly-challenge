package main

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/api"
	"log"
)

const (
	ListenAddress = ":8080"
	Storage       = "memory"
	// TODO: add further configuration parameters here ...
)

func main() {
	server := api.NewServer(ListenAddress, Storage)

	if err := server.Run(); err != nil {
		log.Fatal("Could not start server on ", ListenAddress)
	}
}
