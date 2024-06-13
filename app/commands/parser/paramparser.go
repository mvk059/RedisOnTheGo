package parser

import (
	"errors"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"strconv"
	"strings"
	"time"
)

func ParamParser(ignore int, params []string) (data.ParamOptions, error) {
	paramOptions := data.ParamOptions{}
	timeoutParams := params[ignore:]
	fmt.Printf("Command: %v, Slice: %v", paramOptions, timeoutParams)

	for i := 0; i < len(timeoutParams); i++ {
		switch strings.ToLower(timeoutParams[i]) {
		case "px":
			if i+1 < len(timeoutParams) {
				expiry, err := strconv.Atoi(timeoutParams[i+1])
				if err == nil {
					expiryTime := time.Now().UTC().Add(time.Duration(expiry * int(time.Millisecond)))
					parseExpiry(&paramOptions, expiryTime)
					i += 2
				} else {
					return data.ParamOptions{}, errors.New("could not parse expiry")
				}
			}
		}

	}

	return paramOptions, nil
}

func parseExpiry(redisCommand *data.ParamOptions, expiryTime time.Time) *data.ParamOptions {
	redisCommand.Expiry = expiryTime.UTC().UnixNano()
	fmt.Println("Set expiry time successfully")
	return redisCommand
}
