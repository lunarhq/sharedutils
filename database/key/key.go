package key

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/lunarhq/sharedutils/database"
	"github.com/lunarhq/sharedutils/types"
	"google.golang.org/api/iterator"

	"github.com/sethvargo/go-password/password"
)

type Client struct {
	DB  *firestore.Client
	Ctx context.Context
}

func (c *Client) Create(p database.KeyCreateParams) (*types.Key, error) {
	if p.AccountID == nil {
		return nil, errors.New("AccountID required")
	}

	// Generate a password that is 32 characters long with 6 digits, 0 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	token, err := password.Generate(32, 6, 0, false, true)
	if err != nil {
		return nil, errors.New("Err generating secret token:" + err.Error())
	}

	key := types.Key{
		Created:     time.Now(),
		AccountID:   *p.AccountID,
		SecretToken: token,
		Status:      "pending",
		Pro:         false,
	}

	if p.Name != nil {
		key.Name = *p.Name
	}

	if p.SecretToken != nil {
		key.SecretToken = *p.SecretToken
	}

	if p.Status != nil {
		key.Status = *p.Status
	}

	if p.Pro != nil {
		key.Pro = *p.Pro
	}

	doc, _, err := c.DB.Collection("keys").Add(c.Ctx, key)
	if err == nil {
		key.ID = doc.ID
	}
	return &key, err
}

func (c *Client) Update(id string, p database.KeyUpdateParams) error {
	updates := []firestore.Update{
		firestore.Update{Path: "updatedAt", Value: time.Now()},
	}

	if p.Name != nil {
		u := firestore.Update{Path: "name", Value: p.Name}
		updates = append(updates, u)
	}
	if p.SecretToken != nil {
		u := firestore.Update{Path: "secretToken", Value: p.SecretToken}
		updates = append(updates, u)
	}
	if p.AccountID != nil {
		u := firestore.Update{Path: "accountId", Value: p.AccountID}
		updates = append(updates, u)
	}
	if p.Status != nil {
		u := firestore.Update{Path: "status", Value: p.Status}
		updates = append(updates, u)
	}
	if p.Pro != nil {
		u := firestore.Update{Path: "pro", Value: p.Pro}
		updates = append(updates, u)
	}

	_, err := c.DB.Doc("keys/"+id).Update(c.Ctx, updates)
	return err

}

func (c *Client) Delete(id string) error {
	_, err := c.DB.Doc("keys/" + id).Delete(c.Ctx)
	return err
}

func (c *Client) List(p *database.KeyListParams) ([]*types.Key, error) {
	var iter *firestore.DocumentIterator

	coll := c.DB.Collection("keys")
	if p != nil && p.AccountID != nil {
		iter = coll.Where("accountId", "==", p.AccountID).Documents(c.Ctx)
	} else {
		iter = coll.Documents(c.Ctx)
	}

	defer iter.Stop()

	var result []*types.Key

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return result, err
		}
		var pm types.Key
		err = doc.DataTo(&pm)
		if err != nil {
			return result, err
		}
		pm.ID = doc.Ref.ID
		result = append(result, &pm)
	}

	return result, nil
}

func (c *Client) Get(id string) (*types.Key, error) {
	doc, err := c.DB.Doc("keys/" + id).Get(c.Ctx)
	if err != nil {
		return nil, err
	}
	var result types.Key
	err = doc.DataTo(&result)
	result.ID = doc.Ref.ID
	return &result, err
}

func (c *Client) GetBySecretKey(secret string) (*types.Key, error) {
	iter := c.DB.Collection("keys").Where("secretToken", "==", secret).Documents(c.Ctx)
	defer iter.Stop()

	var result []*types.Key

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var pm types.Key
		err = doc.DataTo(&pm)
		if err != nil {
			return nil, err
		}
		pm.ID = doc.Ref.ID
		result = append(result, &pm)
	}

	if len(result) > 1 {
		log.Println(result)
		return nil, errors.New("Multiple result found:")
	}

	if len(result) < 1 {
		return nil, errors.New("No results found")
	}

	return result[0], nil

}
