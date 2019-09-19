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
	schoolID   string  `bson:"school_i_d"`
	schoolName string  `bson:"school_name"`
	eduMode    string  `bson:"edu_mode"`
	address    Address `bson:"address"`
	rating     float32 `bson:"rating"`
}

//Address type is expose services.
type Address struct {
	address string `bson:"address"`
	state   string `bson:"state"`
	city    string `bson:"city"`
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
