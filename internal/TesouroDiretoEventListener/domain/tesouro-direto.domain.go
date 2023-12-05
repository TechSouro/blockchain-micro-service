package OMarketEventListener

type treasuryType int

// Definindo constantes para representar os diferentes tipos de tesouro
const (
	LTN treasuryType = iota
	NTN_F
	LFT
	NTN_B_MAIN
	NTN_B
)

type TreasuryCreatedEvent struct {
	TokenID      uint64
	TotalValue   uint64
	APY          uint64
	Duration     uint64
	TreasuryType treasuryType
}

type EventListener interface {
	ListenEmitTresure() (<-chan TreasuryCreatedEvent, error)
}
