// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "energi.world/core/gen3"
	"energi.world/core/gen3/accounts/abi"
	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/event"
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

// FaucetABI is the input ABI used to generate the binding from.
const FaucetABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Donation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"donationnumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"GivenDonation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"donationnumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ReceiveDonation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestDonation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"requestGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setDonorName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setRecipientName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDonation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawMoney\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCalculatedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getDonationCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_nr\",\"type\":\"uint32\"}],\"name\":\"getDonationDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withdraw\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"}],\"name\":\"getDonorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getDonorBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getDonorByAddress\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"_totalBalance\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_numPayments\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"}],\"name\":\"getDonorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"_totalBalance\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_numPayments\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDonorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getDonorName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getPaymentCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_nr\",\"type\":\"uint32\"}],\"name\":\"getPaymentDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withdraw\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"}],\"name\":\"getRecipientAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getRecipientBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getRecipientByAddress\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"_totalBalance\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_numPayments\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"}],\"name\":\"getRecipientByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"_totalBalance\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_numPayments\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getRecipientName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecipientsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Faucet is an auto generated Go binding around an Ethereum contract.
type Faucet struct {
	FaucetCaller     // Read-only binding to the contract
	FaucetTransactor // Write-only binding to the contract
	FaucetFilterer   // Log filterer for contract events
}

// FaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaucetSession struct {
	Contract     *Faucet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaucetCallerSession struct {
	Contract *FaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaucetTransactorSession struct {
	Contract     *FaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaucetRaw struct {
	Contract *Faucet // Generic contract binding to access the raw methods on
}

// FaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaucetCallerRaw struct {
	Contract *FaucetCaller // Generic read-only contract binding to access the raw methods on
}

// FaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaucetTransactorRaw struct {
	Contract *FaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaucet creates a new instance of Faucet, bound to a specific deployed contract.
func NewFaucet(address common.Address, backend bind.ContractBackend) (*Faucet, error) {
	contract, err := bindFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Faucet{FaucetCaller: FaucetCaller{contract: contract}, FaucetTransactor: FaucetTransactor{contract: contract}, FaucetFilterer: FaucetFilterer{contract: contract}}, nil
}

// NewFaucetCaller creates a new read-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetCaller(address common.Address, caller bind.ContractCaller) (*FaucetCaller, error) {
	contract, err := bindFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetCaller{contract: contract}, nil
}

// NewFaucetTransactor creates a new write-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*FaucetTransactor, error) {
	contract, err := bindFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetTransactor{contract: contract}, nil
}

// NewFaucetFilterer creates a new log filterer instance of Faucet, bound to a specific deployed contract.
func NewFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*FaucetFilterer, error) {
	contract, err := bindFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaucetFilterer{contract: contract}, nil
}

// bindFaucet binds a generic wrapper to an already deployed contract.
func bindFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FaucetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.FaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transact(opts, method, params...)
}

// Donation is a paid mutator transaction binding the contract method 0x2a12a623.
//
// Solidity: function Donation() returns()
func (_Faucet *FaucetTransactor) Donation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "Donation")
}

// Donation is a paid mutator transaction binding the contract method 0x2a12a623.
//
// Solidity: function Donation() returns()
func (_Faucet *FaucetSession) Donation() (*types.Transaction, error) {
	return _Faucet.Contract.Donation(&_Faucet.TransactOpts)
}

