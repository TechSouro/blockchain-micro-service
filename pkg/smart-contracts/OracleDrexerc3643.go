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

// OracleDrexerc3643MetaData contains all meta data concerning the OracleDrexerc3643 contract.
var OracleDrexerc3643MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_isFrozen\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"AddressFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_compliance\",\"type\":\"address\"}],\"name\":\"ComplianceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_identityRegistry\",\"type\":\"address\"}],\"name\":\"IdentityRegistryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_lostWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_investorOnchainID\",\"type\":\"address\"}],\"name\":\"RecoverySuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TokensFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TokensUnfrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_newName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_newSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_newDecimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_newVersion\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOnchainID\",\"type\":\"address\"}],\"name\":\"UpdatedTokenInformation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"addAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_Addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approveExternal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_userAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_fromList\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_toList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchForcedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_userAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchFreezePartialTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_toList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_userAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_freeze\",\"type\":\"bool[]\"}],\"name\":\"batchSetAddressFrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_toList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_userAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchUnfreezePartialTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compliance\",\"outputs\":[{\"internalType\":\"contractIModularCompliance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"forcedTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"freezePartialTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"getFrozenTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"identityRegistry\",\"outputs\":[{\"internalType\":\"contractIIdentityRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_compliance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_onchainID\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"isAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onchainID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lostWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_investorOnchainID\",\"type\":\"address\"}],\"name\":\"recoveryAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_freeze\",\"type\":\"bool\"}],\"name\":\"setAddressFrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_compliance\",\"type\":\"address\"}],\"name\":\"setCompliance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityRegistry\",\"type\":\"address\"}],\"name\":\"setIdentityRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_onchainID\",\"type\":\"address\"}],\"name\":\"setOnchainID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"setSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unfreezePartialTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// OracleDrexerc3643ABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleDrexerc3643MetaData.ABI instead.
var OracleDrexerc3643ABI = OracleDrexerc3643MetaData.ABI

// OracleDrexerc3643 is an auto generated Go binding around an Ethereum contract.
type OracleDrexerc3643 struct {
	OracleDrexerc3643Caller     // Read-only binding to the contract
	OracleDrexerc3643Transactor // Write-only binding to the contract
	OracleDrexerc3643Filterer   // Log filterer for contract events
}

