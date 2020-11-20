package database

type PaymentMethod struct {
	ID               string `bson:"_id" json:"id"`
	AccountId        string `bson:"accountId" json:"accountId"`
	StripeCustomerId string `bson:"stripeCustomerId" json:"-"`
	Brand            string `bson:"brand" json:"brand"`
	Last4            string `bson:"last4" json:"last4"`
	Expiry           string `bson:"expiry" json:"expiry"`
	IsDefault        bool   `bson:"isDefault" json:"isDefault"`
	Status           string `bson:"status" json:"status"`
}

type PaymentMethodListParams struct {
	AccountID *string
}
