package user

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

func (r *MongoRepo) Create(u *User) (entity.ID, error) {
	_, err := r.collection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
	}
	return u.ID, err
}

func (r *MongoRepo) Update(u *User) error {
	filter := bson.D{{"id", u.ID}}
	update := bson.D{
		{"email", u.Email},
		{"password", u.Password},
	}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *MongoRepo) Delete(id entity.ID) error {
	_, err := r.collection.DeleteMany(context.TODO(), bson.D{{"id", id}})
	return err
}

func (r *MongoRepo) Get(id entity.ID) (*User, error) {
	filter := bson.D{{"id", id}}
	var user User
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return &user, err
}

func (r *MongoRepo) List() ([]*User, error) {
	var userList []*User
	cur, err := r.collection.Find(context.TODO(), bson.D{{}})
	for cur.Next(context.TODO()){
		var user User
		_ = cur.Decode(&user)
		userList = append(userList, &user)
	}
	cur.Close(context.TODO())
	return userList, err

}

func (r *MongoRepo) GetByMail(email string) (*User, error) {
	var user User
	filter := bson.D{{"email", email}}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	return &user, err
}
