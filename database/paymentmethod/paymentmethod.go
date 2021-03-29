package paymentmethod

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/lunarhq/sharedutils/database"
	"github.com/lunarhq/sharedutils/types"
	"google.golang.org/api/iterator"
)

type Client struct {
	DB  *firestore.Client
	Ctx context.Context
}

func (c *Client) Create(pm *types.PaymentMethod) error {
	pm.CreatedAt = time.Now()
	//@Todo make sure this id doesn't exist already
	docRef := c.DB.Collection("payment_methods").Doc(pm.ID)
	_, err := docRef.Set(c.Ctx, pm)
	return err
}

func (c *Client) Update(id string, p database.PaymentMethodUpdateParams) error {
	updates := []firestore.Update{
		firestore.Update{Path: "updatedAt", Value: time.Now()},
	}

	if p.AccountID != nil {
		u := firestore.Update{Path: "accountId", Value: p.AccountID}
		updates = append(updates, u)
	}
	if p.StripeCustomerID != nil {
		u := firestore.Update{Path: "stripeCustomerId", Value: p.StripeCustomerID}
		updates = append(updates, u)
	}
	if p.Brand != nil {
		u := firestore.Update{Path: "brand", Value: p.Brand}
		updates = append(updates, u)
	}
	if p.Last4 != nil {
		u := firestore.Update{Path: "last4", Value: p.Last4}
		updates = append(updates, u)
	}
	if p.Expiry != nil {
		u := firestore.Update{Path: "expiry", Value: p.Expiry}
		updates = append(updates, u)
	}
	if p.Status != nil {
		u := firestore.Update{Path: "status", Value: p.Status}
		updates = append(updates, u)
	}

	_, err := c.DB.Doc("payment_methods/"+id).Update(c.Ctx, updates)
	return err
}

func (c *Client) Delete(id string) error {
	_, err := c.DB.Doc("payment_methods/" + id).Delete(c.Ctx)
	return err
}

func (c *Client) List(p *database.PaymentMethodListParams) ([]*types.PaymentMethod, error) {
	var iter *firestore.DocumentIterator

	coll := c.DB.Collection("payment_methods")
	if p != nil && p.AccountID != nil {
		iter = coll.Where("accountId", "==", p.AccountID).Documents(c.Ctx)
	} else {
		iter = coll.Documents(c.Ctx)
	}

	defer iter.Stop()

	var result []*types.PaymentMethod

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return result, err
		}
		var pm types.PaymentMethod
		err = doc.DataTo(&pm)
		if err != nil {
			return result, err
		}
		result = append(result, &pm)
	}

	return result, nil
}

func (c *Client) Get(id string) (*types.PaymentMethod, error) {
	doc, err := c.DB.Doc("payment_methods/" + id).Get(c.Ctx)
	if err != nil {
		return nil, err
	}
	var result types.PaymentMethod
	err = doc.DataTo(&result)
	return &result, err
}
