package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/parser"
	"io"
	"net"
)

var (
	listen = flag.String("listen", ":6379", "address to listen to")
)

func main() {
	flag.Parse()

	start()
}

func start() {
	fmt.Printf("Begin Process\n")
	listener, err := net.Listen("tcp", *listen)
	if err != nil {
		fmt.Println("\033[31mFailed to bind to port 6379\n\033[0m")
		return
	}
	fmt.Printf("Listening on %s\n", listener.Addr())

	storage := data.NewStorage()
	createConnection(listener, storage)
}

func createConnection(listener net.Listener, storage data.Storage) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("\033[31mError accepting connection: \033[0m", err.Error())
			return
		}
		go handleConnection(conn, storage)
	}
}

func handleConnection(conn net.Conn, storage data.Storage) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if errors.Is(err, io.EOF) {
			fmt.Println("Redis client connection closed")
			break
		} else if err != nil {
			fmt.Println("\033[31mError reading input\033[0m", err.Error())
			return
		}

		redisCommand, err := parser.ParseRedisCommand(buf)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		parser.Parse(conn, storage, *redisCommand)
	}
}
