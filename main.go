package main

import (
	"fypm.com/domain/entity/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"context"
	"fmt"

)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	userRepo := user.NewMongoRepo(client.Database("kvbroker").Collection("user"))
	u := user.User{}
	userRepo.Create(&u)
	fmt.Print(userRepo.Get(u.ID))
}

