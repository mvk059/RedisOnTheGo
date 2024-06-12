package errors

import (
	"fmt"
	"net"
)

func InvalidArgumentLengthError(conn net.Conn) {
	_, err := conn.Write([]byte("+(error) ERR wrong number of arguments for command\r\n"))
	if err != nil {
		fmt.Println("\033[31mError writing data\033[0m", err.Error())
	}
}

func HandleWritingError(err error, message string) {
	var errorMessage = message
	if len(message) == 0 {
		errorMessage = "Error writing data"
	}
	fmt.Printf("\033[31m%s\nError:%s\033[0m\n", errorMessage, err.Error())
}
