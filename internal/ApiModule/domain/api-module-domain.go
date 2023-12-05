// ApiModule/domain/api_types.go
package ApiModule

import (
	"math/big"
)

type GetTreasuryResponse struct {
	TokenID      uint64 `json:"token_id"`
	TotalValue   uint64 `json:"total_value"`
	APY          uint64 `json:"apy"`
	Duration     uint64 `json:"duration"`
	TreasuryType uint8  `json:"treasury_type"`
}

type Bigint struct {
	*big.Int
}

type TreasuryCreatedTable struct {
	TokenID      uint64
	APY          uint64
	Duration     uint64
	TreasuryType uint64
}

type PrimaryCreatedTable struct {
	TokenID   uint64 `json:"token_id"`
	Available uint64 `json:"available"`
	Price     string `json:"price"`
}

type SecondaryCreatedTable struct {
	TokenID   uint64 `json:"token_id"`
	Seller    string `json:"seller"`
	Available uint64 `json:"available"`
	Price     string `json:"price"`
}

type TransferCreatedTable struct {
	ID        uint64 `json:"id"`
	TxOrigin  string `json:"tx_origin"`
	Receiver  string `json:"receiver"`
	Amount    uint64 `json:"amount"`
	Timestamp string `json:"timestamp"`
}
