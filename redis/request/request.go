package request

import (
	"github.com/go-redis/redis"
	"github.com/lunarhq/sharedutils/types"
)

type Client struct {
	R *redis.Client
}

func (c *Client) Get(path string, req types.RosettaRequest) (*types.Response, error) {
	return nil, nil
}

func (c *Client) Store(path string, req types.RosettaRequest, res types.RosettaResponse) error {
	return nil
}
