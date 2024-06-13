package instructions

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/commands/parser"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"io"
)

func Set(rw io.ReadWriter, storage data.StorageHelper, dataItems []string) {
	paramOptions, err := parser.ParamParser(2, dataItems)
	if err != nil {
		fmt.Printf("Error parsing params: %s\n", err)
		return
	}
	value := data.Data{
		Value:          dataItems[1],
		ExpiryTimeNano: paramOptions.Expiry,
	}
	fmt.Printf("SET %s %v\n", dataItems[0], value)
	storage.Set(dataItems[0], value)

	_, err = rw.Write([]byte(fmt.Sprintf("+%s\r\n", "OK")))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
