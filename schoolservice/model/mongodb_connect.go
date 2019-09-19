package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var db *mongo.Client
var Studentdb *mongo.Collection
var mongoCtx context.Context

func db_connect() {

	fmt.Println("Connecting to MongoDB...")
	mongoCtx, _ = context.WithTimeout(context.Background(), 2*time.Second)

	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0-hsyh0.gcp.mongodb.net"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}

	Studentdb = db.Database("schoolservice").Collection("school")
}