// Donation is a paid mutator transaction binding the contract method 0x2a12a623.
//
// Solidity: function Donation() returns()
func (_Faucet *FaucetTransactorSession) Donation() (*types.Transaction, error) {
	return _Faucet.Contract.Donation(&_Faucet.TransactOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Faucet *FaucetTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Faucet *FaucetSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.ChangeOwner(&_Faucet.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Faucet *FaucetTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.ChangeOwner(&_Faucet.TransactOpts, newOwner)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Faucet *FaucetTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Faucet *FaucetSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.DestroySmartContract(&_Faucet.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Faucet *FaucetTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.DestroySmartContract(&_Faucet.TransactOpts, _to)
}

// GetBalance is a paid mutator transaction binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() returns(uint256)
func (_Faucet *FaucetTransactor) GetBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getBalance")
}

// GetBalance is a paid mutator transaction binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() returns(uint256)
func (_Faucet *FaucetSession) GetBalance() (*types.Transaction, error) {
	return _Faucet.Contract.GetBalance(&_Faucet.TransactOpts)
}

// GetBalance is a paid mutator transaction binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() returns(uint256)
func (_Faucet *FaucetTransactorSession) GetBalance() (*types.Transaction, error) {
	return _Faucet.Contract.GetBalance(&_Faucet.TransactOpts)
}

// GetCalculatedBalance is a paid mutator transaction binding the contract method 0xa65ac151.
//
// Solidity: function getCalculatedBalance() returns(uint256)
func (_Faucet *FaucetTransactor) GetCalculatedBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getCalculatedBalance")
}

// GetCalculatedBalance is a paid mutator transaction binding the contract method 0xa65ac151.
//
// Solidity: function getCalculatedBalance() returns(uint256)
func (_Faucet *FaucetSession) GetCalculatedBalance() (*types.Transaction, error) {
	return _Faucet.Contract.GetCalculatedBalance(&_Faucet.TransactOpts)
}

// GetCalculatedBalance is a paid mutator transaction binding the contract method 0xa65ac151.
//
// Solidity: function getCalculatedBalance() returns(uint256)
func (_Faucet *FaucetTransactorSession) GetCalculatedBalance() (*types.Transaction, error) {
	return _Faucet.Contract.GetCalculatedBalance(&_Faucet.TransactOpts)
}

// GetDonationCounts is a paid mutator transaction binding the contract method 0x47745d7c.
//
// Solidity: function getDonationCounts(address _from) returns(uint256)
func (_Faucet *FaucetTransactor) GetDonationCounts(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getDonationCounts", _from)
}

// GetDonationCounts is a paid mutator transaction binding the contract method 0x47745d7c.
//
// Solidity: function getDonationCounts(address _from) returns(uint256)
func (_Faucet *FaucetSession) GetDonationCounts(_from common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonationCounts(&_Faucet.TransactOpts, _from)
}

// GetDonationCounts is a paid mutator transaction binding the contract method 0x47745d7c.
//
// Solidity: function getDonationCounts(address _from) returns(uint256)
func (_Faucet *FaucetTransactorSession) GetDonationCounts(_from common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonationCounts(&_Faucet.TransactOpts, _from)
}

// GetDonationDetails is a paid mutator transaction binding the contract method 0x005ecae4.
//
// Solidity: function getDonationDetails(address _from, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Faucet *FaucetTransactor) GetDonationDetails(opts *bind.TransactOpts, _from common.Address, _nr uint32) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getDonationDetails", _from, _nr)
}

// GetDonationDetails is a paid mutator transaction binding the contract method 0x005ecae4.
//
// Solidity: function getDonationDetails(address _from, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Faucet *FaucetSession) GetDonationDetails(_from common.Address, _nr uint32) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonationDetails(&_Faucet.TransactOpts, _from, _nr)
}

// GetDonationDetails is a paid mutator transaction binding the contract method 0x005ecae4.
//
// Solidity: function getDonationDetails(address _from, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Faucet *FaucetTransactorSession) GetDonationDetails(_from common.Address, _nr uint32) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonationDetails(&_Faucet.TransactOpts, _from, _nr)
}

// GetDonorAddress is a paid mutator transaction binding the contract method 0xba872e62.
//
// Solidity: function getDonorAddress(uint256 _i) returns(address)
func (_Faucet *FaucetTransactor) GetDonorAddress(opts *bind.TransactOpts, _i *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getDonorAddress", _i)
}

// GetDonorAddress is a paid mutator transaction binding the contract method 0xba872e62.
//
// Solidity: function getDonorAddress(uint256 _i) returns(address)
func (_Faucet *FaucetSession) GetDonorAddress(_i *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorAddress(&_Faucet.TransactOpts, _i)
}

