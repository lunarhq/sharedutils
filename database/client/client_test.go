package client

import (
	"context"
	"log"
	"testing"

	"github.com/lunarhq/sharedutils/database"
)

func TestClient(t *testing.T) {
	ctx := context.Background()

	c, err := New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// _, err := c.Accounts.List()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	_, err = c.Accounts.Get("ZU5A3yRwWcVzZ4zLaZyefXHmzRY2")
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Accounts.Get("unknown")
	if err == nil {
		t.Fatal("Expected error")
	}

	acc, err := c.Accounts.Create(database.AccountCreateParams{
		Name:                    database.String("name"),
		Email:                   database.String("email"),
		Pro:                     database.Bool(true),
		Internal:                database.Bool(false),
		StripeCustomerID:        database.String("cust_1212"),
		StripeSubscriptionID:    database.String("sub_1212"),
		StripeSubscriptionItems: &map[string]string{"mainnet": "item_1212"},
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Println(acc)
}
