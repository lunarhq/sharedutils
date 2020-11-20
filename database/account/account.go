package account

import (
	"context"

	"github.com/lunarhq/sharedutils/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//@Todo better context

type Client struct {
	DB *mongo.Database
}

func (c *Client) Create(acc *database.Account) error {
	ctx := context.Background()
	_, err := c.DB.Collection("accounts").InsertOne(ctx, acc)
	return err
}

func (c *Client) Update(id string, p *database.AccountUpdateParams) error {
	ctx := context.Background()
	payload := bson.M{}
	if p.Blocked != nil {
		payload["blocked"] = p.Blocked
	}
	if p.Name != nil {
		payload["name"] = p.Name
	}
	if p.Internal != nil {
		payload["internal"] = p.Internal
	}
	if p.StripeCustomerId != nil {
		payload["stripe.customerId"] = p.StripeCustomerId
	}
	if p.StripeSubscriptionId != nil {
		payload["stripe.subscriptionId"] = p.StripeSubscriptionId
	}
	if p.StripeSubscriptionItems != nil {
		payload["stripe.subscriptionItems"] = p.StripeSubscriptionItems
	}

	_, err := c.DB.Collection("accounts").UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

func (c *Client) Delete(id string) error {
	ctx := context.Background()
	_, err := c.DB.Collection("accounts").DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (c *Client) List() ([]*database.Account, error) {
	ctx := context.Background()

	cur, err := c.DB.Collection("accounts").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var result []*database.Account
	err = cur.All(ctx, &result)
	return result, err
}

func (c *Client) Get(id string) (*database.Account, error) {
	ctx := context.Background()
	res := c.DB.Collection("accounts").FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var result database.Account
	err := res.Decode(&result)
	return &result, err
}
