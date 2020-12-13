// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

// ./abigen -abi ../solidity/artifacts/faucet106.abi.json --pkg main > faucet.go

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

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"donationnumber\",\"type\":\"uint256\"}],\"name\":\"GivenDonation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"donationnumber\",\"type\":\"uint256\"}],\"name\":\"ReceivedDonation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawMoney\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes8\",\"name\":\"_country\",\"type\":\"bytes8\"}],\"name\":\"Donation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCalculatedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getDonationCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_nr\",\"type\":\"uint32\"}],\"name\":\"getDonationDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withdraw\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"}],\"name\":\"getDonorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getDonorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getDonorByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numPayments\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes8\",\"name\":\"_country\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"}],\"name\":\"getDonorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numPayments\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes8\",\"name\":\"_country\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDonorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getDonorCountry\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"getDonorName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getPaymentCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_nr\",\"type\":\"uint32\"}],\"name\":\"getPaymentDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withdraw\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"}],\"name\":\"getRecipientAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getRecipientBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getRecipientByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numPayments\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes8\",\"name\":\"_country\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"}],\"name\":\"getRecipientByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numPayments\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes8\",\"name\":\"_country\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getRecipientCountry\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"getRecipientName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecipientsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"giveBootstrapGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes8\",\"name\":\"_country\",\"type\":\"bytes8\"}],\"name\":\"requestDonation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"_country\",\"type\":\"bytes8\"}],\"name\":\"setDonorCountry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setDonorName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"_country\",\"type\":\"bytes8\"}],\"name\":\"setRecipientCountry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setRecipientName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDonation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Donation is a paid mutator transaction binding the contract method 0x6c6376e4.
//
// Solidity: function Donation(string _name, bytes8 _country) returns()
func (_Main *MainTransactor) Donation(opts *bind.TransactOpts, _name string, _country [8]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "Donation", _name, _country)
}

// Donation is a paid mutator transaction binding the contract method 0x6c6376e4.
//
// Solidity: function Donation(string _name, bytes8 _country) returns()
func (_Main *MainSession) Donation(_name string, _country [8]byte) (*types.Transaction, error) {
	return _Main.Contract.Donation(&_Main.TransactOpts, _name, _country)
}

