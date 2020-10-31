// Package horoscope handles Horoscope RPC calls
package horoscope

import (
	"log"
	"fmt"

	"golang.org/x/net/context"
	pb "github.com/hodamohammadi/grpc-example/proto/horoscope"
)

type Server struct {}

func (s *Server) GetHoroscope(ctx context.Context, in *pb.GetHoroscopeRequest) (*pb.GetHoroscopeResponse, error) {
	log.Printf("Receive request from client: %s", in)
	if in.Sign 
	s := fmt.Sprintf("A new cycle is beginning for you in which you may find yourself throwing away old beliefs and mental processes, %s.", in.Sign.String()) +
	 "Out with the old and in with the new. This time of housecleaning is extremely important, for you will find that the same tired " +
	 "old speech that you've been working on is suddenly defunct. Pull your resources together and construct a new platform that makes you proud."
	return &pb.GetHoroscopeResponse{Text: s}, nil
}