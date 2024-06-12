package parser

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/errors"
	"net"
	"strings"
)

func Parse(conn net.Conn, storage data.Storage, cmd RedisCommand) {
	switch instruction := strings.ToUpper(cmd.Command); instruction {
	case "PING":
		commands.Ping(conn)
	case "ECHO":
		if cmd.ArgsLength != 1 {
			errors.InvalidArgumentLengthError(conn)
			break
		}
		commands.Echo(conn, strings.Join(cmd.Args, " "))
	case "SET":
		if cmd.ArgsLength != 2 {
			errors.InvalidArgumentLengthError(conn)
			break
		}
		commands.Set(conn, storage, cmd.Args)
	case "GET":
		if cmd.ArgsLength != 1 {
			errors.InvalidArgumentLengthError(conn)
			break
		}
		commands.Get(conn, storage, cmd.Args[0])
	default:
		fmt.Printf("%s: command not found\n", instruction)
		errMessage := fmt.Sprintf("+COMMAND NOT RECOGNISED: %s.\r\n", instruction)
		_, err := conn.Write([]byte(errMessage))
		errors.HandleWritingError(err, "")
	}
}