// Donation is a paid mutator transaction binding the contract method 0x6c6376e4.
//
// Solidity: function Donation(string _name, bytes8 _country) returns()
func (_Main *MainTransactorSession) Donation(_name string, _country [8]byte) (*types.Transaction, error) {
	return _Main.Contract.Donation(&_Main.TransactOpts, _name, _country)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Main *MainTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Main *MainSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.ChangeOwner(&_Main.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Main *MainTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.ChangeOwner(&_Main.TransactOpts, newOwner)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x85474728.
//
// Solidity: function destroySmartContract() returns()
func (_Main *MainTransactor) DestroySmartContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "destroySmartContract")
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x85474728.
//
// Solidity: function destroySmartContract() returns()
func (_Main *MainSession) DestroySmartContract() (*types.Transaction, error) {
	return _Main.Contract.DestroySmartContract(&_Main.TransactOpts)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x85474728.
//
// Solidity: function destroySmartContract() returns()
func (_Main *MainTransactorSession) DestroySmartContract() (*types.Transaction, error) {
	return _Main.Contract.DestroySmartContract(&_Main.TransactOpts)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(address)
func (_Main *MainTransactor) GetAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getAddress")
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(address)
func (_Main *MainSession) GetAddress() (*types.Transaction, error) {
	return _Main.Contract.GetAddress(&_Main.TransactOpts)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(address)
func (_Main *MainTransactorSession) GetAddress() (*types.Transaction, error) {
	return _Main.Contract.GetAddress(&_Main.TransactOpts)
}

// GetBalance is a paid mutator transaction binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() returns(uint256)
func (_Main *MainTransactor) GetBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getBalance")
}

// GetBalance is a paid mutator transaction binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() returns(uint256)
func (_Main *MainSession) GetBalance() (*types.Transaction, error) {
	return _Main.Contract.GetBalance(&_Main.TransactOpts)
}

// GetBalance is a paid mutator transaction binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() returns(uint256)
func (_Main *MainTransactorSession) GetBalance() (*types.Transaction, error) {
	return _Main.Contract.GetBalance(&_Main.TransactOpts)
}

// GetCalculatedBalance is a paid mutator transaction binding the contract method 0xa65ac151.
//
// Solidity: function getCalculatedBalance() returns(uint256)
func (_Main *MainTransactor) GetCalculatedBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getCalculatedBalance")
}

// GetCalculatedBalance is a paid mutator transaction binding the contract method 0xa65ac151.
//
// Solidity: function getCalculatedBalance() returns(uint256)
func (_Main *MainSession) GetCalculatedBalance() (*types.Transaction, error) {
	return _Main.Contract.GetCalculatedBalance(&_Main.TransactOpts)
}

// GetCalculatedBalance is a paid mutator transaction binding the contract method 0xa65ac151.
//
// Solidity: function getCalculatedBalance() returns(uint256)
func (_Main *MainTransactorSession) GetCalculatedBalance() (*types.Transaction, error) {
	return _Main.Contract.GetCalculatedBalance(&_Main.TransactOpts)
}

// GetDonationCounts is a paid mutator transaction binding the contract method 0x47745d7c.
//
// Solidity: function getDonationCounts(address _from) returns(uint256)
func (_Main *MainTransactor) GetDonationCounts(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getDonationCounts", _from)
}

// GetDonationCounts is a paid mutator transaction binding the contract method 0x47745d7c.
//
// Solidity: function getDonationCounts(address _from) returns(uint256)
func (_Main *MainSession) GetDonationCounts(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonationCounts(&_Main.TransactOpts, _from)
}

// GetDonationCounts is a paid mutator transaction binding the contract method 0x47745d7c.
//
// Solidity: function getDonationCounts(address _from) returns(uint256)
func (_Main *MainTransactorSession) GetDonationCounts(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonationCounts(&_Main.TransactOpts, _from)
}

// GetDonationDetails is a paid mutator transaction binding the contract method 0x005ecae4.
//
// Solidity: function getDonationDetails(address _from, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Main *MainTransactor) GetDonationDetails(opts *bind.TransactOpts, _from common.Address, _nr uint32) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getDonationDetails", _from, _nr)
}

// GetDonationDetails is a paid mutator transaction binding the contract method 0x005ecae4.
//
// Solidity: function getDonationDetails(address _from, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Main *MainSession) GetDonationDetails(_from common.Address, _nr uint32) (*types.Transaction, error) {
	return _Main.Contract.GetDonationDetails(&_Main.TransactOpts, _from, _nr)
}

// GetDonationDetails is a paid mutator transaction binding the contract method 0x005ecae4.
//
// Solidity: function getDonationDetails(address _from, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Main *MainTransactorSession) GetDonationDetails(_from common.Address, _nr uint32) (*types.Transaction, error) {
	return _Main.Contract.GetDonationDetails(&_Main.TransactOpts, _from, _nr)
}

// GetDonorAddress is a paid mutator transaction binding the contract method 0xba872e62.
//
// Solidity: function getDonorAddress(uint256 _i) returns(address)
func (_Main *MainTransactor) GetDonorAddress(opts *bind.TransactOpts, _i *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getDonorAddress", _i)
}

// GetDonorAddress is a paid mutator transaction binding the contract method 0xba872e62.
//
// Solidity: function getDonorAddress(uint256 _i) returns(address)
func (_Main *MainSession) GetDonorAddress(_i *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GetDonorAddress(&_Main.TransactOpts, _i)
}

// GetDonorAddress is a paid mutator transaction binding the contract method 0xba872e62.
//
// Solidity: function getDonorAddress(uint256 _i) returns(address)
func (_Main *MainTransactorSession) GetDonorAddress(_i *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GetDonorAddress(&_Main.TransactOpts, _i)
}

// GetDonorBalance is a paid mutator transaction binding the contract method 0x4ffa7d50.
//
// Solidity: function getDonorBalance(address _from) returns(uint256)
func (_Main *MainTransactor) GetDonorBalance(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getDonorBalance", _from)
}

// GetDonorBalance is a paid mutator transaction binding the contract method 0x4ffa7d50.
//
// Solidity: function getDonorBalance(address _from) returns(uint256)
func (_Main *MainSession) GetDonorBalance(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonorBalance(&_Main.TransactOpts, _from)
}

// GetDonorBalance is a paid mutator transaction binding the contract method 0x4ffa7d50.
//
// Solidity: function getDonorBalance(address _from) returns(uint256)
func (_Main *MainTransactorSession) GetDonorBalance(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonorBalance(&_Main.TransactOpts, _from)
}

