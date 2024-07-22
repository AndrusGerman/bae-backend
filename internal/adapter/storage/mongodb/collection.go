package mongodb

import (
	"bae-backend/internal/core/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// fluent API o method chaining | struct methods
type Collection struct {
	db              *DB
	mongoCollection *mongo.Collection
	collectionName  string
}

func NewCollection(db *DB, collectionName string) *Collection {
	var collection = &Collection{db: db}
	return collection.Collection(collectionName)
}

func (c *Collection) Collection(collectionName string) *Collection {
	c.collectionName = collectionName
	c.mongoCollection = c.db.GetMongoCollection(collectionName)
	return c
}

func (c *Collection) InsertOne(src any) error {
	var err error
	_, err = c.mongoCollection.InsertOne(context.TODO(), src)
	return err
}

func (c *Collection) FindOneById(Id domain.Id, output any) error {
	return c.FindOne(bson.M{"_id": Id}, output)
}

func (c *Collection) FindOne(filter interface{}, output any) error {
	var result = c.mongoCollection.FindOne(context.TODO(), filter)
	return result.Decode(output)
}

func (c *Collection) FindMany(filter interface{}, output any) error {
	var cur, err = c.mongoCollection.Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	return cur.All(context.TODO(), output)
}
