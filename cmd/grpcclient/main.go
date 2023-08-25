package main

import (
	"context"
	"io"
	"log"

	pb "local/internal/userinterface/grpcapi"

	"google.golang.org/grpc"
)

func main() {
	// dial server
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewProjectTrackerClient(conn)

	// save 2
	inProject := pb.Project{
		Id:    2,
		Label: "test",
	}
	_, err = client.Save(context.Background(), &inProject)
	if err != nil {
		log.Fatalf("save error %v", err)
	}
	log.Printf("save 2: %+v", inProject)

	inProject = pb.Project{
		Id:          3,
		Description: "test",
	}
	_, err = client.Save(context.Background(), &inProject)
	if err != nil {
		log.Fatalf("save error %v", err)
	}
	log.Printf("save 3: %+v", inProject)

	// get all
	in := &pb.GetAllParams{}
	stream, err := client.GetAll(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("resp received: %+v", resp)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished getAll")

	// get by id 1
	inID := pb.ID{
		Id: 2,
	}
	project, err := client.GetByID(context.Background(), &inID)
	if err != nil {
		log.Fatalf("getByID error %v", err)
	}
	log.Printf("getByID: %+v", project)
}
