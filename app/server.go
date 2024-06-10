package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

var (
	listen = flag.String("listen", ":6379", "address to listen to")
)

func main() {
	flag.Parse()

	err := start()
	if err != nil {
		fmt.Printf("\033[31mError: %v\033[0m\n", err)
		os.Exit(1)
	}
}

func start() (err error) {
	l, err := net.Listen("tcp", *listen)
	if err != nil {
		fmt.Println("\033[31mFailed to bind to port 6379\n\033[0m")
		return err
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("\033[31mError accepting connection: \033[0m", err.Error())
		return err
	}

	defer conn.Close()

	buf := make([]byte, 128)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("\033[31mError reading input\033[0m", err.Error())
		return err
	}

	fmt.Printf("Read: %s\n", buf)

	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		fmt.Println("\033[31mError writing data\033[0m", err.Error())
		return err
	}

	return nil
}
