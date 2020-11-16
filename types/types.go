package types

import (
	"github.com/coinbase/rosetta-sdk-go/types"
)

type RosettaRequest struct {
	NetworkIdentifier      *types.NetworkIdentifier     `json:"network_identifier,omitempty"`
	AccountIdentifier      *types.AccountIdentifier     `json:"account_identifier,omitempty"`
	BlockIdentifier        *types.BlockIdentifier       `json:"block_identifier,omitempty"`
	TransactionIdentifier  *types.TransactionIdentifier `json:"transaction_identifier,omitempty"`
	PublicKey              *types.PublicKey             `json:"public_key,omitempty"`
	PublicKeys             []types.PublicKey            `json:"public_keys,omitempty"`
	SignedTransaction      string                       `json:"signed_transaction,omitempty"`
	UnsignedTransaction    string                       `json:"unsigned_transaction,omitempty"`
	Signatures             []types.Signature            `json:"signatures,omitempty"`
	Signed                 bool                         `json:"signed,omitempty"`
	Transaction            string                       `json:"transaction,omitempty"`
	Operations             []types.Operation            `json:"operations,omitempty"`
	SuggestedFeeMultiplier *float64                     `json:"suggested_fee_multiplier,omitempty"`
	MaxFee                 []types.Amount               `json:"max_fee,omitempty"`
	Options                map[string]interface{}       `json:"options,omitempty"`
	Metadata               map[string]interface{}       `json:"metadata,omitempty"`
	IncludeMempool         bool                         `json:"include_mempool,omitempty"`
	Currencies             []types.Currency             `json:"currencies,omitempty"`
}

//@Todo in future we can also put requesters headers/ip so we can use that for say Blacklisting etc.
//This is the payload that is sent whenever there is an api request on api-gw
type PubSubApiRequestPayload struct {
	ID         string          `json:"id,omitempty"`
	StatusCode int             `json:"status,omitempty"`
	Request    *RosettaRequest `json:"request,omitempty"`
	Response   interface{}     `json:"response,omitempty"`
	ApiKey     string          `json:"key,omitempty"`
	Timestamp  int64           `json:"timestamp,omitempty"`
	Path       string          `json:"path,omitempty"`
}
