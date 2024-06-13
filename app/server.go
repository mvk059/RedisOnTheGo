package main

import (
	"flag"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/server"
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
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println("\033[31mFailed to bind to port 6379\n\033[0m")
		return
	}
	fmt.Printf("Listening on %s\n", listener.Addr())

	storage := data.NewStorage()
	server.CreateConnection(listener, storage)
}
