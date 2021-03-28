package types

import "time"

type StripeData struct {
	CustomerID        string            `json:"customerId" firestore:"customerId"`
	SubscriptionItems map[string]string `json:"subscriptionItems" firestore:"subscriptionItems"`
	SubscriptionID    string            `json:"subscriptionId" firestore:"subscriptionId"`
}

type Account struct {
	ID        string      `json:"id" firestore:"id"`
	Name      string      `json:"name" firestore:"name"`
	Email     string      `json:"email" firestore:"email"`
	Pro       bool        `json:"pro" firestore:"pro"`
	Internal  bool        `json:"internal" firestore:"internal"`
	Blocked   bool        `json:"blocked" firestore:"blocked"`
	Stripe    *StripeData `json:"stripe" firestore:"stripe"`
	CreatedAt time.Time   `json:"createdAt" firestore:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt" firestore:"updateAt"`
}
