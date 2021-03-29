package account

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/lunarhq/sharedutils/database"
	"github.com/lunarhq/sharedutils/types"
	"github.com/segmentio/ksuid"
	"google.golang.org/api/iterator"
)

//@Todo better context

type Client struct {
	DB  *firestore.Client
	Ctx context.Context
}

func (c *Client) Create(p database.AccountCreateParams) (*types.Account, error) {
	acc := types.Account{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now(),
		Blocked:   false,
		Internal:  false,
		Pro:       false,
		Stripe:    &types.StripeData{},
	}

	if p.ID != nil {
		acc.ID = *p.ID
	}
	if p.Name != nil {
		acc.Name = *p.Name
	}
	if p.Email != nil {
		acc.Email = *p.Email
	}
	if p.Pro != nil {
		acc.Pro = *p.Pro
	}
	if p.Internal != nil {
		acc.Internal = *p.Internal
	}
	if p.Blocked != nil {
		acc.Blocked = *p.Blocked
	}
	if p.StripeCustomerID != nil {
		acc.Stripe.CustomerID = *p.StripeCustomerID
	}
	if p.StripeSubscriptionID != nil {
		acc.Stripe.SubscriptionID = *p.StripeSubscriptionID
	}
	if p.StripeSubscriptionItems != nil {
		acc.Stripe.SubscriptionItems = *p.StripeSubscriptionItems
	}

	//@Todo make sure doc doesn't already exist
	docRef := c.DB.Collection("accounts").Doc(acc.ID)
	_, err := docRef.Set(c.Ctx, acc)
	return &acc, err
}

func (c *Client) Update(id string, p *database.AccountUpdateParams) error {
	updates := []firestore.Update{
		firestore.Update{Path: "updatedAt", Value: time.Now()},
	}

	if p.Blocked != nil {
		u := firestore.Update{Path: "blocked", Value: p.Blocked}
		updates = append(updates, u)
	}
	if p.Name != nil {
		u := firestore.Update{Path: "name", Value: p.Name}
		updates = append(updates, u)
	}
	if p.Pro != nil {
		u := firestore.Update{Path: "pro", Value: p.Pro}
		updates = append(updates, u)
	}
	if p.Internal != nil {
		u := firestore.Update{Path: "internal", Value: p.Internal}
		updates = append(updates, u)
	}
	if p.StripeCustomerID != nil {
		u := firestore.Update{Path: "stripe.customerId", Value: p.StripeCustomerID}
		updates = append(updates, u)
	}
	if p.StripeSubscriptionID != nil {
		u := firestore.Update{Path: "stripe.subscriptionId", Value: p.StripeSubscriptionID}
		updates = append(updates, u)
	}
	if p.StripeSubscriptionItems != nil {
		u := firestore.Update{Path: "stripe.subscriptionItems", Value: p.StripeSubscriptionItems}
		updates = append(updates, u)
	}

	_, err := c.DB.Doc("accounts/"+id).Update(c.Ctx, updates)
	return err
}

func (c *Client) Delete(id string) error {
	_, err := c.DB.Doc("accounts/" + id).Delete(c.Ctx)
	return err
}

func (c *Client) List() ([]*types.Account, error) {
	var iter *firestore.DocumentIterator

	coll := c.DB.Collection("accounts")
	iter = coll.Documents(c.Ctx)

	defer iter.Stop()

	var result []*types.Account

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return result, err
		}
		var pm types.Account
		err = doc.DataTo(&pm)
		if err != nil {
			return result, err
		}
		pm.ID = doc.Ref.ID
		result = append(result, &pm)
	}

	return result, nil
}

func (c *Client) Get(id string) (*types.Account, error) {
	doc, err := c.DB.Doc("accounts/" + id).Get(c.Ctx)
	if err != nil {
		return nil, err
	}
	var result types.Account
	err = doc.DataTo(&result)
	result.ID = doc.Ref.ID
	return &result, err
}

func (c *Client) GetByStripeID(stripeID string) (*types.Account, error) {
	iter := c.DB.Collection("accounts").Where("stripe.customerId", "==", stripeID).Documents(c.Ctx)
	defer iter.Stop()
	var result []*types.Account

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var pm types.Account
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
