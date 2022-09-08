package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	sugar "github.com/kirandesimone/mflix/movie-app/internal"
)

const (
	collection = "movies"
)

var (
	filter = bson.M{}
)

type MongoStore struct {
	db *mongo.Database
}

func ConnectMongoDb(uri string, name string) *MongoStore {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	e := client.Ping(ctx, readpref.Primary())
	if e != nil {
		log.Fatal(e)
	}
	return &MongoStore{db: client.Database(name)}
}

func (ms *MongoStore) Disconnect() error {
	return ms.db.Client().Disconnect(context.Background())
}

func (ms *MongoStore) FindAll(ctx context.Context, results interface{}) error {
	coll := ms.db.Collection(collection)
	curr, err := coll.Find(ctx, filter)
	sugar.Logger.Info("Made db call to Movies collection")
	if err != nil {
		log.Fatal(err)
	}
	return curr.All(ctx, results)
}
