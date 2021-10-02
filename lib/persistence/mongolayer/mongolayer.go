package mongolayer

import (
	"context"
	"fmt"
	"time"

	"github.com/eventbooking/lib/persistence"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoLayer struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func InitMongoLayer(connection string) *MongoLayer {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected")

	mongoLayer := &MongoLayer{client: client, ctx: ctx, cancel: cancel}

	return mongoLayer

}

func (ml *MongoLayer) Close() {

	defer ml.cancel()

	defer func() {

		if err := ml.client.Disconnect(ml.ctx); err != nil {

			panic(err)

		}
		fmt.Println("Connection Closed")

	}()

}

func (ml *MongoLayer) AddEvent(e persistence.Event) (interface{}, error) {

	// if !e.ID.Valid() {
	// 	e.ID = bson.NewObjectId()
	// }
	// if !e.Location.ID.Valid() {
	// 	e.Location.ID = bson.NewObjectId()
	// }
	opts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	session, err := ml.client.StartSession(opts)
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ml.ctx)

	txnOpts := options.Transaction().
		SetReadPreference(readpref.PrimaryPreferred())

	res, err := session.WithTransaction(ml.ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

		result, err := ml.client.Database(DB).Collection(EVENTS).InsertOne(sessCtx, e)
		if err != nil {
			return nil, err
		}

		return result.InsertedID, err

	}, txnOpts)

	return res, err

}

const (
	DB     = "myevents"
	USER   = "mongodb"
	EVENTS = "events"
)
