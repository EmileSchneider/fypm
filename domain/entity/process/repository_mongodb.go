package process


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

func (r *MongoRepo) Create(p *Process) (entity.ID, error) {
	_, err := r.collection.InsertOne(context.TODO(), p)
	if err != nil {
		log.Fatal(err)
	}
	return p.ID, err
}


func (r *MongoRepo) Update(p *Process) error {
	filter := bson.D{{"id", p.ID}}
	list := bson.A{}
	list = append(list, p.Processes)
	update := bson.D{
		{"$set", bson.D{
			{"firstname", p.FirstName},
			{"lastname", p.LastName},
			{"processes", list}}},
	}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *MongoRepo) Delete(id entity.ID) error {
	_, err := r.collection.DeleteMany(context.TODO(), bson.D{{"id", id}})
	return err
}

func (r *MongoRepo) Get(id entity.ID) (*Process, error) {
	filter := bson.D{{"id", id}}
	var process Process
	err := r.collection.FindOne(context.TODO(), filter).Decode(&process)
	if err != nil {
		log.Fatal(err)
	}
	return &process, err
}

func (r *MongoRepo) List() ([]*Process, error) {
	var processList []*Process
	cur, err := r.collection.Find(context.TODO(), bson.D{{}})
	for cur.Next(context.TODO()){
		var process Process
		_ = cur.Decode(&process)
		processList = append(processList, &process)
	}
	cur.Close(context.TODO())
	return processList, err
}
