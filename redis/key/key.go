package key

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/lunarhq/sharedutils/types"
)

type Client struct {
	R *redis.Client
}

func (c *Client) Get(secretToken string) (*types.Key, error) {
	b, err := c.R.Get(context.Background(), "token-"+secretToken).Bytes()
	if err != nil {
		return nil, err
	}

	var result types.Key
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Store(key types.Key) error {
	payload, err := json.Marshal(key)
	if err != nil {
		return err
	}

	_, err = c.R.Set(context.Background(), "token-"+key.SecretToken, payload, 0).Result()
	return err
}

func (c *Client) Delete(secretToken string) error {
	_, err := c.R.Del(context.Background(), "token-"+secretToken).Result()
	return err
}

//Lists all keys
func (c *Client) List() ([]string, error) {
	//@Todo this should be "token-*" I think
	return c.R.Keys(context.Background(), "*").Result()
}
