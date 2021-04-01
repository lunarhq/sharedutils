package usage

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/lunarhq/sharedutils/types"
)

type Client struct {
	DB  *firestore.Client
	Ctx context.Context
}

func (c *Client) Get(accId string, date string) (*types.Usage, error) {
	key := "/api_usage/" + accId + "/items/" + date
	log.Println("key:", key)
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