// OracleDrexerc3643Caller is an auto generated read-only Go binding around an Ethereum contract.
type OracleDrexerc3643Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleDrexerc3643Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleDrexerc3643Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleDrexerc3643Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleDrexerc3643Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleDrexerc3643Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleDrexerc3643Session struct {
	Contract     *OracleDrexerc3643 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OracleDrexerc3643CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleDrexerc3643CallerSession struct {
	Contract *OracleDrexerc3643Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OracleDrexerc3643TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleDrexerc3643TransactorSession struct {
	Contract     *OracleDrexerc3643Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OracleDrexerc3643Raw is an auto generated low-level Go binding around an Ethereum contract.
type OracleDrexerc3643Raw struct {
	Contract *OracleDrexerc3643 // Generic contract binding to access the raw methods on
}

// OracleDrexerc3643CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleDrexerc3643CallerRaw struct {
	Contract *OracleDrexerc3643Caller // Generic read-only contract binding to access the raw methods on
}

// OracleDrexerc3643TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleDrexerc3643TransactorRaw struct {
	Contract *OracleDrexerc3643Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleDrexerc3643 creates a new instance of OracleDrexerc3643, bound to a specific deployed contract.
func NewOracleDrexerc3643(address common.Address, backend bind.ContractBackend) (*OracleDrexerc3643, error) {
	contract, err := bindOracleDrexerc3643(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643{OracleDrexerc3643Caller: OracleDrexerc3643Caller{contract: contract}, OracleDrexerc3643Transactor: OracleDrexerc3643Transactor{contract: contract}, OracleDrexerc3643Filterer: OracleDrexerc3643Filterer{contract: contract}}, nil
}

// NewOracleDrexerc3643Caller creates a new read-only instance of OracleDrexerc3643, bound to a specific deployed contract.
func NewOracleDrexerc3643Caller(address common.Address, caller bind.ContractCaller) (*OracleDrexerc3643Caller, error) {
	contract, err := bindOracleDrexerc3643(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643Caller{contract: contract}, nil
}

// NewOracleDrexerc3643Transactor creates a new write-only instance of OracleDrexerc3643, bound to a specific deployed contract.
func NewOracleDrexerc3643Transactor(address common.Address, transactor bind.ContractTransactor) (*OracleDrexerc3643Transactor, error) {
	contract, err := bindOracleDrexerc3643(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643Transactor{contract: contract}, nil
}

// NewOracleDrexerc3643Filterer creates a new log filterer instance of OracleDrexerc3643, bound to a specific deployed contract.
func NewOracleDrexerc3643Filterer(address common.Address, filterer bind.ContractFilterer) (*OracleDrexerc3643Filterer, error) {
	contract, err := bindOracleDrexerc3643(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643Filterer{contract: contract}, nil
}

// bindOracleDrexerc3643 binds a generic wrapper to an already deployed contract.
func bindOracleDrexerc3643(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleDrexerc3643MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleDrexerc3643 *OracleDrexerc3643Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleDrexerc3643.Contract.OracleDrexerc3643Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleDrexerc3643 *OracleDrexerc3643Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.OracleDrexerc3643Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleDrexerc3643 *OracleDrexerc3643Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.OracleDrexerc3643Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleDrexerc3643 *OracleDrexerc3643CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleDrexerc3643.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _OracleDrexerc3643.Contract.Allowance(&_OracleDrexerc3643.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _OracleDrexerc3643.Contract.Allowance(&_OracleDrexerc3643.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) BalanceOf(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "balanceOf", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) BalanceOf(_userAddress common.Address) (*big.Int, error) {
	return _OracleDrexerc3643.Contract.BalanceOf(&_OracleDrexerc3643.CallOpts, _userAddress)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) BalanceOf(_userAddress common.Address) (*big.Int, error) {
	return _OracleDrexerc3643.Contract.BalanceOf(&_OracleDrexerc3643.CallOpts, _userAddress)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) Compliance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "compliance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Compliance() (common.Address, error) {
	return _OracleDrexerc3643.Contract.Compliance(&_OracleDrexerc3643.CallOpts)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) Compliance() (common.Address, error) {
	return _OracleDrexerc3643.Contract.Compliance(&_OracleDrexerc3643.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Decimals() (uint8, error) {
	return _OracleDrexerc3643.Contract.Decimals(&_OracleDrexerc3643.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) Decimals() (uint8, error) {
	return _OracleDrexerc3643.Contract.Decimals(&_OracleDrexerc3643.CallOpts)
}

// GetFrozenTokens is a free data retrieval call binding the contract method 0x158b1a57.
//
// Solidity: function getFrozenTokens(address _userAddress) view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) GetFrozenTokens(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "getFrozenTokens", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFrozenTokens is a free data retrieval call binding the contract method 0x158b1a57.
//
// Solidity: function getFrozenTokens(address _userAddress) view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) GetFrozenTokens(_userAddress common.Address) (*big.Int, error) {
	return _OracleDrexerc3643.Contract.GetFrozenTokens(&_OracleDrexerc3643.CallOpts, _userAddress)
}

// GetFrozenTokens is a free data retrieval call binding the contract method 0x158b1a57.
//
// Solidity: function getFrozenTokens(address _userAddress) view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) GetFrozenTokens(_userAddress common.Address) (*big.Int, error) {
	return _OracleDrexerc3643.Contract.GetFrozenTokens(&_OracleDrexerc3643.CallOpts, _userAddress)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) IdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "identityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) IdentityRegistry() (common.Address, error) {
	return _OracleDrexerc3643.Contract.IdentityRegistry(&_OracleDrexerc3643.CallOpts)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) IdentityRegistry() (common.Address, error) {
	return _OracleDrexerc3643.Contract.IdentityRegistry(&_OracleDrexerc3643.CallOpts)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) IsAgent(opts *bind.CallOpts, _agent common.Address) (bool, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "isAgent", _agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) IsAgent(_agent common.Address) (bool, error) {
	return _OracleDrexerc3643.Contract.IsAgent(&_OracleDrexerc3643.CallOpts, _agent)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) IsAgent(_agent common.Address) (bool, error) {
	return _OracleDrexerc3643.Contract.IsAgent(&_OracleDrexerc3643.CallOpts, _agent)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address _userAddress) view returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) IsFrozen(opts *bind.CallOpts, _userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "isFrozen", _userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address _userAddress) view returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) IsFrozen(_userAddress common.Address) (bool, error) {
	return _OracleDrexerc3643.Contract.IsFrozen(&_OracleDrexerc3643.CallOpts, _userAddress)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address _userAddress) view returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) IsFrozen(_userAddress common.Address) (bool, error) {
	return _OracleDrexerc3643.Contract.IsFrozen(&_OracleDrexerc3643.CallOpts, _userAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Name() (string, error) {
	return _OracleDrexerc3643.Contract.Name(&_OracleDrexerc3643.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) Name() (string, error) {
	return _OracleDrexerc3643.Contract.Name(&_OracleDrexerc3643.CallOpts)
}

// OnchainID is a free data retrieval call binding the contract method 0xaba63705.
//
// Solidity: function onchainID() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) OnchainID(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "onchainID")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnchainID is a free data retrieval call binding the contract method 0xaba63705.
//
// Solidity: function onchainID() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) OnchainID() (common.Address, error) {
	return _OracleDrexerc3643.Contract.OnchainID(&_OracleDrexerc3643.CallOpts)
}

// OnchainID is a free data retrieval call binding the contract method 0xaba63705.
//
// Solidity: function onchainID() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) OnchainID() (common.Address, error) {
	return _OracleDrexerc3643.Contract.OnchainID(&_OracleDrexerc3643.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Owner() (common.Address, error) {
	return _OracleDrexerc3643.Contract.Owner(&_OracleDrexerc3643.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) Owner() (common.Address, error) {
	return _OracleDrexerc3643.Contract.Owner(&_OracleDrexerc3643.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Paused() (bool, error) {
	return _OracleDrexerc3643.Contract.Paused(&_OracleDrexerc3643.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) Paused() (bool, error) {
	return _OracleDrexerc3643.Contract.Paused(&_OracleDrexerc3643.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Symbol() (string, error) {
	return _OracleDrexerc3643.Contract.Symbol(&_OracleDrexerc3643.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) Symbol() (string, error) {
	return _OracleDrexerc3643.Contract.Symbol(&_OracleDrexerc3643.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) TotalSupply() (*big.Int, error) {
	return _OracleDrexerc3643.Contract.TotalSupply(&_OracleDrexerc3643.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) TotalSupply() (*big.Int, error) {
	return _OracleDrexerc3643.Contract.TotalSupply(&_OracleDrexerc3643.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_OracleDrexerc3643 *OracleDrexerc3643Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OracleDrexerc3643.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Version() (string, error) {
	return _OracleDrexerc3643.Contract.Version(&_OracleDrexerc3643.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_OracleDrexerc3643 *OracleDrexerc3643CallerSession) Version() (string, error) {
	return _OracleDrexerc3643.Contract.Version(&_OracleDrexerc3643.CallOpts)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) AddAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "addAgent", _agent)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.AddAgent(&_OracleDrexerc3643.TransactOpts, _agent)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.AddAgent(&_OracleDrexerc3643.TransactOpts, _agent)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Approve(&_OracleDrexerc3643.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Approve(&_OracleDrexerc3643.TransactOpts, _spender, _amount)
}

// ApproveExternal is a paid mutator transaction binding the contract method 0xd7699698.
//
// Solidity: function approveExternal(address _Addr, uint256 _value) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) ApproveExternal(opts *bind.TransactOpts, _Addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "approveExternal", _Addr, _value)
}

// ApproveExternal is a paid mutator transaction binding the contract method 0xd7699698.
//
// Solidity: function approveExternal(address _Addr, uint256 _value) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) ApproveExternal(_Addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.ApproveExternal(&_OracleDrexerc3643.TransactOpts, _Addr, _value)
}

// ApproveExternal is a paid mutator transaction binding the contract method 0xd7699698.
//
// Solidity: function approveExternal(address _Addr, uint256 _value) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) ApproveExternal(_Addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.ApproveExternal(&_OracleDrexerc3643.TransactOpts, _Addr, _value)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x4a6cc677.
//
// Solidity: function batchBurn(address[] _userAddresses, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) BatchBurn(opts *bind.TransactOpts, _userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "batchBurn", _userAddresses, _amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x4a6cc677.
//
// Solidity: function batchBurn(address[] _userAddresses, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) BatchBurn(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchBurn(&_OracleDrexerc3643.TransactOpts, _userAddresses, _amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x4a6cc677.
//
// Solidity: function batchBurn(address[] _userAddresses, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) BatchBurn(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchBurn(&_OracleDrexerc3643.TransactOpts, _userAddresses, _amounts)
}

// BatchForcedTransfer is a paid mutator transaction binding the contract method 0x42a47abc.
//
// Solidity: function batchForcedTransfer(address[] _fromList, address[] _toList, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) BatchForcedTransfer(opts *bind.TransactOpts, _fromList []common.Address, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "batchForcedTransfer", _fromList, _toList, _amounts)
}

// BatchForcedTransfer is a paid mutator transaction binding the contract method 0x42a47abc.
//
// Solidity: function batchForcedTransfer(address[] _fromList, address[] _toList, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) BatchForcedTransfer(_fromList []common.Address, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchForcedTransfer(&_OracleDrexerc3643.TransactOpts, _fromList, _toList, _amounts)
}

// BatchForcedTransfer is a paid mutator transaction binding the contract method 0x42a47abc.
//
// Solidity: function batchForcedTransfer(address[] _fromList, address[] _toList, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) BatchForcedTransfer(_fromList []common.Address, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchForcedTransfer(&_OracleDrexerc3643.TransactOpts, _fromList, _toList, _amounts)
}

// BatchFreezePartialTokens is a paid mutator transaction binding the contract method 0xfc7e5fa8.
//
// Solidity: function batchFreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) BatchFreezePartialTokens(opts *bind.TransactOpts, _userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "batchFreezePartialTokens", _userAddresses, _amounts)
}

// BatchFreezePartialTokens is a paid mutator transaction binding the contract method 0xfc7e5fa8.
//
// Solidity: function batchFreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) BatchFreezePartialTokens(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchFreezePartialTokens(&_OracleDrexerc3643.TransactOpts, _userAddresses, _amounts)
}

// BatchFreezePartialTokens is a paid mutator transaction binding the contract method 0xfc7e5fa8.
//
// Solidity: function batchFreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) BatchFreezePartialTokens(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchFreezePartialTokens(&_OracleDrexerc3643.TransactOpts, _userAddresses, _amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] _toList, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) BatchMint(opts *bind.TransactOpts, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "batchMint", _toList, _amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] _toList, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) BatchMint(_toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchMint(&_OracleDrexerc3643.TransactOpts, _toList, _amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] _toList, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) BatchMint(_toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchMint(&_OracleDrexerc3643.TransactOpts, _toList, _amounts)
}

// BatchSetAddressFrozen is a paid mutator transaction binding the contract method 0x1a7af379.
//
// Solidity: function batchSetAddressFrozen(address[] _userAddresses, bool[] _freeze) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) BatchSetAddressFrozen(opts *bind.TransactOpts, _userAddresses []common.Address, _freeze []bool) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "batchSetAddressFrozen", _userAddresses, _freeze)
}

// BatchSetAddressFrozen is a paid mutator transaction binding the contract method 0x1a7af379.
//
// Solidity: function batchSetAddressFrozen(address[] _userAddresses, bool[] _freeze) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) BatchSetAddressFrozen(_userAddresses []common.Address, _freeze []bool) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchSetAddressFrozen(&_OracleDrexerc3643.TransactOpts, _userAddresses, _freeze)
}

// BatchSetAddressFrozen is a paid mutator transaction binding the contract method 0x1a7af379.
//
// Solidity: function batchSetAddressFrozen(address[] _userAddresses, bool[] _freeze) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) BatchSetAddressFrozen(_userAddresses []common.Address, _freeze []bool) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchSetAddressFrozen(&_OracleDrexerc3643.TransactOpts, _userAddresses, _freeze)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] _toList, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) BatchTransfer(opts *bind.TransactOpts, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "batchTransfer", _toList, _amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] _toList, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) BatchTransfer(_toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchTransfer(&_OracleDrexerc3643.TransactOpts, _toList, _amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] _toList, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) BatchTransfer(_toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchTransfer(&_OracleDrexerc3643.TransactOpts, _toList, _amounts)
}

// BatchUnfreezePartialTokens is a paid mutator transaction binding the contract method 0x4710362d.
//
// Solidity: function batchUnfreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) BatchUnfreezePartialTokens(opts *bind.TransactOpts, _userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "batchUnfreezePartialTokens", _userAddresses, _amounts)
}

// BatchUnfreezePartialTokens is a paid mutator transaction binding the contract method 0x4710362d.
//
// Solidity: function batchUnfreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) BatchUnfreezePartialTokens(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchUnfreezePartialTokens(&_OracleDrexerc3643.TransactOpts, _userAddresses, _amounts)
}

// BatchUnfreezePartialTokens is a paid mutator transaction binding the contract method 0x4710362d.
//
// Solidity: function batchUnfreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) BatchUnfreezePartialTokens(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.BatchUnfreezePartialTokens(&_OracleDrexerc3643.TransactOpts, _userAddresses, _amounts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _userAddress, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) Burn(opts *bind.TransactOpts, _userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "burn", _userAddress, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _userAddress, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Burn(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Burn(&_OracleDrexerc3643.TransactOpts, _userAddress, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _userAddress, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) Burn(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Burn(&_OracleDrexerc3643.TransactOpts, _userAddress, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.DecreaseAllowance(&_OracleDrexerc3643.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.DecreaseAllowance(&_OracleDrexerc3643.TransactOpts, _spender, _subtractedValue)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) ForcedTransfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "forcedTransfer", _from, _to, _amount)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) ForcedTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.ForcedTransfer(&_OracleDrexerc3643.TransactOpts, _from, _to, _amount)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) ForcedTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.ForcedTransfer(&_OracleDrexerc3643.TransactOpts, _from, _to, _amount)
}

// FreezePartialTokens is a paid mutator transaction binding the contract method 0x125c4a33.
//
// Solidity: function freezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) FreezePartialTokens(opts *bind.TransactOpts, _userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "freezePartialTokens", _userAddress, _amount)
}

// FreezePartialTokens is a paid mutator transaction binding the contract method 0x125c4a33.
//
// Solidity: function freezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) FreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.FreezePartialTokens(&_OracleDrexerc3643.TransactOpts, _userAddress, _amount)
}

// FreezePartialTokens is a paid mutator transaction binding the contract method 0x125c4a33.
//
// Solidity: function freezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) FreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.FreezePartialTokens(&_OracleDrexerc3643.TransactOpts, _userAddress, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.IncreaseAllowance(&_OracleDrexerc3643.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.IncreaseAllowance(&_OracleDrexerc3643.TransactOpts, _spender, _addedValue)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address _identityRegistry, address _compliance, address _onchainID) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) Init(opts *bind.TransactOpts, _identityRegistry common.Address, _compliance common.Address, _onchainID common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "init", _identityRegistry, _compliance, _onchainID)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address _identityRegistry, address _compliance, address _onchainID) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Init(_identityRegistry common.Address, _compliance common.Address, _onchainID common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Init(&_OracleDrexerc3643.TransactOpts, _identityRegistry, _compliance, _onchainID)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address _identityRegistry, address _compliance, address _onchainID) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) Init(_identityRegistry common.Address, _compliance common.Address, _onchainID common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Init(&_OracleDrexerc3643.TransactOpts, _identityRegistry, _compliance, _onchainID)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Mint(&_OracleDrexerc3643.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Mint(&_OracleDrexerc3643.TransactOpts, _to, _amount)
}

