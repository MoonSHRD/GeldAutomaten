// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KNS

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

// KNSABI is the input ABI used to generate the binding from.
const KNSABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tel\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"Jid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"new_wallet\",\"type\":\"address\"}],\"name\":\"LostedKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prime_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"Jid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tel\",\"type\":\"string\"}],\"name\":\"Registred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prime_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"JID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tel\",\"type\":\"string\"}],\"name\":\"RegistredHuman\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"Jid\",\"type\":\"string\"}],\"name\":\"GetOwnerByJid\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"Jid\",\"type\":\"string\"}],\"name\":\"GetWalletByJid\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tel\",\"type\":\"string\"}],\"name\":\"GetWalletByTel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"Jid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"new_wallet\",\"type\":\"address\"}],\"name\":\"LostKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"prime_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"Jid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tel\",\"type\":\"string\"}],\"name\":\"Register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"RegistryJ\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"prime_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tel\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"RegistryT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"prime_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tel\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KNS is an auto generated Go binding around an Ethereum contract.
type KNS struct {
	KNSCaller     // Read-only binding to the contract
	KNSTransactor // Write-only binding to the contract
	KNSFilterer   // Log filterer for contract events
}

// KNSCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNSSession struct {
	Contract     *KNS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNSCallerSession struct {
	Contract *KNSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KNSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNSTransactorSession struct {
	Contract     *KNSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNSRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNSRaw struct {
	Contract *KNS // Generic contract binding to access the raw methods on
}

// KNSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNSCallerRaw struct {
	Contract *KNSCaller // Generic read-only contract binding to access the raw methods on
}

// KNSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNSTransactorRaw struct {
	Contract *KNSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNS creates a new instance of KNS, bound to a specific deployed contract.
func NewKNS(address common.Address, backend bind.ContractBackend) (*KNS, error) {
	contract, err := bindKNS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNS{KNSCaller: KNSCaller{contract: contract}, KNSTransactor: KNSTransactor{contract: contract}, KNSFilterer: KNSFilterer{contract: contract}}, nil
}

// NewKNSCaller creates a new read-only instance of KNS, bound to a specific deployed contract.
func NewKNSCaller(address common.Address, caller bind.ContractCaller) (*KNSCaller, error) {
	contract, err := bindKNS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNSCaller{contract: contract}, nil
}

// NewKNSTransactor creates a new write-only instance of KNS, bound to a specific deployed contract.
func NewKNSTransactor(address common.Address, transactor bind.ContractTransactor) (*KNSTransactor, error) {
	contract, err := bindKNS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNSTransactor{contract: contract}, nil
}

// NewKNSFilterer creates a new log filterer instance of KNS, bound to a specific deployed contract.
func NewKNSFilterer(address common.Address, filterer bind.ContractFilterer) (*KNSFilterer, error) {
	contract, err := bindKNS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNSFilterer{contract: contract}, nil
}

// bindKNS binds a generic wrapper to an already deployed contract.
func bindKNS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNS *KNSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNS.Contract.KNSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNS *KNSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNS.Contract.KNSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNS *KNSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNS.Contract.KNSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNS *KNSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNS *KNSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNS *KNSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNS.Contract.contract.Transact(opts, method, params...)
}

// GetOwnerByJid is a free data retrieval call binding the contract method 0xee5bd245.
//
// Solidity: function GetOwnerByJid(string Jid) view returns(address owner)
func (_KNS *KNSCaller) GetOwnerByJid(opts *bind.CallOpts, Jid string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNS.contract.Call(opts, out, "GetOwnerByJid", Jid)
	return *ret0, err
}

// GetOwnerByJid is a free data retrieval call binding the contract method 0xee5bd245.
//
// Solidity: function GetOwnerByJid(string Jid) view returns(address owner)
func (_KNS *KNSSession) GetOwnerByJid(Jid string) (common.Address, error) {
	return _KNS.Contract.GetOwnerByJid(&_KNS.CallOpts, Jid)
}

// GetOwnerByJid is a free data retrieval call binding the contract method 0xee5bd245.
//
// Solidity: function GetOwnerByJid(string Jid) view returns(address owner)
func (_KNS *KNSCallerSession) GetOwnerByJid(Jid string) (common.Address, error) {
	return _KNS.Contract.GetOwnerByJid(&_KNS.CallOpts, Jid)
}

