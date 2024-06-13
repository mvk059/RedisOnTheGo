package replication

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/settings"
	"io"
)

const PING = "*1\r\n$4\r\nPING\r\n"

func Handshake(settings settings.ServerSettings) {
	conn := CreateConnection(settings)
	response := sendMessage(conn, PING)
	fmt.Println("Response from ping: ", response)
	response = sendMessage(conn, fmt.Sprintf("*3\r\n$8\r\nREPLCONF\r\n$14\r\nlistening-port\r\n$4\r\n%d\r\n", settings.Port))
	fmt.Println("Response from REPLCONF listening port: ", response)
	response = sendMessage(conn, fmt.Sprintf("*3\r\n$8\r\nREPLCONF\r\n$4\r\ncapa\r\n$6\r\npsync2\r\n"))
	fmt.Println("Response from REPLCONF config: ", response)
}

func sendMessage(rw io.ReadWriter, message string) string {
	_, err := fmt.Fprintf(rw, message)
	if err != nil {
		panic("Could not send message")
	}

	readBuffer := make([]byte, 1024)
	_, readErr := rw.Read(readBuffer)
	if readErr != nil {
		panic("Could not read response message")
	}
	return string(readBuffer)
}
