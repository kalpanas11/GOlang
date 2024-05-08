package main

import (
	"context"
	"log"
	"time"

	pb "go-gRPC/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewExampleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Kalpana"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGreeting())
}
