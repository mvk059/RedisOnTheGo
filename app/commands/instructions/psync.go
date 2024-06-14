package instructions

import (
	"encoding/hex"
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
	replicate(rw)
}

func replicate(rw io.ReadWriter) {
	var rdb, _ = hex.DecodeString("524544495330303131fa0972656469732d76657205372e322e30fa0a72656469732d62697473c040fa056374696d65c26d08bc65fa08757365642d6d656dc2b0c41000fa08616f662d62617365c000fff06e3bfec0ff5aa2")
	_, err := rw.Write(append([]byte(fmt.Sprintf("$%d\r\n", len(rdb))), rdb...))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