// MintTest is a paid mutator transaction binding the contract method 0xf0f142f2.
//
// Solidity: function mintTest(address _to, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) MintTest(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "mintTest", _to, _amount)
}

// MintTest is a paid mutator transaction binding the contract method 0xf0f142f2.
//
// Solidity: function mintTest(address _to, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) MintTest(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.MintTest(&_OracleDrexerc3643.TransactOpts, _to, _amount)
}

// MintTest is a paid mutator transaction binding the contract method 0xf0f142f2.
//
// Solidity: function mintTest(address _to, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) MintTest(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.MintTest(&_OracleDrexerc3643.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Pause() (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Pause(&_OracleDrexerc3643.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) Pause() (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Pause(&_OracleDrexerc3643.TransactOpts)
}

// RecoveryAddress is a paid mutator transaction binding the contract method 0x9285948a.
//
// Solidity: function recoveryAddress(address _lostWallet, address _newWallet, address _investorOnchainID) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) RecoveryAddress(opts *bind.TransactOpts, _lostWallet common.Address, _newWallet common.Address, _investorOnchainID common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "recoveryAddress", _lostWallet, _newWallet, _investorOnchainID)
}

// RecoveryAddress is a paid mutator transaction binding the contract method 0x9285948a.
//
// Solidity: function recoveryAddress(address _lostWallet, address _newWallet, address _investorOnchainID) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) RecoveryAddress(_lostWallet common.Address, _newWallet common.Address, _investorOnchainID common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.RecoveryAddress(&_OracleDrexerc3643.TransactOpts, _lostWallet, _newWallet, _investorOnchainID)
}

