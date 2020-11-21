package database

import "time"

type Key struct {
	ID          string     `bson:"_id" json:"id,omitempty"`
	Name        string     `bson:"name" json:"name"`
	SecretToken string     `bson:"secretToken" json:"secretToken"`
	LastUsed    *time.Time `bson:"lastUsed" json:"lastUsed"`
	Created     time.Time  `bson:"created" json:"created"`
	AccountID   string     `bson:"accountId" json:"accountId"`
	Status      string     `bson:"status" json:"status"`
	Pro         bool       `bson:"pro" json:"pro"`
}

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
