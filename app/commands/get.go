package commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"net"
)

func Get(conn net.Conn, storage data.Storage, data string) {
	get, err := storage.Get(data)
	if err != nil {
		_, _ = conn.Write([]byte("$-1\r\n"))
		if err != nil {
			errors.HandleWritingError(err, "")
			return
		}
		return
	}
	_, err = conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(get), get)))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}

}
