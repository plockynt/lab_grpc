package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"local/internal/application"
	"local/internal/infrastructure/persistance/fs"
	pb "local/internal/userinterface/grpcapi"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	persistance, err := fs.New(context.Background(), "/tmp/test.json")
	if err != nil {
		log.Fatalf("failed to create persistance: %v", err)
	}

	projectsSvc := application.NewProjectsSvc(persistance)

	s := pb.NewServer(projectsSvc)

	grpcServer := grpc.NewServer()

	pb.RegisterProjectTrackerServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
