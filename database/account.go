package database

type AccountCreateParams struct {
	ID                      *string
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
