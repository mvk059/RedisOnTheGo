package parser

import (
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"reflect"
	"testing"
	"time"
)

func TestParseRedisCommand(t *testing.T) {
	echoCommands := []struct {
		name         string
		input        string
		want         *data.RedisCommand
		wantErr      bool
		sleepTimeout time.Duration
	}{
		{
			name:  "ECHO command",
			input: "*2\r\n$4\r\nECHO\r\n$6\r\nbanana\r\n",
			want: &data.RedisCommand{
				Command:    "ECHO",
				Args:       []string{"banana"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
		{
			name:  "ECHO command with spaces",
			input: "*2\r\n$4\r\nECHO\r\n$12\r\nbanana hello\r\n",
			want: &data.RedisCommand{
				Command:    "ECHO",
				Args:       []string{"banana hello"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
	}

	getSetSimpleCommands := []struct {
		name         string
		input        string
		want         *data.RedisCommand
		wantErr      bool
		sleepTimeout time.Duration
	}{
		{
			name:  "GET command",
			input: "*2\r\n$3\r\nGET\r\n$9\r\nblueberry\r\n",
			want: &data.RedisCommand{
				Command:    "GET",
				Args:       []string{"blueberry"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
		{
			name:  "SET command",
			input: "*3\r\n$3\r\nSET\r\n$5\r\nfruit\r\n$6\r\nbanana\r\n",
			want: &data.RedisCommand{
				Command:    "SET",
				Args:       []string{"fruit", "banana"},
				ArgsLength: 2,
			},
			wantErr: false,
		},
		{
			name:  "SET command",
			input: "*3\r\n$3\r\nSET\r\n$6\r\nbanana\r\n$9\r\npineapple\r\n",
			want: &data.RedisCommand{
				Command:    "SET",
				Args:       []string{"banana", "pineapple"},
				ArgsLength: 2,
			},
			wantErr: false,
		},
		{
			name:  "GET command",
			input: "*2\r\n$3\r\nGET\r\n$6\r\nbanana\r\n",
			want: &data.RedisCommand{
				Command:    "GET",
				Args:       []string{"banana"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
		{
			name:  "SET command with spaces",
			input: "*3\r\n$3\r\nSET\r\n$7\r\nmessage\r\n$12\r\nbanana hello\r\n",
			want: &data.RedisCommand{
				Command:    "SET",
				Args:       []string{"message", "banana hello"},
				ArgsLength: 2,
			},
			wantErr: false,
		},
	}

	getSetPxCommands := []struct {
		name         string
		input        string
		want         *data.RedisCommand
		wantErr      bool
		sleepTimeout time.Duration
	}{
		{
			name:  "SET command with PX",
			input: "*5\r\n$3\r\nSET\r\n$6\r\nbanana\r\n$9\r\npineapple\r\n$2\r\nPX\r\n$4\r\n1000\r\n",
			want: &data.RedisCommand{
				Command:    "SET",
				Args:       []string{"banana", "pineapple", "PX", "1000"},
				ArgsLength: 4,
			},
			wantErr: false,
		},
		{
			name:  "GET command after SET with PX",
			input: "*2\r\n$3\r\nGET\r\n$6\r\nbanana\r\n",
			want: &data.RedisCommand{
				Command:    "GET",
				Args:       []string{"banana"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
		{
			name:  "SET command with PX and spaces",
			input: "*5\r\n$3\r\nSET\r\n$7\r\nmessage\r\n$12\r\nbanana hello\r\n$2\r\nPX\r\n$4\r\n1000\r\n",
			want: &data.RedisCommand{
				Command:    "SET",
				Args:       []string{"message", "banana hello", "PX", "1000"},
				ArgsLength: 4,
			},
			wantErr: false,
		},
		{
			name:  "GET command after SET with PX and spaces",
			input: "*2\r\n$3\r\nGET\r\n$7\r\nmessage\r\n",
			want: &data.RedisCommand{
				Command:    "GET",
				Args:       []string{"message"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
		{
			name:  "GET command after expiry",
			input: "*2\r\n$3\r\nGET\r\n$7\r\nmessage\r\n",
			want: &data.RedisCommand{
				Command:    "GET",
				Args:       []string{"message"},
				ArgsLength: 1,
			},
			wantErr:      false,
			sleepTimeout: 1 * time.Second,
		},
		{
			name:  "SET command with short expiry",
			input: "*5\r\n$3\r\nSET\r\n$9\r\nshortlife\r\n$5\r\nhello\r\n$2\r\nPX\r\n$3\r\n500\r\n",
			want: &data.RedisCommand{
				Command:    "SET",
				Args:       []string{"shortlife", "hello", "PX", "500"},
				ArgsLength: 4,
			},
			wantErr: false,
		},
		{
			name:  "GET command immediately after SET with short expiry",
			input: "*2\r\n$3\r\nGET\r\n$9\r\nshortlife\r\n",
			want: &data.RedisCommand{
				Command:    "GET",
				Args:       []string{"shortlife"},
				ArgsLength: 1,
			},
			wantErr:      false,
			sleepTimeout: 0,
		},
		{
			name:  "GET command after delay exceeding short expiry",
			input: "*2\r\n$3\r\nGET\r\n$9\r\nshortlife\r\n",
			want: &data.RedisCommand{
				Command:    "GET",
				Args:       []string{"shortlife"},
				ArgsLength: 1,
			},
			wantErr:      false,
			sleepTimeout: 1 * time.Second,
		},
	}

	invalidCommands := []struct {
		name         string
		input        string
		want         *data.RedisCommand
		wantErr      bool
		sleepTimeout time.Duration
	}{
		{
			name:    "Invalid command format",
			input:   "*2\r\n$4\r\nECHO\r\n",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid number of arguments format",
			input:   "&2\r\n$4\r\nECHO\r\n$6\r\nbanana\r\n",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid length argument format",
			input:   "*2\r\n4\r\nECHO\r\n$6\r\nbanana\r\n",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Argument length mismatch",
			input:   "*2\r\n$4\r\nECHO\r\n$3\r\nbanana\r\n",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid number of arguments",
			input:   "*3\r\n$4\r\nECHO\r\n$6\r\nbanana\r\n",
			want:    nil,
			wantErr: true,
		},
	}

	tests := []struct {
		name         string
		input        string
		want         *data.RedisCommand
		wantErr      bool
		sleepTimeout time.Duration
	}{
		{
			name:  "PING command",
			input: "*1\r\n$4\r\nPING\r\n",
			want: &data.RedisCommand{
				Command:    "PING",
				Args:       []string{},
				ArgsLength: 0,
			},
			wantErr: false,
		},
	}

	tests = append(tests, echoCommands...)
	tests = append(tests, getSetSimpleCommands...)
	tests = append(tests, getSetPxCommands...)
	tests = append(tests, invalidCommands...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRedisCommand([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRedisCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRedisCommand() = %v, want %v", got, tt.want)
			}
			if tt.sleepTimeout > 0 {
				time.Sleep(tt.sleepTimeout)
				if got.Command != tt.want.Command {
					t.Errorf("Expected %q, but got %q", tt.want.Command, got.Command)
				}
			}
		})
	}
}
