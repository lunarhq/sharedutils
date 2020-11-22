package database

type PaymentMethodUpdateParams struct {
	AccountID        *string
	StripeCustomerID *string
	Brand            *string
	Last4            *string
	Expiry           *string
	Status           *string
}

type PaymentMethodListParams struct {
	AccountID *string
}

type PaymentMethodGetParams struct {
	AccountID *string
}
