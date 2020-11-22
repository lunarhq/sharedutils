package account

import (
	"context"
	"time"

	"github.com/lunarhq/sharedutils/database"
	"github.com/lunarhq/sharedutils/types"
	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//@Todo better context

type Client struct {
	DB *mongo.Database
}

func (c *Client) Create(p database.AccountCreateParams) (*types.Account, error) {
	acc := types.Account{
		ID:        "acc_" + ksuid.New().String(),
		CreatedAt: time.Now(),
		Blocked:   false,
		Internal:  false,
		Pro:       false,
	}

	if p.Name != nil {
		acc.Name = *p.Name
	}
	if p.Email != nil {
		acc.Email = *p.Email
	}
	if p.Pro != nil {
		acc.Pro = *p.Pro
	}
	if p.Internal != nil {
		acc.Internal = *p.Internal
	}
	if p.Blocked != nil {
		acc.Blocked = *p.Blocked
	}
	if p.StripeCustomerID != nil {
		acc.Stripe.CustomerID = *p.StripeCustomerID
	}
	if p.StripeSubscriptionID != nil {
		acc.Stripe.SubscriptionID = *p.StripeSubscriptionID
	}
	if p.StripeSubscriptionItems != nil {
		acc.Stripe.SubscriptionItems = *p.StripeSubscriptionItems
	}

	ctx := context.Background()
	_, err := c.DB.Collection("accounts").InsertOne(ctx, acc)
	return &acc, err
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
	if p.StripeCustomerID != nil {
		payload["stripe.customerId"] = p.StripeCustomerID
	}
	if p.StripeSubscriptionID != nil {
		payload["stripe.subscriptionId"] = p.StripeSubscriptionID
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

func (c *Client) List() ([]*types.Account, error) {
	ctx := context.Background()

	cur, err := c.DB.Collection("accounts").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var result []*types.Account
	err = cur.All(ctx, &result)
	return result, err
}

func (c *Client) Get(id string) (*types.Account, error) {
	ctx := context.Background()
	res := c.DB.Collection("accounts").FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var result types.Account
	err := res.Decode(&result)
	return &result, err
}
