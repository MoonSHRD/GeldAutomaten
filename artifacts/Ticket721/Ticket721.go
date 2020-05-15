// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Ticket721

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Ticket721ABI is the input ABI used to generate the binding from.
const Ticket721ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ticket_sale\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"}],\"name\":\"EventIdReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ticket_sale\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"}],\"name\":\"EventIdReservedHuman\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"visitor_wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ticket_id\",\"type\":\"uint256\"}],\"name\":\"TicketBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"visitor_wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ticket_id\",\"type\":\"uint256\"}],\"name\":\"TicketBoughtHuman\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"visitor_wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ticket_id\",\"type\":\"uint256\"}],\"name\":\"TicketFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"visitor_wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ticket_id\",\"type\":\"uint256\"}],\"name\":\"TicketFulfilledHuman\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"JIDs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"_transferFromTicket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ticketAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ticket_type\",\"type\":\"uint256\"}],\"name\":\"buyTicket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventsales\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ticket_id\",\"type\":\"uint256\"}],\"name\":\"getJidByTicketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"jid\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"}],\"name\":\"getJidbyEventID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"jid\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getTicketByOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"}],\"name\":\"getTicketSales\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"}],\"name\":\"getTicketTypeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"orginizer\",\"type\":\"address\"}],\"name\":\"plugSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"visitor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"}],\"name\":\"redeemTicket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"orginizer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"jid\",\"type\":\"string\"}],\"name\":\"reserveEventId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"event_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ticketIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ticketInfoStorage\",\"outputs\":[{\"internalType\":\"enumTicket721.TicketState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ticket_type\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"event_JID\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ticket721 is an auto generated Go binding around an Ethereum contract.
type Ticket721 struct {
	Ticket721Caller     // Read-only binding to the contract
	Ticket721Transactor // Write-only binding to the contract
	Ticket721Filterer   // Log filterer for contract events
}

// Ticket721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ticket721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ticket721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ticket721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ticket721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ticket721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ticket721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ticket721Session struct {
	Contract     *Ticket721        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ticket721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ticket721CallerSession struct {
	Contract *Ticket721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Ticket721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ticket721TransactorSession struct {
	Contract     *Ticket721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Ticket721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ticket721Raw struct {
	Contract *Ticket721 // Generic contract binding to access the raw methods on
}

// Ticket721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ticket721CallerRaw struct {
	Contract *Ticket721Caller // Generic read-only contract binding to access the raw methods on
}

// Ticket721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ticket721TransactorRaw struct {
	Contract *Ticket721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTicket721 creates a new instance of Ticket721, bound to a specific deployed contract.
func NewTicket721(address common.Address, backend bind.ContractBackend) (*Ticket721, error) {
	contract, err := bindTicket721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ticket721{Ticket721Caller: Ticket721Caller{contract: contract}, Ticket721Transactor: Ticket721Transactor{contract: contract}, Ticket721Filterer: Ticket721Filterer{contract: contract}}, nil
}

// NewTicket721Caller creates a new read-only instance of Ticket721, bound to a specific deployed contract.
func NewTicket721Caller(address common.Address, caller bind.ContractCaller) (*Ticket721Caller, error) {
	contract, err := bindTicket721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ticket721Caller{contract: contract}, nil
}

// NewTicket721Transactor creates a new write-only instance of Ticket721, bound to a specific deployed contract.
func NewTicket721Transactor(address common.Address, transactor bind.ContractTransactor) (*Ticket721Transactor, error) {
	contract, err := bindTicket721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ticket721Transactor{contract: contract}, nil
}

// NewTicket721Filterer creates a new log filterer instance of Ticket721, bound to a specific deployed contract.
func NewTicket721Filterer(address common.Address, filterer bind.ContractFilterer) (*Ticket721Filterer, error) {
	contract, err := bindTicket721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ticket721Filterer{contract: contract}, nil
}

// bindTicket721 binds a generic wrapper to an already deployed contract.
func bindTicket721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ticket721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ticket721 *Ticket721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ticket721.Contract.Ticket721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ticket721 *Ticket721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ticket721.Contract.Ticket721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ticket721 *Ticket721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ticket721.Contract.Ticket721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ticket721 *Ticket721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ticket721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ticket721 *Ticket721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ticket721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ticket721 *Ticket721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ticket721.Contract.contract.Transact(opts, method, params...)
}

// JIDs is a free data retrieval call binding the contract method 0xad140b1d.
//
// Solidity: function JIDs(uint256 ) view returns(string)
func (_Ticket721 *Ticket721Caller) JIDs(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "JIDs", arg0)
	return *ret0, err
}

// JIDs is a free data retrieval call binding the contract method 0xad140b1d.
//
// Solidity: function JIDs(uint256 ) view returns(string)
func (_Ticket721 *Ticket721Session) JIDs(arg0 *big.Int) (string, error) {
	return _Ticket721.Contract.JIDs(&_Ticket721.CallOpts, arg0)
}

