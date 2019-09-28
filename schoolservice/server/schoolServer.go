package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/sathishkumar64/grpc_samples/schoolservice/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//SchoolItem is expose services.
type SchoolItem struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	schoolID   string             `bson:"school_i_d"`
	schoolName string             `bson:"school_name"`
	eduMode    string             `bson:"edu_mode"`
	address    Address            `bson:"address"`
	rating     float32            `bson:"rating"`
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

// ListSchool
func (s SchoolServiceServer) ListSchool(req *model.Void, srv model.SchoolService_ListSchoolServer) error {

	data := &SchoolItem{}

	cursor, err := model.Studentdb.Find(context.Background(), bson.M{})

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		log.Println(data)

		// If no error is found send blog over stream
		response := &model.ListSchoolRes{
			School: &model.School{

				SchoolId:   data.schoolID,
				EduMode:    data.eduMode,
				SchoolName: data.schoolName,
			},
		}
		log.Println("The respones is", response.School)
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

func main() {
	fmt.Println("Hey I'm initializing grpc server.")
	srv := grpc.NewServer()
	schoolService := &SchoolServiceServer{}
	model.RegisterSchoolServiceServer(srv, schoolService)
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen to :8888 :%v", err)
	}
	model.DbConnect()
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	c := make(chan os.Signal)
	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)
	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c
	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	srv.Stop()
	lis.Close()
	fmt.Println("Closing MongoDB connection")
	model.DB.Disconnect(model.MongoCtx)
	fmt.Println("Done.")

}
