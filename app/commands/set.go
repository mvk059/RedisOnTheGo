package commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"io"
)

func Set(rw io.ReadWriter, storage data.StorageHelper, data []string) {
	fmt.Printf("SET %s %s\n", data[0], data[1])
	storage.Set(data[0], data[1])
	_, err := rw.Write([]byte(fmt.Sprintf("+%s\r\n", "OK")))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
