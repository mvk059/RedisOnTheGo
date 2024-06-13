package main

import (
	"flag"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/server"
	"github.com/codecrafters-io/redis-starter-go/app/settings"
	"net"
)

var (
	port      = flag.Int("port", 6379, "address to listen to")
	replicaOf = flag.String("replicaof", "", "Replicate to another server")
)

func main() {
	flag.Parse()
	start()
}

func start() {
	fmt.Printf("Begin Process\n")

	serverSettings := settings.ServerSettings{
		Port:   *port,
		Master: true,
	}
	if replicaOf != nil && *replicaOf != "" {
		fmt.Println("This is a Replica Server")
		serverSettings.Master = false
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println("\033[31mFailed to bind to port 6379\n\033[0m")
		return
	}
	fmt.Printf("Listening on %s\n", listener.Addr())

	storage := data.NewStorage()
	server.CreateConnection(listener, storage, serverSettings)
}
