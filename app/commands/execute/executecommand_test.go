package execute_test

import (
	"bytes"
	parser2 "github.com/codecrafters-io/redis-starter-go/app/commands/execute"
	"testing"

	"github.com/codecrafters-io/redis-starter-go/app/data"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name     string
		cmd      data.RedisCommand
		setup    func(storage data.StorageHelper)
		expected string
	}{
		{
			name:     "PING",
			cmd:      data.RedisCommand{Command: "PING"},
			expected: "+PONG\r\n",
		},
		{
			name:     "ECHO",
			cmd:      data.RedisCommand{Command: "ECHO", Args: []string{"hello"}, ArgsLength: 1},
			expected: "+hello\r\n",
		},
		{
			name:     "ECHO with invalid argument length",
			cmd:      data.RedisCommand{Command: "ECHO", Args: []string{"hello", "world"}, ArgsLength: 2},
			expected: "-(error) ERR wrong number of arguments for command\r\n",
		},
		{
			name:     "SET",
			cmd:      data.RedisCommand{Command: "SET", Args: []string{"key", "value"}, ArgsLength: 2},
			expected: "+OK\r\n",
		},
		{
			name:     "SET with timeout",
			cmd:      data.RedisCommand{Command: "SET", Args: []string{"key", "value", "px", "100"}, ArgsLength: 4},
			expected: "+OK\r\n",
		},
		{
			name:     "SET with invalid argument length",
			cmd:      data.RedisCommand{Command: "SET", Args: []string{"key1"}, ArgsLength: 1},
			expected: "-(error) ERR wrong number of arguments for command\r\n",
		},
		{
			name: "GET",
			cmd:  data.RedisCommand{Command: "GET", Args: []string{"key"}, ArgsLength: 1},
			setup: func(storage data.StorageHelper) {
				storage.Set("key", data.Data{
					Value:          "value",
					ExpiryTimeNano: 0,
				})
			},
			expected: "$5\r\nvalue\r\n",
		},
		{
			name:     "GET with invalid argument length",
			cmd:      data.RedisCommand{Command: "GET", Args: []string{"key", "extra"}, ArgsLength: 2},
			expected: "-(error) ERR wrong number of arguments for command\r\n",
		},
		{
			name:     "Unknown command",
			cmd:      data.RedisCommand{Command: "UNKNOWN"},
			expected: "+COMMAND NOT RECOGNISED: UNKNOWN.\r\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			storage := data.NewStorage()
			if tc.setup != nil {
				tc.setup(storage)
			}
			var buf bytes.Buffer
			parser2.Execute(&buf, storage, tc.cmd)
			result := buf.String()

			if result != tc.expected {
				t.Errorf("Expected: %q, Got: %q", tc.expected, result)
			}
		})
	}
}
