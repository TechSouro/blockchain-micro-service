// ApiModule/domain/api_types.go
package ApiModule

type GetTreasuryResponse struct {
	TokenID      uint64 `json:"token_id"`
	TotalValue   uint64 `json:"total_value"`
	APY          uint64 `json:"apy"`
	Duration     uint64 `json:"duration"`
	TreasuryType uint8  `json:"treasury_type"`
}

type TreasuryCreatedTable struct {
	TokenID      uint64 `gorm:"column:token_id"`
	TotalValue   uint64 `gorm:"column:total_value"`
	APY          uint64 `gorm:"column:apy"`
	Duration     uint64 `gorm:"column:duration"`
	TreasuryType uint8  `gorm:"column:treasury_type"`
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
