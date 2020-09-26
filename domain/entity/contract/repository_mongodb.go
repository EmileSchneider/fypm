package contract

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

func (r *MongoRepo) Get(id entity.ID) (*Contract, error) {
	filter := bson.D{{"id", id}}
	var contract Contract
	err := r.collection.FindOne(context.TODO(), filter).Decode(&contract)
	if err != nil {
		log.Fatal(err)
	}
	return &contract, err

}

func (r *MongoRepo) Create(c *Contract) (entity.ID, error){
	_, err := r.collection.InsertOne(context.TODO(), c)
	if err != nil {
		log.Fatal(err)
	}
	return c.ID, err
}

func (r *MongoRepo) Update(c *Contract) error {
	filter := bson.D{{"id", c.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"premiumincents", c.PremiumInCents},
			{"courtageinpercent", c.CourtageInPercent},
			{"insurer", c.Insurer}}},
	}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err

}

func (r *MongoRepo) Delete(c *Contract) error {
	_, err := r.collection.DeleteMany(context.TODO(), bson.D{{"id", c.id}})
	return err
}

func (r *MongoRepo) List() ([]*Contract, error){
	var contractList []*Contract
	cur, err := r.collection.Find(context.TODO(), bson.D{{}})
	for cur.Next(context.TODO()){
		var contract Contract
		_ = cur.Decode(&contract)
		contractList = append(contractList, &contract)
	}
	cur.Close(context.TODO())
	return contractList, err
}
