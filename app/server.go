package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"strings"
)

var (
	listen = flag.String("listen", ":6379", "address to listen to")
)

func main() {
	flag.Parse()

	start()
}

func start() {
	listener, err := net.Listen("tcp", *listen)
	if err != nil {
		fmt.Println("\033[31mFailed to bind to port 6379\n\033[0m")
		return
	}
	//defer l.Close()

	createConnection(listener)
}

func createConnection(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("\033[31mError accepting connection: \033[0m", err.Error())
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
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

		parts := strings.Split(string(buf), "\r\n")
		fmt.Printf("Parts: %s\n", parts)

		switch instruction := strings.ToUpper(parts[2]); instruction {
		case "PING":
			_, err = conn.Write([]byte("+PONG\r\n"))
			if err != nil {
				fmt.Println("\033[31mError writing data\033[0m", err.Error())
				return
			}
		case "ECHO":
			var echo string
			if len(parts) == 6 {
				echo = fmt.Sprintf("+%s\r\n", parts[4])
			} else {
				echo = "+(error) ERR wrong number of arguments for command\r\n"
			}
			_, err := conn.Write([]byte(echo))
			if err != nil {
				fmt.Println("\033[31mError writing data\033[0m", err.Error())
				return
			}
		default:
			fmt.Printf("%s: command not found\n", instruction)
			errMessage := fmt.Sprintf("+COMMAND NOT RECOGNISED: %s.\r\n", instruction)
			_, _ = conn.Write([]byte(errMessage))
		}

	}
}