// RecoveryAddress is a paid mutator transaction binding the contract method 0x9285948a.
//
// Solidity: function recoveryAddress(address _lostWallet, address _newWallet, address _investorOnchainID) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) RecoveryAddress(_lostWallet common.Address, _newWallet common.Address, _investorOnchainID common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.RecoveryAddress(&_OracleDrexerc3643.TransactOpts, _lostWallet, _newWallet, _investorOnchainID)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) RemoveAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "removeAgent", _agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.RemoveAgent(&_OracleDrexerc3643.TransactOpts, _agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.RemoveAgent(&_OracleDrexerc3643.TransactOpts, _agent)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) RenounceOwnership() (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.RenounceOwnership(&_OracleDrexerc3643.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.RenounceOwnership(&_OracleDrexerc3643.TransactOpts)
}

// SetAddressFrozen is a paid mutator transaction binding the contract method 0xc69c09cf.
//
// Solidity: function setAddressFrozen(address _userAddress, bool _freeze) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) SetAddressFrozen(opts *bind.TransactOpts, _userAddress common.Address, _freeze bool) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "setAddressFrozen", _userAddress, _freeze)
}

// SetAddressFrozen is a paid mutator transaction binding the contract method 0xc69c09cf.
//
// Solidity: function setAddressFrozen(address _userAddress, bool _freeze) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) SetAddressFrozen(_userAddress common.Address, _freeze bool) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetAddressFrozen(&_OracleDrexerc3643.TransactOpts, _userAddress, _freeze)
}

// SetAddressFrozen is a paid mutator transaction binding the contract method 0xc69c09cf.
//
// Solidity: function setAddressFrozen(address _userAddress, bool _freeze) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) SetAddressFrozen(_userAddress common.Address, _freeze bool) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetAddressFrozen(&_OracleDrexerc3643.TransactOpts, _userAddress, _freeze)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address _compliance) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) SetCompliance(opts *bind.TransactOpts, _compliance common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "setCompliance", _compliance)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address _compliance) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) SetCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetCompliance(&_OracleDrexerc3643.TransactOpts, _compliance)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address _compliance) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) SetCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetCompliance(&_OracleDrexerc3643.TransactOpts, _compliance)
}

// SetIdentityRegistry is a paid mutator transaction binding the contract method 0xcbf3f861.
//
// Solidity: function setIdentityRegistry(address _identityRegistry) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) SetIdentityRegistry(opts *bind.TransactOpts, _identityRegistry common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "setIdentityRegistry", _identityRegistry)
}

// SetIdentityRegistry is a paid mutator transaction binding the contract method 0xcbf3f861.
//
// Solidity: function setIdentityRegistry(address _identityRegistry) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) SetIdentityRegistry(_identityRegistry common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetIdentityRegistry(&_OracleDrexerc3643.TransactOpts, _identityRegistry)
}

// SetIdentityRegistry is a paid mutator transaction binding the contract method 0xcbf3f861.
//
// Solidity: function setIdentityRegistry(address _identityRegistry) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) SetIdentityRegistry(_identityRegistry common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetIdentityRegistry(&_OracleDrexerc3643.TransactOpts, _identityRegistry)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) SetName(_name string) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetName(&_OracleDrexerc3643.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetName(&_OracleDrexerc3643.TransactOpts, _name)
}

// SetOnchainID is a paid mutator transaction binding the contract method 0x3d1ddc5b.
//
// Solidity: function setOnchainID(address _onchainID) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) SetOnchainID(opts *bind.TransactOpts, _onchainID common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "setOnchainID", _onchainID)
}

// SetOnchainID is a paid mutator transaction binding the contract method 0x3d1ddc5b.
//
// Solidity: function setOnchainID(address _onchainID) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) SetOnchainID(_onchainID common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetOnchainID(&_OracleDrexerc3643.TransactOpts, _onchainID)
}

