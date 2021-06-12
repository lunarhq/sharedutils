package client

import (
	"context"
	"fmt"
	"log"

	"github.com/lunarhq/sharedutils/database/account"
	"github.com/lunarhq/sharedutils/database/key"
	"github.com/lunarhq/sharedutils/database/paymentmethod"
	"github.com/lunarhq/sharedutils/database/proxy_rule"
	"github.com/lunarhq/sharedutils/database/usage"

	firebase "firebase.google.com/go/v4"
)

type DB struct {
	Keys           *key.Client
	Accounts       *account.Client
	PaymentMethods *paymentmethod.Client
	Usage          *usage.Client
	ProxyRules     *proxy_rule.Client
}

func New(ctx context.Context) (*DB, error) {
	fb, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatal(fmt.Errorf("error initializing Firebase: %v", err))
	}

	firestoreClient, err := fb.Firestore(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("error initializing Firstore: %v", err))
	}

	d := &DB{
		Keys:           &key.Client{firestoreClient, ctx},
		Accounts:       &account.Client{firestoreClient, ctx},
		PaymentMethods: &paymentmethod.Client{firestoreClient, ctx},
		Usage:          &usage.Client{firestoreClient, ctx},
		ProxyRules:     &proxy_rule.Client{firestoreClient, ctx},
	}

	return d, nil
}