// JIDs is a free data retrieval call binding the contract method 0xad140b1d.
//
// Solidity: function JIDs(uint256 ) view returns(string)
func (_Ticket721 *Ticket721CallerSession) JIDs(arg0 *big.Int) (string, error) {
	return _Ticket721.Contract.JIDs(&_Ticket721.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ticket721 *Ticket721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ticket721 *Ticket721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ticket721.Contract.BalanceOf(&_Ticket721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ticket721 *Ticket721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ticket721.Contract.BalanceOf(&_Ticket721.CallOpts, owner)
}

// Eventsales is a free data retrieval call binding the contract method 0x9b1d8baf.
//
// Solidity: function eventsales(uint256 , uint256 ) view returns(address)
func (_Ticket721 *Ticket721Caller) Eventsales(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "eventsales", arg0, arg1)
	return *ret0, err
}

// Eventsales is a free data retrieval call binding the contract method 0x9b1d8baf.
//
// Solidity: function eventsales(uint256 , uint256 ) view returns(address)
func (_Ticket721 *Ticket721Session) Eventsales(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Ticket721.Contract.Eventsales(&_Ticket721.CallOpts, arg0, arg1)
}

// Eventsales is a free data retrieval call binding the contract method 0x9b1d8baf.
//
// Solidity: function eventsales(uint256 , uint256 ) view returns(address)
func (_Ticket721 *Ticket721CallerSession) Eventsales(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Ticket721.Contract.Eventsales(&_Ticket721.CallOpts, arg0, arg1)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ticket721 *Ticket721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ticket721 *Ticket721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ticket721.Contract.GetApproved(&_Ticket721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ticket721 *Ticket721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ticket721.Contract.GetApproved(&_Ticket721.CallOpts, tokenId)
}

// GetJidByTicketId is a free data retrieval call binding the contract method 0x1eb8098e.
//
// Solidity: function getJidByTicketId(uint256 ticket_id) view returns(string jid)
func (_Ticket721 *Ticket721Caller) GetJidByTicketId(opts *bind.CallOpts, ticket_id *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "getJidByTicketId", ticket_id)
	return *ret0, err
}

// GetJidByTicketId is a free data retrieval call binding the contract method 0x1eb8098e.
//
// Solidity: function getJidByTicketId(uint256 ticket_id) view returns(string jid)
func (_Ticket721 *Ticket721Session) GetJidByTicketId(ticket_id *big.Int) (string, error) {
	return _Ticket721.Contract.GetJidByTicketId(&_Ticket721.CallOpts, ticket_id)
}

// GetJidByTicketId is a free data retrieval call binding the contract method 0x1eb8098e.
//
// Solidity: function getJidByTicketId(uint256 ticket_id) view returns(string jid)
func (_Ticket721 *Ticket721CallerSession) GetJidByTicketId(ticket_id *big.Int) (string, error) {
	return _Ticket721.Contract.GetJidByTicketId(&_Ticket721.CallOpts, ticket_id)
}

// GetJidbyEventID is a free data retrieval call binding the contract method 0x85c393f8.
//
// Solidity: function getJidbyEventID(uint256 event_id) view returns(string jid)
func (_Ticket721 *Ticket721Caller) GetJidbyEventID(opts *bind.CallOpts, event_id *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "getJidbyEventID", event_id)
	return *ret0, err
}

// GetJidbyEventID is a free data retrieval call binding the contract method 0x85c393f8.
//
// Solidity: function getJidbyEventID(uint256 event_id) view returns(string jid)
func (_Ticket721 *Ticket721Session) GetJidbyEventID(event_id *big.Int) (string, error) {
	return _Ticket721.Contract.GetJidbyEventID(&_Ticket721.CallOpts, event_id)
}

// GetJidbyEventID is a free data retrieval call binding the contract method 0x85c393f8.
//
// Solidity: function getJidbyEventID(uint256 event_id) view returns(string jid)
func (_Ticket721 *Ticket721CallerSession) GetJidbyEventID(event_id *big.Int) (string, error) {
	return _Ticket721.Contract.GetJidbyEventID(&_Ticket721.CallOpts, event_id)
}

// GetTicketByOwner is a free data retrieval call binding the contract method 0x376b029a.
//
// Solidity: function getTicketByOwner(address _owner) view returns(uint256[])
func (_Ticket721 *Ticket721Caller) GetTicketByOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "getTicketByOwner", _owner)
	return *ret0, err
}

// GetTicketByOwner is a free data retrieval call binding the contract method 0x376b029a.
//
// Solidity: function getTicketByOwner(address _owner) view returns(uint256[])
func (_Ticket721 *Ticket721Session) GetTicketByOwner(_owner common.Address) ([]*big.Int, error) {
	return _Ticket721.Contract.GetTicketByOwner(&_Ticket721.CallOpts, _owner)
}

// GetTicketByOwner is a free data retrieval call binding the contract method 0x376b029a.
//
// Solidity: function getTicketByOwner(address _owner) view returns(uint256[])
func (_Ticket721 *Ticket721CallerSession) GetTicketByOwner(_owner common.Address) ([]*big.Int, error) {
	return _Ticket721.Contract.GetTicketByOwner(&_Ticket721.CallOpts, _owner)
}

// GetTicketSales is a free data retrieval call binding the contract method 0x44bb8aff.
//
// Solidity: function getTicketSales(uint256 event_id) view returns(address[])
func (_Ticket721 *Ticket721Caller) GetTicketSales(opts *bind.CallOpts, event_id *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "getTicketSales", event_id)
	return *ret0, err
}

// GetTicketSales is a free data retrieval call binding the contract method 0x44bb8aff.
//
// Solidity: function getTicketSales(uint256 event_id) view returns(address[])
func (_Ticket721 *Ticket721Session) GetTicketSales(event_id *big.Int) ([]common.Address, error) {
	return _Ticket721.Contract.GetTicketSales(&_Ticket721.CallOpts, event_id)
}

// GetTicketSales is a free data retrieval call binding the contract method 0x44bb8aff.
//
// Solidity: function getTicketSales(uint256 event_id) view returns(address[])
func (_Ticket721 *Ticket721CallerSession) GetTicketSales(event_id *big.Int) ([]common.Address, error) {
	return _Ticket721.Contract.GetTicketSales(&_Ticket721.CallOpts, event_id)
}

// GetTicketTypeCount is a free data retrieval call binding the contract method 0xec95d2d0.
//
// Solidity: function getTicketTypeCount(uint256 event_id) view returns(uint256)
func (_Ticket721 *Ticket721Caller) GetTicketTypeCount(opts *bind.CallOpts, event_id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "getTicketTypeCount", event_id)
	return *ret0, err
}

// GetTicketTypeCount is a free data retrieval call binding the contract method 0xec95d2d0.
//
// Solidity: function getTicketTypeCount(uint256 event_id) view returns(uint256)
func (_Ticket721 *Ticket721Session) GetTicketTypeCount(event_id *big.Int) (*big.Int, error) {
	return _Ticket721.Contract.GetTicketTypeCount(&_Ticket721.CallOpts, event_id)
}