// GetDonorAddress is a paid mutator transaction binding the contract method 0xba872e62.
//
// Solidity: function getDonorAddress(uint256 _i) returns(address)
func (_Faucet *FaucetTransactorSession) GetDonorAddress(_i *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorAddress(&_Faucet.TransactOpts, _i)
}

// GetDonorBalance is a paid mutator transaction binding the contract method 0x4ffa7d50.
//
// Solidity: function getDonorBalance(address _from) returns(int256)
func (_Faucet *FaucetTransactor) GetDonorBalance(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getDonorBalance", _from)
}

// GetDonorBalance is a paid mutator transaction binding the contract method 0x4ffa7d50.
//
// Solidity: function getDonorBalance(address _from) returns(int256)
func (_Faucet *FaucetSession) GetDonorBalance(_from common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorBalance(&_Faucet.TransactOpts, _from)
}

// GetDonorBalance is a paid mutator transaction binding the contract method 0x4ffa7d50.
//
// Solidity: function getDonorBalance(address _from) returns(int256)
func (_Faucet *FaucetTransactorSession) GetDonorBalance(_from common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorBalance(&_Faucet.TransactOpts, _from)
}

// GetDonorByAddress is a paid mutator transaction binding the contract method 0x2cc5329a.
//
// Solidity: function getDonorByAddress(address _from) returns(int256 _totalBalance, uint256 _numPayments)
func (_Faucet *FaucetTransactor) GetDonorByAddress(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getDonorByAddress", _from)
}

// GetDonorByAddress is a paid mutator transaction binding the contract method 0x2cc5329a.
//
// Solidity: function getDonorByAddress(address _from) returns(int256 _totalBalance, uint256 _numPayments)
func (_Faucet *FaucetSession) GetDonorByAddress(_from common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorByAddress(&_Faucet.TransactOpts, _from)
}

// GetDonorByAddress is a paid mutator transaction binding the contract method 0x2cc5329a.
//
// Solidity: function getDonorByAddress(address _from) returns(int256 _totalBalance, uint256 _numPayments)
func (_Faucet *FaucetTransactorSession) GetDonorByAddress(_from common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorByAddress(&_Faucet.TransactOpts, _from)
}

// GetDonorByIndex is a paid mutator transaction binding the contract method 0x57075257.
//
// Solidity: function getDonorByIndex(uint256 _i) returns(address _from, int256 _totalBalance, uint256 _numPayments, string _name)
func (_Faucet *FaucetTransactor) GetDonorByIndex(opts *bind.TransactOpts, _i *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getDonorByIndex", _i)
}

// GetDonorByIndex is a paid mutator transaction binding the contract method 0x57075257.
//
// Solidity: function getDonorByIndex(uint256 _i) returns(address _from, int256 _totalBalance, uint256 _numPayments, string _name)
func (_Faucet *FaucetSession) GetDonorByIndex(_i *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorByIndex(&_Faucet.TransactOpts, _i)
}

// GetDonorByIndex is a paid mutator transaction binding the contract method 0x57075257.
//
// Solidity: function getDonorByIndex(uint256 _i) returns(address _from, int256 _totalBalance, uint256 _numPayments, string _name)
func (_Faucet *FaucetTransactorSession) GetDonorByIndex(_i *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorByIndex(&_Faucet.TransactOpts, _i)
}

// GetDonorCount is a paid mutator transaction binding the contract method 0x69bc2f1e.
//
// Solidity: function getDonorCount() returns(uint256)
func (_Faucet *FaucetTransactor) GetDonorCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getDonorCount")
}

// GetDonorCount is a paid mutator transaction binding the contract method 0x69bc2f1e.
//
// Solidity: function getDonorCount() returns(uint256)
func (_Faucet *FaucetSession) GetDonorCount() (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorCount(&_Faucet.TransactOpts)
}

// GetDonorCount is a paid mutator transaction binding the contract method 0x69bc2f1e.
//
// Solidity: function getDonorCount() returns(uint256)
func (_Faucet *FaucetTransactorSession) GetDonorCount() (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorCount(&_Faucet.TransactOpts)
}

