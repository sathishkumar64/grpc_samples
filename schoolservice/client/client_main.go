package main

import (
	"context"
	"fmt"
	"github.com/sathishkumar64/grpc_samples/schoolservice/model"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	con, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Server is not reachable :%v", err)
	}
	client := model.NewSchoolServiceClient(con)
	err = list(context.Background(), client)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func list(ctx context.Context, client model.SchoolServiceClient) error {
	_, err := client.ListSchool(ctx, &model.Void{})

	log.Println("We had error::::::",err)

	if err != nil {
		return fmt.Errorf("Could not fetch task: %v", err)
	}

	/*for _, t := range list.School {
		fmt.Printf("%s\n", t)
	}*/
	return nil
}
