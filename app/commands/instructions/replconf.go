package instructions

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"github.com/codecrafters-io/redis-starter-go/app/settings"
	"io"
	"strconv"
	"strings"
)

func ReplConf(rw io.ReadWriter, settings *settings.ServerSettings, params []string) {
	command := strings.ToUpper(params[0])

	switch command {
	case "LISTENING-PORT":
		if len(params) != 2 {
			fmt.Println("listening-port requires a port number")
			return
		}

		port, err := strconv.Atoi(params[1])
		if err != nil {
			fmt.Println("listening-port is not a number")
		}

		settings.MasterPort = port
		_, err = rw.Write([]byte(fmt.Sprintf("+%s\r\n", "OK")))
		if err != nil {
			errors.HandleWritingError(err, "")
			return
		}
	case "CAPA":
		_, err := rw.Write([]byte(fmt.Sprintf("+%s\r\n", "OK")))
		if err != nil {
			errors.HandleWritingError(err, "")
			return
		}
	default:
		fmt.Printf("Command not found: %s", params[0])
	}
}
