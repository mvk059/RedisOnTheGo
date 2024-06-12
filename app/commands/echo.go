package commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"net"
)

func Echo(conn net.Conn, message string) {
	_, err := conn.Write([]byte(fmt.Sprintf("+%s\r\n", message)))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
