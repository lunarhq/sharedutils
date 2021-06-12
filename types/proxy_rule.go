package types

import "time"

type ProxyRule struct {
	ID        string    `firestore:"-" json:"id"`
	Name      string    `firestore:"name" json:"name"`
	AccountID string    `firestore:"accountId" json:"accountId"`
	Created   time.Time `firestore:"createdAt" json:"created"`

	Endpoint             string `firestore:"endpoint" json:"endpoint"`
	HeaderApiKeyKey      string `firestore:"headerApiKeyKey" json:"headerApiKeyKey"`
	HeaderApiKeyValue    string `firestore:"headerApiKeyValue" json:"headerApiKeyValue"`
	HeaderApiSecretKey   string `firestore:"headerApiSecretKey" json:"headerApiSecretKey"`
	HeaderApiSecretValue string `firestore:"headerApiSecretValue" json:"headerApiSecretValue"`
	WhitelistedDomains   string `firestore:"whitelistedDomains" json:"whitelistedDomains"`
}
