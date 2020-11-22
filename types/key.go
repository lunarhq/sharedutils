package types

import "time"

type Key struct {
	ID          string     `bson:"_id" json:"id"`
	Name        string     `bson:"name" json:"name"`
	SecretToken string     `bson:"secretToken" json:"secretToken"`
	LastUsed    *time.Time `bson:"lastUsed" json:"lastUsed"`
	Created     time.Time  `bson:"created" json:"created"`
	AccountID   string     `bson:"accountId" json:"accountId"`
	Status      string     `bson:"status" json:"status"`
	Pro         bool       `bson:"pro" json:"pro"`
}
