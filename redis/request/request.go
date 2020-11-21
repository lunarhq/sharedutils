package request

import (
	gr "github.com/go-redis/redis"
	"github.com/lunarhq/sharedutils/redis"
)

type Client struct {
	R *gr.Client
}

func (c *Client) Get(req redis.Request) (*redis.Response, error) {
	return nil, nil
}

func (c *Client) Store(req redis.Request, res redis.Response) error {
	return nil
}
