package instructions

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"github.com/codecrafters-io/redis-starter-go/app/settings"
	"io"
)

func Psync(rw io.ReadWriter, serverSettings settings.ServerSettings) {
	response := fmt.Sprintf("+FULLRESYNC %s 0\r\n", serverSettings.MasterReplId)
	_, err := rw.Write([]byte(response))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
