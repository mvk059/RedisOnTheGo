package util

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func GetEncodedStringFromMap(input map[string]string) string {
	var builder strings.Builder
	for k, v := range input {
		builder.WriteString(k)
		builder.WriteString(":")
		builder.WriteString(v)
		builder.WriteString("\r\n")
	}
	finalString := builder.String()
	return toBulkString(finalString)
}

func toBulkString(input ...string) string {
	// "$" + len(input) + "\r\n" + input + "\r\n"
	totalLength := 0
	finalString := ""
	for index, v := range input {
		finalString += v + "\r\n"
		if index == 0 {
			totalLength += len(v)
		} else {
			totalLength += len(v) + 2
		}
	}
	return "$" + fmt.Sprint(totalLength) + "\r\n" + finalString
}

func ParseReplicaParams(params string) (string, int, error) {
	parts := strings.Split(params, " ")
	if len(parts) != 2 {
		return "", 0, errors.New("invalid replica parameters")
	}
	var host string
	ip := net.ParseIP(parts[0])
	if ip != nil {
		host = ip.String()
	} else {
		host = parts[0]
	}

	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, errors.New("port parsing issue")
	}

	return host, port, nil
}
