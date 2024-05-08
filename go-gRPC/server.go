package main

import (
	"context"
	"log"
	"net"

	pb "go-gRPC/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Greeting: "Hello, " + req.GetName()}, nil
}

func main() {
	//listen on port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterExampleServiceServer(grpcServer, &server{})
	log.Printf("Server started at %v", lis.Addr())

	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
