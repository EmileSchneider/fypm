package broker

import(
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

func (r *MongoRepo) Create(b *Broker) (entity.ID, error) {
	_, err := r.collection.InsertOne(context.TODO(), b)
	if err != nil {
		log.Fatal(err)
	}
	return b.ID, err
}


func (r *MongoRepo) Update(b *Broker) error {
	filter := bson.D{{"id", b.ID}}
	update := bson.D{
		{"firstname", b.FirstName},
		{"lastname", b.LastName},
		{"processes", bson.A{}},
	}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *MongoRepo) Delete(id entity.ID) error {
	_, err := r.collection.DeleteMany(context.TODO(), bson.D{{"id", id}})
	return err
}

func (r *MongoRepo) Get(id entity.ID) (*Broker, error) {
	filter := bson.D{{"id", id}}
	var broker Broker
	err := r.collection.FindOne(context.TODO(), filter).Decode(&broker)
	if err != nil {
		log.Fatal(err)
	}
	return &broker, err
}

func (r *MongoRepo) List() ([]*Broker, error) {
	var brokerList []*Broker
	cur, err := r.collection.Find(context.TODO(), bson.D{{}})
	for cur.Next(context.TODO()){
		var broker Broker
		_ = cur.Decode(&broker)
		brokerList = append(brokerList, &broker)
	}
	cur.Close(context.TODO())
	return brokerList, err

}