// GetDonorName is a paid mutator transaction binding the contract method 0x0cf4bbfd.
//
// Solidity: function getDonorName(address _from) returns(string)
func (_Faucet *FaucetTransactor) GetDonorName(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getDonorName", _from)
}

// GetDonorName is a paid mutator transaction binding the contract method 0x0cf4bbfd.
//
// Solidity: function getDonorName(address _from) returns(string)
func (_Faucet *FaucetSession) GetDonorName(_from common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorName(&_Faucet.TransactOpts, _from)
}

// GetDonorName is a paid mutator transaction binding the contract method 0x0cf4bbfd.
//
// Solidity: function getDonorName(address _from) returns(string)
func (_Faucet *FaucetTransactorSession) GetDonorName(_from common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetDonorName(&_Faucet.TransactOpts, _from)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_Faucet *FaucetTransactor) GetOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getOwner")
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_Faucet *FaucetSession) GetOwner() (*types.Transaction, error) {
	return _Faucet.Contract.GetOwner(&_Faucet.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_Faucet *FaucetTransactorSession) GetOwner() (*types.Transaction, error) {
	return _Faucet.Contract.GetOwner(&_Faucet.TransactOpts)
}

// GetPaymentCounts is a paid mutator transaction binding the contract method 0x5f9ea90a.
//
// Solidity: function getPaymentCounts(address _to) returns(uint256)
func (_Faucet *FaucetTransactor) GetPaymentCounts(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getPaymentCounts", _to)
}

// GetPaymentCounts is a paid mutator transaction binding the contract method 0x5f9ea90a.
//
// Solidity: function getPaymentCounts(address _to) returns(uint256)
func (_Faucet *FaucetSession) GetPaymentCounts(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetPaymentCounts(&_Faucet.TransactOpts, _to)
}

// GetPaymentCounts is a paid mutator transaction binding the contract method 0x5f9ea90a.
//
// Solidity: function getPaymentCounts(address _to) returns(uint256)
func (_Faucet *FaucetTransactorSession) GetPaymentCounts(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetPaymentCounts(&_Faucet.TransactOpts, _to)
}

// GetPaymentDetails is a paid mutator transaction binding the contract method 0x655c9575.
//
// Solidity: function getPaymentDetails(address _to, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Faucet *FaucetTransactor) GetPaymentDetails(opts *bind.TransactOpts, _to common.Address, _nr uint32) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getPaymentDetails", _to, _nr)
}

// GetPaymentDetails is a paid mutator transaction binding the contract method 0x655c9575.
//
// Solidity: function getPaymentDetails(address _to, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Faucet *FaucetSession) GetPaymentDetails(_to common.Address, _nr uint32) (*types.Transaction, error) {
	return _Faucet.Contract.GetPaymentDetails(&_Faucet.TransactOpts, _to, _nr)
}

// GetPaymentDetails is a paid mutator transaction binding the contract method 0x655c9575.
//
// Solidity: function getPaymentDetails(address _to, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Faucet *FaucetTransactorSession) GetPaymentDetails(_to common.Address, _nr uint32) (*types.Transaction, error) {
	return _Faucet.Contract.GetPaymentDetails(&_Faucet.TransactOpts, _to, _nr)
}

// GetRecipientAddress is a paid mutator transaction binding the contract method 0x18337012.
//
// Solidity: function getRecipientAddress(uint256 _i) returns(address)
func (_Faucet *FaucetTransactor) GetRecipientAddress(opts *bind.TransactOpts, _i *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getRecipientAddress", _i)
}

// GetRecipientAddress is a paid mutator transaction binding the contract method 0x18337012.
//
// Solidity: function getRecipientAddress(uint256 _i) returns(address)
func (_Faucet *FaucetSession) GetRecipientAddress(_i *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientAddress(&_Faucet.TransactOpts, _i)
}

// GetRecipientAddress is a paid mutator transaction binding the contract method 0x18337012.
//
// Solidity: function getRecipientAddress(uint256 _i) returns(address)
func (_Faucet *FaucetTransactorSession) GetRecipientAddress(_i *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientAddress(&_Faucet.TransactOpts, _i)
}

