package user

import (
	"go.mongodb.org/mongo-driver/mongo"
	"fypm.com/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"context"
)

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(collection *mongo.Collection) *MongoRepo {
	return &MongoRepo{
		collection: collection,
	}
}

func (r *MongoRepo) Create(u *User) (entity.ID, error) {
	_, err := r.collection.InsertOne(context.TODO(), u)
	if err != nil{
		log.Fatal(err)
	}
	return u.ID, err
}

func (r *MongoRepo) Get(id entity.ID) (*User, error){
	filter  := bson.D{
	}
	var user User
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return &user, err
}

