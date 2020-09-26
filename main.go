package main

import (
	"context"
	"fmt"
	"fypm.com/domain/entity/user"
	"fypm.com/domain/entity/broker"
	"fypm.com/domain/usecase/testdbsetup"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	userRepo := user.NewMongoRepo(client.Database("kvbroker").Collection("user"))
	userManager := user.NewManager(userRepo)
	brokerRepo := broker.NewMongoRepo(client.Database("kvbroker").Collection("broker"))
	brokerManager := broker.NewManager(brokerRepo)

	usecase := testdbsetup.NewUseCase(userManager, brokerManager)
	usecase.Setup()
	usecase.Test()
	fmt.Println("running")
}
