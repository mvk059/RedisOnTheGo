package main

import (
	"flag"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/data"
	"github.com/codecrafters-io/redis-starter-go/app/replication"
	"github.com/codecrafters-io/redis-starter-go/app/server"
	"github.com/codecrafters-io/redis-starter-go/app/settings"
	"github.com/codecrafters-io/redis-starter-go/app/util"
	"net"
)

var (
	port      = flag.Int("port", 6379, "address to listen to")
	replicaOf = flag.String("replicaof", "", "Replicate to another server")
)

func main() {
	flag.Parse()
	start()
}

func start() {
	fmt.Printf("Begin Process\n")

	serverSettings := settings.ServerSettings{
		Port:               *port,
		Master:             true,
		MasterReplId:       "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb",
		MasterReplIdOffset: 0,
	}

	if replicaOf != nil && *replicaOf != "" {
		setupReplica(&serverSettings)
		//fmt.Println("This is a Replica Server")
		//host, port, err := parser.ParseReplicaParams(*replicaOf)
		//if err != nil {
		//	panic(fmt.Sprintf("Could not parse replicaOf params: %s", err))
		//}
		//serverSettings.Master = false
		//serverSettings.MasterHost = host
		//serverSettings.MasterPort = port
		replication.Handshake(serverSettings)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println("\033[31mFailed to bind to port 6379\n\033[0m")
		return
	}
	fmt.Printf("Listening on %s\n", listener.Addr())

	storage := data.NewStorage()
	server.CreateConnection(listener, storage, serverSettings)
}

func setupReplica(serverSettings *settings.ServerSettings) {
	fmt.Println("This is a Replica Server")
	host, port, err := util.ParseReplicaParams(*replicaOf)
	if err != nil {
		panic(fmt.Sprintf("Could not parse replicaOf params: %s", err))
	}
	serverSettings.Master = false
	serverSettings.MasterHost = host
	serverSettings.MasterPort = port
}