// GetDonorByAddress is a paid mutator transaction binding the contract method 0x2cc5329a.
//
// Solidity: function getDonorByAddress(address _from) returns(uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainTransactor) GetDonorByAddress(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getDonorByAddress", _from)
}

// GetDonorByAddress is a paid mutator transaction binding the contract method 0x2cc5329a.
//
// Solidity: function getDonorByAddress(address _from) returns(uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainSession) GetDonorByAddress(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonorByAddress(&_Main.TransactOpts, _from)
}

// GetDonorByAddress is a paid mutator transaction binding the contract method 0x2cc5329a.
//
// Solidity: function getDonorByAddress(address _from) returns(uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainTransactorSession) GetDonorByAddress(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonorByAddress(&_Main.TransactOpts, _from)
}

// GetDonorByIndex is a paid mutator transaction binding the contract method 0x57075257.
//
// Solidity: function getDonorByIndex(uint256 _i) returns(address _from, uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainTransactor) GetDonorByIndex(opts *bind.TransactOpts, _i *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getDonorByIndex", _i)
}

// GetDonorByIndex is a paid mutator transaction binding the contract method 0x57075257.
//
// Solidity: function getDonorByIndex(uint256 _i) returns(address _from, uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainSession) GetDonorByIndex(_i *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GetDonorByIndex(&_Main.TransactOpts, _i)
}

// GetDonorByIndex is a paid mutator transaction binding the contract method 0x57075257.
//
// Solidity: function getDonorByIndex(uint256 _i) returns(address _from, uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainTransactorSession) GetDonorByIndex(_i *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GetDonorByIndex(&_Main.TransactOpts, _i)
}

// GetDonorCount is a paid mutator transaction binding the contract method 0x69bc2f1e.
//
// Solidity: function getDonorCount() returns(uint256)
func (_Main *MainTransactor) GetDonorCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getDonorCount")
}

// GetDonorCount is a paid mutator transaction binding the contract method 0x69bc2f1e.
//
// Solidity: function getDonorCount() returns(uint256)
func (_Main *MainSession) GetDonorCount() (*types.Transaction, error) {
	return _Main.Contract.GetDonorCount(&_Main.TransactOpts)
}

// GetDonorCount is a paid mutator transaction binding the contract method 0x69bc2f1e.
//
// Solidity: function getDonorCount() returns(uint256)
func (_Main *MainTransactorSession) GetDonorCount() (*types.Transaction, error) {
	return _Main.Contract.GetDonorCount(&_Main.TransactOpts)
}

// GetDonorCountry is a paid mutator transaction binding the contract method 0xdf2b8aa8.
//
// Solidity: function getDonorCountry(address _from) returns(bytes8)
func (_Main *MainTransactor) GetDonorCountry(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getDonorCountry", _from)
}

// GetDonorCountry is a paid mutator transaction binding the contract method 0xdf2b8aa8.
//
// Solidity: function getDonorCountry(address _from) returns(bytes8)
func (_Main *MainSession) GetDonorCountry(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonorCountry(&_Main.TransactOpts, _from)
}

// GetDonorCountry is a paid mutator transaction binding the contract method 0xdf2b8aa8.
//
// Solidity: function getDonorCountry(address _from) returns(bytes8)
func (_Main *MainTransactorSession) GetDonorCountry(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonorCountry(&_Main.TransactOpts, _from)
}

// GetDonorName is a paid mutator transaction binding the contract method 0x0cf4bbfd.
//
// Solidity: function getDonorName(address _from) returns(string)
func (_Main *MainTransactor) GetDonorName(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getDonorName", _from)
}

// GetDonorName is a paid mutator transaction binding the contract method 0x0cf4bbfd.
//
// Solidity: function getDonorName(address _from) returns(string)
func (_Main *MainSession) GetDonorName(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonorName(&_Main.TransactOpts, _from)
}

