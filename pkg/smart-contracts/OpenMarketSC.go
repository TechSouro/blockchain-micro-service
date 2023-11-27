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

// OpenMarketMetaData contains all meta data concerning the OpenMarket contract.
var OpenMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_payment\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__union\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NotKYCed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NotTreasuryAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"paymentNotMade\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"retrievalFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"primarySale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"publicOrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"retrievalsucceed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_units\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"secondaryForSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_units\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"secondarySold\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_KYCed\",\"type\":\"address\"}],\"name\":\"KYC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_units\",\"type\":\"uint256\"}],\"name\":\"buySecondary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getOpenBuy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"getSecondaryMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openBuy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_avlb\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"purchasePrimary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"retrieveInvestment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"secondaryMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_avlb\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_units\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"sellMyUnits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OpenMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use OpenMarketMetaData.ABI instead.
var OpenMarketABI = OpenMarketMetaData.ABI

// OpenMarket is an auto generated Go binding around an Ethereum contract.
type OpenMarket struct {
	OpenMarketCaller     // Read-only binding to the contract
	OpenMarketTransactor // Write-only binding to the contract
	OpenMarketFilterer   // Log filterer for contract events
}

// OpenMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenMarketSession struct {
	Contract     *OpenMarket       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpenMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenMarketCallerSession struct {
	Contract *OpenMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OpenMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenMarketTransactorSession struct {
	Contract     *OpenMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OpenMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenMarketRaw struct {
	Contract *OpenMarket // Generic contract binding to access the raw methods on
}

// OpenMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenMarketCallerRaw struct {
	Contract *OpenMarketCaller // Generic read-only contract binding to access the raw methods on
}

// OpenMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenMarketTransactorRaw struct {
	Contract *OpenMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpenMarket creates a new instance of OpenMarket, bound to a specific deployed contract.
func NewOpenMarket(address common.Address, backend bind.ContractBackend) (*OpenMarket, error) {
	contract, err := bindOpenMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpenMarket{OpenMarketCaller: OpenMarketCaller{contract: contract}, OpenMarketTransactor: OpenMarketTransactor{contract: contract}, OpenMarketFilterer: OpenMarketFilterer{contract: contract}}, nil
}

// NewOpenMarketCaller creates a new read-only instance of OpenMarket, bound to a specific deployed contract.
func NewOpenMarketCaller(address common.Address, caller bind.ContractCaller) (*OpenMarketCaller, error) {
	contract, err := bindOpenMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenMarketCaller{contract: contract}, nil
}

// NewOpenMarketTransactor creates a new write-only instance of OpenMarket, bound to a specific deployed contract.
func NewOpenMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenMarketTransactor, error) {
	contract, err := bindOpenMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenMarketTransactor{contract: contract}, nil
}

// NewOpenMarketFilterer creates a new log filterer instance of OpenMarket, bound to a specific deployed contract.
func NewOpenMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenMarketFilterer, error) {
	contract, err := bindOpenMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenMarketFilterer{contract: contract}, nil
}

