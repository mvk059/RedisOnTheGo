package parser

import (
	"reflect"
	"testing"
)

func TestParseRedisCommand(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *RedisCommand
		wantErr bool
	}{
		{
			name:  "GET command",
			input: "*2\r\n$3\r\nGET\r\n$9\r\nblueberry\r\n",
			want: &RedisCommand{
				Command:    "GET",
				Args:       []string{"blueberry"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
		{
			name:  "SET command",
			input: "*3\r\n$3\r\nSET\r\n$5\r\nfruit\r\n$6\r\nbanana\r\n",
			want: &RedisCommand{
				Command:    "SET",
				Args:       []string{"fruit", "banana"},
				ArgsLength: 2,
			},
			wantErr: false,
		},
		{
			name:  "SET command",
			input: "*3\r\n$3\r\nSET\r\n$6\r\nbanana\r\n$9\r\npineapple\r\n",
			want: &RedisCommand{
				Command:    "SET",
				Args:       []string{"banana", "pineapple"},
				ArgsLength: 2,
			},
			wantErr: false,
		},
		{
			name:  "GET command",
			input: "*2\r\n$3\r\nGET\r\n$6\r\nbanana\r\n",
			want: &RedisCommand{
				Command:    "GET",
				Args:       []string{"banana"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
		{
			name:  "SET command with spaces",
			input: "*3\r\n$3\r\nSET\r\n$7\r\nmessage\r\n$12\r\nbanana hello\r\n",
			want: &RedisCommand{
				Command:    "SET",
				Args:       []string{"message", "banana hello"},
				ArgsLength: 2,
			},
			wantErr: false,
		},
		{
			name:  "PING command",
			input: "*1\r\n$4\r\nPING\r\n",
			want: &RedisCommand{
				Command:    "PING",
				Args:       []string{},
				ArgsLength: 0,
			},
			wantErr: false,
		},
		{
			name:  "ECHO command with spaces",
			input: "*2\r\n$4\r\nECHO\r\n$12\r\nbanana hello\r\n",
			want: &RedisCommand{
				Command:    "ECHO",
				Args:       []string{"banana hello"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
		{
			name:  "ECHO command",
			input: "*2\r\n$4\r\nECHO\r\n$6\r\nbanana\r\n",
			want: &RedisCommand{
				Command:    "ECHO",
				Args:       []string{"banana"},
				ArgsLength: 1,
			},
			wantErr: false,
		},
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
		})
	}
}
