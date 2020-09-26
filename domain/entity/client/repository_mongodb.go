package client

import (
	"context"
	"fypm.com/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(collection *mongo.Collection) *MongoRepo {
	return &MongoRepo{
		collection: collection,
	}
}

func (r *MongoRepo) Get(id entity.ID) (*Client, error) {
	filter := bson.D{{"id", id}}
	var client Client
	err := r.collection.FindOne(context.TODO(), filter).Decode(&client)
	if err != nil {
		log.Fatal(err)
	}
	return &client, err

}

func (r *MongoRepo) Create(c *Client) (entity.ID, error){
	_, err := r.collection.InsertOne(context.TODO(), c)
	if err != nil {
		log.Fatal(err)
	}
	return c.ID, err
}

func (r *MongoRepo) Update(c *Client) error {
	filter := bson.D{{"id", c.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"firstname", c.FirstName},
			{"lastname", c.LastName},
			{"address", c.Address},
			{"country", c.Country},
		},},
	}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err

}

func (r *MongoRepo) Delete(c *Client) error {
	_, err := r.collection.DeleteMany(context.TODO(), bson.D{{"id", c.id}})
	return err
}

func (r *MongoRepo) List() ([]*Client, error){
	var clientList []*Client
	cur, err := r.collection.Find(context.TODO(), bson.D{{}})
	for cur.Next(context.TODO()){
		var client Client
		_ = cur.Decode(&client)
		clientList = append(clientList, &client)
	}
	cur.Close(context.TODO())
	return clientList, err
}
