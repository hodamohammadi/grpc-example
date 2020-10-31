package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/hodamohammadi/grpc-example/horoscope"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := horoscope.NewHoroscopeServiceClient(conn)

	response, err := c.GetHoroscope(context.Background(), &horoscope.GetHoroscopeRequest{Sign: "Scorpio"})
	if err != nil {
		log.Fatalf("Error when calling GetHoroscope: %s", err)
	}
	log.Printf("Response from server: %s", response.text)

}