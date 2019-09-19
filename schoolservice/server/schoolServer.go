package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sathishkumar64/grpc_samples/schoolservice/model"
	"google.golang.org/grpc"
)


//School type is expose services.
type School struct {
	schoolID   string 
	schoolname string 
	eduMode   string 
	address  Address 
	rating  float32 
}

//Address type is expose services.
type Address struct {
	address   string 
	state string 
	city   string 
}

// SchoolServiceServer as service to expose
type SchoolServiceServer struct {
}

// ListSchool to list out all schools
func (s SchoolServiceServer) ListSchool(ctx context.Context, void *model.Void) (*model.ListSchoolRes, error) {
	var list []model.ListSchoolRes
	err := model.studentdb.Find(bson.D{}).All(&list)
	 for _, response := range responses {
		log.Printf("all docs %v\n", response)
	}
  
	data := &models.ListSchoolRes{
		School: responses,
	}
	return data, err
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
