package proxy_rule

import (
	"context"
	"errors"
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

func (c *Client) Create(p database.ProxyRuleCreateParams) (*types.ProxyRule, error) {
	if p.AccountID == nil {
		return nil, errors.New("AccountID required")
	}

	rule := types.ProxyRule{
		Created:   time.Now(),
		AccountID: *p.AccountID,
	}

	if p.Name != nil {
		rule.Name = *p.Name
	}
	if p.Endpoint != nil {
		rule.Endpoint = *p.Endpoint
	}
	if p.HeaderApiKeyKey != nil {
		rule.HeaderApiKeyKey = *p.HeaderApiKeyKey
	}
	if p.HeaderApiKeyValue != nil {
		rule.HeaderApiKeyValue = *p.HeaderApiKeyValue
	}
	if p.HeaderApiSecretKey != nil {
		rule.HeaderApiSecretKey = *p.HeaderApiSecretKey
	}
	if p.HeaderApiSecretValue != nil {
		rule.HeaderApiSecretValue = *p.HeaderApiSecretValue
	}
	if p.WhitelistedDomains != nil {
		rule.WhitelistedDomains = *p.WhitelistedDomains
	}

	doc, _, err := c.DB.Collection("proxy_rules").Add(c.Ctx, rule)
	if err == nil {
		rule.ID = doc.ID
	}
	return &rule, err
}

func (c *Client) Delete(id string) error {
	_, err := c.DB.Doc("proxy_rules/" + id).Delete(c.Ctx)
	return err
}

func (c *Client) List(p *database.ProxyRuleListParams) ([]*types.ProxyRule, error) {
	var iter *firestore.DocumentIterator

	coll := c.DB.Collection("proxy_rules")
	if p != nil && p.AccountID != nil {
		iter = coll.Where("accountId", "==", p.AccountID).Documents(c.Ctx)
	} else {
		iter = coll.Documents(c.Ctx)
	}

	defer iter.Stop()

	var result []*types.ProxyRule

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return result, err
		}
		var pm types.ProxyRule
		err = doc.DataTo(&pm)
		if err != nil {
			return result, err
		}
		pm.ID = doc.Ref.ID
		result = append(result, &pm)
	}

	return result, nil
}

func (c *Client) Get(id string) (*types.ProxyRule, error) {
	doc, err := c.DB.Doc("proxy_rules/" + id).Get(c.Ctx)
	if err != nil {
		return nil, err
	}
	var result types.ProxyRule
	err = doc.DataTo(&result)
	result.ID = doc.Ref.ID
	return &result, err
}
