package parser

import (
	"fmt"
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
