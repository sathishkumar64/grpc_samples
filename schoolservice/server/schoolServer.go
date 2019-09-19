package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/sathishkumar64/grpc_samples/schoolservice/model"
)

//School type is expose services.
type School struct {
	schoolID   string
	schoolName string
	eduMode    string
	address    Address
	rating     float32
}

//Address type is expose services.
type Address struct {
	address string
	state   string
	city    string
}

// SchoolServiceServer as service to expose
type SchoolServiceServer struct {
}

// ListSchool to list out all schools
func (s SchoolServiceServer) ListSchool(ctx context.Context, void *model.Void) (*model.ListSchoolRes, error) {
	var list []model.ListSchoolRes
	_, err := model.Studentdb.Find(ctx, nil)
	for _, response := range list {
		log.Printf("all docs %v\n", response)
	}
	return nil, err
}

func main() {
	fmt.Println("Hey I'm initializing grpc server.")
	srv := grpc.NewServer()
	var schoolService SchoolServiceServer
	model.RegisterSchoolServiceServer(srv, schoolService)
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen to :8888 :%v", err)
	}
	log.Fatal(srv.Serve(lis))
}
