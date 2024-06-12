package server

import (
	"errors"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/parser"
	"io"
	"net"
)

func CreateConnection(listener net.Listener, storage data.StorageHelper) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("\033[31mError accepting connection: \033[0m", err.Error())
			return
		}
		//go handleConnection(conn, storage)
		go func(conn net.Conn) {
			defer conn.Close()
			handleConnection(conn, storage)
		}(conn)
	}
}

func handleConnection(rw io.ReadWriter, storage data.StorageHelper) {
	buf := make([]byte, 1024)
	for {
		_, err := rw.Read(buf)
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
		parser.Parse(rw, storage, *redisCommand)
	}
}
