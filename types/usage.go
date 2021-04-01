package types

type Usage struct {
	ID         string `json:"id" firestore:"id"`
	AccountID  string `json:"accountId" firestore:"accountId"`
	Key        string `json:"key" firestore:"key"`
	Blockchain string `json:"blockchain" firestore:"blockchain"`
	Network    string `json:"network" firestore:"network"`
	IsMainnet  bool   `json:"isMainnet" firestore:"isMainnet"`
	Date       string `json:"date" firestore:"date"`
	Hits       int64  `json:"hits" firestore:"hits"`
}
