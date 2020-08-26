package main

import (
	"fmt"
	"log"

	"github.com/irohit427/go_grpc_course/greet/greetpb"
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
}