// GetWalletByJid is a free data retrieval call binding the contract method 0x2e37ccb1.
//
// Solidity: function GetWalletByJid(string Jid) view returns(address wallet)
func (_KNS *KNSCaller) GetWalletByJid(opts *bind.CallOpts, Jid string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNS.contract.Call(opts, out, "GetWalletByJid", Jid)
	return *ret0, err
}

// GetWalletByJid is a free data retrieval call binding the contract method 0x2e37ccb1.
//
// Solidity: function GetWalletByJid(string Jid) view returns(address wallet)
func (_KNS *KNSSession) GetWalletByJid(Jid string) (common.Address, error) {
	return _KNS.Contract.GetWalletByJid(&_KNS.CallOpts, Jid)
}

// GetWalletByJid is a free data retrieval call binding the contract method 0x2e37ccb1.
//
// Solidity: function GetWalletByJid(string Jid) view returns(address wallet)
func (_KNS *KNSCallerSession) GetWalletByJid(Jid string) (common.Address, error) {
	return _KNS.Contract.GetWalletByJid(&_KNS.CallOpts, Jid)
}

// GetWalletByTel is a free data retrieval call binding the contract method 0xc0ea7fa8.
//
// Solidity: function GetWalletByTel(string tel) view returns(address wallet)
func (_KNS *KNSCaller) GetWalletByTel(opts *bind.CallOpts, tel string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNS.contract.Call(opts, out, "GetWalletByTel", tel)
	return *ret0, err
}

// GetWalletByTel is a free data retrieval call binding the contract method 0xc0ea7fa8.
//
// Solidity: function GetWalletByTel(string tel) view returns(address wallet)
func (_KNS *KNSSession) GetWalletByTel(tel string) (common.Address, error) {
	return _KNS.Contract.GetWalletByTel(&_KNS.CallOpts, tel)
}

// GetWalletByTel is a free data retrieval call binding the contract method 0xc0ea7fa8.
//
// Solidity: function GetWalletByTel(string tel) view returns(address wallet)
func (_KNS *KNSCallerSession) GetWalletByTel(tel string) (common.Address, error) {
	return _KNS.Contract.GetWalletByTel(&_KNS.CallOpts, tel)
}

// RegistryJ is a free data retrieval call binding the contract method 0xddecfe27.
//
// Solidity: function RegistryJ(string ) view returns(address prime_owner, address wallet, string tel)
func (_KNS *KNSCaller) RegistryJ(opts *bind.CallOpts, arg0 string) (struct {
	PrimeOwner common.Address
	Wallet     common.Address
	Tel        string
}, error) {
	ret := new(struct {
		PrimeOwner common.Address
		Wallet     common.Address
		Tel        string
	})
	out := ret
	err := _KNS.contract.Call(opts, out, "RegistryJ", arg0)
	return *ret, err
}

// RegistryJ is a free data retrieval call binding the contract method 0xddecfe27.
//
// Solidity: function RegistryJ(string ) view returns(address prime_owner, address wallet, string tel)
func (_KNS *KNSSession) RegistryJ(arg0 string) (struct {
	PrimeOwner common.Address
	Wallet     common.Address
	Tel        string
}, error) {
	return _KNS.Contract.RegistryJ(&_KNS.CallOpts, arg0)
}

// RegistryJ is a free data retrieval call binding the contract method 0xddecfe27.
//
// Solidity: function RegistryJ(string ) view returns(address prime_owner, address wallet, string tel)
func (_KNS *KNSCallerSession) RegistryJ(arg0 string) (struct {
	PrimeOwner common.Address
	Wallet     common.Address
	Tel        string
}, error) {
	return _KNS.Contract.RegistryJ(&_KNS.CallOpts, arg0)
}

// RegistryT is a free data retrieval call binding the contract method 0xeab81808.
//
// Solidity: function RegistryT(string ) view returns(address prime_owner, address wallet, string tel)
func (_KNS *KNSCaller) RegistryT(opts *bind.CallOpts, arg0 string) (struct {
	PrimeOwner common.Address
	Wallet     common.Address
	Tel        string
}, error) {
	ret := new(struct {
		PrimeOwner common.Address
		Wallet     common.Address
		Tel        string
	})
	out := ret
	err := _KNS.contract.Call(opts, out, "RegistryT", arg0)
	return *ret, err
}

