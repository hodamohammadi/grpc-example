// Package main runs the horoscope service
package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/hodamohammadi/grpc-example/horoscope"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := horoscope.NewServer{}
	grpcServer := grpc.NewServer()

	horoscope.RegisterHoroscopeServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}