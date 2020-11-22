package types

import "time"

type PaymentMethod struct {
	ID               string    `bson:"_id" json:"id"`
	AccountID        string    `bson:"accountId" json:"accountId"`
	StripeCustomerID string    `bson:"stripeCustomerId" json:"-"`
	Brand            string    `bson:"brand" json:"brand"`
	Last4            string    `bson:"last4" json:"last4"`
	Expiry           string    `bson:"expiry" json:"expiry"`
	Status           string    `bson:"status" json:"status"`
	UpdatedAt        time.Time `bson:"updatedAt" json:"updatedAt"`
	CreatedAt        time.Time `bson:"createdAt" json:"createdAt"`
}
