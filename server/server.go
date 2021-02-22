package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ayeniblessing101/calculate-break-even/breakeven"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := breakeven.Server{}

	grpcServer := grpc.NewServer()

	breakeven.RegisterBreakEvenServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