// GetDonorName is a paid mutator transaction binding the contract method 0x0cf4bbfd.
//
// Solidity: function getDonorName(address _from) returns(string)
func (_Main *MainTransactorSession) GetDonorName(_from common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetDonorName(&_Main.TransactOpts, _from)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_Main *MainTransactor) GetOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getOwner")
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_Main *MainSession) GetOwner() (*types.Transaction, error) {
	return _Main.Contract.GetOwner(&_Main.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_Main *MainTransactorSession) GetOwner() (*types.Transaction, error) {
	return _Main.Contract.GetOwner(&_Main.TransactOpts)
}

// GetPaymentCounts is a paid mutator transaction binding the contract method 0x5f9ea90a.
//
// Solidity: function getPaymentCounts(address _to) returns(uint256)
func (_Main *MainTransactor) GetPaymentCounts(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getPaymentCounts", _to)
}

// GetPaymentCounts is a paid mutator transaction binding the contract method 0x5f9ea90a.
//
// Solidity: function getPaymentCounts(address _to) returns(uint256)
func (_Main *MainSession) GetPaymentCounts(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetPaymentCounts(&_Main.TransactOpts, _to)
}

// GetPaymentCounts is a paid mutator transaction binding the contract method 0x5f9ea90a.
//
// Solidity: function getPaymentCounts(address _to) returns(uint256)
func (_Main *MainTransactorSession) GetPaymentCounts(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetPaymentCounts(&_Main.TransactOpts, _to)
}

// GetPaymentDetails is a paid mutator transaction binding the contract method 0x655c9575.
//
// Solidity: function getPaymentDetails(address _to, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Main *MainTransactor) GetPaymentDetails(opts *bind.TransactOpts, _to common.Address, _nr uint32) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getPaymentDetails", _to, _nr)
}

// GetPaymentDetails is a paid mutator transaction binding the contract method 0x655c9575.
//
// Solidity: function getPaymentDetails(address _to, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Main *MainSession) GetPaymentDetails(_to common.Address, _nr uint32) (*types.Transaction, error) {
	return _Main.Contract.GetPaymentDetails(&_Main.TransactOpts, _to, _nr)
}

// GetPaymentDetails is a paid mutator transaction binding the contract method 0x655c9575.
//
// Solidity: function getPaymentDetails(address _to, uint32 _nr) returns(uint256 _amount, bool _withdraw, uint256 _timestamp)
func (_Main *MainTransactorSession) GetPaymentDetails(_to common.Address, _nr uint32) (*types.Transaction, error) {
	return _Main.Contract.GetPaymentDetails(&_Main.TransactOpts, _to, _nr)
}

// GetRecipientAddress is a paid mutator transaction binding the contract method 0x18337012.
//
// Solidity: function getRecipientAddress(uint256 _i) returns(address)
func (_Main *MainTransactor) GetRecipientAddress(opts *bind.TransactOpts, _i *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getRecipientAddress", _i)
}

// GetRecipientAddress is a paid mutator transaction binding the contract method 0x18337012.
//
// Solidity: function getRecipientAddress(uint256 _i) returns(address)
func (_Main *MainSession) GetRecipientAddress(_i *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientAddress(&_Main.TransactOpts, _i)
}

// GetRecipientAddress is a paid mutator transaction binding the contract method 0x18337012.
//
// Solidity: function getRecipientAddress(uint256 _i) returns(address)
func (_Main *MainTransactorSession) GetRecipientAddress(_i *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientAddress(&_Main.TransactOpts, _i)
}

// GetRecipientBalance is a paid mutator transaction binding the contract method 0x8dff5c5f.
//
// Solidity: function getRecipientBalance(address _to) returns(uint256)
func (_Main *MainTransactor) GetRecipientBalance(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getRecipientBalance", _to)
}

// GetRecipientBalance is a paid mutator transaction binding the contract method 0x8dff5c5f.
//
// Solidity: function getRecipientBalance(address _to) returns(uint256)
func (_Main *MainSession) GetRecipientBalance(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientBalance(&_Main.TransactOpts, _to)
}

// GetRecipientBalance is a paid mutator transaction binding the contract method 0x8dff5c5f.
//
// Solidity: function getRecipientBalance(address _to) returns(uint256)
func (_Main *MainTransactorSession) GetRecipientBalance(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientBalance(&_Main.TransactOpts, _to)
}

// GetRecipientByAddress is a paid mutator transaction binding the contract method 0x2342c260.
//
// Solidity: function getRecipientByAddress(address _to) returns(uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainTransactor) GetRecipientByAddress(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getRecipientByAddress", _to)
}

// GetRecipientByAddress is a paid mutator transaction binding the contract method 0x2342c260.
//
// Solidity: function getRecipientByAddress(address _to) returns(uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainSession) GetRecipientByAddress(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientByAddress(&_Main.TransactOpts, _to)
}

// GetRecipientByAddress is a paid mutator transaction binding the contract method 0x2342c260.
//
// Solidity: function getRecipientByAddress(address _to) returns(uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainTransactorSession) GetRecipientByAddress(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientByAddress(&_Main.TransactOpts, _to)
}