// GetRecipientBalance is a paid mutator transaction binding the contract method 0x8dff5c5f.
//
// Solidity: function getRecipientBalance(address _to) returns(int256)
func (_Faucet *FaucetTransactor) GetRecipientBalance(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getRecipientBalance", _to)
}

// GetRecipientBalance is a paid mutator transaction binding the contract method 0x8dff5c5f.
//
// Solidity: function getRecipientBalance(address _to) returns(int256)
func (_Faucet *FaucetSession) GetRecipientBalance(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientBalance(&_Faucet.TransactOpts, _to)
}

// GetRecipientBalance is a paid mutator transaction binding the contract method 0x8dff5c5f.
//
// Solidity: function getRecipientBalance(address _to) returns(int256)
func (_Faucet *FaucetTransactorSession) GetRecipientBalance(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientBalance(&_Faucet.TransactOpts, _to)
}

// GetRecipientByAddress is a paid mutator transaction binding the contract method 0x2342c260.
//
// Solidity: function getRecipientByAddress(address _to) returns(int256 _totalBalance, uint256 _numPayments)
func (_Faucet *FaucetTransactor) GetRecipientByAddress(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getRecipientByAddress", _to)
}

// GetRecipientByAddress is a paid mutator transaction binding the contract method 0x2342c260.
//
// Solidity: function getRecipientByAddress(address _to) returns(int256 _totalBalance, uint256 _numPayments)
func (_Faucet *FaucetSession) GetRecipientByAddress(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientByAddress(&_Faucet.TransactOpts, _to)
}

// GetRecipientByAddress is a paid mutator transaction binding the contract method 0x2342c260.
//
// Solidity: function getRecipientByAddress(address _to) returns(int256 _totalBalance, uint256 _numPayments)
func (_Faucet *FaucetTransactorSession) GetRecipientByAddress(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientByAddress(&_Faucet.TransactOpts, _to)
}

// GetRecipientByIndex is a paid mutator transaction binding the contract method 0x5af84d5c.
//
// Solidity: function getRecipientByIndex(uint256 _i) returns(address _recipient, int256 _totalBalance, uint256 _numPayments, string _name)
func (_Faucet *FaucetTransactor) GetRecipientByIndex(opts *bind.TransactOpts, _i *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getRecipientByIndex", _i)
}

// GetRecipientByIndex is a paid mutator transaction binding the contract method 0x5af84d5c.
//
// Solidity: function getRecipientByIndex(uint256 _i) returns(address _recipient, int256 _totalBalance, uint256 _numPayments, string _name)
func (_Faucet *FaucetSession) GetRecipientByIndex(_i *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientByIndex(&_Faucet.TransactOpts, _i)
}

// GetRecipientByIndex is a paid mutator transaction binding the contract method 0x5af84d5c.
//
// Solidity: function getRecipientByIndex(uint256 _i) returns(address _recipient, int256 _totalBalance, uint256 _numPayments, string _name)
func (_Faucet *FaucetTransactorSession) GetRecipientByIndex(_i *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientByIndex(&_Faucet.TransactOpts, _i)
}

// GetRecipientName is a paid mutator transaction binding the contract method 0x6d347482.
//
// Solidity: function getRecipientName(address _to) returns(string)
func (_Faucet *FaucetTransactor) GetRecipientName(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getRecipientName", _to)
}

// GetRecipientName is a paid mutator transaction binding the contract method 0x6d347482.
//
// Solidity: function getRecipientName(address _to) returns(string)
func (_Faucet *FaucetSession) GetRecipientName(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientName(&_Faucet.TransactOpts, _to)
}

// GetRecipientName is a paid mutator transaction binding the contract method 0x6d347482.
//
// Solidity: function getRecipientName(address _to) returns(string)
func (_Faucet *FaucetTransactorSession) GetRecipientName(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientName(&_Faucet.TransactOpts, _to)
}

