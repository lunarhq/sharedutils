package paymentmethod

import (
	"context"
	"time"

	"github.com/lunarhq/sharedutils/database"
	"github.com/lunarhq/sharedutils/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB *mongo.Database
}

func (c *Client) Create(pm *types.PaymentMethod) error {
	ctx := context.Background()
	pm.CreatedAt = time.Now()
	_, err := c.DB.Collection("payment_methods").InsertOne(ctx, pm)
	return err
}

func (c *Client) Update(id string, p database.PaymentMethodUpdateParams) error {
	ctx := context.Background()
	payload := bson.M{}

	if p.AccountID != nil {
		payload["accountId"] = p.AccountID
	}
	if p.StripeCustomerID != nil {
		payload["stripeCustomerId"] = p.StripeCustomerID
	}
	if p.Brand != nil {
		payload["brand"] = p.Brand
	}
	if p.Last4 != nil {
		payload["last4"] = p.Last4
	}
	if p.Expiry != nil {
		payload["expiry"] = p.Expiry
	}
	if p.Status != nil {
		payload["status"] = p.Status
	}

	payload["updatedAt"] = time.Now()

	_, err := c.DB.Collection("payment_methods").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": payload})
	return err
}

func (c *Client) Delete(id string) error {
	ctx := context.Background()
	_, err := c.DB.Collection("payment_methods").DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (c *Client) List(p *database.PaymentMethodListParams) ([]*types.PaymentMethod, error) {
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

	var result []*types.PaymentMethod
	err = cur.All(ctx, &result)
	return result, err
}

func (c *Client) Get(id string) (*types.PaymentMethod, error) {
	ctx := context.Background()
	res := c.DB.Collection("payment_methods").FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var result types.PaymentMethod
	err := res.Decode(&result)
	return &result, err
}
