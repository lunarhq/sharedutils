package types

type Usage struct {
	AccountID string `json:"accountId" firestore:"accountId"`
	Date      string `json:"date" firestore:"date"`
	Mainnet   int64  `json:"mainnet" firestore:"mainnet"`
	Testnet   int64  `json:"testnet" firestore:"testnet"`
}
