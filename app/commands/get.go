package commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"io"
)

func Get(rw io.ReadWriter, storage data.StorageHelper, data string) {
	get, err := storage.Get(data)
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("Get: %s\n", get)
	if err != nil {
		_, _ = rw.Write([]byte("$-1\r\n"))
		if err != nil {
			errors.HandleWritingError(err, "")
			return
		}
		return
	}
	_, err = rw.Write([]byte(fmt.Sprintf("%s\r\n", get)))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}

}
