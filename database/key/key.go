package key

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB *mongo.Database
}

func (c *Client) Create() error           { return nil }
func (c *Client) Update() error           { return nil }
func (c *Client) Delete() error           { return nil }
func (c *Client) List() error             { return nil }
func (c *Client) Get() error              { return nil }
func (c *Client) GetBySecretToken() error { return nil }
