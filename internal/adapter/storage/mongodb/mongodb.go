package mongodb

import (
	"bae-backend/internal/adapter/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	MongoDatabase *mongo.Database
	Client        *mongo.Client
}

func New(dbConfig *config.DB) (*DB, error) {

	var uri = fmt.Sprintf("mongodb://%s:%s@%s:%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port) //fmt.Sprintf("mongodb://root:root@localhost:27017/")

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	return &DB{
		MongoDatabase: client.Database(dbConfig.Name),
		Client:        client,
	}, nil
}

func (db *DB) NewCollection(collectionName string) *Collection {
	return NewCollection(db, collectionName)
}

func (db *DB) GetMongoCollection(collectionName string) *mongo.Collection {
	return db.MongoDatabase.Collection(collectionName)
}

func (db *DB) Close() error {
	return db.MongoDatabase.Client().Disconnect(context.TODO())
}