// GetRecipientByIndex is a paid mutator transaction binding the contract method 0x5af84d5c.
//
// Solidity: function getRecipientByIndex(uint256 _i) returns(address _recipient, uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainTransactor) GetRecipientByIndex(opts *bind.TransactOpts, _i *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getRecipientByIndex", _i)
}

// GetRecipientByIndex is a paid mutator transaction binding the contract method 0x5af84d5c.
//
// Solidity: function getRecipientByIndex(uint256 _i) returns(address _recipient, uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainSession) GetRecipientByIndex(_i *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientByIndex(&_Main.TransactOpts, _i)
}

// GetRecipientByIndex is a paid mutator transaction binding the contract method 0x5af84d5c.
//
// Solidity: function getRecipientByIndex(uint256 _i) returns(address _recipient, uint256 _totalBalance, uint256 _numPayments, string _name, bytes8 _country)
func (_Main *MainTransactorSession) GetRecipientByIndex(_i *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientByIndex(&_Main.TransactOpts, _i)
}

// GetRecipientCountry is a paid mutator transaction binding the contract method 0x48509539.
//
// Solidity: function getRecipientCountry(address _to) returns(bytes8)
func (_Main *MainTransactor) GetRecipientCountry(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getRecipientCountry", _to)
}

// GetRecipientCountry is a paid mutator transaction binding the contract method 0x48509539.
//
// Solidity: function getRecipientCountry(address _to) returns(bytes8)
func (_Main *MainSession) GetRecipientCountry(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientCountry(&_Main.TransactOpts, _to)
}

// GetRecipientCountry is a paid mutator transaction binding the contract method 0x48509539.
//
// Solidity: function getRecipientCountry(address _to) returns(bytes8)
func (_Main *MainTransactorSession) GetRecipientCountry(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientCountry(&_Main.TransactOpts, _to)
}

// GetRecipientName is a paid mutator transaction binding the contract method 0x6d347482.
//
// Solidity: function getRecipientName(address _to) returns(string)
func (_Main *MainTransactor) GetRecipientName(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getRecipientName", _to)
}

// GetRecipientName is a paid mutator transaction binding the contract method 0x6d347482.
//
// Solidity: function getRecipientName(address _to) returns(string)
func (_Main *MainSession) GetRecipientName(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientName(&_Main.TransactOpts, _to)
}

// GetRecipientName is a paid mutator transaction binding the contract method 0x6d347482.
//
// Solidity: function getRecipientName(address _to) returns(string)
func (_Main *MainTransactorSession) GetRecipientName(_to common.Address) (*types.Transaction, error) {
	return _Main.Contract.GetRecipientName(&_Main.TransactOpts, _to)
}

// GetRecipientsCount is a paid mutator transaction binding the contract method 0x221b17db.
//
// Solidity: function getRecipientsCount() returns(uint256)
func (_Main *MainTransactor) GetRecipientsCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getRecipientsCount")
}

// GetRecipientsCount is a paid mutator transaction binding the contract method 0x221b17db.
//
// Solidity: function getRecipientsCount() returns(uint256)
func (_Main *MainSession) GetRecipientsCount() (*types.Transaction, error) {
	return _Main.Contract.GetRecipientsCount(&_Main.TransactOpts)
}

// GetRecipientsCount is a paid mutator transaction binding the contract method 0x221b17db.
//
// Solidity: function getRecipientsCount() returns(uint256)
func (_Main *MainTransactorSession) GetRecipientsCount() (*types.Transaction, error) {
	return _Main.Contract.GetRecipientsCount(&_Main.TransactOpts)
}

// GetVersion is a paid mutator transaction binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() returns(uint16)
func (_Main *MainTransactor) GetVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "getVersion")
}

// GetVersion is a paid mutator transaction binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() returns(uint16)
func (_Main *MainSession) GetVersion() (*types.Transaction, error) {
	return _Main.Contract.GetVersion(&_Main.TransactOpts)
}

// GetVersion is a paid mutator transaction binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() returns(uint16)
func (_Main *MainTransactorSession) GetVersion() (*types.Transaction, error) {
	return _Main.Contract.GetVersion(&_Main.TransactOpts)
}

