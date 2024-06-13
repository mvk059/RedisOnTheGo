package instructions

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"io"
	"time"
)

func Get(rw io.ReadWriter, storage data.StorageHelper, data string) {
	dataObject, err := storage.Get(data)
	if err != nil || isTimeExpired(dataObject.ExpiryTimeNano) {
		_, _ = rw.Write([]byte("$-1\r\n"))
		if err != nil {
			errors.HandleWritingError(err, "")
			return
		}
		return
	}
	_, err = rw.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(dataObject.Value), dataObject.Value)))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}

}

func isTimeExpired(expiryTime int64) bool {
	return expiryTime != 0 && time.Now().UnixNano() > expiryTime
}
