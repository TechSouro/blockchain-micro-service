// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pkg

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TesouroDiretotreasuryData is an auto generated low-level Go binding around an user-defined struct.
type TesouroDiretotreasuryData struct {
	Type          uint8
	Apy           *big.Int
	MinInvestment *big.Int
	ValidThru     *big.Int
	AvlbTokens    *big.Int
	Creation      *big.Int
}

// TesouroDiretoMetaData contains all meta data concerning the TesouroDireto contract.
var TesouroDiretoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_openMarket\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wdrex\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotKYCed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidTreasuryType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"notOpenMarketContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_totalValue\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_apy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumtesouroDireto.treasuryType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"treasuryCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumtesouroDireto.treasuryType\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"_apy\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"_minInvestment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validThru\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_avlbTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creation\",\"type\":\"uint256\"}],\"internalType\":\"structtesouroDireto.treasuryData\",\"name\":\"_data\",\"type\":\"tuple\"}],\"name\":\"emitTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getPriceAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_apy\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"_validThru\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"openPublicOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"retrievePartialInvestment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"retriveFullInvestment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newemmit\",\"type\":\"address\"}],\"name\":\"setEmmiter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenInfo\",\"outputs\":[{\"internalType\":\"enumtesouroDireto.treasuryType\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"_apy\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"_minInvestment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validThru\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_avlbTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TesouroDiretoABI is the input ABI used to generate the binding from.
// Deprecated: Use TesouroDiretoMetaData.ABI instead.
var TesouroDiretoABI = TesouroDiretoMetaData.ABI

// TesouroDireto is an auto generated Go binding around an Ethereum contract.
type TesouroDireto struct {
	TesouroDiretoCaller     // Read-only binding to the contract
	TesouroDiretoTransactor // Write-only binding to the contract
	TesouroDiretoFilterer   // Log filterer for contract events
}

// TesouroDiretoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TesouroDiretoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TesouroDiretoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TesouroDiretoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TesouroDiretoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TesouroDiretoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TesouroDiretoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TesouroDiretoSession struct {
	Contract     *TesouroDireto    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TesouroDiretoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TesouroDiretoCallerSession struct {
	Contract *TesouroDiretoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TesouroDiretoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TesouroDiretoTransactorSession struct {
	Contract     *TesouroDiretoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TesouroDiretoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TesouroDiretoRaw struct {
	Contract *TesouroDireto // Generic contract binding to access the raw methods on
}

// TesouroDiretoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TesouroDiretoCallerRaw struct {
	Contract *TesouroDiretoCaller // Generic read-only contract binding to access the raw methods on
}

// TesouroDiretoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TesouroDiretoTransactorRaw struct {
	Contract *TesouroDiretoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTesouroDireto creates a new instance of TesouroDireto, bound to a specific deployed contract.
func NewTesouroDireto(address common.Address, backend bind.ContractBackend) (*TesouroDireto, error) {
	contract, err := bindTesouroDireto(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TesouroDireto{TesouroDiretoCaller: TesouroDiretoCaller{contract: contract}, TesouroDiretoTransactor: TesouroDiretoTransactor{contract: contract}, TesouroDiretoFilterer: TesouroDiretoFilterer{contract: contract}}, nil
}

// NewTesouroDiretoCaller creates a new read-only instance of TesouroDireto, bound to a specific deployed contract.
func NewTesouroDiretoCaller(address common.Address, caller bind.ContractCaller) (*TesouroDiretoCaller, error) {
	contract, err := bindTesouroDireto(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TesouroDiretoCaller{contract: contract}, nil
}

// NewTesouroDiretoTransactor creates a new write-only instance of TesouroDireto, bound to a specific deployed contract.
func NewTesouroDiretoTransactor(address common.Address, transactor bind.ContractTransactor) (*TesouroDiretoTransactor, error) {
	contract, err := bindTesouroDireto(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TesouroDiretoTransactor{contract: contract}, nil
}

// NewTesouroDiretoFilterer creates a new log filterer instance of TesouroDireto, bound to a specific deployed contract.
func NewTesouroDiretoFilterer(address common.Address, filterer bind.ContractFilterer) (*TesouroDiretoFilterer, error) {
	contract, err := bindTesouroDireto(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TesouroDiretoFilterer{contract: contract}, nil
}

// bindTesouroDireto binds a generic wrapper to an already deployed contract.
func bindTesouroDireto(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TesouroDiretoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TesouroDireto *TesouroDiretoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TesouroDireto.Contract.TesouroDiretoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TesouroDireto *TesouroDiretoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TesouroDireto.Contract.TesouroDiretoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TesouroDireto *TesouroDiretoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TesouroDireto.Contract.TesouroDiretoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TesouroDireto *TesouroDiretoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TesouroDireto.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TesouroDireto *TesouroDiretoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TesouroDireto.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TesouroDireto *TesouroDiretoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TesouroDireto.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TesouroDireto *TesouroDiretoCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TesouroDireto *TesouroDiretoSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TesouroDireto.Contract.BalanceOf(&_TesouroDireto.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TesouroDireto *TesouroDiretoCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TesouroDireto.Contract.BalanceOf(&_TesouroDireto.CallOpts, owner)
}

// Emitter is a free data retrieval call binding the contract method 0xdce11375.
//
// Solidity: function emitter() view returns(address)
func (_TesouroDireto *TesouroDiretoCaller) Emitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "emitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Emitter is a free data retrieval call binding the contract method 0xdce11375.
//
// Solidity: function emitter() view returns(address)
func (_TesouroDireto *TesouroDiretoSession) Emitter() (common.Address, error) {
	return _TesouroDireto.Contract.Emitter(&_TesouroDireto.CallOpts)
}

// Emitter is a free data retrieval call binding the contract method 0xdce11375.
//
// Solidity: function emitter() view returns(address)
func (_TesouroDireto *TesouroDiretoCallerSession) Emitter() (common.Address, error) {
	return _TesouroDireto.Contract.Emitter(&_TesouroDireto.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TesouroDireto *TesouroDiretoCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TesouroDireto *TesouroDiretoSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TesouroDireto.Contract.GetApproved(&_TesouroDireto.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TesouroDireto *TesouroDiretoCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TesouroDireto.Contract.GetApproved(&_TesouroDireto.CallOpts, tokenId)
}

// GetPriceAmount is a free data retrieval call binding the contract method 0xb176c59d.
//
// Solidity: function getPriceAmount(uint256 _tokenId) view returns(uint256 _price, uint256 _amount)
func (_TesouroDireto *TesouroDiretoCaller) GetPriceAmount(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Price  *big.Int
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "getPriceAmount", _tokenId)

	outstruct := new(struct {
		Price  *big.Int
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPriceAmount is a free data retrieval call binding the contract method 0xb176c59d.
//
// Solidity: function getPriceAmount(uint256 _tokenId) view returns(uint256 _price, uint256 _amount)
func (_TesouroDireto *TesouroDiretoSession) GetPriceAmount(_tokenId *big.Int) (struct {
	Price  *big.Int
	Amount *big.Int
}, error) {
	return _TesouroDireto.Contract.GetPriceAmount(&_TesouroDireto.CallOpts, _tokenId)
}

// GetPriceAmount is a free data retrieval call binding the contract method 0xb176c59d.
//
// Solidity: function getPriceAmount(uint256 _tokenId) view returns(uint256 _price, uint256 _amount)
func (_TesouroDireto *TesouroDiretoCallerSession) GetPriceAmount(_tokenId *big.Int) (struct {
	Price  *big.Int
	Amount *big.Int
}, error) {
	return _TesouroDireto.Contract.GetPriceAmount(&_TesouroDireto.CallOpts, _tokenId)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 _tokenId) view returns(uint24 _apy, uint256 _validThru, uint256 _price, uint256 _creation)
func (_TesouroDireto *TesouroDiretoCaller) GetTokenInfo(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Apy       *big.Int
	ValidThru *big.Int
	Price     *big.Int
	Creation  *big.Int
}, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "getTokenInfo", _tokenId)

	outstruct := new(struct {
		Apy       *big.Int
		ValidThru *big.Int
		Price     *big.Int
		Creation  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Apy = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidThru = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Creation = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 _tokenId) view returns(uint24 _apy, uint256 _validThru, uint256 _price, uint256 _creation)
func (_TesouroDireto *TesouroDiretoSession) GetTokenInfo(_tokenId *big.Int) (struct {
	Apy       *big.Int
	ValidThru *big.Int
	Price     *big.Int
	Creation  *big.Int
}, error) {
	return _TesouroDireto.Contract.GetTokenInfo(&_TesouroDireto.CallOpts, _tokenId)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 _tokenId) view returns(uint24 _apy, uint256 _validThru, uint256 _price, uint256 _creation)
func (_TesouroDireto *TesouroDiretoCallerSession) GetTokenInfo(_tokenId *big.Int) (struct {
	Apy       *big.Int
	ValidThru *big.Int
	Price     *big.Int
	Creation  *big.Int
}, error) {
	return _TesouroDireto.Contract.GetTokenInfo(&_TesouroDireto.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TesouroDireto *TesouroDiretoCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TesouroDireto *TesouroDiretoSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TesouroDireto.Contract.IsApprovedForAll(&_TesouroDireto.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TesouroDireto *TesouroDiretoCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TesouroDireto.Contract.IsApprovedForAll(&_TesouroDireto.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TesouroDireto *TesouroDiretoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TesouroDireto *TesouroDiretoSession) Name() (string, error) {
	return _TesouroDireto.Contract.Name(&_TesouroDireto.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TesouroDireto *TesouroDiretoCallerSession) Name() (string, error) {
	return _TesouroDireto.Contract.Name(&_TesouroDireto.CallOpts)
}

// OpenMarket is a free data retrieval call binding the contract method 0x3606f5b9.
//
// Solidity: function openMarket() view returns(address)
func (_TesouroDireto *TesouroDiretoCaller) OpenMarket(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "openMarket")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OpenMarket is a free data retrieval call binding the contract method 0x3606f5b9.
//
// Solidity: function openMarket() view returns(address)
func (_TesouroDireto *TesouroDiretoSession) OpenMarket() (common.Address, error) {
	return _TesouroDireto.Contract.OpenMarket(&_TesouroDireto.CallOpts)
}

// OpenMarket is a free data retrieval call binding the contract method 0x3606f5b9.
//
// Solidity: function openMarket() view returns(address)
func (_TesouroDireto *TesouroDiretoCallerSession) OpenMarket() (common.Address, error) {
	return _TesouroDireto.Contract.OpenMarket(&_TesouroDireto.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TesouroDireto *TesouroDiretoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TesouroDireto *TesouroDiretoSession) Owner() (common.Address, error) {
	return _TesouroDireto.Contract.Owner(&_TesouroDireto.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TesouroDireto *TesouroDiretoCallerSession) Owner() (common.Address, error) {
	return _TesouroDireto.Contract.Owner(&_TesouroDireto.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TesouroDireto *TesouroDiretoCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TesouroDireto *TesouroDiretoSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TesouroDireto.Contract.OwnerOf(&_TesouroDireto.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TesouroDireto *TesouroDiretoCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TesouroDireto.Contract.OwnerOf(&_TesouroDireto.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TesouroDireto *TesouroDiretoCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TesouroDireto *TesouroDiretoSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TesouroDireto.Contract.SupportsInterface(&_TesouroDireto.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TesouroDireto *TesouroDiretoCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TesouroDireto.Contract.SupportsInterface(&_TesouroDireto.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TesouroDireto *TesouroDiretoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TesouroDireto *TesouroDiretoSession) Symbol() (string, error) {
	return _TesouroDireto.Contract.Symbol(&_TesouroDireto.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TesouroDireto *TesouroDiretoCallerSession) Symbol() (string, error) {
	return _TesouroDireto.Contract.Symbol(&_TesouroDireto.CallOpts)
}

// TokenInfo is a free data retrieval call binding the contract method 0xcc33c875.
//
// Solidity: function tokenInfo(uint256 ) view returns(uint8 _type, uint24 _apy, uint256 _minInvestment, uint256 _validThru, uint256 _avlbTokens, uint256 _creation)
func (_TesouroDireto *TesouroDiretoCaller) TokenInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Type          uint8
	Apy           *big.Int
	MinInvestment *big.Int
	ValidThru     *big.Int
	AvlbTokens    *big.Int
	Creation      *big.Int
}, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "tokenInfo", arg0)

	outstruct := new(struct {
		Type          uint8
		Apy           *big.Int
		MinInvestment *big.Int
		ValidThru     *big.Int
		AvlbTokens    *big.Int
		Creation      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Type = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Apy = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MinInvestment = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidThru = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AvlbTokens = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Creation = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenInfo is a free data retrieval call binding the contract method 0xcc33c875.
//
// Solidity: function tokenInfo(uint256 ) view returns(uint8 _type, uint24 _apy, uint256 _minInvestment, uint256 _validThru, uint256 _avlbTokens, uint256 _creation)
func (_TesouroDireto *TesouroDiretoSession) TokenInfo(arg0 *big.Int) (struct {
	Type          uint8
	Apy           *big.Int
	MinInvestment *big.Int
	ValidThru     *big.Int
	AvlbTokens    *big.Int
	Creation      *big.Int
}, error) {
	return _TesouroDireto.Contract.TokenInfo(&_TesouroDireto.CallOpts, arg0)
}

// TokenInfo is a free data retrieval call binding the contract method 0xcc33c875.
//
// Solidity: function tokenInfo(uint256 ) view returns(uint8 _type, uint24 _apy, uint256 _minInvestment, uint256 _validThru, uint256 _avlbTokens, uint256 _creation)
func (_TesouroDireto *TesouroDiretoCallerSession) TokenInfo(arg0 *big.Int) (struct {
	Type          uint8
	Apy           *big.Int
	MinInvestment *big.Int
	ValidThru     *big.Int
	AvlbTokens    *big.Int
	Creation      *big.Int
}, error) {
	return _TesouroDireto.Contract.TokenInfo(&_TesouroDireto.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TesouroDireto *TesouroDiretoCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TesouroDireto *TesouroDiretoSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TesouroDireto.Contract.TokenURI(&_TesouroDireto.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TesouroDireto *TesouroDiretoCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TesouroDireto.Contract.TokenURI(&_TesouroDireto.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TesouroDireto *TesouroDiretoCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TesouroDireto.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TesouroDireto *TesouroDiretoSession) TotalSupply() (*big.Int, error) {
	return _TesouroDireto.Contract.TotalSupply(&_TesouroDireto.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TesouroDireto *TesouroDiretoCallerSession) TotalSupply() (*big.Int, error) {
	return _TesouroDireto.Contract.TotalSupply(&_TesouroDireto.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_TesouroDireto *TesouroDiretoTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_TesouroDireto *TesouroDiretoSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.Approve(&_TesouroDireto.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.Approve(&_TesouroDireto.TransactOpts, to, tokenId)
}

// EmitTreasury is a paid mutator transaction binding the contract method 0xa768ef65.
//
// Solidity: function emitTreasury((uint8,uint24,uint256,uint256,uint256,uint256) _data) returns()
func (_TesouroDireto *TesouroDiretoTransactor) EmitTreasury(opts *bind.TransactOpts, _data TesouroDiretotreasuryData) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "emitTreasury", _data)
}

// EmitTreasury is a paid mutator transaction binding the contract method 0xa768ef65.
//
// Solidity: function emitTreasury((uint8,uint24,uint256,uint256,uint256,uint256) _data) returns()
func (_TesouroDireto *TesouroDiretoSession) EmitTreasury(_data TesouroDiretotreasuryData) (*types.Transaction, error) {
	return _TesouroDireto.Contract.EmitTreasury(&_TesouroDireto.TransactOpts, _data)
}

// EmitTreasury is a paid mutator transaction binding the contract method 0xa768ef65.
//
// Solidity: function emitTreasury((uint8,uint24,uint256,uint256,uint256,uint256) _data) returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) EmitTreasury(_data TesouroDiretotreasuryData) (*types.Transaction, error) {
	return _TesouroDireto.Contract.EmitTreasury(&_TesouroDireto.TransactOpts, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_TesouroDireto *TesouroDiretoTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_TesouroDireto *TesouroDiretoSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TesouroDireto.Contract.OnERC721Received(&_TesouroDireto.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_TesouroDireto *TesouroDiretoTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TesouroDireto.Contract.OnERC721Received(&_TesouroDireto.TransactOpts, operator, from, tokenId, data)
}

// OpenPublicOffer is a paid mutator transaction binding the contract method 0x5cd0ea89.
//
// Solidity: function openPublicOffer(uint256 _tokenId) returns()
func (_TesouroDireto *TesouroDiretoTransactor) OpenPublicOffer(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "openPublicOffer", _tokenId)
}

// OpenPublicOffer is a paid mutator transaction binding the contract method 0x5cd0ea89.
//
// Solidity: function openPublicOffer(uint256 _tokenId) returns()
func (_TesouroDireto *TesouroDiretoSession) OpenPublicOffer(_tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.OpenPublicOffer(&_TesouroDireto.TransactOpts, _tokenId)
}

// OpenPublicOffer is a paid mutator transaction binding the contract method 0x5cd0ea89.
//
// Solidity: function openPublicOffer(uint256 _tokenId) returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) OpenPublicOffer(_tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.OpenPublicOffer(&_TesouroDireto.TransactOpts, _tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TesouroDireto *TesouroDiretoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TesouroDireto *TesouroDiretoSession) RenounceOwnership() (*types.Transaction, error) {
	return _TesouroDireto.Contract.RenounceOwnership(&_TesouroDireto.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TesouroDireto.Contract.RenounceOwnership(&_TesouroDireto.TransactOpts)
}

// RetrievePartialInvestment is a paid mutator transaction binding the contract method 0xfdddaff3.
//
// Solidity: function retrievePartialInvestment(uint256 _tokenId, uint256 _amount) returns(bool)
func (_TesouroDireto *TesouroDiretoTransactor) RetrievePartialInvestment(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "retrievePartialInvestment", _tokenId, _amount)
}

// RetrievePartialInvestment is a paid mutator transaction binding the contract method 0xfdddaff3.
//
// Solidity: function retrievePartialInvestment(uint256 _tokenId, uint256 _amount) returns(bool)
func (_TesouroDireto *TesouroDiretoSession) RetrievePartialInvestment(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.RetrievePartialInvestment(&_TesouroDireto.TransactOpts, _tokenId, _amount)
}

// RetrievePartialInvestment is a paid mutator transaction binding the contract method 0xfdddaff3.
//
// Solidity: function retrievePartialInvestment(uint256 _tokenId, uint256 _amount) returns(bool)
func (_TesouroDireto *TesouroDiretoTransactorSession) RetrievePartialInvestment(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.RetrievePartialInvestment(&_TesouroDireto.TransactOpts, _tokenId, _amount)
}

// RetriveFullInvestment is a paid mutator transaction binding the contract method 0xd6ceaa70.
//
// Solidity: function retriveFullInvestment(uint256 _tokenId, uint256 _amount) returns(bool)
func (_TesouroDireto *TesouroDiretoTransactor) RetriveFullInvestment(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "retriveFullInvestment", _tokenId, _amount)
}

// RetriveFullInvestment is a paid mutator transaction binding the contract method 0xd6ceaa70.
//
// Solidity: function retriveFullInvestment(uint256 _tokenId, uint256 _amount) returns(bool)
func (_TesouroDireto *TesouroDiretoSession) RetriveFullInvestment(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.RetriveFullInvestment(&_TesouroDireto.TransactOpts, _tokenId, _amount)
}

// RetriveFullInvestment is a paid mutator transaction binding the contract method 0xd6ceaa70.
//
// Solidity: function retriveFullInvestment(uint256 _tokenId, uint256 _amount) returns(bool)
func (_TesouroDireto *TesouroDiretoTransactorSession) RetriveFullInvestment(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.RetriveFullInvestment(&_TesouroDireto.TransactOpts, _tokenId, _amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_TesouroDireto *TesouroDiretoTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_TesouroDireto *TesouroDiretoSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.SafeTransferFrom(&_TesouroDireto.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.SafeTransferFrom(&_TesouroDireto.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_TesouroDireto *TesouroDiretoTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_TesouroDireto *TesouroDiretoSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TesouroDireto.Contract.SafeTransferFrom0(&_TesouroDireto.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TesouroDireto.Contract.SafeTransferFrom0(&_TesouroDireto.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TesouroDireto *TesouroDiretoTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TesouroDireto *TesouroDiretoSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TesouroDireto.Contract.SetApprovalForAll(&_TesouroDireto.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TesouroDireto.Contract.SetApprovalForAll(&_TesouroDireto.TransactOpts, operator, approved)
}

// SetEmmiter is a paid mutator transaction binding the contract method 0x181b2235.
//
// Solidity: function setEmmiter(address _newemmit) returns()
func (_TesouroDireto *TesouroDiretoTransactor) SetEmmiter(opts *bind.TransactOpts, _newemmit common.Address) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "setEmmiter", _newemmit)
}

// SetEmmiter is a paid mutator transaction binding the contract method 0x181b2235.
//
// Solidity: function setEmmiter(address _newemmit) returns()
func (_TesouroDireto *TesouroDiretoSession) SetEmmiter(_newemmit common.Address) (*types.Transaction, error) {
	return _TesouroDireto.Contract.SetEmmiter(&_TesouroDireto.TransactOpts, _newemmit)
}

// SetEmmiter is a paid mutator transaction binding the contract method 0x181b2235.
//
// Solidity: function setEmmiter(address _newemmit) returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) SetEmmiter(_newemmit common.Address) (*types.Transaction, error) {
	return _TesouroDireto.Contract.SetEmmiter(&_TesouroDireto.TransactOpts, _newemmit)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_TesouroDireto *TesouroDiretoTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_TesouroDireto *TesouroDiretoSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.TransferFrom(&_TesouroDireto.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TesouroDireto.Contract.TransferFrom(&_TesouroDireto.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TesouroDireto *TesouroDiretoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TesouroDireto.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TesouroDireto *TesouroDiretoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TesouroDireto.Contract.TransferOwnership(&_TesouroDireto.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TesouroDireto *TesouroDiretoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TesouroDireto.Contract.TransferOwnership(&_TesouroDireto.TransactOpts, newOwner)
}

// TesouroDiretoApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TesouroDireto contract.
type TesouroDiretoApprovalIterator struct {
	Event *TesouroDiretoApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TesouroDiretoApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TesouroDiretoApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TesouroDiretoApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TesouroDiretoApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TesouroDiretoApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TesouroDiretoApproval represents a Approval event raised by the TesouroDireto contract.
type TesouroDiretoApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TesouroDireto *TesouroDiretoFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TesouroDiretoApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TesouroDireto.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TesouroDiretoApprovalIterator{contract: _TesouroDireto.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TesouroDireto *TesouroDiretoFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TesouroDiretoApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TesouroDireto.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TesouroDiretoApproval)
				if err := _TesouroDireto.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TesouroDireto *TesouroDiretoFilterer) ParseApproval(log types.Log) (*TesouroDiretoApproval, error) {
	event := new(TesouroDiretoApproval)
	if err := _TesouroDireto.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TesouroDiretoApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TesouroDireto contract.
type TesouroDiretoApprovalForAllIterator struct {
	Event *TesouroDiretoApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TesouroDiretoApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TesouroDiretoApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TesouroDiretoApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TesouroDiretoApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TesouroDiretoApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TesouroDiretoApprovalForAll represents a ApprovalForAll event raised by the TesouroDireto contract.
type TesouroDiretoApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TesouroDireto *TesouroDiretoFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TesouroDiretoApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TesouroDireto.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TesouroDiretoApprovalForAllIterator{contract: _TesouroDireto.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TesouroDireto *TesouroDiretoFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TesouroDiretoApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TesouroDireto.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TesouroDiretoApprovalForAll)
				if err := _TesouroDireto.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TesouroDireto *TesouroDiretoFilterer) ParseApprovalForAll(log types.Log) (*TesouroDiretoApprovalForAll, error) {
	event := new(TesouroDiretoApprovalForAll)
	if err := _TesouroDireto.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TesouroDiretoConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the TesouroDireto contract.
type TesouroDiretoConsecutiveTransferIterator struct {
	Event *TesouroDiretoConsecutiveTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TesouroDiretoConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TesouroDiretoConsecutiveTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TesouroDiretoConsecutiveTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TesouroDiretoConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TesouroDiretoConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TesouroDiretoConsecutiveTransfer represents a ConsecutiveTransfer event raised by the TesouroDireto contract.
type TesouroDiretoConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_TesouroDireto *TesouroDiretoFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*TesouroDiretoConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TesouroDireto.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TesouroDiretoConsecutiveTransferIterator{contract: _TesouroDireto.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_TesouroDireto *TesouroDiretoFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *TesouroDiretoConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TesouroDireto.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TesouroDiretoConsecutiveTransfer)
				if err := _TesouroDireto.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_TesouroDireto *TesouroDiretoFilterer) ParseConsecutiveTransfer(log types.Log) (*TesouroDiretoConsecutiveTransfer, error) {
	event := new(TesouroDiretoConsecutiveTransfer)
	if err := _TesouroDireto.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TesouroDiretoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TesouroDireto contract.
type TesouroDiretoOwnershipTransferredIterator struct {
	Event *TesouroDiretoOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TesouroDiretoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TesouroDiretoOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TesouroDiretoOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TesouroDiretoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TesouroDiretoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TesouroDiretoOwnershipTransferred represents a OwnershipTransferred event raised by the TesouroDireto contract.
type TesouroDiretoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TesouroDireto *TesouroDiretoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TesouroDiretoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TesouroDireto.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TesouroDiretoOwnershipTransferredIterator{contract: _TesouroDireto.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TesouroDireto *TesouroDiretoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TesouroDiretoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TesouroDireto.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TesouroDiretoOwnershipTransferred)
				if err := _TesouroDireto.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TesouroDireto *TesouroDiretoFilterer) ParseOwnershipTransferred(log types.Log) (*TesouroDiretoOwnershipTransferred, error) {
	event := new(TesouroDiretoOwnershipTransferred)
	if err := _TesouroDireto.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TesouroDiretoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TesouroDireto contract.
type TesouroDiretoTransferIterator struct {
	Event *TesouroDiretoTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TesouroDiretoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TesouroDiretoTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TesouroDiretoTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TesouroDiretoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TesouroDiretoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TesouroDiretoTransfer represents a Transfer event raised by the TesouroDireto contract.
type TesouroDiretoTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TesouroDireto *TesouroDiretoFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TesouroDiretoTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TesouroDireto.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TesouroDiretoTransferIterator{contract: _TesouroDireto.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TesouroDireto *TesouroDiretoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TesouroDiretoTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TesouroDireto.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TesouroDiretoTransfer)
				if err := _TesouroDireto.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TesouroDireto *TesouroDiretoFilterer) ParseTransfer(log types.Log) (*TesouroDiretoTransfer, error) {
	event := new(TesouroDiretoTransfer)
	if err := _TesouroDireto.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TesouroDiretoTreasuryCreatedIterator is returned from FilterTreasuryCreated and is used to iterate over the raw logs and unpacked data for TreasuryCreated events raised by the TesouroDireto contract.
type TesouroDiretoTreasuryCreatedIterator struct {
	Event *TesouroDiretoTreasuryCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TesouroDiretoTreasuryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TesouroDiretoTreasuryCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TesouroDiretoTreasuryCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TesouroDiretoTreasuryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TesouroDiretoTreasuryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TesouroDiretoTreasuryCreated represents a TreasuryCreated event raised by the TesouroDireto contract.
type TesouroDiretoTreasuryCreated struct {
	TokenId    *big.Int
	TotalValue *big.Int
	Apy        *big.Int
	Duration   *big.Int
	Arg4       uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTreasuryCreated is a free log retrieval operation binding the contract event 0xf83eeebb88a585a0bf111d55674f1c4cff9319871b905a62e7ede317bc032149.
//
// Solidity: event treasuryCreated(uint256 indexed tokenId, uint256 indexed _totalValue, uint256 indexed _apy, uint256 _duration, uint8 arg4)
func (_TesouroDireto *TesouroDiretoFilterer) FilterTreasuryCreated(opts *bind.FilterOpts, tokenId []*big.Int, _totalValue []*big.Int, _apy []*big.Int) (*TesouroDiretoTreasuryCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var _totalValueRule []interface{}
	for _, _totalValueItem := range _totalValue {
		_totalValueRule = append(_totalValueRule, _totalValueItem)
	}
	var _apyRule []interface{}
	for _, _apyItem := range _apy {
		_apyRule = append(_apyRule, _apyItem)
	}

	logs, sub, err := _TesouroDireto.contract.FilterLogs(opts, "treasuryCreated", tokenIdRule, _totalValueRule, _apyRule)
	if err != nil {
		return nil, err
	}
	return &TesouroDiretoTreasuryCreatedIterator{contract: _TesouroDireto.contract, event: "treasuryCreated", logs: logs, sub: sub}, nil
}

// WatchTreasuryCreated is a free log subscription operation binding the contract event 0xf83eeebb88a585a0bf111d55674f1c4cff9319871b905a62e7ede317bc032149.
//
// Solidity: event treasuryCreated(uint256 indexed tokenId, uint256 indexed _totalValue, uint256 indexed _apy, uint256 _duration, uint8 arg4)
func (_TesouroDireto *TesouroDiretoFilterer) WatchTreasuryCreated(opts *bind.WatchOpts, sink chan<- *TesouroDiretoTreasuryCreated, tokenId []*big.Int, _totalValue []*big.Int, _apy []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var _totalValueRule []interface{}
	for _, _totalValueItem := range _totalValue {
		_totalValueRule = append(_totalValueRule, _totalValueItem)
	}
	var _apyRule []interface{}
	for _, _apyItem := range _apy {
		_apyRule = append(_apyRule, _apyItem)
	}

	logs, sub, err := _TesouroDireto.contract.WatchLogs(opts, "treasuryCreated", tokenIdRule, _totalValueRule, _apyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TesouroDiretoTreasuryCreated)
				if err := _TesouroDireto.contract.UnpackLog(event, "treasuryCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasuryCreated is a log parse operation binding the contract event 0xf83eeebb88a585a0bf111d55674f1c4cff9319871b905a62e7ede317bc032149.
//
// Solidity: event treasuryCreated(uint256 indexed tokenId, uint256 indexed _totalValue, uint256 indexed _apy, uint256 _duration, uint8 arg4)
func (_TesouroDireto *TesouroDiretoFilterer) ParseTreasuryCreated(log types.Log) (*TesouroDiretoTreasuryCreated, error) {
	event := new(TesouroDiretoTreasuryCreated)
	if err := _TesouroDireto.contract.UnpackLog(event, "treasuryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