// bindOpenMarket binds a generic wrapper to an already deployed contract.
func bindOpenMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OpenMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenMarket *OpenMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenMarket.Contract.OpenMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenMarket *OpenMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenMarket.Contract.OpenMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenMarket *OpenMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenMarket.Contract.OpenMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenMarket *OpenMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenMarket *OpenMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenMarket *OpenMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenMarket.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_OpenMarket *OpenMarketCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_OpenMarket *OpenMarketSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _OpenMarket.Contract.BalanceOf(&_OpenMarket.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_OpenMarket *OpenMarketCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _OpenMarket.Contract.BalanceOf(&_OpenMarket.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_OpenMarket *OpenMarketCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_OpenMarket *OpenMarketSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _OpenMarket.Contract.BalanceOfBatch(&_OpenMarket.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_OpenMarket *OpenMarketCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _OpenMarket.Contract.BalanceOfBatch(&_OpenMarket.CallOpts, accounts, ids)
}

// Buyer is a free data retrieval call binding the contract method 0xb7748208.
//
// Solidity: function buyer(address ) view returns(bool)
func (_OpenMarket *OpenMarketCaller) Buyer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "buyer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0xb7748208.
//
// Solidity: function buyer(address ) view returns(bool)
func (_OpenMarket *OpenMarketSession) Buyer(arg0 common.Address) (bool, error) {
	return _OpenMarket.Contract.Buyer(&_OpenMarket.CallOpts, arg0)
}

// Buyer is a free data retrieval call binding the contract method 0xb7748208.
//
// Solidity: function buyer(address ) view returns(bool)
func (_OpenMarket *OpenMarketCallerSession) Buyer(arg0 common.Address) (bool, error) {
	return _OpenMarket.Contract.Buyer(&_OpenMarket.CallOpts, arg0)
}

// GetOpenBuy is a free data retrieval call binding the contract method 0xd98499f3.
//
// Solidity: function getOpenBuy(uint256 _tokenId) view returns(uint256, uint256)
func (_OpenMarket *OpenMarketCaller) GetOpenBuy(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "getOpenBuy", _tokenId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetOpenBuy is a free data retrieval call binding the contract method 0xd98499f3.
//
// Solidity: function getOpenBuy(uint256 _tokenId) view returns(uint256, uint256)
func (_OpenMarket *OpenMarketSession) GetOpenBuy(_tokenId *big.Int) (*big.Int, *big.Int, error) {
	return _OpenMarket.Contract.GetOpenBuy(&_OpenMarket.CallOpts, _tokenId)
}

// GetOpenBuy is a free data retrieval call binding the contract method 0xd98499f3.
//
// Solidity: function getOpenBuy(uint256 _tokenId) view returns(uint256, uint256)
func (_OpenMarket *OpenMarketCallerSession) GetOpenBuy(_tokenId *big.Int) (*big.Int, *big.Int, error) {
	return _OpenMarket.Contract.GetOpenBuy(&_OpenMarket.CallOpts, _tokenId)
}

// GetSecondaryMarket is a free data retrieval call binding the contract method 0x5bc4dacd.
//
// Solidity: function getSecondaryMarket(uint256 _tokenId, address _seller) view returns(uint256, uint256)
func (_OpenMarket *OpenMarketCaller) GetSecondaryMarket(opts *bind.CallOpts, _tokenId *big.Int, _seller common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "getSecondaryMarket", _tokenId, _seller)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSecondaryMarket is a free data retrieval call binding the contract method 0x5bc4dacd.
//
// Solidity: function getSecondaryMarket(uint256 _tokenId, address _seller) view returns(uint256, uint256)
func (_OpenMarket *OpenMarketSession) GetSecondaryMarket(_tokenId *big.Int, _seller common.Address) (*big.Int, *big.Int, error) {
	return _OpenMarket.Contract.GetSecondaryMarket(&_OpenMarket.CallOpts, _tokenId, _seller)
}

// GetSecondaryMarket is a free data retrieval call binding the contract method 0x5bc4dacd.
//
// Solidity: function getSecondaryMarket(uint256 _tokenId, address _seller) view returns(uint256, uint256)
func (_OpenMarket *OpenMarketCallerSession) GetSecondaryMarket(_tokenId *big.Int, _seller common.Address) (*big.Int, *big.Int, error) {
	return _OpenMarket.Contract.GetSecondaryMarket(&_OpenMarket.CallOpts, _tokenId, _seller)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_OpenMarket *OpenMarketCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_OpenMarket *OpenMarketSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _OpenMarket.Contract.IsApprovedForAll(&_OpenMarket.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_OpenMarket *OpenMarketCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _OpenMarket.Contract.IsApprovedForAll(&_OpenMarket.CallOpts, account, operator)
}

// OpenBuy is a free data retrieval call binding the contract method 0x0e372fc5.
//
// Solidity: function openBuy(uint256 ) view returns(uint256 _price, uint256 _avlb)
func (_OpenMarket *OpenMarketCaller) OpenBuy(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Price *big.Int
	Avlb  *big.Int
}, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "openBuy", arg0)

	outstruct := new(struct {
		Price *big.Int
		Avlb  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Avlb = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OpenBuy is a free data retrieval call binding the contract method 0x0e372fc5.
//
// Solidity: function openBuy(uint256 ) view returns(uint256 _price, uint256 _avlb)
func (_OpenMarket *OpenMarketSession) OpenBuy(arg0 *big.Int) (struct {
	Price *big.Int
	Avlb  *big.Int
}, error) {
	return _OpenMarket.Contract.OpenBuy(&_OpenMarket.CallOpts, arg0)
}

// OpenBuy is a free data retrieval call binding the contract method 0x0e372fc5.
//
// Solidity: function openBuy(uint256 ) view returns(uint256 _price, uint256 _avlb)
func (_OpenMarket *OpenMarketCallerSession) OpenBuy(arg0 *big.Int) (struct {
	Price *big.Int
	Avlb  *big.Int
}, error) {
	return _OpenMarket.Contract.OpenBuy(&_OpenMarket.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpenMarket *OpenMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpenMarket *OpenMarketSession) Owner() (common.Address, error) {
	return _OpenMarket.Contract.Owner(&_OpenMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpenMarket *OpenMarketCallerSession) Owner() (common.Address, error) {
	return _OpenMarket.Contract.Owner(&_OpenMarket.CallOpts)
}

// SecondaryMarket is a free data retrieval call binding the contract method 0x5ad2dfed.
//
// Solidity: function secondaryMarket(uint256 , address ) view returns(uint256 _price, uint256 _avlb)
func (_OpenMarket *OpenMarketCaller) SecondaryMarket(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Price *big.Int
	Avlb  *big.Int
}, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "secondaryMarket", arg0, arg1)

	outstruct := new(struct {
		Price *big.Int
		Avlb  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Avlb = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SecondaryMarket is a free data retrieval call binding the contract method 0x5ad2dfed.
//
// Solidity: function secondaryMarket(uint256 , address ) view returns(uint256 _price, uint256 _avlb)
func (_OpenMarket *OpenMarketSession) SecondaryMarket(arg0 *big.Int, arg1 common.Address) (struct {
	Price *big.Int
	Avlb  *big.Int
}, error) {
	return _OpenMarket.Contract.SecondaryMarket(&_OpenMarket.CallOpts, arg0, arg1)
}

// SecondaryMarket is a free data retrieval call binding the contract method 0x5ad2dfed.
//
// Solidity: function secondaryMarket(uint256 , address ) view returns(uint256 _price, uint256 _avlb)
func (_OpenMarket *OpenMarketCallerSession) SecondaryMarket(arg0 *big.Int, arg1 common.Address) (struct {
	Price *big.Int
	Avlb  *big.Int
}, error) {
	return _OpenMarket.Contract.SecondaryMarket(&_OpenMarket.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpenMarket *OpenMarketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpenMarket *OpenMarketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OpenMarket.Contract.SupportsInterface(&_OpenMarket.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpenMarket *OpenMarketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OpenMarket.Contract.SupportsInterface(&_OpenMarket.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_OpenMarket *OpenMarketCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _OpenMarket.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_OpenMarket *OpenMarketSession) Uri(arg0 *big.Int) (string, error) {
	return _OpenMarket.Contract.Uri(&_OpenMarket.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_OpenMarket *OpenMarketCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _OpenMarket.Contract.Uri(&_OpenMarket.CallOpts, arg0)
}

// KYC is a paid mutator transaction binding the contract method 0xbd3b1046.
//
// Solidity: function KYC(address _KYCed) returns()
func (_OpenMarket *OpenMarketTransactor) KYC(opts *bind.TransactOpts, _KYCed common.Address) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "KYC", _KYCed)
}

// KYC is a paid mutator transaction binding the contract method 0xbd3b1046.
//
// Solidity: function KYC(address _KYCed) returns()
func (_OpenMarket *OpenMarketSession) KYC(_KYCed common.Address) (*types.Transaction, error) {
	return _OpenMarket.Contract.KYC(&_OpenMarket.TransactOpts, _KYCed)
}

// KYC is a paid mutator transaction binding the contract method 0xbd3b1046.
//
// Solidity: function KYC(address _KYCed) returns()
func (_OpenMarket *OpenMarketTransactorSession) KYC(_KYCed common.Address) (*types.Transaction, error) {
	return _OpenMarket.Contract.KYC(&_OpenMarket.TransactOpts, _KYCed)
}

// BuySecondary is a paid mutator transaction binding the contract method 0x3efcb0f5.
//
// Solidity: function buySecondary(address _seller, uint256 _tokenId, uint256 _units) returns()
func (_OpenMarket *OpenMarketTransactor) BuySecondary(opts *bind.TransactOpts, _seller common.Address, _tokenId *big.Int, _units *big.Int) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "buySecondary", _seller, _tokenId, _units)
}

// BuySecondary is a paid mutator transaction binding the contract method 0x3efcb0f5.
//
// Solidity: function buySecondary(address _seller, uint256 _tokenId, uint256 _units) returns()
func (_OpenMarket *OpenMarketSession) BuySecondary(_seller common.Address, _tokenId *big.Int, _units *big.Int) (*types.Transaction, error) {
	return _OpenMarket.Contract.BuySecondary(&_OpenMarket.TransactOpts, _seller, _tokenId, _units)
}

// BuySecondary is a paid mutator transaction binding the contract method 0x3efcb0f5.
//
// Solidity: function buySecondary(address _seller, uint256 _tokenId, uint256 _units) returns()
func (_OpenMarket *OpenMarketTransactorSession) BuySecondary(_seller common.Address, _tokenId *big.Int, _units *big.Int) (*types.Transaction, error) {
	return _OpenMarket.Contract.BuySecondary(&_OpenMarket.TransactOpts, _seller, _tokenId, _units)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_OpenMarket *OpenMarketTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_OpenMarket *OpenMarketSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.OnERC1155BatchReceived(&_OpenMarket.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_OpenMarket *OpenMarketTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.OnERC1155BatchReceived(&_OpenMarket.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_OpenMarket *OpenMarketTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_OpenMarket *OpenMarketSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.OnERC1155Received(&_OpenMarket.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_OpenMarket *OpenMarketTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.OnERC1155Received(&_OpenMarket.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_OpenMarket *OpenMarketTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_OpenMarket *OpenMarketSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.OnERC721Received(&_OpenMarket.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_OpenMarket *OpenMarketTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.OnERC721Received(&_OpenMarket.TransactOpts, operator, from, tokenId, data)
}

// PurchasePrimary is a paid mutator transaction binding the contract method 0xd0a19852.
//
// Solidity: function purchasePrimary(uint256 _tokenId, uint256 _amount) returns()
func (_OpenMarket *OpenMarketTransactor) PurchasePrimary(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "purchasePrimary", _tokenId, _amount)
}

// PurchasePrimary is a paid mutator transaction binding the contract method 0xd0a19852.
//
// Solidity: function purchasePrimary(uint256 _tokenId, uint256 _amount) returns()
func (_OpenMarket *OpenMarketSession) PurchasePrimary(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _OpenMarket.Contract.PurchasePrimary(&_OpenMarket.TransactOpts, _tokenId, _amount)
}

// PurchasePrimary is a paid mutator transaction binding the contract method 0xd0a19852.
//
// Solidity: function purchasePrimary(uint256 _tokenId, uint256 _amount) returns()
func (_OpenMarket *OpenMarketTransactorSession) PurchasePrimary(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _OpenMarket.Contract.PurchasePrimary(&_OpenMarket.TransactOpts, _tokenId, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpenMarket *OpenMarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpenMarket *OpenMarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _OpenMarket.Contract.RenounceOwnership(&_OpenMarket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpenMarket *OpenMarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OpenMarket.Contract.RenounceOwnership(&_OpenMarket.TransactOpts)
}

// RetrieveInvestment is a paid mutator transaction binding the contract method 0xfeb8dc7c.
//
// Solidity: function retrieveInvestment(uint256 _tokenId, uint256 _amount) returns()
func (_OpenMarket *OpenMarketTransactor) RetrieveInvestment(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "retrieveInvestment", _tokenId, _amount)
}

// RetrieveInvestment is a paid mutator transaction binding the contract method 0xfeb8dc7c.
//
// Solidity: function retrieveInvestment(uint256 _tokenId, uint256 _amount) returns()
func (_OpenMarket *OpenMarketSession) RetrieveInvestment(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _OpenMarket.Contract.RetrieveInvestment(&_OpenMarket.TransactOpts, _tokenId, _amount)
}

// RetrieveInvestment is a paid mutator transaction binding the contract method 0xfeb8dc7c.
//
// Solidity: function retrieveInvestment(uint256 _tokenId, uint256 _amount) returns()
func (_OpenMarket *OpenMarketTransactorSession) RetrieveInvestment(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _OpenMarket.Contract.RetrieveInvestment(&_OpenMarket.TransactOpts, _tokenId, _amount)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_OpenMarket *OpenMarketTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_OpenMarket *OpenMarketSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.SafeBatchTransferFrom(&_OpenMarket.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_OpenMarket *OpenMarketTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.SafeBatchTransferFrom(&_OpenMarket.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_OpenMarket *OpenMarketTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_OpenMarket *OpenMarketSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.SafeTransferFrom(&_OpenMarket.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_OpenMarket *OpenMarketTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OpenMarket.Contract.SafeTransferFrom(&_OpenMarket.TransactOpts, from, to, id, value, data)
}

// SellMyUnits is a paid mutator transaction binding the contract method 0x3c4ce768.
//
// Solidity: function sellMyUnits(uint256 _tokenId, uint256 _units, uint256 _price) returns()
func (_OpenMarket *OpenMarketTransactor) SellMyUnits(opts *bind.TransactOpts, _tokenId *big.Int, _units *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "sellMyUnits", _tokenId, _units, _price)
}

// SellMyUnits is a paid mutator transaction binding the contract method 0x3c4ce768.
//
// Solidity: function sellMyUnits(uint256 _tokenId, uint256 _units, uint256 _price) returns()
func (_OpenMarket *OpenMarketSession) SellMyUnits(_tokenId *big.Int, _units *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _OpenMarket.Contract.SellMyUnits(&_OpenMarket.TransactOpts, _tokenId, _units, _price)
}

// SellMyUnits is a paid mutator transaction binding the contract method 0x3c4ce768.
//
// Solidity: function sellMyUnits(uint256 _tokenId, uint256 _units, uint256 _price) returns()
func (_OpenMarket *OpenMarketTransactorSession) SellMyUnits(_tokenId *big.Int, _units *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _OpenMarket.Contract.SellMyUnits(&_OpenMarket.TransactOpts, _tokenId, _units, _price)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_OpenMarket *OpenMarketTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_OpenMarket *OpenMarketSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _OpenMarket.Contract.SetApprovalForAll(&_OpenMarket.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_OpenMarket *OpenMarketTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _OpenMarket.Contract.SetApprovalForAll(&_OpenMarket.TransactOpts, operator, approved)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _addr) returns()
func (_OpenMarket *OpenMarketTransactor) SetTreasury(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "setTreasury", _addr)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _addr) returns()
func (_OpenMarket *OpenMarketSession) SetTreasury(_addr common.Address) (*types.Transaction, error) {
	return _OpenMarket.Contract.SetTreasury(&_OpenMarket.TransactOpts, _addr)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _addr) returns()
func (_OpenMarket *OpenMarketTransactorSession) SetTreasury(_addr common.Address) (*types.Transaction, error) {
	return _OpenMarket.Contract.SetTreasury(&_OpenMarket.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpenMarket *OpenMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OpenMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpenMarket *OpenMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OpenMarket.Contract.TransferOwnership(&_OpenMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpenMarket *OpenMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OpenMarket.Contract.TransferOwnership(&_OpenMarket.TransactOpts, newOwner)
}

// OpenMarketApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the OpenMarket contract.
type OpenMarketApprovalForAllIterator struct {
	Event *OpenMarketApprovalForAll // Event containing the contract specifics and raw log

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
func (it *OpenMarketApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketApprovalForAll)
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
		it.Event = new(OpenMarketApprovalForAll)
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
func (it *OpenMarketApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketApprovalForAll represents a ApprovalForAll event raised by the OpenMarket contract.
type OpenMarketApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_OpenMarket *OpenMarketFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*OpenMarketApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &OpenMarketApprovalForAllIterator{contract: _OpenMarket.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_OpenMarket *OpenMarketFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *OpenMarketApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketApprovalForAll)
				if err := _OpenMarket.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_OpenMarket *OpenMarketFilterer) ParseApprovalForAll(log types.Log) (*OpenMarketApprovalForAll, error) {
	event := new(OpenMarketApprovalForAll)
	if err := _OpenMarket.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OpenMarket contract.
type OpenMarketOwnershipTransferredIterator struct {
	Event *OpenMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OpenMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketOwnershipTransferred)
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
		it.Event = new(OpenMarketOwnershipTransferred)
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
func (it *OpenMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketOwnershipTransferred represents a OwnershipTransferred event raised by the OpenMarket contract.
type OpenMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OpenMarket *OpenMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OpenMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OpenMarketOwnershipTransferredIterator{contract: _OpenMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OpenMarket *OpenMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OpenMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketOwnershipTransferred)
				if err := _OpenMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OpenMarket *OpenMarketFilterer) ParseOwnershipTransferred(log types.Log) (*OpenMarketOwnershipTransferred, error) {
	event := new(OpenMarketOwnershipTransferred)
	if err := _OpenMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenMarketTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the OpenMarket contract.
type OpenMarketTransferBatchIterator struct {
	Event *OpenMarketTransferBatch // Event containing the contract specifics and raw log

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
func (it *OpenMarketTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketTransferBatch)
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
		it.Event = new(OpenMarketTransferBatch)
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
func (it *OpenMarketTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketTransferBatch represents a TransferBatch event raised by the OpenMarket contract.
type OpenMarketTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_OpenMarket *OpenMarketFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*OpenMarketTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OpenMarketTransferBatchIterator{contract: _OpenMarket.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_OpenMarket *OpenMarketFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *OpenMarketTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketTransferBatch)
				if err := _OpenMarket.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_OpenMarket *OpenMarketFilterer) ParseTransferBatch(log types.Log) (*OpenMarketTransferBatch, error) {
	event := new(OpenMarketTransferBatch)
	if err := _OpenMarket.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenMarketTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the OpenMarket contract.
type OpenMarketTransferSingleIterator struct {
	Event *OpenMarketTransferSingle // Event containing the contract specifics and raw log

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
func (it *OpenMarketTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketTransferSingle)
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
		it.Event = new(OpenMarketTransferSingle)
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
func (it *OpenMarketTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketTransferSingle represents a TransferSingle event raised by the OpenMarket contract.
type OpenMarketTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_OpenMarket *OpenMarketFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*OpenMarketTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OpenMarketTransferSingleIterator{contract: _OpenMarket.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_OpenMarket *OpenMarketFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *OpenMarketTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketTransferSingle)
				if err := _OpenMarket.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_OpenMarket *OpenMarketFilterer) ParseTransferSingle(log types.Log) (*OpenMarketTransferSingle, error) {
	event := new(OpenMarketTransferSingle)
	if err := _OpenMarket.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenMarketURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the OpenMarket contract.
type OpenMarketURIIterator struct {
	Event *OpenMarketURI // Event containing the contract specifics and raw log

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
func (it *OpenMarketURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketURI)
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
		it.Event = new(OpenMarketURI)
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
func (it *OpenMarketURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketURI represents a URI event raised by the OpenMarket contract.
type OpenMarketURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_OpenMarket *OpenMarketFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*OpenMarketURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &OpenMarketURIIterator{contract: _OpenMarket.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_OpenMarket *OpenMarketFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *OpenMarketURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketURI)
				if err := _OpenMarket.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_OpenMarket *OpenMarketFilterer) ParseURI(log types.Log) (*OpenMarketURI, error) {
	event := new(OpenMarketURI)
	if err := _OpenMarket.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenMarketPrimarySaleIterator is returned from FilterPrimarySale and is used to iterate over the raw logs and unpacked data for PrimarySale events raised by the OpenMarket contract.
type OpenMarketPrimarySaleIterator struct {
	Event *OpenMarketPrimarySale // Event containing the contract specifics and raw log

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
func (it *OpenMarketPrimarySaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketPrimarySale)
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
		it.Event = new(OpenMarketPrimarySale)
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
func (it *OpenMarketPrimarySaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketPrimarySaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketPrimarySale represents a PrimarySale event raised by the OpenMarket contract.
type OpenMarketPrimarySale struct {
	Sender  common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPrimarySale is a free log retrieval operation binding the contract event 0x578e1297d3ef1c181a1ca672374ed2a4472844362b1dde5923fa2a2c12347e49.
//
// Solidity: event primarySale(address indexed _sender, uint256 indexed _tokenId, uint256 indexed _amount)
func (_OpenMarket *OpenMarketFilterer) FilterPrimarySale(opts *bind.FilterOpts, _sender []common.Address, _tokenId []*big.Int, _amount []*big.Int) (*OpenMarketPrimarySaleIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "primarySale", _senderRule, _tokenIdRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &OpenMarketPrimarySaleIterator{contract: _OpenMarket.contract, event: "primarySale", logs: logs, sub: sub}, nil
}

// WatchPrimarySale is a free log subscription operation binding the contract event 0x578e1297d3ef1c181a1ca672374ed2a4472844362b1dde5923fa2a2c12347e49.
//
// Solidity: event primarySale(address indexed _sender, uint256 indexed _tokenId, uint256 indexed _amount)
func (_OpenMarket *OpenMarketFilterer) WatchPrimarySale(opts *bind.WatchOpts, sink chan<- *OpenMarketPrimarySale, _sender []common.Address, _tokenId []*big.Int, _amount []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "primarySale", _senderRule, _tokenIdRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketPrimarySale)
				if err := _OpenMarket.contract.UnpackLog(event, "primarySale", log); err != nil {
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

// ParsePrimarySale is a log parse operation binding the contract event 0x578e1297d3ef1c181a1ca672374ed2a4472844362b1dde5923fa2a2c12347e49.
//
// Solidity: event primarySale(address indexed _sender, uint256 indexed _tokenId, uint256 indexed _amount)
func (_OpenMarket *OpenMarketFilterer) ParsePrimarySale(log types.Log) (*OpenMarketPrimarySale, error) {
	event := new(OpenMarketPrimarySale)
	if err := _OpenMarket.contract.UnpackLog(event, "primarySale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenMarketPublicOrderCreatedIterator is returned from FilterPublicOrderCreated and is used to iterate over the raw logs and unpacked data for PublicOrderCreated events raised by the OpenMarket contract.
type OpenMarketPublicOrderCreatedIterator struct {
	Event *OpenMarketPublicOrderCreated // Event containing the contract specifics and raw log

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
func (it *OpenMarketPublicOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketPublicOrderCreated)
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
		it.Event = new(OpenMarketPublicOrderCreated)
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
func (it *OpenMarketPublicOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketPublicOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketPublicOrderCreated represents a PublicOrderCreated event raised by the OpenMarket contract.
type OpenMarketPublicOrderCreated struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPublicOrderCreated is a free log retrieval operation binding the contract event 0x94dfc0d447c08076d42633196f4738eab6bffb6e26a858bbc09ae254b12101c2.
//
// Solidity: event publicOrderCreated(uint256 indexed _tokenId)
func (_OpenMarket *OpenMarketFilterer) FilterPublicOrderCreated(opts *bind.FilterOpts, _tokenId []*big.Int) (*OpenMarketPublicOrderCreatedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "publicOrderCreated", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OpenMarketPublicOrderCreatedIterator{contract: _OpenMarket.contract, event: "publicOrderCreated", logs: logs, sub: sub}, nil
}

// WatchPublicOrderCreated is a free log subscription operation binding the contract event 0x94dfc0d447c08076d42633196f4738eab6bffb6e26a858bbc09ae254b12101c2.
//
// Solidity: event publicOrderCreated(uint256 indexed _tokenId)
func (_OpenMarket *OpenMarketFilterer) WatchPublicOrderCreated(opts *bind.WatchOpts, sink chan<- *OpenMarketPublicOrderCreated, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "publicOrderCreated", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketPublicOrderCreated)
				if err := _OpenMarket.contract.UnpackLog(event, "publicOrderCreated", log); err != nil {
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

// ParsePublicOrderCreated is a log parse operation binding the contract event 0x94dfc0d447c08076d42633196f4738eab6bffb6e26a858bbc09ae254b12101c2.
//
// Solidity: event publicOrderCreated(uint256 indexed _tokenId)
func (_OpenMarket *OpenMarketFilterer) ParsePublicOrderCreated(log types.Log) (*OpenMarketPublicOrderCreated, error) {
	event := new(OpenMarketPublicOrderCreated)
	if err := _OpenMarket.contract.UnpackLog(event, "publicOrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenMarketRetrievalsucceedIterator is returned from FilterRetrievalsucceed and is used to iterate over the raw logs and unpacked data for Retrievalsucceed events raised by the OpenMarket contract.
type OpenMarketRetrievalsucceedIterator struct {
	Event *OpenMarketRetrievalsucceed // Event containing the contract specifics and raw log

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
func (it *OpenMarketRetrievalsucceedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketRetrievalsucceed)
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
		it.Event = new(OpenMarketRetrievalsucceed)
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
func (it *OpenMarketRetrievalsucceedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketRetrievalsucceedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketRetrievalsucceed represents a Retrievalsucceed event raised by the OpenMarket contract.
type OpenMarketRetrievalsucceed struct {
	Sender  common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRetrievalsucceed is a free log retrieval operation binding the contract event 0x5fee19f5bf8eb8b0af1645ea5d38435af21e550792658dd02025644dc1373ebf.
//
// Solidity: event retrievalsucceed(address _sender, uint256 _tokenId, uint256 _amount)
func (_OpenMarket *OpenMarketFilterer) FilterRetrievalsucceed(opts *bind.FilterOpts) (*OpenMarketRetrievalsucceedIterator, error) {

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "retrievalsucceed")
	if err != nil {
		return nil, err
	}
	return &OpenMarketRetrievalsucceedIterator{contract: _OpenMarket.contract, event: "retrievalsucceed", logs: logs, sub: sub}, nil
}

// WatchRetrievalsucceed is a free log subscription operation binding the contract event 0x5fee19f5bf8eb8b0af1645ea5d38435af21e550792658dd02025644dc1373ebf.
//
// Solidity: event retrievalsucceed(address _sender, uint256 _tokenId, uint256 _amount)
func (_OpenMarket *OpenMarketFilterer) WatchRetrievalsucceed(opts *bind.WatchOpts, sink chan<- *OpenMarketRetrievalsucceed) (event.Subscription, error) {

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "retrievalsucceed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketRetrievalsucceed)
				if err := _OpenMarket.contract.UnpackLog(event, "retrievalsucceed", log); err != nil {
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

// ParseRetrievalsucceed is a log parse operation binding the contract event 0x5fee19f5bf8eb8b0af1645ea5d38435af21e550792658dd02025644dc1373ebf.
//
// Solidity: event retrievalsucceed(address _sender, uint256 _tokenId, uint256 _amount)
func (_OpenMarket *OpenMarketFilterer) ParseRetrievalsucceed(log types.Log) (*OpenMarketRetrievalsucceed, error) {
	event := new(OpenMarketRetrievalsucceed)
	if err := _OpenMarket.contract.UnpackLog(event, "retrievalsucceed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenMarketSecondaryForSaleIterator is returned from FilterSecondaryForSale and is used to iterate over the raw logs and unpacked data for SecondaryForSale events raised by the OpenMarket contract.
type OpenMarketSecondaryForSaleIterator struct {
	Event *OpenMarketSecondaryForSale // Event containing the contract specifics and raw log

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
func (it *OpenMarketSecondaryForSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketSecondaryForSale)
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
		it.Event = new(OpenMarketSecondaryForSale)
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
func (it *OpenMarketSecondaryForSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketSecondaryForSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketSecondaryForSale represents a SecondaryForSale event raised by the OpenMarket contract.
type OpenMarketSecondaryForSale struct {
	Seller  common.Address
	TokenId *big.Int
	Units   *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSecondaryForSale is a free log retrieval operation binding the contract event 0x7aa393e2de6e760269612f345e5741eb8713b054ad41e6f84de83f3ee5edd00c.
//
// Solidity: event secondaryForSale(address indexed _seller, uint256 indexed _tokenId, uint256 indexed _units, uint256 _price)
func (_OpenMarket *OpenMarketFilterer) FilterSecondaryForSale(opts *bind.FilterOpts, _seller []common.Address, _tokenId []*big.Int, _units []*big.Int) (*OpenMarketSecondaryForSaleIterator, error) {

	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _unitsRule []interface{}
	for _, _unitsItem := range _units {
		_unitsRule = append(_unitsRule, _unitsItem)
	}

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "secondaryForSale", _sellerRule, _tokenIdRule, _unitsRule)
	if err != nil {
		return nil, err
	}
	return &OpenMarketSecondaryForSaleIterator{contract: _OpenMarket.contract, event: "secondaryForSale", logs: logs, sub: sub}, nil
}

// WatchSecondaryForSale is a free log subscription operation binding the contract event 0x7aa393e2de6e760269612f345e5741eb8713b054ad41e6f84de83f3ee5edd00c.
//
// Solidity: event secondaryForSale(address indexed _seller, uint256 indexed _tokenId, uint256 indexed _units, uint256 _price)
func (_OpenMarket *OpenMarketFilterer) WatchSecondaryForSale(opts *bind.WatchOpts, sink chan<- *OpenMarketSecondaryForSale, _seller []common.Address, _tokenId []*big.Int, _units []*big.Int) (event.Subscription, error) {

	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _unitsRule []interface{}
	for _, _unitsItem := range _units {
		_unitsRule = append(_unitsRule, _unitsItem)
	}

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "secondaryForSale", _sellerRule, _tokenIdRule, _unitsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketSecondaryForSale)
				if err := _OpenMarket.contract.UnpackLog(event, "secondaryForSale", log); err != nil {
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

// ParseSecondaryForSale is a log parse operation binding the contract event 0x7aa393e2de6e760269612f345e5741eb8713b054ad41e6f84de83f3ee5edd00c.
//
// Solidity: event secondaryForSale(address indexed _seller, uint256 indexed _tokenId, uint256 indexed _units, uint256 _price)
func (_OpenMarket *OpenMarketFilterer) ParseSecondaryForSale(log types.Log) (*OpenMarketSecondaryForSale, error) {
	event := new(OpenMarketSecondaryForSale)
	if err := _OpenMarket.contract.UnpackLog(event, "secondaryForSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenMarketSecondarySoldIterator is returned from FilterSecondarySold and is used to iterate over the raw logs and unpacked data for SecondarySold events raised by the OpenMarket contract.
type OpenMarketSecondarySoldIterator struct {
	Event *OpenMarketSecondarySold // Event containing the contract specifics and raw log

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
func (it *OpenMarketSecondarySoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenMarketSecondarySold)
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
		it.Event = new(OpenMarketSecondarySold)
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
func (it *OpenMarketSecondarySoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenMarketSecondarySoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenMarketSecondarySold represents a SecondarySold event raised by the OpenMarket contract.
type OpenMarketSecondarySold struct {
	Seller  common.Address
	Buyer   common.Address
	Units   *big.Int
	Price   *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSecondarySold is a free log retrieval operation binding the contract event 0x2c939b612f6e497b4f3e7beac4db3de1aa2fdaa818c78408dfd361db5b195dcf.
//
// Solidity: event secondarySold(address indexed _seller, address indexed _buyer, uint256 indexed _units, uint256 _price, uint256 _tokenId)
func (_OpenMarket *OpenMarketFilterer) FilterSecondarySold(opts *bind.FilterOpts, _seller []common.Address, _buyer []common.Address, _units []*big.Int) (*OpenMarketSecondarySoldIterator, error) {

	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _unitsRule []interface{}
	for _, _unitsItem := range _units {
		_unitsRule = append(_unitsRule, _unitsItem)
	}

	logs, sub, err := _OpenMarket.contract.FilterLogs(opts, "secondarySold", _sellerRule, _buyerRule, _unitsRule)
	if err != nil {
		return nil, err
	}
	return &OpenMarketSecondarySoldIterator{contract: _OpenMarket.contract, event: "secondarySold", logs: logs, sub: sub}, nil
}

// WatchSecondarySold is a free log subscription operation binding the contract event 0x2c939b612f6e497b4f3e7beac4db3de1aa2fdaa818c78408dfd361db5b195dcf.
//
// Solidity: event secondarySold(address indexed _seller, address indexed _buyer, uint256 indexed _units, uint256 _price, uint256 _tokenId)
func (_OpenMarket *OpenMarketFilterer) WatchSecondarySold(opts *bind.WatchOpts, sink chan<- *OpenMarketSecondarySold, _seller []common.Address, _buyer []common.Address, _units []*big.Int) (event.Subscription, error) {

	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _unitsRule []interface{}
	for _, _unitsItem := range _units {
		_unitsRule = append(_unitsRule, _unitsItem)
	}

	logs, sub, err := _OpenMarket.contract.WatchLogs(opts, "secondarySold", _sellerRule, _buyerRule, _unitsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenMarketSecondarySold)
				if err := _OpenMarket.contract.UnpackLog(event, "secondarySold", log); err != nil {
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

// ParseSecondarySold is a log parse operation binding the contract event 0x2c939b612f6e497b4f3e7beac4db3de1aa2fdaa818c78408dfd361db5b195dcf.
//
// Solidity: event secondarySold(address indexed _seller, address indexed _buyer, uint256 indexed _units, uint256 _price, uint256 _tokenId)
func (_OpenMarket *OpenMarketFilterer) ParseSecondarySold(log types.Log) (*OpenMarketSecondarySold, error) {
	event := new(OpenMarketSecondarySold)
	if err := _OpenMarket.contract.UnpackLog(event, "secondarySold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
