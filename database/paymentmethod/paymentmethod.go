package paymentmethod

import (
	"context"

	"github.com/lunarhq/sharedutils/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB *mongo.Database
}

func (c *Client) Create(pm *database.PaymentMethod) error {
	ctx := context.Background()
	_, err := c.DB.Collection("payment_methods").InsertOne(ctx, pm)
	return err
}

func (c *Client) Update() error { return nil }

func (c *Client) Delete(id string) error {
	ctx := context.Background()
	_, err := c.DB.Collection("payment_methods").DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (c *Client) List(p *database.PaymentMethodListParams) ([]*database.PaymentMethod, error) {
	filter := bson.M{}
	if p != nil && p.AccountID != nil {
		filter["accountId"] = p.AccountID
	}

	ctx := context.Background()
	cur, err := c.DB.Collection("payment_methods").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var result []*database.PaymentMethod
	err = cur.All(ctx, &result)
	return result, err
}

func (c *Client) Get(id string) (*database.PaymentMethod, error) {
	ctx := context.Background()
	res := c.DB.Collection("payment_methods").FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var result database.PaymentMethod
	err := res.Decode(&result)
	return &result, err
}
