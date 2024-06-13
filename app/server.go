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
	port = flag.Int("port", 6379, "address to listen to")
)

func main() {
	flag.Parse()
	start()
}

func start() {
	fmt.Printf("Begin Process\n")

	serverSettings := settings.ServerSettings{
		Port:   6379,
		Master: true,
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