// GetRecipientsCount is a paid mutator transaction binding the contract method 0x221b17db.
//
// Solidity: function getRecipientsCount() returns(uint256)
func (_Faucet *FaucetTransactor) GetRecipientsCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "getRecipientsCount")
}

// GetRecipientsCount is a paid mutator transaction binding the contract method 0x221b17db.
//
// Solidity: function getRecipientsCount() returns(uint256)
func (_Faucet *FaucetSession) GetRecipientsCount() (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientsCount(&_Faucet.TransactOpts)
}

// GetRecipientsCount is a paid mutator transaction binding the contract method 0x221b17db.
//
// Solidity: function getRecipientsCount() returns(uint256)
func (_Faucet *FaucetTransactorSession) GetRecipientsCount() (*types.Transaction, error) {
	return _Faucet.Contract.GetRecipientsCount(&_Faucet.TransactOpts)
}

// RequestDonation is a paid mutator transaction binding the contract method 0xcdf1ec0f.
//
// Solidity: function requestDonation(address _to, uint256 _amount) returns()
func (_Faucet *FaucetTransactor) RequestDonation(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "requestDonation", _to, _amount)
}

// RequestDonation is a paid mutator transaction binding the contract method 0xcdf1ec0f.
//
// Solidity: function requestDonation(address _to, uint256 _amount) returns()
func (_Faucet *FaucetSession) RequestDonation(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.RequestDonation(&_Faucet.TransactOpts, _to, _amount)
}

// RequestDonation is a paid mutator transaction binding the contract method 0xcdf1ec0f.
//
// Solidity: function requestDonation(address _to, uint256 _amount) returns()
func (_Faucet *FaucetTransactorSession) RequestDonation(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.RequestDonation(&_Faucet.TransactOpts, _to, _amount)
}

// RequestGas is a paid mutator transaction binding the contract method 0xe5ee3197.
//
// Solidity: function requestGas(address _to) returns()
func (_Faucet *FaucetTransactor) RequestGas(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "requestGas", _to)
}

// RequestGas is a paid mutator transaction binding the contract method 0xe5ee3197.
//
// Solidity: function requestGas(address _to) returns()
func (_Faucet *FaucetSession) RequestGas(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.RequestGas(&_Faucet.TransactOpts, _to)
}

// RequestGas is a paid mutator transaction binding the contract method 0xe5ee3197.
//
// Solidity: function requestGas(address _to) returns()
func (_Faucet *FaucetTransactorSession) RequestGas(_to common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.RequestGas(&_Faucet.TransactOpts, _to)
}

// SetDonorName is a paid mutator transaction binding the contract method 0x34607994.
//
// Solidity: function setDonorName(string _name) returns()
func (_Faucet *FaucetTransactor) SetDonorName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setDonorName", _name)
}

// SetDonorName is a paid mutator transaction binding the contract method 0x34607994.
//
// Solidity: function setDonorName(string _name) returns()
func (_Faucet *FaucetSession) SetDonorName(_name string) (*types.Transaction, error) {
	return _Faucet.Contract.SetDonorName(&_Faucet.TransactOpts, _name)
}

// SetDonorName is a paid mutator transaction binding the contract method 0x34607994.
//
// Solidity: function setDonorName(string _name) returns()
func (_Faucet *FaucetTransactorSession) SetDonorName(_name string) (*types.Transaction, error) {
	return _Faucet.Contract.SetDonorName(&_Faucet.TransactOpts, _name)
}

// SetRecipientName is a paid mutator transaction binding the contract method 0x4655458d.
//
// Solidity: function setRecipientName(string _name) returns()
func (_Faucet *FaucetTransactor) SetRecipientName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setRecipientName", _name)
}

// SetRecipientName is a paid mutator transaction binding the contract method 0x4655458d.
//
// Solidity: function setRecipientName(string _name) returns()
func (_Faucet *FaucetSession) SetRecipientName(_name string) (*types.Transaction, error) {
	return _Faucet.Contract.SetRecipientName(&_Faucet.TransactOpts, _name)
}

