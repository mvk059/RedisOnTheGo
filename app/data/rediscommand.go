package data

// RedisCommand represents a parsed Redis command.
type RedisCommand struct {
	Command    string   // Name of the Redis command
	Args       []string // Arguments of the Redis command
	ArgsLength int      // Length of the arguments
}

type ParamOptions struct {
	Expiry int64
}
