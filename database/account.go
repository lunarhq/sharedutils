package database

import "time"

type StripeData struct {
	CustomerID        string            `bson:"customerId" json:"customerId,omitempty"`
	SubscriptionItems map[string]string `bson:"subscriptionItems" json:"subscriptionItems,omitempty"`
	SubscriptionID    string            `bson:"subscriptionId" json:"subscriptionId,omitempty"`
}

type Account struct {
	ID        string      `bson:"_id" json:"id,omitempty"`
	Name      string      `bson:"name" json:"name,omitempty"`
	Email     string      `bson:"email" json:"email,omitempty"`
	Pro       bool        `bson:"pro" json:"pro"`
	Internal  bool        `bson:"internal" json:"internal,omitempty"`
	Blocked   bool        `bson:"blocked" json:"blocked,omitempty"`
	Stripe    *StripeData `bson:"stripe" json:"-"`
	CreatedAt time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time   `bson:"updatedAt" json:"updatedAt,omitempty"`
}

type AccountCreateParams struct {
	Name                    *string
	Email                   *string
	Pro                     *bool
	Internal                *bool
	Blocked                 *bool
	StripeCustomerID        *string
	StripeSubscriptionID    *string
	StripeSubscriptionItems *map[string]string
}

type AccountUpdateParams struct {
	Name                    *string
	Pro                     *bool
	Internal                *bool
	Blocked                 *bool
	StripeCustomerID        *string
	StripeSubscriptionID    *string
	StripeSubscriptionItems *map[string]string
}
