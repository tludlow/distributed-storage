package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/JDJFisher/distributed-storage/master/chain"
	"github.com/JDJFisher/distributed-storage/master/health"
	"github.com/JDJFisher/distributed-storage/master/servers"
	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	// Determine port number
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		port = 6000
	}

	// Go
	serve(port)
}

func serve(port int) {
	// Create a TCP connection for the GRPC server
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to open tcp listener... %v", err.Error())
	}

	// Create a GRPC server
	grpcServer := grpc.NewServer()

	// Create a new chain
	chain := chain.NewChain()

	// Register Chain service
	chainServer := servers.NewChainServer(chain)
	protos.RegisterChainServer(grpcServer, chainServer)

	// Register storage service
	storageServer := servers.NewStorageServer(chain)
	protos.RegisterStorageServer(grpcServer, storageServer)

	// Register a health checking server - uses a new() func to initialize the map
	healthServer := health.NewHealthServer(chain)
	protos.RegisterHealthServer(grpcServer, healthServer)

	// Check the status of nodes every 5 seconds
	go healthServer.CheckNodes(2)

	// Start serving GRPC requests on the open tcp connection
	log.Println("Starting master")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start serving the grpc server %v", err.Error())
	}
	defer grpcServer.GracefulStop()
}
