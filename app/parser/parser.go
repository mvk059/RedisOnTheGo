package parser

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"io"
	"strings"
)

func Parse(rw io.ReadWriter, storage data.StorageHelper, cmd RedisCommand) {
	switch instruction := strings.ToUpper(cmd.Command); instruction {
	case "PING":
		commands.Ping(rw)
	case "ECHO":
		if cmd.ArgsLength != 1 {
			errors.InvalidArgumentLengthError(rw)
			break
		}
		commands.Echo(rw, strings.Join(cmd.Args, " "))
	case "SET":
		if cmd.ArgsLength != 2 {
			errors.InvalidArgumentLengthError(rw)
			break
		}
		commands.Set(rw, storage, cmd.Args)
	case "GET":
		if cmd.ArgsLength != 1 {
			errors.InvalidArgumentLengthError(rw)
			break
		}
		commands.Get(rw, storage, cmd.Args[0])
	default:
		fmt.Printf("%s: command not found\n", instruction)
		errMessage := fmt.Sprintf("+COMMAND NOT RECOGNISED: %s.\r\n", instruction)
		_, err := rw.Write([]byte(errMessage))
		errors.HandleWritingError(err, "")
	}
}
