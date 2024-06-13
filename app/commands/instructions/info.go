package instructions

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/commands/parser"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"github.com/codecrafters-io/redis-starter-go/app/settings"
	"io"
)

func Info(rw io.ReadWriter, serverSettings settings.ServerSettings, data string) {
	roleValue := make(map[string]string)
	roleValue[settings.Role] = settings.GetRoleValue(serverSettings)
	//roleValue["port"] = fmt.Sprintf("%d", serverSettings.Port)
	roleValue[settings.MasterReplId] = serverSettings.MasterReplId
	roleValue[settings.MasterReplIdOffset] = fmt.Sprintf("%d", serverSettings.MasterReplIdOffset)

	formattedInput := parser.GetEncodedStringFromMap(roleValue)
	_, err := rw.Write([]byte(fmt.Sprintf("%s", formattedInput)))
	if err != nil {
		errors.HandleWritingError(err, "")
		return
	}
}