// GiveBootstrapGas is a paid mutator transaction binding the contract method 0x82b66b01.
//
// Solidity: function giveBootstrapGas(address _to, uint256 _amount) returns()
func (_Main *MainTransactor) GiveBootstrapGas(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "giveBootstrapGas", _to, _amount)
}

// GiveBootstrapGas is a paid mutator transaction binding the contract method 0x82b66b01.
//
// Solidity: function giveBootstrapGas(address _to, uint256 _amount) returns()
func (_Main *MainSession) GiveBootstrapGas(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GiveBootstrapGas(&_Main.TransactOpts, _to, _amount)
}

// GiveBootstrapGas is a paid mutator transaction binding the contract method 0x82b66b01.
//
// Solidity: function giveBootstrapGas(address _to, uint256 _amount) returns()
func (_Main *MainTransactorSession) GiveBootstrapGas(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.GiveBootstrapGas(&_Main.TransactOpts, _to, _amount)
}

// RequestDonation is a paid mutator transaction binding the contract method 0xd087f000.
//
// Solidity: function requestDonation(address _to, uint256 _amount, string _name, bytes8 _country) returns()
func (_Main *MainTransactor) RequestDonation(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _name string, _country [8]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "requestDonation", _to, _amount, _name, _country)
}

// RequestDonation is a paid mutator transaction binding the contract method 0xd087f000.
//
// Solidity: function requestDonation(address _to, uint256 _amount, string _name, bytes8 _country) returns()
func (_Main *MainSession) RequestDonation(_to common.Address, _amount *big.Int, _name string, _country [8]byte) (*types.Transaction, error) {
	return _Main.Contract.RequestDonation(&_Main.TransactOpts, _to, _amount, _name, _country)
}

// RequestDonation is a paid mutator transaction binding the contract method 0xd087f000.
//
// Solidity: function requestDonation(address _to, uint256 _amount, string _name, bytes8 _country) returns()
func (_Main *MainTransactorSession) RequestDonation(_to common.Address, _amount *big.Int, _name string, _country [8]byte) (*types.Transaction, error) {
	return _Main.Contract.RequestDonation(&_Main.TransactOpts, _to, _amount, _name, _country)
}

// SetDonorCountry is a paid mutator transaction binding the contract method 0x1639d0af.
//
// Solidity: function setDonorCountry(bytes8 _country) returns()
func (_Main *MainTransactor) SetDonorCountry(opts *bind.TransactOpts, _country [8]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setDonorCountry", _country)
}

// SetDonorCountry is a paid mutator transaction binding the contract method 0x1639d0af.
//
// Solidity: function setDonorCountry(bytes8 _country) returns()
func (_Main *MainSession) SetDonorCountry(_country [8]byte) (*types.Transaction, error) {
	return _Main.Contract.SetDonorCountry(&_Main.TransactOpts, _country)
}

// SetDonorCountry is a paid mutator transaction binding the contract method 0x1639d0af.
//
// Solidity: function setDonorCountry(bytes8 _country) returns()
func (_Main *MainTransactorSession) SetDonorCountry(_country [8]byte) (*types.Transaction, error) {
	return _Main.Contract.SetDonorCountry(&_Main.TransactOpts, _country)
}

// SetDonorName is a paid mutator transaction binding the contract method 0x34607994.
//
// Solidity: function setDonorName(string _name) returns()
func (_Main *MainTransactor) SetDonorName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setDonorName", _name)
}

// SetDonorName is a paid mutator transaction binding the contract method 0x34607994.
//
// Solidity: function setDonorName(string _name) returns()
func (_Main *MainSession) SetDonorName(_name string) (*types.Transaction, error) {
	return _Main.Contract.SetDonorName(&_Main.TransactOpts, _name)
}

// SetDonorName is a paid mutator transaction binding the contract method 0x34607994.
//
// Solidity: function setDonorName(string _name) returns()
func (_Main *MainTransactorSession) SetDonorName(_name string) (*types.Transaction, error) {
	return _Main.Contract.SetDonorName(&_Main.TransactOpts, _name)
}

// SetRecipientCountry is a paid mutator transaction binding the contract method 0xce766b3b.
//
// Solidity: function setRecipientCountry(bytes8 _country) returns()
func (_Main *MainTransactor) SetRecipientCountry(opts *bind.TransactOpts, _country [8]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setRecipientCountry", _country)
}

