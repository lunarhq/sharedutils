package usage

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/lunarhq/sharedutils/types"
)

type Client struct {
	DB  *firestore.Client
	Ctx context.Context
}

func (c *Client) Get(accId string, date string) (*types.Usage, error) {
	key := "api_usage/" + accId + "/items/" + date
	doc, err := c.DB.Doc(key).Get(c.Ctx)
	if err != nil {
		return nil, err
	}

	var item types.Usage
	err = doc.DataTo(&item)
	if err != nil {
		return nil, err
	}
	item.AccountID = accId
	item.Date = date

	return &item, nil
}
func (c *Client) Increment(accId string, date string, networkType string, hits int64) error {
	//1. First check if key exist else init it
	key := "api_usage/" + accId + "/items/" + date
	_, err := c.DB.Doc(key).Get(c.Ctx)
	if err != nil {
		initData := map[string]int64{"mainnet": 0, "testnet": 0}
		_, err = c.DB.Doc(key).Set(c.Ctx, initData)
		if err != nil {
			return err
		}
	}

	updates := []firestore.Update{
		firestore.Update{Path: "mainnet", Value: firestore.Increment(hits)},
	}
	_, err = c.DB.Doc(key).Update(c.Ctx, updates)
	return err
}
