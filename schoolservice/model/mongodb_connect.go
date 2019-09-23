package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	DB *mongo.Client
	//Studentdb
	Studentdb *mongo.Collection
	MongoCtx  context.Context
)

//DB_connect
func DbConnect() {
	fmt.Println("Connecting to MongoDB...")
	//MongoCtx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	MongoCtx = context.Background()
	//DB, err := mongo.Connect(MongoCtx, options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0-hsyh0.gcp.mongodb.net"))
	DB, err := mongo.Connect(MongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping(MongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	Studentdb = DB.Database("schoolservice").Collection("school")
}
