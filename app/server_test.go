package main

import (
	"bufio"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/server"
	"net"
	"strings"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	// Create a test listener
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to create test listener: %v", err)
	}
	defer listener.Close()

	// Create a mock storage
	storage := data.NewStorage()

	// Start the server in a separate goroutine
	go server.CreateConnection(listener, storage)

	// Establish a connection to the test server
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}
	defer conn.Close()

	// Create a buffered reader for reading responses
	reader := bufio.NewReader(conn)

	// Test PING command
	sendCommand(t, conn, "*1\r\n$4\r\nPING\r\n")
	expectResponse(t, reader, "+PONG\r\n")

	// Test ECHO command
	sendCommand(t, conn, "*2\r\n$4\r\nECHO\r\n$6\r\nbanana\r\n")
	expectResponse(t, reader, "+banana\r\n")

	// Test SET command
	sendCommand(t, conn, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n")
	expectResponse(t, reader, "+OK\r\n")

	// Test GET command
	sendCommand(t, conn, "*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n")
	expectResponse(t, reader, "value\r\n")

	// Test invalid command
	sendCommand(t, conn, "*2\r\n$7\r\nINVALID\r\n$3\r\nkey\r\n")
	expectResponse(t, reader, "+COMMAND NOT RECOGNISED: INVALID.\r\n")
}

/*
	func TestEndToEnd(t *testing.T) {
		// Start the Redis server
		go start()

		// Connect to the Redis server
		conn, err := net.Dial("tcp", ":6379")
		if err != nil {
			t.Fatalf("Failed to connect to Redis server: %v", err)
		}
		defer conn.Close()

		// Create a buffered reader for reading responses
		reader := bufio.NewReader(conn)

		// Test PING command
		sendCommand(t, conn, "*1\r\n$4\r\nPING\r\n")
		expectResponse(t, reader, "+PONG\r\n")

		// Test ECHO command
		sendCommand(t, conn, "*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n")
		expectResponse(t, reader, "+hello\r\n")

		// Test SET command
		sendCommand(t, conn, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n")
		expectResponse(t, reader, "+OK\r\n")

		// Test GET command
		sendCommand(t, conn, "*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n")
		expectResponse(t, reader, "value\r\n")

		// Test invalid command
		sendCommand(t, conn, "*2\r\n$7\r\nINVALID\r\n$3\r\nkey\r\n")
		expectResponse(t, reader, "+COMMAND NOT RECOGNISED: INVALID.\r\n")
	}
*/
func sendCommand(t *testing.T, conn net.Conn, command string) {
	_, err := conn.Write([]byte(command))
	if err != nil {
		t.Fatalf("Failed to send command: %v", err)
	}
}

func expectResponse(t *testing.T, reader *bufio.Reader, expected string) {
	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}
	if strings.TrimSpace(response) != strings.TrimSpace(expected) {
		t.Errorf("Unexpected response. Got: %s, Expected: %s", response, expected)
	}
}
