package main

import (
	"flag"
	"fmt"
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
	l, err := net.Listen("tcp", *listen)
	if err != nil {
		fmt.Println("\033[31mFailed to bind to port 6379\n\033[0m")
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("\033[31mError accepting connection: \033[0m", err.Error())
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 128)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("\033[31mError reading input\033[0m", err.Error())
			conn.Close()
			return
		}

		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("\033[31mError writing data\033[0m", err.Error())
			conn.Close()
			return
		}
	}
}
