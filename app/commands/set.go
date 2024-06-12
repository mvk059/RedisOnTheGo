package commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"net"
)

func Set(conn net.Conn, storage data.Storage, data []string) {
	storage.Set(data[0], data[1])
	_, err := conn.Write([]byte(fmt.Sprintf("+%s\r\n", "OK")))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