// SetRecipientName is a paid mutator transaction binding the contract method 0x4655458d.
//
// Solidity: function setRecipientName(string _name) returns()
func (_Faucet *FaucetTransactorSession) SetRecipientName(_name string) (*types.Transaction, error) {
	return _Faucet.Contract.SetRecipientName(&_Faucet.TransactOpts, _name)
}

// Version is a paid mutator transaction binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(uint8)
func (_Faucet *FaucetTransactor) Version(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "version")
}

// Version is a paid mutator transaction binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(uint8)
func (_Faucet *FaucetSession) Version() (*types.Transaction, error) {
	return _Faucet.Contract.Version(&_Faucet.TransactOpts)
}

// Version is a paid mutator transaction binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(uint8)
func (_Faucet *FaucetTransactorSession) Version() (*types.Transaction, error) {
	return _Faucet.Contract.Version(&_Faucet.TransactOpts)
}

// WithdrawDonation is a paid mutator transaction binding the contract method 0xe9a04974.
//
// Solidity: function withdrawDonation(address _to, uint256 _amount) returns()
func (_Faucet *FaucetTransactor) WithdrawDonation(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "withdrawDonation", _to, _amount)
}

// WithdrawDonation is a paid mutator transaction binding the contract method 0xe9a04974.
//
// Solidity: function withdrawDonation(address _to, uint256 _amount) returns()
func (_Faucet *FaucetSession) WithdrawDonation(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.WithdrawDonation(&_Faucet.TransactOpts, _to, _amount)
}

// WithdrawDonation is a paid mutator transaction binding the contract method 0xe9a04974.
//
// Solidity: function withdrawDonation(address _to, uint256 _amount) returns()
func (_Faucet *FaucetTransactorSession) WithdrawDonation(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.WithdrawDonation(&_Faucet.TransactOpts, _to, _amount)
}

