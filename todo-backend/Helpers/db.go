package Helpers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://piotr:piotr1234@cluster0.r2ntk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error while initializing the DB connection", err)
	}

	// client,err := mongo.Connect(context.TODO(),clientOptions)

	// if err != nil{
	// 	log.Fatal(err)
	// }

	if err == nil {
		log.Println("Connection to DB succeeded!")
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Error while pining the DB: ", err)
	}

	DB = client.Database("ToDos")
}
