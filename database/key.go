package database

type KeyCreateParams struct {
	Name        *string
	SecretToken *string
	AccountID   *string
	Status      *string
	Pro         *bool
}

type KeyUpdateParams struct {
	Name        *string
	SecretToken *string
	AccountID   *string
	Status      *string
	Pro         *bool
}

type KeyListParams struct {
	AccountID *string
}