// SetRecipientCountry is a paid mutator transaction binding the contract method 0xce766b3b.
//
// Solidity: function setRecipientCountry(bytes8 _country) returns()
func (_Main *MainSession) SetRecipientCountry(_country [8]byte) (*types.Transaction, error) {
	return _Main.Contract.SetRecipientCountry(&_Main.TransactOpts, _country)
}

// SetRecipientCountry is a paid mutator transaction binding the contract method 0xce766b3b.
//
// Solidity: function setRecipientCountry(bytes8 _country) returns()
func (_Main *MainTransactorSession) SetRecipientCountry(_country [8]byte) (*types.Transaction, error) {
	return _Main.Contract.SetRecipientCountry(&_Main.TransactOpts, _country)
}

// SetRecipientName is a paid mutator transaction binding the contract method 0x4655458d.
//
// Solidity: function setRecipientName(string _name) returns()
func (_Main *MainTransactor) SetRecipientName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setRecipientName", _name)
}

// SetRecipientName is a paid mutator transaction binding the contract method 0x4655458d.
//
// Solidity: function setRecipientName(string _name) returns()
func (_Main *MainSession) SetRecipientName(_name string) (*types.Transaction, error) {
	return _Main.Contract.SetRecipientName(&_Main.TransactOpts, _name)
}

// SetRecipientName is a paid mutator transaction binding the contract method 0x4655458d.
//
// Solidity: function setRecipientName(string _name) returns()
func (_Main *MainTransactorSession) SetRecipientName(_name string) (*types.Transaction, error) {
	return _Main.Contract.SetRecipientName(&_Main.TransactOpts, _name)
}

// Version is a paid mutator transaction binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(uint16)
func (_Main *MainTransactor) Version(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "version")
}

// Version is a paid mutator transaction binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(uint16)
func (_Main *MainSession) Version() (*types.Transaction, error) {
	return _Main.Contract.Version(&_Main.TransactOpts)
}

// Version is a paid mutator transaction binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(uint16)
func (_Main *MainTransactorSession) Version() (*types.Transaction, error) {
	return _Main.Contract.Version(&_Main.TransactOpts)
}

// WithdrawDonation is a paid mutator transaction binding the contract method 0xe9a04974.
//
// Solidity: function withdrawDonation(address _to, uint256 _amount) returns()
func (_Main *MainTransactor) WithdrawDonation(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "withdrawDonation", _to, _amount)
}

// WithdrawDonation is a paid mutator transaction binding the contract method 0xe9a04974.
//
// Solidity: function withdrawDonation(address _to, uint256 _amount) returns()
func (_Main *MainSession) WithdrawDonation(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.WithdrawDonation(&_Main.TransactOpts, _to, _amount)
}

// WithdrawDonation is a paid mutator transaction binding the contract method 0xe9a04974.
//
// Solidity: function withdrawDonation(address _to, uint256 _amount) returns()
func (_Main *MainTransactorSession) WithdrawDonation(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.WithdrawDonation(&_Main.TransactOpts, _to, _amount)
}

