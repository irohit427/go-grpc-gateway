package main

import (
	"context"
	"fmt"
	"io"
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
	fmt.Printf("Created client %v", c)
	doUnary(c)
	//doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Server Streaming")
	req := &greetpb.GreetManyTimesRequest{
		Greet: &greetpb.Greeting{
			FirstName: "Rohit",
			LastName:  "Raj",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//reached end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Println(msg.GetResult())
	}
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
