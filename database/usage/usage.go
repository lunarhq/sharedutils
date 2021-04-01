package usage

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/lunarhq/sharedutils/types"
	"google.golang.org/api/iterator"
)

type Client struct {
	DB  *firestore.Client
	Ctx context.Context
}

func (c *Client) ListByDate(date string) ([]*types.Usage, error) {
	iter := c.DB.Collection("api_usage").Where("date", "==", date).Documents(c.Ctx)
	defer iter.Stop()
	var result []*types.Usage

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var item types.Usage
		err = doc.DataTo(&item)
		if err != nil {
			return nil, err
		}
		item.ID = doc.Ref.ID
		result = append(result, &item)
	}

	return result, nil
}
