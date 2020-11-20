package database

import (
	"context"
	"time"

	"github.com/lunarhq/sharedutils/database/account"
	"github.com/lunarhq/sharedutils/database/invoice"
	"github.com/lunarhq/sharedutils/database/key"
	"github.com/lunarhq/sharedutils/database/paymentmethod"
	"github.com/lunarhq/sharedutils/database/usagerecord"

	"github.com/lunarhq/sharedutils/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDBName = env.Get("MONGO_DB_NAME", "lunardb")
	MongoURI    = env.Get("MONGO_URI", "mongodb://mongo-0.mongo,mongo-1.mongo:27017")
)

type DB struct {
	Keys           *key.Client
	Accounts       *account.Client
	PaymentMethods *paymentmethod.Client
	UsageRecords   *usagerecord.Client
	Invoices       *invoice.Client
}

func NewClient() (*DB, error) {
	mdb, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = mdb.Connect(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	db := mdb.Database(MongoDBName)

	d := &DB{
		Keys:           &key.Client{db},
		Accounts:       &account.Client{db},
		PaymentMethods: &paymentmethod.Client{db},
		UsageRecords:   &usagerecord.Client{db},
		Invoices:       &invoice.Client{db},
	}
	return d, nil
}
