package key

import (
	gr "github.com/go-redis/redis"
	"github.com/lunarhq/sharedutils/redis"
)

type Client struct {
	R *gr.Client
}

func (c *Client) Get(secretToken string) (*redis.Key, error) {
	return nil, nil
}

func (c *Client) Store(key redis.Key) error {
	return nil
}