// SetOnchainID is a paid mutator transaction binding the contract method 0x3d1ddc5b.
//
// Solidity: function setOnchainID(address _onchainID) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) SetOnchainID(_onchainID common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetOnchainID(&_OracleDrexerc3643.TransactOpts, _onchainID)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) SetSymbol(opts *bind.TransactOpts, _symbol string) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "setSymbol", _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetSymbol(&_OracleDrexerc3643.TransactOpts, _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.SetSymbol(&_OracleDrexerc3643.TransactOpts, _symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Transfer(&_OracleDrexerc3643.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Transfer(&_OracleDrexerc3643.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643Session) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.TransferFrom(&_OracleDrexerc3643.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.TransferFrom(&_OracleDrexerc3643.TransactOpts, _from, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.TransferOwnership(&_OracleDrexerc3643.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.TransferOwnership(&_OracleDrexerc3643.TransactOpts, newOwner)
}

// UnfreezePartialTokens is a paid mutator transaction binding the contract method 0x1fe56f7d.
//
// Solidity: function unfreezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) UnfreezePartialTokens(opts *bind.TransactOpts, _userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "unfreezePartialTokens", _userAddress, _amount)
}

// UnfreezePartialTokens is a paid mutator transaction binding the contract method 0x1fe56f7d.
//
// Solidity: function unfreezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) UnfreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.UnfreezePartialTokens(&_OracleDrexerc3643.TransactOpts, _userAddress, _amount)
}

// UnfreezePartialTokens is a paid mutator transaction binding the contract method 0x1fe56f7d.
//
// Solidity: function unfreezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) UnfreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.UnfreezePartialTokens(&_OracleDrexerc3643.TransactOpts, _userAddress, _amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleDrexerc3643.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleDrexerc3643 *OracleDrexerc3643Session) Unpause() (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Unpause(&_OracleDrexerc3643.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleDrexerc3643 *OracleDrexerc3643TransactorSession) Unpause() (*types.Transaction, error) {
	return _OracleDrexerc3643.Contract.Unpause(&_OracleDrexerc3643.TransactOpts)
}

