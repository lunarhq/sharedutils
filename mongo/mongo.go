package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoURI = "mongodb://mongo-0.mongo,mongo-1.mongo:27017"
)

type DB struct {
	*mongo.Client
	name   string
	cancel context.CancelFunc
	ctx    context.Context
}

func (db *DB) Close() {
	db.cancel()
	db.Disconnect(db.ctx)
}

func New(dbName string) (*DB, error) {
	mdb, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	db := DB{mdb, dbName, cancel, ctx}
	err = db.Connect(db.ctx)
	if err != nil {
		db.cancel()
		return nil, err
	}
	return &db, nil
}

func (db *DB) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return db.Database(db.name).Collection(name, opts...)
}
