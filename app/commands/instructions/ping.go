package instructions

import (
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"io"
)

func Ping(rw io.ReadWriter) {
	_, err := rw.Write([]byte("+PONG\r\n"))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