// GetTicketTypeCount is a free data retrieval call binding the contract method 0xec95d2d0.
//
// Solidity: function getTicketTypeCount(uint256 event_id) view returns(uint256)
func (_Ticket721 *Ticket721CallerSession) GetTicketTypeCount(event_id *big.Int) (*big.Int, error) {
	return _Ticket721.Contract.GetTicketTypeCount(&_Ticket721.CallOpts, event_id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ticket721 *Ticket721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ticket721 *Ticket721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ticket721.Contract.IsApprovedForAll(&_Ticket721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ticket721 *Ticket721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ticket721.Contract.IsApprovedForAll(&_Ticket721.CallOpts, owner, operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Ticket721 *Ticket721Caller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Ticket721 *Ticket721Session) IsMinter(account common.Address) (bool, error) {
	return _Ticket721.Contract.IsMinter(&_Ticket721.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Ticket721 *Ticket721CallerSession) IsMinter(account common.Address) (bool, error) {
	return _Ticket721.Contract.IsMinter(&_Ticket721.CallOpts, account)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ticket721 *Ticket721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ticket721 *Ticket721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ticket721.Contract.OwnerOf(&_Ticket721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ticket721 *Ticket721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ticket721.Contract.OwnerOf(&_Ticket721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ticket721 *Ticket721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ticket721 *Ticket721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ticket721.Contract.SupportsInterface(&_Ticket721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ticket721 *Ticket721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ticket721.Contract.SupportsInterface(&_Ticket721.CallOpts, interfaceId)
}

// TicketIds is a free data retrieval call binding the contract method 0x78df3a52.
//
// Solidity: function ticketIds(uint256 , uint256 ) view returns(uint256)
func (_Ticket721 *Ticket721Caller) TicketIds(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "ticketIds", arg0, arg1)
	return *ret0, err
}

// TicketIds is a free data retrieval call binding the contract method 0x78df3a52.
//
// Solidity: function ticketIds(uint256 , uint256 ) view returns(uint256)
func (_Ticket721 *Ticket721Session) TicketIds(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Ticket721.Contract.TicketIds(&_Ticket721.CallOpts, arg0, arg1)
}

// TicketIds is a free data retrieval call binding the contract method 0x78df3a52.
//
// Solidity: function ticketIds(uint256 , uint256 ) view returns(uint256)
func (_Ticket721 *Ticket721CallerSession) TicketIds(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Ticket721.Contract.TicketIds(&_Ticket721.CallOpts, arg0, arg1)
}

// TicketInfoStorage is a free data retrieval call binding the contract method 0x1d663f10.
//
// Solidity: function ticketInfoStorage(uint256 ) view returns(uint8 state, uint256 ticket_type, string event_JID)
func (_Ticket721 *Ticket721Caller) TicketInfoStorage(opts *bind.CallOpts, arg0 *big.Int) (struct {
	State      uint8
	TicketType *big.Int
	EventJID   string
}, error) {
	ret := new(struct {
		State      uint8
		TicketType *big.Int
		EventJID   string
	})
	out := ret
	err := _Ticket721.contract.Call(opts, out, "ticketInfoStorage", arg0)
	return *ret, err
}

// TicketInfoStorage is a free data retrieval call binding the contract method 0x1d663f10.
//
// Solidity: function ticketInfoStorage(uint256 ) view returns(uint8 state, uint256 ticket_type, string event_JID)
func (_Ticket721 *Ticket721Session) TicketInfoStorage(arg0 *big.Int) (struct {
	State      uint8
	TicketType *big.Int
	EventJID   string
}, error) {
	return _Ticket721.Contract.TicketInfoStorage(&_Ticket721.CallOpts, arg0)
}

// TicketInfoStorage is a free data retrieval call binding the contract method 0x1d663f10.
//
// Solidity: function ticketInfoStorage(uint256 ) view returns(uint8 state, uint256 ticket_type, string event_JID)
func (_Ticket721 *Ticket721CallerSession) TicketInfoStorage(arg0 *big.Int) (struct {
	State      uint8
	TicketType *big.Int
	EventJID   string
}, error) {
	return _Ticket721.Contract.TicketInfoStorage(&_Ticket721.CallOpts, arg0)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Ticket721 *Ticket721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Ticket721 *Ticket721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Ticket721.Contract.TokenByIndex(&_Ticket721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Ticket721 *Ticket721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Ticket721.Contract.TokenByIndex(&_Ticket721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Ticket721 *Ticket721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Ticket721 *Ticket721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Ticket721.Contract.TokenOfOwnerByIndex(&_Ticket721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Ticket721 *Ticket721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Ticket721.Contract.TokenOfOwnerByIndex(&_Ticket721.CallOpts, owner, index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ticket721 *Ticket721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ticket721.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ticket721 *Ticket721Session) TotalSupply() (*big.Int, error) {
	return _Ticket721.Contract.TotalSupply(&_Ticket721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ticket721 *Ticket721CallerSession) TotalSupply() (*big.Int, error) {
	return _Ticket721.Contract.TotalSupply(&_Ticket721.CallOpts)
}

// TransferFromTicket is a paid mutator transaction binding the contract method 0xab31520a.
//
// Solidity: function _transferFromTicket(address from, address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721Transactor) TransferFromTicket(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "_transferFromTicket", from, to, tokenId)
}

// TransferFromTicket is a paid mutator transaction binding the contract method 0xab31520a.
//
// Solidity: function _transferFromTicket(address from, address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721Session) TransferFromTicket(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.TransferFromTicket(&_Ticket721.TransactOpts, from, to, tokenId)
}

// TransferFromTicket is a paid mutator transaction binding the contract method 0xab31520a.
//
// Solidity: function _transferFromTicket(address from, address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721TransactorSession) TransferFromTicket(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.TransferFromTicket(&_Ticket721.TransactOpts, from, to, tokenId)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_Ticket721 *Ticket721Transactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_Ticket721 *Ticket721Session) AddMinter(account common.Address) (*types.Transaction, error) {
	return _Ticket721.Contract.AddMinter(&_Ticket721.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_Ticket721 *Ticket721TransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _Ticket721.Contract.AddMinter(&_Ticket721.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.Approve(&_Ticket721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.Approve(&_Ticket721.TransactOpts, to, tokenId)
}

// BuyTicket is a paid mutator transaction binding the contract method 0xcd3a15f8.
//
// Solidity: function buyTicket(address buyer, uint256 ticketAmount, uint256 event_id, uint256 _ticket_type) returns()
func (_Ticket721 *Ticket721Transactor) BuyTicket(opts *bind.TransactOpts, buyer common.Address, ticketAmount *big.Int, event_id *big.Int, _ticket_type *big.Int) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "buyTicket", buyer, ticketAmount, event_id, _ticket_type)
}

// BuyTicket is a paid mutator transaction binding the contract method 0xcd3a15f8.
//
// Solidity: function buyTicket(address buyer, uint256 ticketAmount, uint256 event_id, uint256 _ticket_type) returns()
func (_Ticket721 *Ticket721Session) BuyTicket(buyer common.Address, ticketAmount *big.Int, event_id *big.Int, _ticket_type *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.BuyTicket(&_Ticket721.TransactOpts, buyer, ticketAmount, event_id, _ticket_type)
}

// BuyTicket is a paid mutator transaction binding the contract method 0xcd3a15f8.
//
// Solidity: function buyTicket(address buyer, uint256 ticketAmount, uint256 event_id, uint256 _ticket_type) returns()
func (_Ticket721 *Ticket721TransactorSession) BuyTicket(buyer common.Address, ticketAmount *big.Int, event_id *big.Int, _ticket_type *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.BuyTicket(&_Ticket721.TransactOpts, buyer, ticketAmount, event_id, _ticket_type)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns(bool)
func (_Ticket721 *Ticket721Transactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns(bool)
func (_Ticket721 *Ticket721Session) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.Mint(&_Ticket721.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns(bool)
func (_Ticket721 *Ticket721TransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.Mint(&_Ticket721.TransactOpts, to, tokenId)
}

// PlugSale is a paid mutator transaction binding the contract method 0x645c0fb0.
//
// Solidity: function plugSale(uint256 event_id, address orginizer) returns(uint256)
func (_Ticket721 *Ticket721Transactor) PlugSale(opts *bind.TransactOpts, event_id *big.Int, orginizer common.Address) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "plugSale", event_id, orginizer)
}

// PlugSale is a paid mutator transaction binding the contract method 0x645c0fb0.
//
// Solidity: function plugSale(uint256 event_id, address orginizer) returns(uint256)
func (_Ticket721 *Ticket721Session) PlugSale(event_id *big.Int, orginizer common.Address) (*types.Transaction, error) {
	return _Ticket721.Contract.PlugSale(&_Ticket721.TransactOpts, event_id, orginizer)
}

// PlugSale is a paid mutator transaction binding the contract method 0x645c0fb0.
//
// Solidity: function plugSale(uint256 event_id, address orginizer) returns(uint256)
func (_Ticket721 *Ticket721TransactorSession) PlugSale(event_id *big.Int, orginizer common.Address) (*types.Transaction, error) {
	return _Ticket721.Contract.PlugSale(&_Ticket721.TransactOpts, event_id, orginizer)
}

// RedeemTicket is a paid mutator transaction binding the contract method 0xa27c282e.
//
// Solidity: function redeemTicket(address visitor, uint256 tokenId, uint256 event_id) returns()
func (_Ticket721 *Ticket721Transactor) RedeemTicket(opts *bind.TransactOpts, visitor common.Address, tokenId *big.Int, event_id *big.Int) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "redeemTicket", visitor, tokenId, event_id)
}

// RedeemTicket is a paid mutator transaction binding the contract method 0xa27c282e.
//
// Solidity: function redeemTicket(address visitor, uint256 tokenId, uint256 event_id) returns()
func (_Ticket721 *Ticket721Session) RedeemTicket(visitor common.Address, tokenId *big.Int, event_id *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.RedeemTicket(&_Ticket721.TransactOpts, visitor, tokenId, event_id)
}

// RedeemTicket is a paid mutator transaction binding the contract method 0xa27c282e.
//
// Solidity: function redeemTicket(address visitor, uint256 tokenId, uint256 event_id) returns()
func (_Ticket721 *Ticket721TransactorSession) RedeemTicket(visitor common.Address, tokenId *big.Int, event_id *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.RedeemTicket(&_Ticket721.TransactOpts, visitor, tokenId, event_id)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_Ticket721 *Ticket721Transactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_Ticket721 *Ticket721Session) RenounceMinter() (*types.Transaction, error) {
	return _Ticket721.Contract.RenounceMinter(&_Ticket721.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_Ticket721 *Ticket721TransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _Ticket721.Contract.RenounceMinter(&_Ticket721.TransactOpts)
}

// ReserveEventId is a paid mutator transaction binding the contract method 0xa201b238.
//
// Solidity: function reserveEventId(address orginizer, string jid) returns(uint256 event_id)
func (_Ticket721 *Ticket721Transactor) ReserveEventId(opts *bind.TransactOpts, orginizer common.Address, jid string) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "reserveEventId", orginizer, jid)
}

// ReserveEventId is a paid mutator transaction binding the contract method 0xa201b238.
//
// Solidity: function reserveEventId(address orginizer, string jid) returns(uint256 event_id)
func (_Ticket721 *Ticket721Session) ReserveEventId(orginizer common.Address, jid string) (*types.Transaction, error) {
	return _Ticket721.Contract.ReserveEventId(&_Ticket721.TransactOpts, orginizer, jid)
}

// ReserveEventId is a paid mutator transaction binding the contract method 0xa201b238.
//
// Solidity: function reserveEventId(address orginizer, string jid) returns(uint256 event_id)
func (_Ticket721 *Ticket721TransactorSession) ReserveEventId(orginizer common.Address, jid string) (*types.Transaction, error) {
	return _Ticket721.Contract.ReserveEventId(&_Ticket721.TransactOpts, orginizer, jid)
}

// SafeMint is a paid mutator transaction binding the contract method 0x8832e6e3.
//
// Solidity: function safeMint(address to, uint256 tokenId, bytes _data) returns(bool)
func (_Ticket721 *Ticket721Transactor) SafeMint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "safeMint", to, tokenId, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0x8832e6e3.
//
// Solidity: function safeMint(address to, uint256 tokenId, bytes _data) returns(bool)
func (_Ticket721 *Ticket721Session) SafeMint(to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ticket721.Contract.SafeMint(&_Ticket721.TransactOpts, to, tokenId, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0x8832e6e3.
//
// Solidity: function safeMint(address to, uint256 tokenId, bytes _data) returns(bool)
func (_Ticket721 *Ticket721TransactorSession) SafeMint(to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ticket721.Contract.SafeMint(&_Ticket721.TransactOpts, to, tokenId, _data)
}

// SafeMint0 is a paid mutator transaction binding the contract method 0xa1448194.
//
// Solidity: function safeMint(address to, uint256 tokenId) returns(bool)
func (_Ticket721 *Ticket721Transactor) SafeMint0(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "safeMint0", to, tokenId)
}

// SafeMint0 is a paid mutator transaction binding the contract method 0xa1448194.
//
// Solidity: function safeMint(address to, uint256 tokenId) returns(bool)
func (_Ticket721 *Ticket721Session) SafeMint0(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.SafeMint0(&_Ticket721.TransactOpts, to, tokenId)
}

// SafeMint0 is a paid mutator transaction binding the contract method 0xa1448194.
//
// Solidity: function safeMint(address to, uint256 tokenId) returns(bool)
func (_Ticket721 *Ticket721TransactorSession) SafeMint0(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.SafeMint0(&_Ticket721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.SafeTransferFrom(&_Ticket721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.SafeTransferFrom(&_Ticket721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Ticket721 *Ticket721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Ticket721 *Ticket721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ticket721.Contract.SafeTransferFrom0(&_Ticket721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Ticket721 *Ticket721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ticket721.Contract.SafeTransferFrom0(&_Ticket721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Ticket721 *Ticket721Transactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Ticket721 *Ticket721Session) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _Ticket721.Contract.SetApprovalForAll(&_Ticket721.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Ticket721 *Ticket721TransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _Ticket721.Contract.SetApprovalForAll(&_Ticket721.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.TransferFrom(&_Ticket721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ticket721 *Ticket721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ticket721.Contract.TransferFrom(&_Ticket721.TransactOpts, from, to, tokenId)
}

// Ticket721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ticket721 contract.
type Ticket721ApprovalIterator struct {
	Event *Ticket721Approval // Event containing the contract specifics and raw log

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
func (it *Ticket721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721Approval)
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
		it.Event = new(Ticket721Approval)
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
func (it *Ticket721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721Approval represents a Approval event raised by the Ticket721 contract.
type Ticket721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ticket721 *Ticket721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Ticket721ApprovalIterator, error) {

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

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Ticket721ApprovalIterator{contract: _Ticket721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ticket721 *Ticket721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Ticket721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721Approval)
				if err := _Ticket721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Ticket721 *Ticket721Filterer) ParseApproval(log types.Log) (*Ticket721Approval, error) {
	event := new(Ticket721Approval)
	if err := _Ticket721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Ticket721 contract.
type Ticket721ApprovalForAllIterator struct {
	Event *Ticket721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Ticket721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721ApprovalForAll)
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
		it.Event = new(Ticket721ApprovalForAll)
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
func (it *Ticket721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721ApprovalForAll represents a ApprovalForAll event raised by the Ticket721 contract.
type Ticket721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ticket721 *Ticket721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Ticket721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Ticket721ApprovalForAllIterator{contract: _Ticket721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ticket721 *Ticket721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Ticket721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721ApprovalForAll)
				if err := _Ticket721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Ticket721 *Ticket721Filterer) ParseApprovalForAll(log types.Log) (*Ticket721ApprovalForAll, error) {
	event := new(Ticket721ApprovalForAll)
	if err := _Ticket721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721EventIdReservedIterator is returned from FilterEventIdReserved and is used to iterate over the raw logs and unpacked data for EventIdReserved events raised by the Ticket721 contract.
type Ticket721EventIdReservedIterator struct {
	Event *Ticket721EventIdReserved // Event containing the contract specifics and raw log

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
func (it *Ticket721EventIdReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721EventIdReserved)
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
		it.Event = new(Ticket721EventIdReserved)
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
func (it *Ticket721EventIdReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721EventIdReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721EventIdReserved represents a EventIdReserved event raised by the Ticket721 contract.
type Ticket721EventIdReserved struct {
	TicketSale common.Address
	EventId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventIdReserved is a free log retrieval operation binding the contract event 0xe7d1a453d66d83d6e4ffae7901ed08b56c3b1f53a1c43135c2de683e4529b65b.
//
// Solidity: event EventIdReserved(address indexed ticket_sale, uint256 indexed event_id)
func (_Ticket721 *Ticket721Filterer) FilterEventIdReserved(opts *bind.FilterOpts, ticket_sale []common.Address, event_id []*big.Int) (*Ticket721EventIdReservedIterator, error) {

	var ticket_saleRule []interface{}
	for _, ticket_saleItem := range ticket_sale {
		ticket_saleRule = append(ticket_saleRule, ticket_saleItem)
	}
	var event_idRule []interface{}
	for _, event_idItem := range event_id {
		event_idRule = append(event_idRule, event_idItem)
	}

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "EventIdReserved", ticket_saleRule, event_idRule)
	if err != nil {
		return nil, err
	}
	return &Ticket721EventIdReservedIterator{contract: _Ticket721.contract, event: "EventIdReserved", logs: logs, sub: sub}, nil
}

// WatchEventIdReserved is a free log subscription operation binding the contract event 0xe7d1a453d66d83d6e4ffae7901ed08b56c3b1f53a1c43135c2de683e4529b65b.
//
// Solidity: event EventIdReserved(address indexed ticket_sale, uint256 indexed event_id)
func (_Ticket721 *Ticket721Filterer) WatchEventIdReserved(opts *bind.WatchOpts, sink chan<- *Ticket721EventIdReserved, ticket_sale []common.Address, event_id []*big.Int) (event.Subscription, error) {

	var ticket_saleRule []interface{}
	for _, ticket_saleItem := range ticket_sale {
		ticket_saleRule = append(ticket_saleRule, ticket_saleItem)
	}
	var event_idRule []interface{}
	for _, event_idItem := range event_id {
		event_idRule = append(event_idRule, event_idItem)
	}

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "EventIdReserved", ticket_saleRule, event_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721EventIdReserved)
				if err := _Ticket721.contract.UnpackLog(event, "EventIdReserved", log); err != nil {
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

// ParseEventIdReserved is a log parse operation binding the contract event 0xe7d1a453d66d83d6e4ffae7901ed08b56c3b1f53a1c43135c2de683e4529b65b.
//
// Solidity: event EventIdReserved(address indexed ticket_sale, uint256 indexed event_id)
func (_Ticket721 *Ticket721Filterer) ParseEventIdReserved(log types.Log) (*Ticket721EventIdReserved, error) {
	event := new(Ticket721EventIdReserved)
	if err := _Ticket721.contract.UnpackLog(event, "EventIdReserved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721EventIdReservedHumanIterator is returned from FilterEventIdReservedHuman and is used to iterate over the raw logs and unpacked data for EventIdReservedHuman events raised by the Ticket721 contract.
type Ticket721EventIdReservedHumanIterator struct {
	Event *Ticket721EventIdReservedHuman // Event containing the contract specifics and raw log

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
func (it *Ticket721EventIdReservedHumanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721EventIdReservedHuman)
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
		it.Event = new(Ticket721EventIdReservedHuman)
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
func (it *Ticket721EventIdReservedHumanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721EventIdReservedHumanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721EventIdReservedHuman represents a EventIdReservedHuman event raised by the Ticket721 contract.
type Ticket721EventIdReservedHuman struct {
	TicketSale common.Address
	EventId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventIdReservedHuman is a free log retrieval operation binding the contract event 0x0a554ff72c713dffd0372d9c23cc1e214c1dfefbadda4710f4eee43eb6229a7b.
//
// Solidity: event EventIdReservedHuman(address ticket_sale, uint256 event_id)
func (_Ticket721 *Ticket721Filterer) FilterEventIdReservedHuman(opts *bind.FilterOpts) (*Ticket721EventIdReservedHumanIterator, error) {

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "EventIdReservedHuman")
	if err != nil {
		return nil, err
	}
	return &Ticket721EventIdReservedHumanIterator{contract: _Ticket721.contract, event: "EventIdReservedHuman", logs: logs, sub: sub}, nil
}

// WatchEventIdReservedHuman is a free log subscription operation binding the contract event 0x0a554ff72c713dffd0372d9c23cc1e214c1dfefbadda4710f4eee43eb6229a7b.
//
// Solidity: event EventIdReservedHuman(address ticket_sale, uint256 event_id)
func (_Ticket721 *Ticket721Filterer) WatchEventIdReservedHuman(opts *bind.WatchOpts, sink chan<- *Ticket721EventIdReservedHuman) (event.Subscription, error) {

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "EventIdReservedHuman")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721EventIdReservedHuman)
				if err := _Ticket721.contract.UnpackLog(event, "EventIdReservedHuman", log); err != nil {
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

// ParseEventIdReservedHuman is a log parse operation binding the contract event 0x0a554ff72c713dffd0372d9c23cc1e214c1dfefbadda4710f4eee43eb6229a7b.
//
// Solidity: event EventIdReservedHuman(address ticket_sale, uint256 event_id)
func (_Ticket721 *Ticket721Filterer) ParseEventIdReservedHuman(log types.Log) (*Ticket721EventIdReservedHuman, error) {
	event := new(Ticket721EventIdReservedHuman)
	if err := _Ticket721.contract.UnpackLog(event, "EventIdReservedHuman", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721MinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the Ticket721 contract.
type Ticket721MinterAddedIterator struct {
	Event *Ticket721MinterAdded // Event containing the contract specifics and raw log

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
func (it *Ticket721MinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721MinterAdded)
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
		it.Event = new(Ticket721MinterAdded)
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
func (it *Ticket721MinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721MinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721MinterAdded represents a MinterAdded event raised by the Ticket721 contract.
type Ticket721MinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_Ticket721 *Ticket721Filterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*Ticket721MinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &Ticket721MinterAddedIterator{contract: _Ticket721.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_Ticket721 *Ticket721Filterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *Ticket721MinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721MinterAdded)
				if err := _Ticket721.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_Ticket721 *Ticket721Filterer) ParseMinterAdded(log types.Log) (*Ticket721MinterAdded, error) {
	event := new(Ticket721MinterAdded)
	if err := _Ticket721.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721MinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the Ticket721 contract.
type Ticket721MinterRemovedIterator struct {
	Event *Ticket721MinterRemoved // Event containing the contract specifics and raw log

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
func (it *Ticket721MinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721MinterRemoved)
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
		it.Event = new(Ticket721MinterRemoved)
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
func (it *Ticket721MinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721MinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721MinterRemoved represents a MinterRemoved event raised by the Ticket721 contract.
type Ticket721MinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_Ticket721 *Ticket721Filterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*Ticket721MinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &Ticket721MinterRemovedIterator{contract: _Ticket721.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_Ticket721 *Ticket721Filterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *Ticket721MinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721MinterRemoved)
				if err := _Ticket721.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_Ticket721 *Ticket721Filterer) ParseMinterRemoved(log types.Log) (*Ticket721MinterRemoved, error) {
	event := new(Ticket721MinterRemoved)
	if err := _Ticket721.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721TicketBoughtIterator is returned from FilterTicketBought and is used to iterate over the raw logs and unpacked data for TicketBought events raised by the Ticket721 contract.
type Ticket721TicketBoughtIterator struct {
	Event *Ticket721TicketBought // Event containing the contract specifics and raw log

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
func (it *Ticket721TicketBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721TicketBought)
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
		it.Event = new(Ticket721TicketBought)
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
func (it *Ticket721TicketBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721TicketBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721TicketBought represents a TicketBought event raised by the Ticket721 contract.
type Ticket721TicketBought struct {
	VisitorWallet common.Address
	EventId       *big.Int
	TicketId      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTicketBought is a free log retrieval operation binding the contract event 0x97fac0ce6558f4accbb4696915809c0ef7023ffe7ba9454e94d49fb652aae1d5.
//
// Solidity: event TicketBought(address indexed visitor_wallet, uint256 indexed event_id, uint256 indexed ticket_id)
func (_Ticket721 *Ticket721Filterer) FilterTicketBought(opts *bind.FilterOpts, visitor_wallet []common.Address, event_id []*big.Int, ticket_id []*big.Int) (*Ticket721TicketBoughtIterator, error) {

	var visitor_walletRule []interface{}
	for _, visitor_walletItem := range visitor_wallet {
		visitor_walletRule = append(visitor_walletRule, visitor_walletItem)
	}
	var event_idRule []interface{}
	for _, event_idItem := range event_id {
		event_idRule = append(event_idRule, event_idItem)
	}
	var ticket_idRule []interface{}
	for _, ticket_idItem := range ticket_id {
		ticket_idRule = append(ticket_idRule, ticket_idItem)
	}

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "TicketBought", visitor_walletRule, event_idRule, ticket_idRule)
	if err != nil {
		return nil, err
	}
	return &Ticket721TicketBoughtIterator{contract: _Ticket721.contract, event: "TicketBought", logs: logs, sub: sub}, nil
}

// WatchTicketBought is a free log subscription operation binding the contract event 0x97fac0ce6558f4accbb4696915809c0ef7023ffe7ba9454e94d49fb652aae1d5.
//
// Solidity: event TicketBought(address indexed visitor_wallet, uint256 indexed event_id, uint256 indexed ticket_id)
func (_Ticket721 *Ticket721Filterer) WatchTicketBought(opts *bind.WatchOpts, sink chan<- *Ticket721TicketBought, visitor_wallet []common.Address, event_id []*big.Int, ticket_id []*big.Int) (event.Subscription, error) {

	var visitor_walletRule []interface{}
	for _, visitor_walletItem := range visitor_wallet {
		visitor_walletRule = append(visitor_walletRule, visitor_walletItem)
	}
	var event_idRule []interface{}
	for _, event_idItem := range event_id {
		event_idRule = append(event_idRule, event_idItem)
	}
	var ticket_idRule []interface{}
	for _, ticket_idItem := range ticket_id {
		ticket_idRule = append(ticket_idRule, ticket_idItem)
	}

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "TicketBought", visitor_walletRule, event_idRule, ticket_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721TicketBought)
				if err := _Ticket721.contract.UnpackLog(event, "TicketBought", log); err != nil {
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

// ParseTicketBought is a log parse operation binding the contract event 0x97fac0ce6558f4accbb4696915809c0ef7023ffe7ba9454e94d49fb652aae1d5.
//
// Solidity: event TicketBought(address indexed visitor_wallet, uint256 indexed event_id, uint256 indexed ticket_id)
func (_Ticket721 *Ticket721Filterer) ParseTicketBought(log types.Log) (*Ticket721TicketBought, error) {
	event := new(Ticket721TicketBought)
	if err := _Ticket721.contract.UnpackLog(event, "TicketBought", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721TicketBoughtHumanIterator is returned from FilterTicketBoughtHuman and is used to iterate over the raw logs and unpacked data for TicketBoughtHuman events raised by the Ticket721 contract.
type Ticket721TicketBoughtHumanIterator struct {
	Event *Ticket721TicketBoughtHuman // Event containing the contract specifics and raw log

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
func (it *Ticket721TicketBoughtHumanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721TicketBoughtHuman)
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
		it.Event = new(Ticket721TicketBoughtHuman)
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
func (it *Ticket721TicketBoughtHumanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721TicketBoughtHumanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721TicketBoughtHuman represents a TicketBoughtHuman event raised by the Ticket721 contract.
type Ticket721TicketBoughtHuman struct {
	VisitorWallet common.Address
	EventId       *big.Int
	TicketId      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTicketBoughtHuman is a free log retrieval operation binding the contract event 0x6d651fdc8652e9b93372b6cece4377b2c0f78be48b1ffbf3bdafc7baef0138ab.
//
// Solidity: event TicketBoughtHuman(address visitor_wallet, uint256 event_id, uint256 ticket_id)
func (_Ticket721 *Ticket721Filterer) FilterTicketBoughtHuman(opts *bind.FilterOpts) (*Ticket721TicketBoughtHumanIterator, error) {

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "TicketBoughtHuman")
	if err != nil {
		return nil, err
	}
	return &Ticket721TicketBoughtHumanIterator{contract: _Ticket721.contract, event: "TicketBoughtHuman", logs: logs, sub: sub}, nil
}

// WatchTicketBoughtHuman is a free log subscription operation binding the contract event 0x6d651fdc8652e9b93372b6cece4377b2c0f78be48b1ffbf3bdafc7baef0138ab.
//
// Solidity: event TicketBoughtHuman(address visitor_wallet, uint256 event_id, uint256 ticket_id)
func (_Ticket721 *Ticket721Filterer) WatchTicketBoughtHuman(opts *bind.WatchOpts, sink chan<- *Ticket721TicketBoughtHuman) (event.Subscription, error) {

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "TicketBoughtHuman")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721TicketBoughtHuman)
				if err := _Ticket721.contract.UnpackLog(event, "TicketBoughtHuman", log); err != nil {
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

// ParseTicketBoughtHuman is a log parse operation binding the contract event 0x6d651fdc8652e9b93372b6cece4377b2c0f78be48b1ffbf3bdafc7baef0138ab.
//
// Solidity: event TicketBoughtHuman(address visitor_wallet, uint256 event_id, uint256 ticket_id)
func (_Ticket721 *Ticket721Filterer) ParseTicketBoughtHuman(log types.Log) (*Ticket721TicketBoughtHuman, error) {
	event := new(Ticket721TicketBoughtHuman)
	if err := _Ticket721.contract.UnpackLog(event, "TicketBoughtHuman", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721TicketFulfilledIterator is returned from FilterTicketFulfilled and is used to iterate over the raw logs and unpacked data for TicketFulfilled events raised by the Ticket721 contract.
type Ticket721TicketFulfilledIterator struct {
	Event *Ticket721TicketFulfilled // Event containing the contract specifics and raw log

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
func (it *Ticket721TicketFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721TicketFulfilled)
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
		it.Event = new(Ticket721TicketFulfilled)
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
func (it *Ticket721TicketFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721TicketFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721TicketFulfilled represents a TicketFulfilled event raised by the Ticket721 contract.
type Ticket721TicketFulfilled struct {
	VisitorWallet common.Address
	EventId       *big.Int
	TicketId      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTicketFulfilled is a free log retrieval operation binding the contract event 0x1852380d0869dc478452f3f2bcfac707f842df524f5b8896d2a4f70c756f8184.
//
// Solidity: event TicketFulfilled(address indexed visitor_wallet, uint256 indexed event_id, uint256 indexed ticket_id)
func (_Ticket721 *Ticket721Filterer) FilterTicketFulfilled(opts *bind.FilterOpts, visitor_wallet []common.Address, event_id []*big.Int, ticket_id []*big.Int) (*Ticket721TicketFulfilledIterator, error) {

	var visitor_walletRule []interface{}
	for _, visitor_walletItem := range visitor_wallet {
		visitor_walletRule = append(visitor_walletRule, visitor_walletItem)
	}
	var event_idRule []interface{}
	for _, event_idItem := range event_id {
		event_idRule = append(event_idRule, event_idItem)
	}
	var ticket_idRule []interface{}
	for _, ticket_idItem := range ticket_id {
		ticket_idRule = append(ticket_idRule, ticket_idItem)
	}

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "TicketFulfilled", visitor_walletRule, event_idRule, ticket_idRule)
	if err != nil {
		return nil, err
	}
	return &Ticket721TicketFulfilledIterator{contract: _Ticket721.contract, event: "TicketFulfilled", logs: logs, sub: sub}, nil
}

// WatchTicketFulfilled is a free log subscription operation binding the contract event 0x1852380d0869dc478452f3f2bcfac707f842df524f5b8896d2a4f70c756f8184.
//
// Solidity: event TicketFulfilled(address indexed visitor_wallet, uint256 indexed event_id, uint256 indexed ticket_id)
func (_Ticket721 *Ticket721Filterer) WatchTicketFulfilled(opts *bind.WatchOpts, sink chan<- *Ticket721TicketFulfilled, visitor_wallet []common.Address, event_id []*big.Int, ticket_id []*big.Int) (event.Subscription, error) {

	var visitor_walletRule []interface{}
	for _, visitor_walletItem := range visitor_wallet {
		visitor_walletRule = append(visitor_walletRule, visitor_walletItem)
	}
	var event_idRule []interface{}
	for _, event_idItem := range event_id {
		event_idRule = append(event_idRule, event_idItem)
	}
	var ticket_idRule []interface{}
	for _, ticket_idItem := range ticket_id {
		ticket_idRule = append(ticket_idRule, ticket_idItem)
	}

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "TicketFulfilled", visitor_walletRule, event_idRule, ticket_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721TicketFulfilled)
				if err := _Ticket721.contract.UnpackLog(event, "TicketFulfilled", log); err != nil {
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

// ParseTicketFulfilled is a log parse operation binding the contract event 0x1852380d0869dc478452f3f2bcfac707f842df524f5b8896d2a4f70c756f8184.
//
// Solidity: event TicketFulfilled(address indexed visitor_wallet, uint256 indexed event_id, uint256 indexed ticket_id)
func (_Ticket721 *Ticket721Filterer) ParseTicketFulfilled(log types.Log) (*Ticket721TicketFulfilled, error) {
	event := new(Ticket721TicketFulfilled)
	if err := _Ticket721.contract.UnpackLog(event, "TicketFulfilled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721TicketFulfilledHumanIterator is returned from FilterTicketFulfilledHuman and is used to iterate over the raw logs and unpacked data for TicketFulfilledHuman events raised by the Ticket721 contract.
type Ticket721TicketFulfilledHumanIterator struct {
	Event *Ticket721TicketFulfilledHuman // Event containing the contract specifics and raw log

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
func (it *Ticket721TicketFulfilledHumanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721TicketFulfilledHuman)
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
		it.Event = new(Ticket721TicketFulfilledHuman)
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
func (it *Ticket721TicketFulfilledHumanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721TicketFulfilledHumanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721TicketFulfilledHuman represents a TicketFulfilledHuman event raised by the Ticket721 contract.
type Ticket721TicketFulfilledHuman struct {
	VisitorWallet common.Address
	EventId       *big.Int
	TicketId      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTicketFulfilledHuman is a free log retrieval operation binding the contract event 0xe82baa24ba3534133b9033b3b6013784eaa8e524dc5ef8cf2e3e0945d0074624.
//
// Solidity: event TicketFulfilledHuman(address visitor_wallet, uint256 event_id, uint256 ticket_id)
func (_Ticket721 *Ticket721Filterer) FilterTicketFulfilledHuman(opts *bind.FilterOpts) (*Ticket721TicketFulfilledHumanIterator, error) {

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "TicketFulfilledHuman")
	if err != nil {
		return nil, err
	}
	return &Ticket721TicketFulfilledHumanIterator{contract: _Ticket721.contract, event: "TicketFulfilledHuman", logs: logs, sub: sub}, nil
}

// WatchTicketFulfilledHuman is a free log subscription operation binding the contract event 0xe82baa24ba3534133b9033b3b6013784eaa8e524dc5ef8cf2e3e0945d0074624.
//
// Solidity: event TicketFulfilledHuman(address visitor_wallet, uint256 event_id, uint256 ticket_id)
func (_Ticket721 *Ticket721Filterer) WatchTicketFulfilledHuman(opts *bind.WatchOpts, sink chan<- *Ticket721TicketFulfilledHuman) (event.Subscription, error) {

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "TicketFulfilledHuman")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721TicketFulfilledHuman)
				if err := _Ticket721.contract.UnpackLog(event, "TicketFulfilledHuman", log); err != nil {
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

// ParseTicketFulfilledHuman is a log parse operation binding the contract event 0xe82baa24ba3534133b9033b3b6013784eaa8e524dc5ef8cf2e3e0945d0074624.
//
// Solidity: event TicketFulfilledHuman(address visitor_wallet, uint256 event_id, uint256 ticket_id)
func (_Ticket721 *Ticket721Filterer) ParseTicketFulfilledHuman(log types.Log) (*Ticket721TicketFulfilledHuman, error) {
	event := new(Ticket721TicketFulfilledHuman)
	if err := _Ticket721.contract.UnpackLog(event, "TicketFulfilledHuman", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Ticket721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ticket721 contract.
type Ticket721TransferIterator struct {
	Event *Ticket721Transfer // Event containing the contract specifics and raw log

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
func (it *Ticket721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ticket721Transfer)
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
		it.Event = new(Ticket721Transfer)
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
func (it *Ticket721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ticket721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ticket721Transfer represents a Transfer event raised by the Ticket721 contract.
type Ticket721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ticket721 *Ticket721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Ticket721TransferIterator, error) {

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

	logs, sub, err := _Ticket721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Ticket721TransferIterator{contract: _Ticket721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ticket721 *Ticket721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Ticket721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Ticket721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ticket721Transfer)
				if err := _Ticket721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Ticket721 *Ticket721Filterer) ParseTransfer(log types.Log) (*Ticket721Transfer, error) {
	event := new(Ticket721Transfer)
	if err := _Ticket721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
