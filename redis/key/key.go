package key

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/lunarhq/sharedutils/types"
)

type Client struct {
	R *redis.Client
}

func (c *Client) Get(secretToken string) (*types.Key, error) {
	b, err := c.R.Get("token-" + secretToken).Bytes()
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

	_, err = c.R.Set("token-"+key.SecretToken, payload, 0).Result()
	return err
}

func (c *Client) Delete(secretToken string) error {
	_, err := c.R.Del("token-" + secretToken).Result()
	return err
}