// FaucetGivenDonationIterator is returned from FilterGivenDonation and is used to iterate over the raw logs and unpacked data for GivenDonation events raised by the Faucet contract.
type FaucetGivenDonationIterator struct {
	Event *FaucetGivenDonation // Event containing the contract specifics and raw log

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
func (it *FaucetGivenDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetGivenDonation)
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
		it.Event = new(FaucetGivenDonation)
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
func (it *FaucetGivenDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetGivenDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetGivenDonation represents a GivenDonation event raised by the Faucet contract.
type FaucetGivenDonation struct {
	Recipient      common.Address
	Amount         *big.Int
	Donationnumber *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGivenDonation is a free log retrieval operation binding the contract event 0xcb4723b674e1fe7d58e6bdf336606df0930579a4013e8b7776bd7fa2423c3459.
//
// Solidity: event GivenDonation(address indexed recipient, uint256 amount, uint256 donationnumber, uint256 timestamp)
func (_Faucet *FaucetFilterer) FilterGivenDonation(opts *bind.FilterOpts, recipient []common.Address) (*FaucetGivenDonationIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "GivenDonation", recipientRule)
	if err != nil {
		return nil, err
	}
	return &FaucetGivenDonationIterator{contract: _Faucet.contract, event: "GivenDonation", logs: logs, sub: sub}, nil
}

// WatchGivenDonation is a free log subscription operation binding the contract event 0xcb4723b674e1fe7d58e6bdf336606df0930579a4013e8b7776bd7fa2423c3459.
//
// Solidity: event GivenDonation(address indexed recipient, uint256 amount, uint256 donationnumber, uint256 timestamp)
func (_Faucet *FaucetFilterer) WatchGivenDonation(opts *bind.WatchOpts, sink chan<- *FaucetGivenDonation, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "GivenDonation", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetGivenDonation)
				if err := _Faucet.contract.UnpackLog(event, "GivenDonation", log); err != nil {
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

// FaucetOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the Faucet contract.
type FaucetOwnerSetIterator struct {
	Event *FaucetOwnerSet // Event containing the contract specifics and raw log

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
func (it *FaucetOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetOwnerSet)
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
		it.Event = new(FaucetOwnerSet)
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
func (it *FaucetOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetOwnerSet represents a OwnerSet event raised by the Faucet contract.
type FaucetOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Faucet *FaucetFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*FaucetOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FaucetOwnerSetIterator{contract: _Faucet.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Faucet *FaucetFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *FaucetOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetOwnerSet)
				if err := _Faucet.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// FaucetReceiveDonationIterator is returned from FilterReceiveDonation and is used to iterate over the raw logs and unpacked data for ReceiveDonation events raised by the Faucet contract.
type FaucetReceiveDonationIterator struct {
	Event *FaucetReceiveDonation // Event containing the contract specifics and raw log

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
func (it *FaucetReceiveDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetReceiveDonation)
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
		it.Event = new(FaucetReceiveDonation)
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
func (it *FaucetReceiveDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetReceiveDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetReceiveDonation represents a ReceiveDonation event raised by the Faucet contract.
type FaucetReceiveDonation struct {
	Donor          common.Address
	Amount         *big.Int
	Donationnumber *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReceiveDonation is a free log retrieval operation binding the contract event 0x324679f5b90bc0d0379e554810049f1674e53461b9cfc5e0fb606a3e6d91970c.
//
// Solidity: event ReceiveDonation(address indexed donor, uint256 amount, uint256 donationnumber, uint256 timestamp)
func (_Faucet *FaucetFilterer) FilterReceiveDonation(opts *bind.FilterOpts, donor []common.Address) (*FaucetReceiveDonationIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "ReceiveDonation", donorRule)
	if err != nil {
		return nil, err
	}
	return &FaucetReceiveDonationIterator{contract: _Faucet.contract, event: "ReceiveDonation", logs: logs, sub: sub}, nil
}

// WatchReceiveDonation is a free log subscription operation binding the contract event 0x324679f5b90bc0d0379e554810049f1674e53461b9cfc5e0fb606a3e6d91970c.
//
// Solidity: event ReceiveDonation(address indexed donor, uint256 amount, uint256 donationnumber, uint256 timestamp)
func (_Faucet *FaucetFilterer) WatchReceiveDonation(opts *bind.WatchOpts, sink chan<- *FaucetReceiveDonation, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "ReceiveDonation", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetReceiveDonation)
				if err := _Faucet.contract.UnpackLog(event, "ReceiveDonation", log); err != nil {
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

// FaucetWithdrawMoneyIterator is returned from FilterWithdrawMoney and is used to iterate over the raw logs and unpacked data for WithdrawMoney events raised by the Faucet contract.
type FaucetWithdrawMoneyIterator struct {
	Event *FaucetWithdrawMoney // Event containing the contract specifics and raw log

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
func (it *FaucetWithdrawMoneyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetWithdrawMoney)
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
		it.Event = new(FaucetWithdrawMoney)
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
func (it *FaucetWithdrawMoneyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetWithdrawMoneyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetWithdrawMoney represents a WithdrawMoney event raised by the Faucet contract.
type FaucetWithdrawMoney struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawMoney is a free log retrieval operation binding the contract event 0x9fe2db218b58c00ab3556f2be44b2769b37978368c63313d82d78306568e4776.
//
// Solidity: event WithdrawMoney(address indexed to, uint256 indexed amount)
func (_Faucet *FaucetFilterer) FilterWithdrawMoney(opts *bind.FilterOpts, to []common.Address, amount []*big.Int) (*FaucetWithdrawMoneyIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "WithdrawMoney", toRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &FaucetWithdrawMoneyIterator{contract: _Faucet.contract, event: "WithdrawMoney", logs: logs, sub: sub}, nil
}

// WatchWithdrawMoney is a free log subscription operation binding the contract event 0x9fe2db218b58c00ab3556f2be44b2769b37978368c63313d82d78306568e4776.
//
// Solidity: event WithdrawMoney(address indexed to, uint256 indexed amount)
func (_Faucet *FaucetFilterer) WatchWithdrawMoney(opts *bind.WatchOpts, sink chan<- *FaucetWithdrawMoney, to []common.Address, amount []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "WithdrawMoney", toRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetWithdrawMoney)
				if err := _Faucet.contract.UnpackLog(event, "WithdrawMoney", log); err != nil {
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
