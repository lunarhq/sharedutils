package pubsub

//Key payload passed in msgs
type Key struct {
	ID          string `json:"id"`
	SecretToken string `json:"secretToken"`
	AccountID   string `json:"accountId"`
}
