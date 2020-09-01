package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/irohit427/go_grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreet().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Function was invoked: %v", req)
	firstName := req.GetGreeting().GetFirstName()
	message := "Hello" + firstName
	res := &greetpb.GreetResponse{
		Message: message,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello Server")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
