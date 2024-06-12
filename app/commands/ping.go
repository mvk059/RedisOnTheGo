package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"net"
)

func Ping(conn net.Conn) {
	_, err := conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
