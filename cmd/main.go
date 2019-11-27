package main

import (
	"fmt"
	"log"
	"meetup/api/bistream"
	"meetup/api/clistream"
	"meetup/api/srvstream"
	"net"

	"google.golang.org/grpc"
)

func main() {
	address := fmt.Sprintf("%s:%d", "127.0.0.1", 9090)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	bistream.Register(s)
	clistream.Register(s)
	srvstream.Register(s)

	s.Serve(lis)
}
