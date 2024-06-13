package execute

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/commands/instructions"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"github.com/codecrafters-io/redis-starter-go/app/settings"
	"io"
	"strings"
)

func Execute(rw io.ReadWriter, storage data.StorageHelper, cmd data.RedisCommand, serverSettings settings.ServerSettings) {
	switch instruction := strings.ToUpper(cmd.Command); instruction {
	case "PING":
		instructions.Ping(rw)
	case "ECHO":
		if cmd.ArgsLength != 1 {
			errors.InvalidArgumentLengthError(rw)
			break
		}
		instructions.Echo(rw, strings.Join(cmd.Args, " "))
	case "SET":
		if cmd.ArgsLength < 2 {
			errors.InvalidArgumentLengthError(rw)
			break
		}
		instructions.Set(rw, storage, cmd.Args)
	case "GET":
		if cmd.ArgsLength != 1 {
			errors.InvalidArgumentLengthError(rw)
			break
		}
		instructions.Get(rw, storage, cmd.Args[0])
	case "INFO":
		if cmd.ArgsLength > 1 {
			errors.InvalidArgumentLengthError(rw)
			break
		}
		if cmd.ArgsLength == 1 {
			instructions.Info(rw, serverSettings, cmd.Args[0])
		}
	case "REPLCONF":
		if cmd.ArgsLength < 1 {
			errors.InvalidArgumentLengthError(rw)
			break
		}
		instructions.ReplConf(rw, &serverSettings, cmd.Args)
	default:
		fmt.Printf("%s: command not found\n", instruction)
		errMessage := fmt.Sprintf("+COMMAND NOT RECOGNISED: %s.\r\n", instruction)
		_, err := rw.Write([]byte(errMessage))
		errors.HandleWritingError(err, "")
	}
}
