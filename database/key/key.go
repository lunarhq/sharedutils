package key

import (
	"context"

	"github.com/lunarhq/sharedutils/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB *mongo.Database
}

func (c *Client) Create(key *database.Key) error {
	ctx := context.Background()
	_, err := c.DB.Collection("keys").InsertOne(ctx, key)
	return err
}

func (c *Client) Update() error { return nil }

func (c *Client) Delete(id string) error {
	ctx := context.Background()
	_, err := c.DB.Collection("keys").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		//@Todo hack since some old ids were using primitive ones
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

func (c *Client) List(p *database.KeyListParams) ([]*database.Key, error) {
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

	var result []*database.Key
	err = cur.All(ctx, &result)
	return result, err

}
func (c *Client) Get(id string) (*database.Key, error) {
	ctx := context.Background()
	res := c.DB.Collection("keys").FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var result database.Key
	err := res.Decode(&result)
	return &result, err
}
func (c *Client) GetBySecretToken() error { return nil }
