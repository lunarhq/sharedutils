package key

import (
	"context"
	"errors"
	"time"

	"github.com/lunarhq/sharedutils/database"
	"github.com/lunarhq/sharedutils/types"
	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB *mongo.Database
}

func (c *Client) Create(p database.KeyCreateParams) (*types.Key, error) {
	if p.AccountID == nil {
		return nil, errors.New("AccountID required")
	}

	key := types.Key{
		ID:          ksuid.New().String(),
		Created:     time.Now(),
		AccountID:   *p.AccountID,
		SecretToken: ksuid.New().String(),
		Status:      "pending",
		Pro:         false,
	}

	if p.Name != nil {
		key.Name = *p.Name
	}

	if p.SecretToken != nil {
		key.SecretToken = *p.SecretToken
	}

	if p.Status != nil {
		key.Status = *p.Status
	}

	if p.Pro != nil {
		key.Pro = *p.Pro
	}

	ctx := context.Background()
	_, err := c.DB.Collection("keys").InsertOne(ctx, key)
	return &key, err
}

func (c *Client) Update(id string, p database.KeyUpdateParams) error {
	ctx := context.Background()
	payload := bson.M{}

	//@Todo use reflect.ValueOf instead of checking all pointers
	//and setting manually
	if p.Name != nil {
		payload["name"] = p.Name
	}
	if p.SecretToken != nil {
		payload["secretToken"] = p.SecretToken
	}
	if p.AccountID != nil {
		payload["accountId"] = p.AccountID
	}
	if p.Status != nil {
		payload["status"] = p.Status
	}
	if p.Pro != nil {
		payload["pro"] = p.Pro
	}

	_, err := c.DB.Collection("keys").UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err

}

func (c *Client) Delete(id string) error {
	ctx := context.Background()
	_, err := c.DB.Collection("keys").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		//@Todo hack since some old ids were using primitive ones
		//Remove this after all keys are migrated to new syntax
		pid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		_, err = c.DB.Collection("keys").DeleteOne(ctx, bson.M{"_id": pid})
		if err != nil {
			return err
		}
		return nil

	}
	return nil
}

func (c *Client) List(p *database.KeyListParams) ([]*types.Key, error) {
	filter := bson.M{}
	if p != nil && p.AccountID != nil {
		filter["accountId"] = p.AccountID
	}

	ctx := context.Background()
	cur, err := c.DB.Collection("keys").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var result []*types.Key
	err = cur.All(ctx, &result)
	return result, err

}

func (c *Client) Get(id string) (*types.Key, error) {
	ctx := context.Background()
	res := c.DB.Collection("keys").FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var result types.Key
	err := res.Decode(&result)
	return &result, err
}

func (c *Client) GetBySecretToken() error { return nil }