// OracleDrexerc3643AddressFrozenIterator is returned from FilterAddressFrozen and is used to iterate over the raw logs and unpacked data for AddressFrozen events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643AddressFrozenIterator struct {
	Event *OracleDrexerc3643AddressFrozen // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643AddressFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643AddressFrozen)
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
		it.Event = new(OracleDrexerc3643AddressFrozen)
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
func (it *OracleDrexerc3643AddressFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643AddressFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643AddressFrozen represents a AddressFrozen event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643AddressFrozen struct {
	UserAddress common.Address
	IsFrozen    bool
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddressFrozen is a free log retrieval operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: event AddressFrozen(address indexed _userAddress, bool indexed _isFrozen, address indexed _owner)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterAddressFrozen(opts *bind.FilterOpts, _userAddress []common.Address, _isFrozen []bool, _owner []common.Address) (*OracleDrexerc3643AddressFrozenIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}
	var _isFrozenRule []interface{}
	for _, _isFrozenItem := range _isFrozen {
		_isFrozenRule = append(_isFrozenRule, _isFrozenItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "AddressFrozen", _userAddressRule, _isFrozenRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643AddressFrozenIterator{contract: _OracleDrexerc3643.contract, event: "AddressFrozen", logs: logs, sub: sub}, nil
}

// WatchAddressFrozen is a free log subscription operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: event AddressFrozen(address indexed _userAddress, bool indexed _isFrozen, address indexed _owner)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchAddressFrozen(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643AddressFrozen, _userAddress []common.Address, _isFrozen []bool, _owner []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}
	var _isFrozenRule []interface{}
	for _, _isFrozenItem := range _isFrozen {
		_isFrozenRule = append(_isFrozenRule, _isFrozenItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "AddressFrozen", _userAddressRule, _isFrozenRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643AddressFrozen)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "AddressFrozen", log); err != nil {
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

// ParseAddressFrozen is a log parse operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: event AddressFrozen(address indexed _userAddress, bool indexed _isFrozen, address indexed _owner)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseAddressFrozen(log types.Log) (*OracleDrexerc3643AddressFrozen, error) {
	event := new(OracleDrexerc3643AddressFrozen)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "AddressFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643AgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643AgentAddedIterator struct {
	Event *OracleDrexerc3643AgentAdded // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643AgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643AgentAdded)
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
		it.Event = new(OracleDrexerc3643AgentAdded)
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
func (it *OracleDrexerc3643AgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643AgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643AgentAdded represents a AgentAdded event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643AgentAdded struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterAgentAdded(opts *bind.FilterOpts, _agent []common.Address) (*OracleDrexerc3643AgentAddedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "AgentAdded", _agentRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643AgentAddedIterator{contract: _OracleDrexerc3643.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643AgentAdded, _agent []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "AgentAdded", _agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643AgentAdded)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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

// ParseAgentAdded is a log parse operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseAgentAdded(log types.Log) (*OracleDrexerc3643AgentAdded, error) {
	event := new(OracleDrexerc3643AgentAdded)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643AgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643AgentRemovedIterator struct {
	Event *OracleDrexerc3643AgentRemoved // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643AgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643AgentRemoved)
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
		it.Event = new(OracleDrexerc3643AgentRemoved)
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
func (it *OracleDrexerc3643AgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643AgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643AgentRemoved represents a AgentRemoved event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643AgentRemoved struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterAgentRemoved(opts *bind.FilterOpts, _agent []common.Address) (*OracleDrexerc3643AgentRemovedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "AgentRemoved", _agentRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643AgentRemovedIterator{contract: _OracleDrexerc3643.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643AgentRemoved, _agent []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "AgentRemoved", _agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643AgentRemoved)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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

// ParseAgentRemoved is a log parse operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseAgentRemoved(log types.Log) (*OracleDrexerc3643AgentRemoved, error) {
	event := new(OracleDrexerc3643AgentRemoved)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643ApprovalIterator struct {
	Event *OracleDrexerc3643Approval // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643Approval)
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
		it.Event = new(OracleDrexerc3643Approval)
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
func (it *OracleDrexerc3643ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643Approval represents a Approval event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OracleDrexerc3643ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643ApprovalIterator{contract: _OracleDrexerc3643.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643Approval)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseApproval(log types.Log) (*OracleDrexerc3643Approval, error) {
	event := new(OracleDrexerc3643Approval)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643ComplianceAddedIterator is returned from FilterComplianceAdded and is used to iterate over the raw logs and unpacked data for ComplianceAdded events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643ComplianceAddedIterator struct {
	Event *OracleDrexerc3643ComplianceAdded // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643ComplianceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643ComplianceAdded)
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
		it.Event = new(OracleDrexerc3643ComplianceAdded)
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
func (it *OracleDrexerc3643ComplianceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643ComplianceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643ComplianceAdded represents a ComplianceAdded event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643ComplianceAdded struct {
	Compliance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterComplianceAdded is a free log retrieval operation binding the contract event 0x7f3a888862559648ec01d97deb7b5012bff86dc91e654a1de397170db40e35b6.
//
// Solidity: event ComplianceAdded(address indexed _compliance)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterComplianceAdded(opts *bind.FilterOpts, _compliance []common.Address) (*OracleDrexerc3643ComplianceAddedIterator, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "ComplianceAdded", _complianceRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643ComplianceAddedIterator{contract: _OracleDrexerc3643.contract, event: "ComplianceAdded", logs: logs, sub: sub}, nil
}

// WatchComplianceAdded is a free log subscription operation binding the contract event 0x7f3a888862559648ec01d97deb7b5012bff86dc91e654a1de397170db40e35b6.
//
// Solidity: event ComplianceAdded(address indexed _compliance)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchComplianceAdded(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643ComplianceAdded, _compliance []common.Address) (event.Subscription, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "ComplianceAdded", _complianceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643ComplianceAdded)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "ComplianceAdded", log); err != nil {
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

// ParseComplianceAdded is a log parse operation binding the contract event 0x7f3a888862559648ec01d97deb7b5012bff86dc91e654a1de397170db40e35b6.
//
// Solidity: event ComplianceAdded(address indexed _compliance)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseComplianceAdded(log types.Log) (*OracleDrexerc3643ComplianceAdded, error) {
	event := new(OracleDrexerc3643ComplianceAdded)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "ComplianceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643IdentityRegistryAddedIterator is returned from FilterIdentityRegistryAdded and is used to iterate over the raw logs and unpacked data for IdentityRegistryAdded events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643IdentityRegistryAddedIterator struct {
	Event *OracleDrexerc3643IdentityRegistryAdded // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643IdentityRegistryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643IdentityRegistryAdded)
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
		it.Event = new(OracleDrexerc3643IdentityRegistryAdded)
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
func (it *OracleDrexerc3643IdentityRegistryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643IdentityRegistryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643IdentityRegistryAdded represents a IdentityRegistryAdded event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643IdentityRegistryAdded struct {
	IdentityRegistry common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIdentityRegistryAdded is a free log retrieval operation binding the contract event 0xd2be862d755bca7e0d39772b2cab3a5578da9c285f69199f4c063c2294a7f36c.
//
// Solidity: event IdentityRegistryAdded(address indexed _identityRegistry)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterIdentityRegistryAdded(opts *bind.FilterOpts, _identityRegistry []common.Address) (*OracleDrexerc3643IdentityRegistryAddedIterator, error) {

	var _identityRegistryRule []interface{}
	for _, _identityRegistryItem := range _identityRegistry {
		_identityRegistryRule = append(_identityRegistryRule, _identityRegistryItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "IdentityRegistryAdded", _identityRegistryRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643IdentityRegistryAddedIterator{contract: _OracleDrexerc3643.contract, event: "IdentityRegistryAdded", logs: logs, sub: sub}, nil
}

// WatchIdentityRegistryAdded is a free log subscription operation binding the contract event 0xd2be862d755bca7e0d39772b2cab3a5578da9c285f69199f4c063c2294a7f36c.
//
// Solidity: event IdentityRegistryAdded(address indexed _identityRegistry)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchIdentityRegistryAdded(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643IdentityRegistryAdded, _identityRegistry []common.Address) (event.Subscription, error) {

	var _identityRegistryRule []interface{}
	for _, _identityRegistryItem := range _identityRegistry {
		_identityRegistryRule = append(_identityRegistryRule, _identityRegistryItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "IdentityRegistryAdded", _identityRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643IdentityRegistryAdded)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "IdentityRegistryAdded", log); err != nil {
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

// ParseIdentityRegistryAdded is a log parse operation binding the contract event 0xd2be862d755bca7e0d39772b2cab3a5578da9c285f69199f4c063c2294a7f36c.
//
// Solidity: event IdentityRegistryAdded(address indexed _identityRegistry)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseIdentityRegistryAdded(log types.Log) (*OracleDrexerc3643IdentityRegistryAdded, error) {
	event := new(OracleDrexerc3643IdentityRegistryAdded)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "IdentityRegistryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643InitializedIterator struct {
	Event *OracleDrexerc3643Initialized // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643Initialized)
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
		it.Event = new(OracleDrexerc3643Initialized)
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
func (it *OracleDrexerc3643InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643Initialized represents a Initialized event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterInitialized(opts *bind.FilterOpts) (*OracleDrexerc3643InitializedIterator, error) {

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643InitializedIterator{contract: _OracleDrexerc3643.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643Initialized) (event.Subscription, error) {

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643Initialized)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseInitialized(log types.Log) (*OracleDrexerc3643Initialized, error) {
	event := new(OracleDrexerc3643Initialized)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643OwnershipTransferredIterator struct {
	Event *OracleDrexerc3643OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643OwnershipTransferred)
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
		it.Event = new(OracleDrexerc3643OwnershipTransferred)
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
func (it *OracleDrexerc3643OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643OwnershipTransferred represents a OwnershipTransferred event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleDrexerc3643OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643OwnershipTransferredIterator{contract: _OracleDrexerc3643.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643OwnershipTransferred)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseOwnershipTransferred(log types.Log) (*OracleDrexerc3643OwnershipTransferred, error) {
	event := new(OracleDrexerc3643OwnershipTransferred)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643PausedIterator struct {
	Event *OracleDrexerc3643Paused // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643Paused)
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
		it.Event = new(OracleDrexerc3643Paused)
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
func (it *OracleDrexerc3643PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643Paused represents a Paused event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643Paused struct {
	UserAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address _userAddress)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterPaused(opts *bind.FilterOpts) (*OracleDrexerc3643PausedIterator, error) {

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643PausedIterator{contract: _OracleDrexerc3643.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address _userAddress)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643Paused) (event.Subscription, error) {

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643Paused)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address _userAddress)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParsePaused(log types.Log) (*OracleDrexerc3643Paused, error) {
	event := new(OracleDrexerc3643Paused)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643RecoverySuccessIterator is returned from FilterRecoverySuccess and is used to iterate over the raw logs and unpacked data for RecoverySuccess events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643RecoverySuccessIterator struct {
	Event *OracleDrexerc3643RecoverySuccess // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643RecoverySuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643RecoverySuccess)
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
		it.Event = new(OracleDrexerc3643RecoverySuccess)
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
func (it *OracleDrexerc3643RecoverySuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643RecoverySuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643RecoverySuccess represents a RecoverySuccess event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643RecoverySuccess struct {
	LostWallet        common.Address
	NewWallet         common.Address
	InvestorOnchainID common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRecoverySuccess is a free log retrieval operation binding the contract event 0xf0c9129a94f30f1caaceb63e44b9811d0a3edf1d6c23757f346093af5553fed0.
//
// Solidity: event RecoverySuccess(address indexed _lostWallet, address indexed _newWallet, address indexed _investorOnchainID)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterRecoverySuccess(opts *bind.FilterOpts, _lostWallet []common.Address, _newWallet []common.Address, _investorOnchainID []common.Address) (*OracleDrexerc3643RecoverySuccessIterator, error) {

	var _lostWalletRule []interface{}
	for _, _lostWalletItem := range _lostWallet {
		_lostWalletRule = append(_lostWalletRule, _lostWalletItem)
	}
	var _newWalletRule []interface{}
	for _, _newWalletItem := range _newWallet {
		_newWalletRule = append(_newWalletRule, _newWalletItem)
	}
	var _investorOnchainIDRule []interface{}
	for _, _investorOnchainIDItem := range _investorOnchainID {
		_investorOnchainIDRule = append(_investorOnchainIDRule, _investorOnchainIDItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "RecoverySuccess", _lostWalletRule, _newWalletRule, _investorOnchainIDRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643RecoverySuccessIterator{contract: _OracleDrexerc3643.contract, event: "RecoverySuccess", logs: logs, sub: sub}, nil
}

// WatchRecoverySuccess is a free log subscription operation binding the contract event 0xf0c9129a94f30f1caaceb63e44b9811d0a3edf1d6c23757f346093af5553fed0.
//
// Solidity: event RecoverySuccess(address indexed _lostWallet, address indexed _newWallet, address indexed _investorOnchainID)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchRecoverySuccess(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643RecoverySuccess, _lostWallet []common.Address, _newWallet []common.Address, _investorOnchainID []common.Address) (event.Subscription, error) {

	var _lostWalletRule []interface{}
	for _, _lostWalletItem := range _lostWallet {
		_lostWalletRule = append(_lostWalletRule, _lostWalletItem)
	}
	var _newWalletRule []interface{}
	for _, _newWalletItem := range _newWallet {
		_newWalletRule = append(_newWalletRule, _newWalletItem)
	}
	var _investorOnchainIDRule []interface{}
	for _, _investorOnchainIDItem := range _investorOnchainID {
		_investorOnchainIDRule = append(_investorOnchainIDRule, _investorOnchainIDItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "RecoverySuccess", _lostWalletRule, _newWalletRule, _investorOnchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643RecoverySuccess)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "RecoverySuccess", log); err != nil {
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

// ParseRecoverySuccess is a log parse operation binding the contract event 0xf0c9129a94f30f1caaceb63e44b9811d0a3edf1d6c23757f346093af5553fed0.
//
// Solidity: event RecoverySuccess(address indexed _lostWallet, address indexed _newWallet, address indexed _investorOnchainID)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseRecoverySuccess(log types.Log) (*OracleDrexerc3643RecoverySuccess, error) {
	event := new(OracleDrexerc3643RecoverySuccess)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "RecoverySuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643TokensFrozenIterator is returned from FilterTokensFrozen and is used to iterate over the raw logs and unpacked data for TokensFrozen events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643TokensFrozenIterator struct {
	Event *OracleDrexerc3643TokensFrozen // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643TokensFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643TokensFrozen)
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
		it.Event = new(OracleDrexerc3643TokensFrozen)
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
func (it *OracleDrexerc3643TokensFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643TokensFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643TokensFrozen represents a TokensFrozen event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643TokensFrozen struct {
	UserAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensFrozen is a free log retrieval operation binding the contract event 0xa065e63c631c86f1b9f66a4a2f63f2093bf1c2168d23290259dbd969e0222a45.
//
// Solidity: event TokensFrozen(address indexed _userAddress, uint256 _amount)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterTokensFrozen(opts *bind.FilterOpts, _userAddress []common.Address) (*OracleDrexerc3643TokensFrozenIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "TokensFrozen", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643TokensFrozenIterator{contract: _OracleDrexerc3643.contract, event: "TokensFrozen", logs: logs, sub: sub}, nil
}

// WatchTokensFrozen is a free log subscription operation binding the contract event 0xa065e63c631c86f1b9f66a4a2f63f2093bf1c2168d23290259dbd969e0222a45.
//
// Solidity: event TokensFrozen(address indexed _userAddress, uint256 _amount)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchTokensFrozen(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643TokensFrozen, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "TokensFrozen", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643TokensFrozen)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "TokensFrozen", log); err != nil {
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

// ParseTokensFrozen is a log parse operation binding the contract event 0xa065e63c631c86f1b9f66a4a2f63f2093bf1c2168d23290259dbd969e0222a45.
//
// Solidity: event TokensFrozen(address indexed _userAddress, uint256 _amount)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseTokensFrozen(log types.Log) (*OracleDrexerc3643TokensFrozen, error) {
	event := new(OracleDrexerc3643TokensFrozen)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "TokensFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643TokensUnfrozenIterator is returned from FilterTokensUnfrozen and is used to iterate over the raw logs and unpacked data for TokensUnfrozen events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643TokensUnfrozenIterator struct {
	Event *OracleDrexerc3643TokensUnfrozen // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643TokensUnfrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643TokensUnfrozen)
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
		it.Event = new(OracleDrexerc3643TokensUnfrozen)
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
func (it *OracleDrexerc3643TokensUnfrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643TokensUnfrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643TokensUnfrozen represents a TokensUnfrozen event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643TokensUnfrozen struct {
	UserAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensUnfrozen is a free log retrieval operation binding the contract event 0x9bed35cb62ad0dba04f9d5bfee4b5bc91443e77da8a65c4c84834c51bb08b0d6.
//
// Solidity: event TokensUnfrozen(address indexed _userAddress, uint256 _amount)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterTokensUnfrozen(opts *bind.FilterOpts, _userAddress []common.Address) (*OracleDrexerc3643TokensUnfrozenIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "TokensUnfrozen", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643TokensUnfrozenIterator{contract: _OracleDrexerc3643.contract, event: "TokensUnfrozen", logs: logs, sub: sub}, nil
}

// WatchTokensUnfrozen is a free log subscription operation binding the contract event 0x9bed35cb62ad0dba04f9d5bfee4b5bc91443e77da8a65c4c84834c51bb08b0d6.
//
// Solidity: event TokensUnfrozen(address indexed _userAddress, uint256 _amount)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchTokensUnfrozen(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643TokensUnfrozen, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "TokensUnfrozen", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643TokensUnfrozen)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "TokensUnfrozen", log); err != nil {
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

// ParseTokensUnfrozen is a log parse operation binding the contract event 0x9bed35cb62ad0dba04f9d5bfee4b5bc91443e77da8a65c4c84834c51bb08b0d6.
//
// Solidity: event TokensUnfrozen(address indexed _userAddress, uint256 _amount)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseTokensUnfrozen(log types.Log) (*OracleDrexerc3643TokensUnfrozen, error) {
	event := new(OracleDrexerc3643TokensUnfrozen)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "TokensUnfrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643TransferIterator struct {
	Event *OracleDrexerc3643Transfer // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643Transfer)
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
		it.Event = new(OracleDrexerc3643Transfer)
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
func (it *OracleDrexerc3643TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643Transfer represents a Transfer event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OracleDrexerc3643TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643TransferIterator{contract: _OracleDrexerc3643.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643Transfer)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseTransfer(log types.Log) (*OracleDrexerc3643Transfer, error) {
	event := new(OracleDrexerc3643Transfer)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643UnpausedIterator struct {
	Event *OracleDrexerc3643Unpaused // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643Unpaused)
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
		it.Event = new(OracleDrexerc3643Unpaused)
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
func (it *OracleDrexerc3643UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643Unpaused represents a Unpaused event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643Unpaused struct {
	UserAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address _userAddress)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterUnpaused(opts *bind.FilterOpts) (*OracleDrexerc3643UnpausedIterator, error) {

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643UnpausedIterator{contract: _OracleDrexerc3643.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address _userAddress)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643Unpaused) (event.Subscription, error) {

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643Unpaused)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address _userAddress)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseUnpaused(log types.Log) (*OracleDrexerc3643Unpaused, error) {
	event := new(OracleDrexerc3643Unpaused)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleDrexerc3643UpdatedTokenInformationIterator is returned from FilterUpdatedTokenInformation and is used to iterate over the raw logs and unpacked data for UpdatedTokenInformation events raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643UpdatedTokenInformationIterator struct {
	Event *OracleDrexerc3643UpdatedTokenInformation // Event containing the contract specifics and raw log

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
func (it *OracleDrexerc3643UpdatedTokenInformationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleDrexerc3643UpdatedTokenInformation)
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
		it.Event = new(OracleDrexerc3643UpdatedTokenInformation)
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
func (it *OracleDrexerc3643UpdatedTokenInformationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleDrexerc3643UpdatedTokenInformationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleDrexerc3643UpdatedTokenInformation represents a UpdatedTokenInformation event raised by the OracleDrexerc3643 contract.
type OracleDrexerc3643UpdatedTokenInformation struct {
	NewName      common.Hash
	NewSymbol    common.Hash
	NewDecimals  uint8
	NewVersion   string
	NewOnchainID common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenInformation is a free log retrieval operation binding the contract event 0x6a1105ac8148a3c319adbc369f9072573e8a11d3a3d195e067e7c40767ec54d1.
//
// Solidity: event UpdatedTokenInformation(string indexed _newName, string indexed _newSymbol, uint8 _newDecimals, string _newVersion, address indexed _newOnchainID)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) FilterUpdatedTokenInformation(opts *bind.FilterOpts, _newName []string, _newSymbol []string, _newOnchainID []common.Address) (*OracleDrexerc3643UpdatedTokenInformationIterator, error) {

	var _newNameRule []interface{}
	for _, _newNameItem := range _newName {
		_newNameRule = append(_newNameRule, _newNameItem)
	}
	var _newSymbolRule []interface{}
	for _, _newSymbolItem := range _newSymbol {
		_newSymbolRule = append(_newSymbolRule, _newSymbolItem)
	}

	var _newOnchainIDRule []interface{}
	for _, _newOnchainIDItem := range _newOnchainID {
		_newOnchainIDRule = append(_newOnchainIDRule, _newOnchainIDItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.FilterLogs(opts, "UpdatedTokenInformation", _newNameRule, _newSymbolRule, _newOnchainIDRule)
	if err != nil {
		return nil, err
	}
	return &OracleDrexerc3643UpdatedTokenInformationIterator{contract: _OracleDrexerc3643.contract, event: "UpdatedTokenInformation", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenInformation is a free log subscription operation binding the contract event 0x6a1105ac8148a3c319adbc369f9072573e8a11d3a3d195e067e7c40767ec54d1.
//
// Solidity: event UpdatedTokenInformation(string indexed _newName, string indexed _newSymbol, uint8 _newDecimals, string _newVersion, address indexed _newOnchainID)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) WatchUpdatedTokenInformation(opts *bind.WatchOpts, sink chan<- *OracleDrexerc3643UpdatedTokenInformation, _newName []string, _newSymbol []string, _newOnchainID []common.Address) (event.Subscription, error) {

	var _newNameRule []interface{}
	for _, _newNameItem := range _newName {
		_newNameRule = append(_newNameRule, _newNameItem)
	}
	var _newSymbolRule []interface{}
	for _, _newSymbolItem := range _newSymbol {
		_newSymbolRule = append(_newSymbolRule, _newSymbolItem)
	}

	var _newOnchainIDRule []interface{}
	for _, _newOnchainIDItem := range _newOnchainID {
		_newOnchainIDRule = append(_newOnchainIDRule, _newOnchainIDItem)
	}

	logs, sub, err := _OracleDrexerc3643.contract.WatchLogs(opts, "UpdatedTokenInformation", _newNameRule, _newSymbolRule, _newOnchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleDrexerc3643UpdatedTokenInformation)
				if err := _OracleDrexerc3643.contract.UnpackLog(event, "UpdatedTokenInformation", log); err != nil {
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

// ParseUpdatedTokenInformation is a log parse operation binding the contract event 0x6a1105ac8148a3c319adbc369f9072573e8a11d3a3d195e067e7c40767ec54d1.
//
// Solidity: event UpdatedTokenInformation(string indexed _newName, string indexed _newSymbol, uint8 _newDecimals, string _newVersion, address indexed _newOnchainID)
func (_OracleDrexerc3643 *OracleDrexerc3643Filterer) ParseUpdatedTokenInformation(log types.Log) (*OracleDrexerc3643UpdatedTokenInformation, error) {
	event := new(OracleDrexerc3643UpdatedTokenInformation)
	if err := _OracleDrexerc3643.contract.UnpackLog(event, "UpdatedTokenInformation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
