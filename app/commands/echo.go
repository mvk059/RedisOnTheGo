package commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"io"
)

func Echo(rw io.ReadWriter, message string) {
	_, err := rw.Write([]byte(fmt.Sprintf("+%s\r\n", message)))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
