package database

type StripeData struct {
	CustomerId        string            `bson:"customerId" json:"customerId,omitempty"`
	SubscriptionItems map[string]string `bson:"subscriptionItems" json:"subscriptionItems,omitempty"`
	SubscriptionId    string            `bson:"subscriptionId" json:"subscriptionId,omitempty"`
}

type Account struct {
	ID       string      `bson:"_id" json:"id,omitempty"`
	Name     string      `bson:"name" json:"name,omitempty"`
	Email    string      `bson:"email" json:"email,omitempty"`
	Pro      bool        `bson:"pro" json:"pro"`
	Internal bool        `bson:"internal" json:"internal,omitempty"`
	Blocked  bool        `bson:"blocked" json:"blocked,omitempty"`
	Stripe   *StripeData `bson:"stripe" json:"-"`
}

type AccountUpdateParams struct {
	Name                    *string
	Pro                     *bool
	Internal                *bool
	Blocked                 *bool
	StripeCustomerId        *string
	StripeSubscriptionId    *string
	StripeSubscriptionItems *map[string]string
}
