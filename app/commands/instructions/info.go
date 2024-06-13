package instructions

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"github.com/codecrafters-io/redis-starter-go/app/settings"
	"io"
	"strings"
)

func Info(rw io.ReadWriter, serverSettings settings.ServerSettings, data string) {
	var roleValue []string
	roleValue = append(roleValue, "role:")
	roleValue = append(roleValue, getRoleValue(serverSettings))

	res := strings.Join(roleValue, "")
	_, err := rw.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(res), res)))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}

func getRoleValue(serverSettings settings.ServerSettings) string {
	if serverSettings.Master {
		return "master"
	}
	return "slave"
}
