package main

import (
	"context"
	"fmt"
	"log"

	"github.com/irohit427/go_grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Hello, Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rohit",
			LastName:  "Raj",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet %v", err)
	}
	log.Printf("Response form greet: %v", res.Message)
}
