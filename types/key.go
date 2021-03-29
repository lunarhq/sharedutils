package types

import "time"

type Key struct {
	ID          string     `firestore:"-" json:"id"`
	Name        string     `firestore:"name" json:"name"`
	SecretToken string     `firestore:"secretToken" json:"secretToken"`
	LastUsed    *time.Time `firestore:"lastUsed" json:"lastUsed"`
	Created     time.Time  `firestore:"createdAt" json:"created"`
	AccountID   string     `firestore:"accountId" json:"accountId"`
	Status      string     `firestore:"status" json:"status"`
	Pro         bool       `firestore:"pro" json:"pro"`
}