// RegistryT is a free data retrieval call binding the contract method 0xeab81808.
//
// Solidity: function RegistryT(string ) view returns(address prime_owner, address wallet, string tel)
func (_KNS *KNSSession) RegistryT(arg0 string) (struct {
	PrimeOwner common.Address
	Wallet     common.Address
	Tel        string
}, error) {
	return _KNS.Contract.RegistryT(&_KNS.CallOpts, arg0)
}

// RegistryT is a free data retrieval call binding the contract method 0xeab81808.
//
// Solidity: function RegistryT(string ) view returns(address prime_owner, address wallet, string tel)
func (_KNS *KNSCallerSession) RegistryT(arg0 string) (struct {
	PrimeOwner common.Address
	Wallet     common.Address
	Tel        string
}, error) {
	return _KNS.Contract.RegistryT(&_KNS.CallOpts, arg0)
}

// LostKey is a paid mutator transaction binding the contract method 0xa94f7966.
//
// Solidity: function LostKey(string Jid, address new_wallet) returns()
func (_KNS *KNSTransactor) LostKey(opts *bind.TransactOpts, Jid string, new_wallet common.Address) (*types.Transaction, error) {
	return _KNS.contract.Transact(opts, "LostKey", Jid, new_wallet)
}

// LostKey is a paid mutator transaction binding the contract method 0xa94f7966.
//
// Solidity: function LostKey(string Jid, address new_wallet) returns()
func (_KNS *KNSSession) LostKey(Jid string, new_wallet common.Address) (*types.Transaction, error) {
	return _KNS.Contract.LostKey(&_KNS.TransactOpts, Jid, new_wallet)
}

// LostKey is a paid mutator transaction binding the contract method 0xa94f7966.
//
// Solidity: function LostKey(string Jid, address new_wallet) returns()
func (_KNS *KNSTransactorSession) LostKey(Jid string, new_wallet common.Address) (*types.Transaction, error) {
	return _KNS.Contract.LostKey(&_KNS.TransactOpts, Jid, new_wallet)
}

// Register is a paid mutator transaction binding the contract method 0x6702b431.
//
// Solidity: function Register(address prime_owner, address wallet, string Jid, string tel) returns()
func (_KNS *KNSTransactor) Register(opts *bind.TransactOpts, prime_owner common.Address, wallet common.Address, Jid string, tel string) (*types.Transaction, error) {
	return _KNS.contract.Transact(opts, "Register", prime_owner, wallet, Jid, tel)
}

// Register is a paid mutator transaction binding the contract method 0x6702b431.
//
// Solidity: function Register(address prime_owner, address wallet, string Jid, string tel) returns()
func (_KNS *KNSSession) Register(prime_owner common.Address, wallet common.Address, Jid string, tel string) (*types.Transaction, error) {
	return _KNS.Contract.Register(&_KNS.TransactOpts, prime_owner, wallet, Jid, tel)
}

// Register is a paid mutator transaction binding the contract method 0x6702b431.
//
// Solidity: function Register(address prime_owner, address wallet, string Jid, string tel) returns()
func (_KNS *KNSTransactorSession) Register(prime_owner common.Address, wallet common.Address, Jid string, tel string) (*types.Transaction, error) {
	return _KNS.Contract.Register(&_KNS.TransactOpts, prime_owner, wallet, Jid, tel)
}

