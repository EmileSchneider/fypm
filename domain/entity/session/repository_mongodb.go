package session


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

func (r *MongoRepo) Create(s *Session) error {
	_ , err := r.collection.InsertOne(context.TODO(), s)
	return err
}

func (r *MongoRepo) Delete(id entity.ID){
	 _, err := r.collection.DeleteMany(context.TODO(), bson.D{{"id", id}})
	return err
}

func (r *MongoRepo) List() ([]*Session, error){
	var sessions []*Session
	cur, err := r.collection.Find(context.TODO(), bson.D{{}})
	for cur.Next(context.TODO()){
		var session Session
		_ = cur.Decode(&session)
		sessions = append(sessions, &session)
	}
	cur.Close(context.TODO())
	return sessions, err
}
