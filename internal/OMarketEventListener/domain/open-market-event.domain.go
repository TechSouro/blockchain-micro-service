package OMarketEventListener

import "github.com/ethereum/go-ethereum/common"

/**
* @dev soldity events signature
* event publicOrderCreated(uint256 indexed _tokenId,uint256 indexed _price, uint256 indexed _available);
	event primarySale(address indexed _sender, uint256 indexed _tokenId, uint256 indexed _amount);
	event secondaryForSale(
		address indexed _seller,
		uint256 indexed _tokenId,
		uint256 indexed _units,
		uint256 _price
		);
	event secondarySold(address indexed _seller, address indexed _buyer, uint256 indexed _units, uint
*/

type PublicOrderCreated struct {
	TokenID   uint64
	Price     uint64
	Available uint64
}

type PrimarySale struct {
	Sender  common.Address
	TokenID uint64
	Amount  uint64
}

type SecondaryForSale struct {
	Seller  common.Address
	TokenID uint64
	Units   uint64
	Price   uint64
}

type SecondarySold struct {
	Seller common.Address
	Buyer  common.Address
	Units  uint64
	Price  uint64
}

type EventListener interface {
	ListenPublicOrderCreated() (<-chan PublicOrderCreated, error)
	ListenPrimarySale() (<-chan PrimarySale, error)
	ListenSecondaryForSale() (<-chan SecondaryForSale, error)
	ListenSecondarySold() (<-chan SecondarySold, error)
}