// KNSLostedKeyIterator is returned from FilterLostedKey and is used to iterate over the raw logs and unpacked data for LostedKey events raised by the KNS contract.
type KNSLostedKeyIterator struct {
	Event *KNSLostedKey // Event containing the contract specifics and raw log

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
func (it *KNSLostedKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KNSLostedKey)
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
		it.Event = new(KNSLostedKey)
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
func (it *KNSLostedKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KNSLostedKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KNSLostedKey represents a LostedKey event raised by the KNS contract.
type KNSLostedKey struct {
	Tel       common.Hash
	Jid       common.Hash
	NewWallet common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLostedKey is a free log retrieval operation binding the contract event 0x515f8f170e144c8fa03149925175e551812aeabc54b31fe6a45e0577e509d97e.
//
// Solidity: event LostedKey(string indexed tel, string indexed Jid, address indexed new_wallet)
func (_KNS *KNSFilterer) FilterLostedKey(opts *bind.FilterOpts, tel []string, Jid []string, new_wallet []common.Address) (*KNSLostedKeyIterator, error) {

	var telRule []interface{}
	for _, telItem := range tel {
		telRule = append(telRule, telItem)
	}
	var JidRule []interface{}
	for _, JidItem := range Jid {
		JidRule = append(JidRule, JidItem)
	}
	var new_walletRule []interface{}
	for _, new_walletItem := range new_wallet {
		new_walletRule = append(new_walletRule, new_walletItem)
	}

	logs, sub, err := _KNS.contract.FilterLogs(opts, "LostedKey", telRule, JidRule, new_walletRule)
	if err != nil {
		return nil, err
	}
	return &KNSLostedKeyIterator{contract: _KNS.contract, event: "LostedKey", logs: logs, sub: sub}, nil
}

// WatchLostedKey is a free log subscription operation binding the contract event 0x515f8f170e144c8fa03149925175e551812aeabc54b31fe6a45e0577e509d97e.
//
// Solidity: event LostedKey(string indexed tel, string indexed Jid, address indexed new_wallet)
func (_KNS *KNSFilterer) WatchLostedKey(opts *bind.WatchOpts, sink chan<- *KNSLostedKey, tel []string, Jid []string, new_wallet []common.Address) (event.Subscription, error) {

	var telRule []interface{}
	for _, telItem := range tel {
		telRule = append(telRule, telItem)
	}
	var JidRule []interface{}
	for _, JidItem := range Jid {
		JidRule = append(JidRule, JidItem)
	}
	var new_walletRule []interface{}
	for _, new_walletItem := range new_wallet {
		new_walletRule = append(new_walletRule, new_walletItem)
	}

	logs, sub, err := _KNS.contract.WatchLogs(opts, "LostedKey", telRule, JidRule, new_walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KNSLostedKey)
				if err := _KNS.contract.UnpackLog(event, "LostedKey", log); err != nil {
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

// ParseLostedKey is a log parse operation binding the contract event 0x515f8f170e144c8fa03149925175e551812aeabc54b31fe6a45e0577e509d97e.
//
// Solidity: event LostedKey(string indexed tel, string indexed Jid, address indexed new_wallet)
func (_KNS *KNSFilterer) ParseLostedKey(log types.Log) (*KNSLostedKey, error) {
	event := new(KNSLostedKey)
	if err := _KNS.contract.UnpackLog(event, "LostedKey", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KNSRegistredIterator is returned from FilterRegistred and is used to iterate over the raw logs and unpacked data for Registred events raised by the KNS contract.
type KNSRegistredIterator struct {
	Event *KNSRegistred // Event containing the contract specifics and raw log

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
func (it *KNSRegistredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KNSRegistred)
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
		it.Event = new(KNSRegistred)
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
func (it *KNSRegistredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KNSRegistredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KNSRegistred represents a Registred event raised by the KNS contract.
type KNSRegistred struct {
	PrimeOwner common.Address
	Wallet     common.Address
	Jid        common.Hash
	Tel        common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistred is a free log retrieval operation binding the contract event 0x460bd0e304b19a98768709c98d7f5943be773145f506c0f490c2f47f703f4757.
//
// Solidity: event Registred(address prime_owner, address wallet, string indexed Jid, string indexed tel)
func (_KNS *KNSFilterer) FilterRegistred(opts *bind.FilterOpts, Jid []string, tel []string) (*KNSRegistredIterator, error) {

	var JidRule []interface{}
	for _, JidItem := range Jid {
		JidRule = append(JidRule, JidItem)
	}
	var telRule []interface{}
	for _, telItem := range tel {
		telRule = append(telRule, telItem)
	}

	logs, sub, err := _KNS.contract.FilterLogs(opts, "Registred", JidRule, telRule)
	if err != nil {
		return nil, err
	}
	return &KNSRegistredIterator{contract: _KNS.contract, event: "Registred", logs: logs, sub: sub}, nil
}

// WatchRegistred is a free log subscription operation binding the contract event 0x460bd0e304b19a98768709c98d7f5943be773145f506c0f490c2f47f703f4757.
//
// Solidity: event Registred(address prime_owner, address wallet, string indexed Jid, string indexed tel)
func (_KNS *KNSFilterer) WatchRegistred(opts *bind.WatchOpts, sink chan<- *KNSRegistred, Jid []string, tel []string) (event.Subscription, error) {

	var JidRule []interface{}
	for _, JidItem := range Jid {
		JidRule = append(JidRule, JidItem)
	}
	var telRule []interface{}
	for _, telItem := range tel {
		telRule = append(telRule, telItem)
	}

	logs, sub, err := _KNS.contract.WatchLogs(opts, "Registred", JidRule, telRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KNSRegistred)
				if err := _KNS.contract.UnpackLog(event, "Registred", log); err != nil {
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

// ParseRegistred is a log parse operation binding the contract event 0x460bd0e304b19a98768709c98d7f5943be773145f506c0f490c2f47f703f4757.
//
// Solidity: event Registred(address prime_owner, address wallet, string indexed Jid, string indexed tel)
func (_KNS *KNSFilterer) ParseRegistred(log types.Log) (*KNSRegistred, error) {
	event := new(KNSRegistred)
	if err := _KNS.contract.UnpackLog(event, "Registred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KNSRegistredHumanIterator is returned from FilterRegistredHuman and is used to iterate over the raw logs and unpacked data for RegistredHuman events raised by the KNS contract.
type KNSRegistredHumanIterator struct {
	Event *KNSRegistredHuman // Event containing the contract specifics and raw log

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
func (it *KNSRegistredHumanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KNSRegistredHuman)
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
		it.Event = new(KNSRegistredHuman)
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
func (it *KNSRegistredHumanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KNSRegistredHumanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KNSRegistredHuman represents a RegistredHuman event raised by the KNS contract.
type KNSRegistredHuman struct {
	PrimeOwner common.Address
	Wallet     common.Address
	JID        string
	Tel        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistredHuman is a free log retrieval operation binding the contract event 0x25db3868baa42d60a73b6f076fceb10a25e6ad50e7e55ffed52513c1d96fa482.
//
// Solidity: event RegistredHuman(address prime_owner, address wallet, string JID, string tel)
func (_KNS *KNSFilterer) FilterRegistredHuman(opts *bind.FilterOpts) (*KNSRegistredHumanIterator, error) {

	logs, sub, err := _KNS.contract.FilterLogs(opts, "RegistredHuman")
	if err != nil {
		return nil, err
	}
	return &KNSRegistredHumanIterator{contract: _KNS.contract, event: "RegistredHuman", logs: logs, sub: sub}, nil
}

// WatchRegistredHuman is a free log subscription operation binding the contract event 0x25db3868baa42d60a73b6f076fceb10a25e6ad50e7e55ffed52513c1d96fa482.
//
// Solidity: event RegistredHuman(address prime_owner, address wallet, string JID, string tel)
func (_KNS *KNSFilterer) WatchRegistredHuman(opts *bind.WatchOpts, sink chan<- *KNSRegistredHuman) (event.Subscription, error) {

	logs, sub, err := _KNS.contract.WatchLogs(opts, "RegistredHuman")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KNSRegistredHuman)
				if err := _KNS.contract.UnpackLog(event, "RegistredHuman", log); err != nil {
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

// ParseRegistredHuman is a log parse operation binding the contract event 0x25db3868baa42d60a73b6f076fceb10a25e6ad50e7e55ffed52513c1d96fa482.
//
// Solidity: event RegistredHuman(address prime_owner, address wallet, string JID, string tel)
func (_KNS *KNSFilterer) ParseRegistredHuman(log types.Log) (*KNSRegistredHuman, error) {
	event := new(KNSRegistredHuman)
	if err := _KNS.contract.UnpackLog(event, "RegistredHuman", log); err != nil {
		return nil, err
	}
	return event, nil
}
