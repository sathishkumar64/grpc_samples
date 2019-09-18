package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sathishkumar64/grpc_samples/schoolservice/model"
	"google.golang.org/grpc"
)

// SchoolServiceServer as service to expose
type SchoolServiceServer struct {
}

// ListSchool to list out all schools
func (s SchoolServiceServer) ListSchool(ctx context.Context, void *model.Void) (*model.ListSchoolRes, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func main() {
	srv := grpc.NewServer()
	var schoolService SchoolServiceServer
	model.RegisterSchoolServiceServer(srv, schoolService)
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen to :8888 :%v", err)
	}
	log.Fatal(srv.Serve(lis))
}
