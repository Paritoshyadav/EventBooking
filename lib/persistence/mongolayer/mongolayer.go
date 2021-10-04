package mongolayer

import (
	"context"
	"fmt"
	"time"

	"github.com/eventbooking/lib/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoLayer struct {
	client *mongo.Client
}

func InitMongoLayer(connection string) (*MongoLayer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected")

	mongoLayer := &MongoLayer{client: client}

	return mongoLayer, nil

}

func (ml *MongoLayer) Close(ctx context.Context, cancel context.CancelFunc) {

	defer func() {

		if err := ml.client.Disconnect(ctx); err != nil {

			panic(err)

		}
		fmt.Println("Connection Closed")

	}()

}

// func (ml *MongoLayer) getFreshSession() (mongo.Session, error) {
// 	opts := options.Session().SetDefaultReadConcern(readconcern.Majority())
// 	s, err := ml.client.StartSession(opts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return s, nil
// }

// with transcation property
// func (ml *MongoLayer) AddEvent(e persistence.Event) (interface{}, error) {
// 	session, err := ml.getFreshSession()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer session.EndSession(ml.ctx)

// 	//mongodb primary Preferred transcation
// 	txnOpts := options.Transaction().
// 		SetReadPreference(readpref.PrimaryPreferred())

// 	res, err := session.WithTransaction(ml.ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

// 		result, err := ml.client.Database(DB).Collection(EVENTS).InsertOne(sessCtx, e)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return result.InsertedID, err

// 	}, txnOpts)

// 	return res, err

// }

func (ml *MongoLayer) AddEvent(e persistence.Event) (interface{}, error) {

	result, err := ml.client.Database(DB).Collection(EVENTS).InsertOne(context.TODO(), e)

	return result.InsertedID, err

}

func (ml *MongoLayer) FindEvent(id string) (persistence.Event, error) {

	e := persistence.Event{}
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return e, err
	}
	err = ml.client.Database(DB).Collection(EVENTS).FindOne(context.TODO(), bson.D{bson.E{Key: "_id", Value: obj_id}}).Decode(&e)
	return e, err
}

func (ml *MongoLayer) FindEventByName(name string) (persistence.Event, error) {
	e := persistence.Event{}
	err := ml.client.Database(DB).Collection(EVENTS).FindOne(context.TODO(), bson.D{bson.E{Key: "name", Value: name}}).Decode(&e)
	return e, err
}

func (ml *MongoLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	e := []persistence.Event{}
	cur, err := ml.client.Database(DB).Collection(EVENTS).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &e)
	if err != nil {
		return nil, err
	}
	return e, err
}

const (
	DB     = "myevents"
	USER   = "mongodb"
	EVENTS = "events"
)
