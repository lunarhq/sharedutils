package client

import (
	"context"
	"testing"
	"time"

	"github.com/lunarhq/sharedutils/database"
	"github.com/lunarhq/sharedutils/types"
)

func TestClient(t *testing.T) {
	ctx := context.Background()

	c, err := New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	//Accounts
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

	_, err = c.Accounts.Get(acc.ID)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Accounts.Update(acc.ID, &database.AccountUpdateParams{Pro: database.Bool(false), Internal: database.Bool(true)})
	if err != nil {
		t.Fatal(err)
	}
	accUpdated, err := c.Accounts.Get(acc.ID)
	if err != nil {
		t.Fatal(err)
	}
	if accUpdated.Pro != false {
		t.Fail()
	}
	if accUpdated.Internal != true {
		t.Fail()
	}

	err = c.Accounts.Delete(acc.ID)
	if err != nil {
		t.Fatal(err)
	}

	//Keys
	_, err = c.Keys.Get("Unknown")
	if err == nil {
		t.Fatal("Expected error")
	}

	key, err := c.Keys.Create(database.KeyCreateParams{
		Name:        database.String("theName"),
		SecretToken: database.String("secret"),
		AccountID:   database.String("accId"),
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Keys.Get(key.ID)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Keys.Update(key.ID, database.KeyUpdateParams{Pro: database.Bool(true)})
	if err != nil {
		t.Fatal(err)
	}
	keyUpdated, err := c.Keys.Get(key.ID)
	if err != nil {
		t.Fatal(err)
	}
	if keyUpdated.Pro != true {
		t.Fail()
	}

	err = c.Keys.Delete(key.ID)
	if err != nil {
		t.Fatal(err)
	}

	//Payment methods
	_, err = c.Keys.Get("Unknown")
	if err == nil {
		t.Fatal("Expected error")
	}

	pm := &types.PaymentMethod{
		ID:               "pm-12121212",
		AccountID:        "accId",
		StripeCustomerID: "stripeCustId",
		Brand:            "bran",
		Last4:            "last4",
		Expiry:           "expiry",
		Status:           "status",
		CreatedAt:        time.Now(),
	}
	err = c.PaymentMethods.Create(pm)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.PaymentMethods.Get(pm.ID)
	if err != nil {
		t.Fatal(err)
	}

	err = c.PaymentMethods.Update(pm.ID, database.PaymentMethodUpdateParams{Brand: database.String("Matercard")})
	if err != nil {
		t.Fatal(err)
	}
	pmUpdated, err := c.PaymentMethods.Get(pm.ID)
	if err != nil {
		t.Fatal(err)
	}
	if pmUpdated.Brand != "Matercard" {
		t.Fail()
	}

	err = c.PaymentMethods.Delete(pm.ID)
	if err != nil {
		t.Fatal(err)
	}
}
