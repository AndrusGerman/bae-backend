package mongodb

import (
	"bae-backend/internal/core/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// fluent API o method chaining | struct methods
type Collection struct {
	db             *DB
	collectionName string
}

func NewCollection(db *DB, collectionName string) *Collection {
	var collection = &Collection{db: db}
	return collection.Collection(collectionName)
}

func (c *Collection) Collection(collectionName string) *Collection {
	c.collectionName = collectionName
	return c
}

func (c *Collection) InsertOne(src any) error {
	var err error
	_, err = c.GetMongoCollection().InsertOne(context.TODO(), src)
	return err
}

func (c *Collection) FindOneById(Id domain.Id, output any) error {
	return c.FindOne(bson.M{"_id": Id}, output)
}

func (c *Collection) FindOne(filter interface{}, output any) error {
	var result = c.GetMongoCollection().FindOne(context.TODO(), filter)
	return result.Decode(output)
}

func (c *Collection) FindMany(filter interface{}, output any) error {
	var cur, err = c.GetMongoCollection().Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	return cur.All(context.TODO(), output)
}
func (c *Collection) GetMongoCollection() *mongo.Collection {
	return c.db.GetMongoCollection(c.collectionName)
}

func (c *Collection) CreateCollection(cco *options.CreateCollectionOptions) error {
	return c.db.CreateCollection(c.collectionName, cco)
}