// MainGivenDonationIterator is returned from FilterGivenDonation and is used to iterate over the raw logs and unpacked data for GivenDonation events raised by the Main contract.
type MainGivenDonationIterator struct {
	Event *MainGivenDonation // Event containing the contract specifics and raw log

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
func (it *MainGivenDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainGivenDonation)
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
		it.Event = new(MainGivenDonation)
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
func (it *MainGivenDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainGivenDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainGivenDonation represents a GivenDonation event raised by the Main contract.
type MainGivenDonation struct {
	Recipient      common.Address
	Amount         *big.Int
	Donationnumber *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGivenDonation is a free log retrieval operation binding the contract event 0x44786e582df7ed007f4d4ec0fa031b3b32bb594c3451da5b954bb2349715e2a5.
//
// Solidity: event GivenDonation(address indexed recipient, uint256 amount, uint256 donationnumber)
func (_Main *MainFilterer) FilterGivenDonation(opts *bind.FilterOpts, recipient []common.Address) (*MainGivenDonationIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "GivenDonation", recipientRule)
	if err != nil {
		return nil, err
	}
	return &MainGivenDonationIterator{contract: _Main.contract, event: "GivenDonation", logs: logs, sub: sub}, nil
}

// WatchGivenDonation is a free log subscription operation binding the contract event 0x44786e582df7ed007f4d4ec0fa031b3b32bb594c3451da5b954bb2349715e2a5.
//
// Solidity: event GivenDonation(address indexed recipient, uint256 amount, uint256 donationnumber)
func (_Main *MainFilterer) WatchGivenDonation(opts *bind.WatchOpts, sink chan<- *MainGivenDonation, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "GivenDonation", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainGivenDonation)
				if err := _Main.contract.UnpackLog(event, "GivenDonation", log); err != nil {
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

// MainOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the Main contract.
type MainOwnerSetIterator struct {
	Event *MainOwnerSet // Event containing the contract specifics and raw log

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
func (it *MainOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainOwnerSet)
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
		it.Event = new(MainOwnerSet)
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
func (it *MainOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainOwnerSet represents a OwnerSet event raised by the Main contract.
type MainOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Main *MainFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*MainOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainOwnerSetIterator{contract: _Main.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Main *MainFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *MainOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainOwnerSet)
				if err := _Main.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// MainReceivedDonationIterator is returned from FilterReceivedDonation and is used to iterate over the raw logs and unpacked data for ReceivedDonation events raised by the Main contract.
type MainReceivedDonationIterator struct {
	Event *MainReceivedDonation // Event containing the contract specifics and raw log

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
func (it *MainReceivedDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainReceivedDonation)
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
		it.Event = new(MainReceivedDonation)
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
func (it *MainReceivedDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainReceivedDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainReceivedDonation represents a ReceivedDonation event raised by the Main contract.
type MainReceivedDonation struct {
	Donor          common.Address
	Amount         *big.Int
	Donationnumber *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReceivedDonation is a free log retrieval operation binding the contract event 0xcd4a13669adcedfe2d4597b3fdf126cf0555673710873d2f5882c1668f1d46f9.
//
// Solidity: event ReceivedDonation(address indexed donor, uint256 amount, uint256 donationnumber)
func (_Main *MainFilterer) FilterReceivedDonation(opts *bind.FilterOpts, donor []common.Address) (*MainReceivedDonationIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ReceivedDonation", donorRule)
	if err != nil {
		return nil, err
	}
	return &MainReceivedDonationIterator{contract: _Main.contract, event: "ReceivedDonation", logs: logs, sub: sub}, nil
}

// WatchReceivedDonation is a free log subscription operation binding the contract event 0xcd4a13669adcedfe2d4597b3fdf126cf0555673710873d2f5882c1668f1d46f9.
//
// Solidity: event ReceivedDonation(address indexed donor, uint256 amount, uint256 donationnumber)
func (_Main *MainFilterer) WatchReceivedDonation(opts *bind.WatchOpts, sink chan<- *MainReceivedDonation, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ReceivedDonation", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainReceivedDonation)
				if err := _Main.contract.UnpackLog(event, "ReceivedDonation", log); err != nil {
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

// MainWithdrawMoneyIterator is returned from FilterWithdrawMoney and is used to iterate over the raw logs and unpacked data for WithdrawMoney events raised by the Main contract.
type MainWithdrawMoneyIterator struct {
	Event *MainWithdrawMoney // Event containing the contract specifics and raw log

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
func (it *MainWithdrawMoneyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainWithdrawMoney)
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
		it.Event = new(MainWithdrawMoney)
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
func (it *MainWithdrawMoneyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainWithdrawMoneyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainWithdrawMoney represents a WithdrawMoney event raised by the Main contract.
type MainWithdrawMoney struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawMoney is a free log retrieval operation binding the contract event 0x9fe2db218b58c00ab3556f2be44b2769b37978368c63313d82d78306568e4776.
//
// Solidity: event WithdrawMoney(address indexed to, uint256 indexed amount)
func (_Main *MainFilterer) FilterWithdrawMoney(opts *bind.FilterOpts, to []common.Address, amount []*big.Int) (*MainWithdrawMoneyIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "WithdrawMoney", toRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MainWithdrawMoneyIterator{contract: _Main.contract, event: "WithdrawMoney", logs: logs, sub: sub}, nil
}

// WatchWithdrawMoney is a free log subscription operation binding the contract event 0x9fe2db218b58c00ab3556f2be44b2769b37978368c63313d82d78306568e4776.
//
// Solidity: event WithdrawMoney(address indexed to, uint256 indexed amount)
func (_Main *MainFilterer) WatchWithdrawMoney(opts *bind.WatchOpts, sink chan<- *MainWithdrawMoney, to []common.Address, amount []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "WithdrawMoney", toRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainWithdrawMoney)
				if err := _Main.contract.UnpackLog(event, "WithdrawMoney", log); err != nil {
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
