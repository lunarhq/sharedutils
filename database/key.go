package database

import "time"

type Key struct {
	ID          string     `bson:"_id" json:"id,omitempty"`
	Name        string     `bson:"name" json:"name"`
	SecretToken string     `bson:"secretToken" json:"secretToken"`
	LastUsed    *time.Time `bson:"lastUsed" json:"lastUsed"`
	Created     time.Time  `bson:"created" json:"created"`
	AccountId   string     `bson:"accountId" json:"accountId"`
	Status      string     `bson:"status" json:"status"`
	Pro         bool       `bson:"pro" json:"pro"`
}

type KeyListParams struct {
	AccountID *string
}
