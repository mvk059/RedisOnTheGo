package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// RedisCommand represents a parsed Redis command.
type RedisCommand struct {
	Command    string   // Name of the Redis command
	Args       []string // Arguments of the Redis command
	ArgsLength int      // Length of the arguments
}

// ParseRedisCommand parses a Redis command string and returns a RedisCommand struct.
//
// The Redis command string should be in the following format:
// *<number of arguments>\r\n
// $<length of command name>\r\n
// <command name>\r\n
// $<length of argument 1>\r\n
// <argument 1>\r\n
// ...
// $<length of argument N>\r\n
// <argument N>\r\n
//
// Example: *2\r\n$4\r\nECHO\r\n$6\r\nbanana\r\n
//
// If the command string is not in the valid format, an error is returned.
func ParseRedisCommand(buf []byte) (*RedisCommand, error) {
	// Split the command string into parts using "\r\n" as the separator
	parts := strings.Split(string(buf), "\r\n")

	// Check if the command string has at least 3 parts
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid command format\n")
	}

	// Check if the command starts with "*"
	if parts[0][:1] != "*" {
		return nil, fmt.Errorf("invalid number of arguments format\n")
	}

	// Parse the number of arguments from the first part of the command string
	numArgs, err := strconv.Atoi(parts[0][1:])
	if err != nil {
		return nil, fmt.Errorf("invalid number of arguments format: %v\n", err)
	}

	// Create a new RedisCommand struct with a pre-allocated slice for arguments
	redisCmd := &RedisCommand{
		Args:       make([]string, 0, numArgs-1),
		ArgsLength: numArgs - 1,
	}
	// Iterate over each argument in the command string.
	for i := 0; i < numArgs; i++ {
		// Check if the argument length part starts with "$"
		if len(parts[2*i+1]) < 2 || !strings.HasPrefix(parts[2*i+1][:1], "$") {
			return nil, fmt.Errorf("invalid length argument format\n")
		}

		// Parse the argument length from the argument length part.
		argLen, err := strconv.Atoi(parts[2*i+1][1:])
		if err != nil || len(parts[2*i+2]) != argLen {
			return nil, fmt.Errorf("argument is not of valid length\n")
		}

		// Set the command name for the first argument.
		if i == 0 {
			redisCmd.Command = parts[2*i+2]
		} else {
			// Append the argument to the Args slice for subsequent arguments.
			redisCmd.Args = append(redisCmd.Args, parts[2*i+2])
		}
	}

	return redisCmd, err
}
