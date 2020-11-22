package key

import (
	"github.com/go-redis/redis"
	"github.com/lunarhq/sharedutils/types"
)

type Client struct {
	R *redis.Client
}

func (c *Client) Get(secretToken string) (*types.Key, error) {
	return nil, nil
}

func (c *Client) Store(key types.Key) error {
	return nil
}
