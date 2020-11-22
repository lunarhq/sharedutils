package types

import "time"

type StripeData struct {
	CustomerID        string            `bson:"customerId" json:"customerId"`
	SubscriptionItems map[string]string `bson:"subscriptionItems" json:"subscriptionItems"`
	SubscriptionID    string            `bson:"subscriptionId" json:"subscriptionId"`
}

type Account struct {
	ID        string      `bson:"_id" json:"id"`
	Name      string      `bson:"name" json:"name"`
	Email     string      `bson:"email" json:"email"`
	Pro       bool        `bson:"pro" json:"pro"`
	Internal  bool        `bson:"internal" json:"internal"`
	Blocked   bool        `bson:"blocked" json:"blocked"`
	Stripe    *StripeData `bson:"stripe" json:"-"`
	CreatedAt time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time   `bson:"updatedAt" json:"updatedAt"`
}
