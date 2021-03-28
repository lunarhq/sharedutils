package types

import "time"

type PaymentMethod struct {
	ID               string    `firestore:"-" json:"id"`
	AccountID        string    `firestore:"accountId" json:"accountId"`
	StripeCustomerID string    `firestore:"stripeCustomerId" json:"-"`
	Brand            string    `firestore:"brand" json:"brand"`
	Last4            string    `firestore:"last4" json:"last4"`
	Expiry           string    `firestore:"expiry" json:"expiry"`
	Status           string    `firestore:"status" json:"status"`
	UpdatedAt        time.Time `firestore:"updatedAt" json:"updatedAt"`
	CreatedAt        time.Time `firestore:"createdAt" json:"createdAt"`
}
