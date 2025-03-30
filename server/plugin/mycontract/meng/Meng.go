// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package meng

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

// MengOprationComprehensiveRecord is an auto generated low-level Go binding around an user-defined struct.
type MengOprationComprehensiveRecord struct {
	RelatedAddress common.Address
	Amount         *big.Int
	Types          uint8
	Timestamp      *big.Int
	IsWithdraw     bool
}

// MycontractMetaData contains all meta data concerning the Mycontract contract.
var MycontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"actionType\",\"type\":\"uint8\"}],\"name\":\"TransferAndActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnAll\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"boxreward\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNodeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cliBoxwith\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"cliUpdateUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecosystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getRecordsByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"relatedAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"types\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isWithdraw\",\"type\":\"bool\"}],\"internalType\":\"structMengOpration.ComprehensiveRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeamount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodereceive\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodereward\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateboxreward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateecosystem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratenodereward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratetechnology\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateupdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recordMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"relatedAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"types\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isWithdraw\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registernum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootaddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresss\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_prizeTypes\",\"type\":\"uint8[]\"}],\"name\":\"serveBoxset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_technology\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ecosystem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodereward\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_boxreward\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodereceive\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ratetechnology\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rateecosystem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ratenodereward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rateboxreward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rateupdate\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeamount\",\"type\":\"uint256\"}],\"name\":\"setRigisternum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"setWithAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"severFromtowith\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_invite\",\"type\":\"address\"}],\"name\":\"severRigister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"technology\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_invite\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_vip\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_node\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_nodereward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_boxreward\",\"type\":\"uint256\"}],\"name\":\"updateUserInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfoMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"myaddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"vip\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"node\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nodereward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeinvite\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boxreward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatereward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invitereward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodewithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boxwithdraw\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userlists\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_words\",\"type\":\"uint256\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawBnb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_words\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MycontractABI is the input ABI used to generate the binding from.
// Deprecated: Use MycontractMetaData.ABI instead.
var MycontractABI = MycontractMetaData.ABI

// Mycontract is an auto generated Go binding around an Ethereum contract.
type Mycontract struct {
	MycontractCaller     // Read-only binding to the contract
	MycontractTransactor // Write-only binding to the contract
	MycontractFilterer   // Log filterer for contract events
}

// MycontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MycontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MycontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MycontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MycontractSession struct {
	Contract     *Mycontract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MycontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MycontractCallerSession struct {
	Contract *MycontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MycontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MycontractTransactorSession struct {
	Contract     *MycontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MycontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MycontractRaw struct {
	Contract *Mycontract // Generic contract binding to access the raw methods on
}

// MycontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MycontractCallerRaw struct {
	Contract *MycontractCaller // Generic read-only contract binding to access the raw methods on
}

// MycontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MycontractTransactorRaw struct {
	Contract *MycontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMycontract creates a new instance of Mycontract, bound to a specific deployed contract.
func NewMycontract(address common.Address, backend bind.ContractBackend) (*Mycontract, error) {
	contract, err := bindMycontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mycontract{MycontractCaller: MycontractCaller{contract: contract}, MycontractTransactor: MycontractTransactor{contract: contract}, MycontractFilterer: MycontractFilterer{contract: contract}}, nil
}

// NewMycontractCaller creates a new read-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractCaller(address common.Address, caller bind.ContractCaller) (*MycontractCaller, error) {
	contract, err := bindMycontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractCaller{contract: contract}, nil
}

// NewMycontractTransactor creates a new write-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractTransactor(address common.Address, transactor bind.ContractTransactor) (*MycontractTransactor, error) {
	contract, err := bindMycontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractTransactor{contract: contract}, nil
}

// NewMycontractFilterer creates a new log filterer instance of Mycontract, bound to a specific deployed contract.
func NewMycontractFilterer(address common.Address, filterer bind.ContractFilterer) (*MycontractFilterer, error) {
	contract, err := bindMycontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MycontractFilterer{contract: contract}, nil
}

// bindMycontract binds a generic wrapper to an already deployed contract.
func bindMycontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MycontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.MycontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transact(opts, method, params...)
}

// Boxreward is a free data retrieval call binding the contract method 0x34e1b07a.
//
// Solidity: function boxreward() view returns(address)
func (_Mycontract *MycontractCaller) Boxreward(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "boxreward")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Boxreward is a free data retrieval call binding the contract method 0x34e1b07a.
//
// Solidity: function boxreward() view returns(address)
func (_Mycontract *MycontractSession) Boxreward() (common.Address, error) {
	return _Mycontract.Contract.Boxreward(&_Mycontract.CallOpts)
}

// Boxreward is a free data retrieval call binding the contract method 0x34e1b07a.
//
// Solidity: function boxreward() view returns(address)
func (_Mycontract *MycontractCallerSession) Boxreward() (common.Address, error) {
	return _Mycontract.Contract.Boxreward(&_Mycontract.CallOpts)
}

// Ecosystem is a free data retrieval call binding the contract method 0x9c74a579.
//
// Solidity: function ecosystem() view returns(address)
func (_Mycontract *MycontractCaller) Ecosystem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "ecosystem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ecosystem is a free data retrieval call binding the contract method 0x9c74a579.
//
// Solidity: function ecosystem() view returns(address)
func (_Mycontract *MycontractSession) Ecosystem() (common.Address, error) {
	return _Mycontract.Contract.Ecosystem(&_Mycontract.CallOpts)
}

// Ecosystem is a free data retrieval call binding the contract method 0x9c74a579.
//
// Solidity: function ecosystem() view returns(address)
func (_Mycontract *MycontractCallerSession) Ecosystem() (common.Address, error) {
	return _Mycontract.Contract.Ecosystem(&_Mycontract.CallOpts)
}

// GetNodeList is a free data retrieval call binding the contract method 0x53f3b713.
//
// Solidity: function getNodeList() view returns(address[])
func (_Mycontract *MycontractCaller) GetNodeList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getNodeList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetNodeList is a free data retrieval call binding the contract method 0x53f3b713.
//
// Solidity: function getNodeList() view returns(address[])
func (_Mycontract *MycontractSession) GetNodeList() ([]common.Address, error) {
	return _Mycontract.Contract.GetNodeList(&_Mycontract.CallOpts)
}

// GetNodeList is a free data retrieval call binding the contract method 0x53f3b713.
//
// Solidity: function getNodeList() view returns(address[])
func (_Mycontract *MycontractCallerSession) GetNodeList() ([]common.Address, error) {
	return _Mycontract.Contract.GetNodeList(&_Mycontract.CallOpts)
}

// GetRecordsByAddress is a free data retrieval call binding the contract method 0xad88bb27.
//
// Solidity: function getRecordsByAddress(address _user) view returns((address,uint256,uint8,uint256,bool)[])
func (_Mycontract *MycontractCaller) GetRecordsByAddress(opts *bind.CallOpts, _user common.Address) ([]MengOprationComprehensiveRecord, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getRecordsByAddress", _user)

	if err != nil {
		return *new([]MengOprationComprehensiveRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]MengOprationComprehensiveRecord)).(*[]MengOprationComprehensiveRecord)

	return out0, err

}

// GetRecordsByAddress is a free data retrieval call binding the contract method 0xad88bb27.
//
// Solidity: function getRecordsByAddress(address _user) view returns((address,uint256,uint8,uint256,bool)[])
func (_Mycontract *MycontractSession) GetRecordsByAddress(_user common.Address) ([]MengOprationComprehensiveRecord, error) {
	return _Mycontract.Contract.GetRecordsByAddress(&_Mycontract.CallOpts, _user)
}

// GetRecordsByAddress is a free data retrieval call binding the contract method 0xad88bb27.
//
// Solidity: function getRecordsByAddress(address _user) view returns((address,uint256,uint8,uint256,bool)[])
func (_Mycontract *MycontractCallerSession) GetRecordsByAddress(_user common.Address) ([]MengOprationComprehensiveRecord, error) {
	return _Mycontract.Contract.GetRecordsByAddress(&_Mycontract.CallOpts, _user)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_Mycontract *MycontractCaller) GetUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_Mycontract *MycontractSession) GetUsers() ([]common.Address, error) {
	return _Mycontract.Contract.GetUsers(&_Mycontract.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_Mycontract *MycontractCallerSession) GetUsers() ([]common.Address, error) {
	return _Mycontract.Contract.GetUsers(&_Mycontract.CallOpts)
}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_Mycontract *MycontractCaller) NodeList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "nodeList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_Mycontract *MycontractSession) NodeList(arg0 *big.Int) (common.Address, error) {
	return _Mycontract.Contract.NodeList(&_Mycontract.CallOpts, arg0)
}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_Mycontract *MycontractCallerSession) NodeList(arg0 *big.Int) (common.Address, error) {
	return _Mycontract.Contract.NodeList(&_Mycontract.CallOpts, arg0)
}

// Nodeamount is a free data retrieval call binding the contract method 0x16d1347d.
//
// Solidity: function nodeamount() view returns(uint256)
func (_Mycontract *MycontractCaller) Nodeamount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "nodeamount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nodeamount is a free data retrieval call binding the contract method 0x16d1347d.
//
// Solidity: function nodeamount() view returns(uint256)
func (_Mycontract *MycontractSession) Nodeamount() (*big.Int, error) {
	return _Mycontract.Contract.Nodeamount(&_Mycontract.CallOpts)
}

// Nodeamount is a free data retrieval call binding the contract method 0x16d1347d.
//
// Solidity: function nodeamount() view returns(uint256)
func (_Mycontract *MycontractCallerSession) Nodeamount() (*big.Int, error) {
	return _Mycontract.Contract.Nodeamount(&_Mycontract.CallOpts)
}

// Nodereceive is a free data retrieval call binding the contract method 0x682a0c59.
//
// Solidity: function nodereceive() view returns(address)
func (_Mycontract *MycontractCaller) Nodereceive(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "nodereceive")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nodereceive is a free data retrieval call binding the contract method 0x682a0c59.
//
// Solidity: function nodereceive() view returns(address)
func (_Mycontract *MycontractSession) Nodereceive() (common.Address, error) {
	return _Mycontract.Contract.Nodereceive(&_Mycontract.CallOpts)
}

// Nodereceive is a free data retrieval call binding the contract method 0x682a0c59.
//
// Solidity: function nodereceive() view returns(address)
func (_Mycontract *MycontractCallerSession) Nodereceive() (common.Address, error) {
	return _Mycontract.Contract.Nodereceive(&_Mycontract.CallOpts)
}

// Nodereward is a free data retrieval call binding the contract method 0x4d41e76f.
//
// Solidity: function nodereward() view returns(address)
func (_Mycontract *MycontractCaller) Nodereward(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "nodereward")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nodereward is a free data retrieval call binding the contract method 0x4d41e76f.
//
// Solidity: function nodereward() view returns(address)
func (_Mycontract *MycontractSession) Nodereward() (common.Address, error) {
	return _Mycontract.Contract.Nodereward(&_Mycontract.CallOpts)
}

// Nodereward is a free data retrieval call binding the contract method 0x4d41e76f.
//
// Solidity: function nodereward() view returns(address)
func (_Mycontract *MycontractCallerSession) Nodereward() (common.Address, error) {
	return _Mycontract.Contract.Nodereward(&_Mycontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractSession) Owner() (common.Address, error) {
	return _Mycontract.Contract.Owner(&_Mycontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractCallerSession) Owner() (common.Address, error) {
	return _Mycontract.Contract.Owner(&_Mycontract.CallOpts)
}

// Rateboxreward is a free data retrieval call binding the contract method 0x5d1f8ae6.
//
// Solidity: function rateboxreward() view returns(uint256)
func (_Mycontract *MycontractCaller) Rateboxreward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "rateboxreward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rateboxreward is a free data retrieval call binding the contract method 0x5d1f8ae6.
//
// Solidity: function rateboxreward() view returns(uint256)
func (_Mycontract *MycontractSession) Rateboxreward() (*big.Int, error) {
	return _Mycontract.Contract.Rateboxreward(&_Mycontract.CallOpts)
}

// Rateboxreward is a free data retrieval call binding the contract method 0x5d1f8ae6.
//
// Solidity: function rateboxreward() view returns(uint256)
func (_Mycontract *MycontractCallerSession) Rateboxreward() (*big.Int, error) {
	return _Mycontract.Contract.Rateboxreward(&_Mycontract.CallOpts)
}

// Rateecosystem is a free data retrieval call binding the contract method 0x6e5d0b3f.
//
// Solidity: function rateecosystem() view returns(uint256)
func (_Mycontract *MycontractCaller) Rateecosystem(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "rateecosystem")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rateecosystem is a free data retrieval call binding the contract method 0x6e5d0b3f.
//
// Solidity: function rateecosystem() view returns(uint256)
func (_Mycontract *MycontractSession) Rateecosystem() (*big.Int, error) {
	return _Mycontract.Contract.Rateecosystem(&_Mycontract.CallOpts)
}

// Rateecosystem is a free data retrieval call binding the contract method 0x6e5d0b3f.
//
// Solidity: function rateecosystem() view returns(uint256)
func (_Mycontract *MycontractCallerSession) Rateecosystem() (*big.Int, error) {
	return _Mycontract.Contract.Rateecosystem(&_Mycontract.CallOpts)
}

// Ratenodereward is a free data retrieval call binding the contract method 0xc63f65cb.
//
// Solidity: function ratenodereward() view returns(uint256)
func (_Mycontract *MycontractCaller) Ratenodereward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "ratenodereward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ratenodereward is a free data retrieval call binding the contract method 0xc63f65cb.
//
// Solidity: function ratenodereward() view returns(uint256)
func (_Mycontract *MycontractSession) Ratenodereward() (*big.Int, error) {
	return _Mycontract.Contract.Ratenodereward(&_Mycontract.CallOpts)
}

// Ratenodereward is a free data retrieval call binding the contract method 0xc63f65cb.
//
// Solidity: function ratenodereward() view returns(uint256)
func (_Mycontract *MycontractCallerSession) Ratenodereward() (*big.Int, error) {
	return _Mycontract.Contract.Ratenodereward(&_Mycontract.CallOpts)
}

// Ratetechnology is a free data retrieval call binding the contract method 0xc4ac72b8.
//
// Solidity: function ratetechnology() view returns(uint256)
func (_Mycontract *MycontractCaller) Ratetechnology(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "ratetechnology")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ratetechnology is a free data retrieval call binding the contract method 0xc4ac72b8.
//
// Solidity: function ratetechnology() view returns(uint256)
func (_Mycontract *MycontractSession) Ratetechnology() (*big.Int, error) {
	return _Mycontract.Contract.Ratetechnology(&_Mycontract.CallOpts)
}

// Ratetechnology is a free data retrieval call binding the contract method 0xc4ac72b8.
//
// Solidity: function ratetechnology() view returns(uint256)
func (_Mycontract *MycontractCallerSession) Ratetechnology() (*big.Int, error) {
	return _Mycontract.Contract.Ratetechnology(&_Mycontract.CallOpts)
}

// Rateupdate is a free data retrieval call binding the contract method 0xd17f7fad.
//
// Solidity: function rateupdate() view returns(uint256)
func (_Mycontract *MycontractCaller) Rateupdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "rateupdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rateupdate is a free data retrieval call binding the contract method 0xd17f7fad.
//
// Solidity: function rateupdate() view returns(uint256)
func (_Mycontract *MycontractSession) Rateupdate() (*big.Int, error) {
	return _Mycontract.Contract.Rateupdate(&_Mycontract.CallOpts)
}

// Rateupdate is a free data retrieval call binding the contract method 0xd17f7fad.
//
// Solidity: function rateupdate() view returns(uint256)
func (_Mycontract *MycontractCallerSession) Rateupdate() (*big.Int, error) {
	return _Mycontract.Contract.Rateupdate(&_Mycontract.CallOpts)
}

// RecordMap is a free data retrieval call binding the contract method 0xfe7d9e27.
//
// Solidity: function recordMap(address , uint256 ) view returns(address relatedAddress, uint256 amount, uint8 types, uint256 timestamp, bool isWithdraw)
func (_Mycontract *MycontractCaller) RecordMap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	RelatedAddress common.Address
	Amount         *big.Int
	Types          uint8
	Timestamp      *big.Int
	IsWithdraw     bool
}, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "recordMap", arg0, arg1)

	outstruct := new(struct {
		RelatedAddress common.Address
		Amount         *big.Int
		Types          uint8
		Timestamp      *big.Int
		IsWithdraw     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RelatedAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Types = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsWithdraw = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// RecordMap is a free data retrieval call binding the contract method 0xfe7d9e27.
//
// Solidity: function recordMap(address , uint256 ) view returns(address relatedAddress, uint256 amount, uint8 types, uint256 timestamp, bool isWithdraw)
func (_Mycontract *MycontractSession) RecordMap(arg0 common.Address, arg1 *big.Int) (struct {
	RelatedAddress common.Address
	Amount         *big.Int
	Types          uint8
	Timestamp      *big.Int
	IsWithdraw     bool
}, error) {
	return _Mycontract.Contract.RecordMap(&_Mycontract.CallOpts, arg0, arg1)
}

// RecordMap is a free data retrieval call binding the contract method 0xfe7d9e27.
//
// Solidity: function recordMap(address , uint256 ) view returns(address relatedAddress, uint256 amount, uint8 types, uint256 timestamp, bool isWithdraw)
func (_Mycontract *MycontractCallerSession) RecordMap(arg0 common.Address, arg1 *big.Int) (struct {
	RelatedAddress common.Address
	Amount         *big.Int
	Types          uint8
	Timestamp      *big.Int
	IsWithdraw     bool
}, error) {
	return _Mycontract.Contract.RecordMap(&_Mycontract.CallOpts, arg0, arg1)
}

// Registernum is a free data retrieval call binding the contract method 0x2b0e5908.
//
// Solidity: function registernum() view returns(uint256)
func (_Mycontract *MycontractCaller) Registernum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "registernum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Registernum is a free data retrieval call binding the contract method 0x2b0e5908.
//
// Solidity: function registernum() view returns(uint256)
func (_Mycontract *MycontractSession) Registernum() (*big.Int, error) {
	return _Mycontract.Contract.Registernum(&_Mycontract.CallOpts)
}

// Registernum is a free data retrieval call binding the contract method 0x2b0e5908.
//
// Solidity: function registernum() view returns(uint256)
func (_Mycontract *MycontractCallerSession) Registernum() (*big.Int, error) {
	return _Mycontract.Contract.Registernum(&_Mycontract.CallOpts)
}

// Rootaddress is a free data retrieval call binding the contract method 0x02f44c02.
//
// Solidity: function rootaddress() view returns(address)
func (_Mycontract *MycontractCaller) Rootaddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "rootaddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rootaddress is a free data retrieval call binding the contract method 0x02f44c02.
//
// Solidity: function rootaddress() view returns(address)
func (_Mycontract *MycontractSession) Rootaddress() (common.Address, error) {
	return _Mycontract.Contract.Rootaddress(&_Mycontract.CallOpts)
}

// Rootaddress is a free data retrieval call binding the contract method 0x02f44c02.
//
// Solidity: function rootaddress() view returns(address)
func (_Mycontract *MycontractCallerSession) Rootaddress() (common.Address, error) {
	return _Mycontract.Contract.Rootaddress(&_Mycontract.CallOpts)
}

// Technology is a free data retrieval call binding the contract method 0xc49160e5.
//
// Solidity: function technology() view returns(address)
func (_Mycontract *MycontractCaller) Technology(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "technology")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Technology is a free data retrieval call binding the contract method 0xc49160e5.
//
// Solidity: function technology() view returns(address)
func (_Mycontract *MycontractSession) Technology() (common.Address, error) {
	return _Mycontract.Contract.Technology(&_Mycontract.CallOpts)
}

// Technology is a free data retrieval call binding the contract method 0xc49160e5.
//
// Solidity: function technology() view returns(address)
func (_Mycontract *MycontractCallerSession) Technology() (common.Address, error) {
	return _Mycontract.Contract.Technology(&_Mycontract.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Mycontract *MycontractCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "tokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Mycontract *MycontractSession) TokenContract() (common.Address, error) {
	return _Mycontract.Contract.TokenContract(&_Mycontract.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Mycontract *MycontractCallerSession) TokenContract() (common.Address, error) {
	return _Mycontract.Contract.TokenContract(&_Mycontract.CallOpts)
}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_Mycontract *MycontractCaller) UsdtToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "usdtToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_Mycontract *MycontractSession) UsdtToken() (common.Address, error) {
	return _Mycontract.Contract.UsdtToken(&_Mycontract.CallOpts)
}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_Mycontract *MycontractCallerSession) UsdtToken() (common.Address, error) {
	return _Mycontract.Contract.UsdtToken(&_Mycontract.CallOpts)
}

// UserInfoMap is a free data retrieval call binding the contract method 0xaa779a53.
//
// Solidity: function userInfoMap(address ) view returns(address myaddr, address inviter, uint8 vip, uint8 node, uint256 nodereward, uint256 nodeinvite, uint256 boxreward, uint256 updatereward, uint256 invitereward, uint256 nodewithdraw, uint256 boxwithdraw)
func (_Mycontract *MycontractCaller) UserInfoMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Myaddr       common.Address
	Inviter      common.Address
	Vip          uint8
	Node         uint8
	Nodereward   *big.Int
	Nodeinvite   *big.Int
	Boxreward    *big.Int
	Updatereward *big.Int
	Invitereward *big.Int
	Nodewithdraw *big.Int
	Boxwithdraw  *big.Int
}, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "userInfoMap", arg0)

	outstruct := new(struct {
		Myaddr       common.Address
		Inviter      common.Address
		Vip          uint8
		Node         uint8
		Nodereward   *big.Int
		Nodeinvite   *big.Int
		Boxreward    *big.Int
		Updatereward *big.Int
		Invitereward *big.Int
		Nodewithdraw *big.Int
		Boxwithdraw  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Myaddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Inviter = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Vip = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Node = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Nodereward = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Nodeinvite = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Boxreward = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Updatereward = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Invitereward = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Nodewithdraw = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Boxwithdraw = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfoMap is a free data retrieval call binding the contract method 0xaa779a53.
//
// Solidity: function userInfoMap(address ) view returns(address myaddr, address inviter, uint8 vip, uint8 node, uint256 nodereward, uint256 nodeinvite, uint256 boxreward, uint256 updatereward, uint256 invitereward, uint256 nodewithdraw, uint256 boxwithdraw)
func (_Mycontract *MycontractSession) UserInfoMap(arg0 common.Address) (struct {
	Myaddr       common.Address
	Inviter      common.Address
	Vip          uint8
	Node         uint8
	Nodereward   *big.Int
	Nodeinvite   *big.Int
	Boxreward    *big.Int
	Updatereward *big.Int
	Invitereward *big.Int
	Nodewithdraw *big.Int
	Boxwithdraw  *big.Int
}, error) {
	return _Mycontract.Contract.UserInfoMap(&_Mycontract.CallOpts, arg0)
}

// UserInfoMap is a free data retrieval call binding the contract method 0xaa779a53.
//
// Solidity: function userInfoMap(address ) view returns(address myaddr, address inviter, uint8 vip, uint8 node, uint256 nodereward, uint256 nodeinvite, uint256 boxreward, uint256 updatereward, uint256 invitereward, uint256 nodewithdraw, uint256 boxwithdraw)
func (_Mycontract *MycontractCallerSession) UserInfoMap(arg0 common.Address) (struct {
	Myaddr       common.Address
	Inviter      common.Address
	Vip          uint8
	Node         uint8
	Nodereward   *big.Int
	Nodeinvite   *big.Int
	Boxreward    *big.Int
	Updatereward *big.Int
	Invitereward *big.Int
	Nodewithdraw *big.Int
	Boxwithdraw  *big.Int
}, error) {
	return _Mycontract.Contract.UserInfoMap(&_Mycontract.CallOpts, arg0)
}

// Userlists is a free data retrieval call binding the contract method 0xeba71d4d.
//
// Solidity: function userlists(uint256 ) view returns(address)
func (_Mycontract *MycontractCaller) Userlists(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "userlists", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Userlists is a free data retrieval call binding the contract method 0xeba71d4d.
//
// Solidity: function userlists(uint256 ) view returns(address)
func (_Mycontract *MycontractSession) Userlists(arg0 *big.Int) (common.Address, error) {
	return _Mycontract.Contract.Userlists(&_Mycontract.CallOpts, arg0)
}

// Userlists is a free data retrieval call binding the contract method 0xeba71d4d.
//
// Solidity: function userlists(uint256 ) view returns(address)
func (_Mycontract *MycontractCallerSession) Userlists(arg0 *big.Int) (common.Address, error) {
	return _Mycontract.Contract.Userlists(&_Mycontract.CallOpts, arg0)
}

// WithAddress is a free data retrieval call binding the contract method 0x6037b001.
//
// Solidity: function withAddress() view returns(address)
func (_Mycontract *MycontractCaller) WithAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "withAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithAddress is a free data retrieval call binding the contract method 0x6037b001.
//
// Solidity: function withAddress() view returns(address)
func (_Mycontract *MycontractSession) WithAddress() (common.Address, error) {
	return _Mycontract.Contract.WithAddress(&_Mycontract.CallOpts)
}

// WithAddress is a free data retrieval call binding the contract method 0x6037b001.
//
// Solidity: function withAddress() view returns(address)
func (_Mycontract *MycontractCallerSession) WithAddress() (common.Address, error) {
	return _Mycontract.Contract.WithAddress(&_Mycontract.CallOpts)
}

// ClaimNodeRewards is a paid mutator transaction binding the contract method 0xa58b603b.
//
// Solidity: function claimNodeRewards() returns()
func (_Mycontract *MycontractTransactor) ClaimNodeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "claimNodeRewards")
}

// ClaimNodeRewards is a paid mutator transaction binding the contract method 0xa58b603b.
//
// Solidity: function claimNodeRewards() returns()
func (_Mycontract *MycontractSession) ClaimNodeRewards() (*types.Transaction, error) {
	return _Mycontract.Contract.ClaimNodeRewards(&_Mycontract.TransactOpts)
}

// ClaimNodeRewards is a paid mutator transaction binding the contract method 0xa58b603b.
//
// Solidity: function claimNodeRewards() returns()
func (_Mycontract *MycontractTransactorSession) ClaimNodeRewards() (*types.Transaction, error) {
	return _Mycontract.Contract.ClaimNodeRewards(&_Mycontract.TransactOpts)
}

// CliBoxwith is a paid mutator transaction binding the contract method 0xa30bd3ad.
//
// Solidity: function cliBoxwith() returns()
func (_Mycontract *MycontractTransactor) CliBoxwith(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "cliBoxwith")
}

// CliBoxwith is a paid mutator transaction binding the contract method 0xa30bd3ad.
//
// Solidity: function cliBoxwith() returns()
func (_Mycontract *MycontractSession) CliBoxwith() (*types.Transaction, error) {
	return _Mycontract.Contract.CliBoxwith(&_Mycontract.TransactOpts)
}

// CliBoxwith is a paid mutator transaction binding the contract method 0xa30bd3ad.
//
// Solidity: function cliBoxwith() returns()
func (_Mycontract *MycontractTransactorSession) CliBoxwith() (*types.Transaction, error) {
	return _Mycontract.Contract.CliBoxwith(&_Mycontract.TransactOpts)
}

// CliUpdateUser is a paid mutator transaction binding the contract method 0x045c96f8.
//
// Solidity: function cliUpdateUser(address _to) returns(bool)
func (_Mycontract *MycontractTransactor) CliUpdateUser(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "cliUpdateUser", _to)
}

// CliUpdateUser is a paid mutator transaction binding the contract method 0x045c96f8.
//
// Solidity: function cliUpdateUser(address _to) returns(bool)
func (_Mycontract *MycontractSession) CliUpdateUser(_to common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.CliUpdateUser(&_Mycontract.TransactOpts, _to)
}

// CliUpdateUser is a paid mutator transaction binding the contract method 0x045c96f8.
//
// Solidity: function cliUpdateUser(address _to) returns(bool)
func (_Mycontract *MycontractTransactorSession) CliUpdateUser(_to common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.CliUpdateUser(&_Mycontract.TransactOpts, _to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mycontract *MycontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mycontract *MycontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mycontract.Contract.RenounceOwnership(&_Mycontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mycontract *MycontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mycontract.Contract.RenounceOwnership(&_Mycontract.TransactOpts)
}

// ServeBoxset is a paid mutator transaction binding the contract method 0x004a94bd.
//
// Solidity: function serveBoxset(address[] _addresss, uint256[] _amounts, uint8[] _prizeTypes) returns()
func (_Mycontract *MycontractTransactor) ServeBoxset(opts *bind.TransactOpts, _addresss []common.Address, _amounts []*big.Int, _prizeTypes []uint8) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "serveBoxset", _addresss, _amounts, _prizeTypes)
}

// ServeBoxset is a paid mutator transaction binding the contract method 0x004a94bd.
//
// Solidity: function serveBoxset(address[] _addresss, uint256[] _amounts, uint8[] _prizeTypes) returns()
func (_Mycontract *MycontractSession) ServeBoxset(_addresss []common.Address, _amounts []*big.Int, _prizeTypes []uint8) (*types.Transaction, error) {
	return _Mycontract.Contract.ServeBoxset(&_Mycontract.TransactOpts, _addresss, _amounts, _prizeTypes)
}

// ServeBoxset is a paid mutator transaction binding the contract method 0x004a94bd.
//
// Solidity: function serveBoxset(address[] _addresss, uint256[] _amounts, uint8[] _prizeTypes) returns()
func (_Mycontract *MycontractTransactorSession) ServeBoxset(_addresss []common.Address, _amounts []*big.Int, _prizeTypes []uint8) (*types.Transaction, error) {
	return _Mycontract.Contract.ServeBoxset(&_Mycontract.TransactOpts, _addresss, _amounts, _prizeTypes)
}

// SetAddress is a paid mutator transaction binding the contract method 0x12c59488.
//
// Solidity: function setAddress(address _technology, address _ecosystem, address _nodereward, address _boxreward, address _nodereceive) returns()
func (_Mycontract *MycontractTransactor) SetAddress(opts *bind.TransactOpts, _technology common.Address, _ecosystem common.Address, _nodereward common.Address, _boxreward common.Address, _nodereceive common.Address) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setAddress", _technology, _ecosystem, _nodereward, _boxreward, _nodereceive)
}

// SetAddress is a paid mutator transaction binding the contract method 0x12c59488.
//
// Solidity: function setAddress(address _technology, address _ecosystem, address _nodereward, address _boxreward, address _nodereceive) returns()
func (_Mycontract *MycontractSession) SetAddress(_technology common.Address, _ecosystem common.Address, _nodereward common.Address, _boxreward common.Address, _nodereceive common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.SetAddress(&_Mycontract.TransactOpts, _technology, _ecosystem, _nodereward, _boxreward, _nodereceive)
}

// SetAddress is a paid mutator transaction binding the contract method 0x12c59488.
//
// Solidity: function setAddress(address _technology, address _ecosystem, address _nodereward, address _boxreward, address _nodereceive) returns()
func (_Mycontract *MycontractTransactorSession) SetAddress(_technology common.Address, _ecosystem common.Address, _nodereward common.Address, _boxreward common.Address, _nodereceive common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.SetAddress(&_Mycontract.TransactOpts, _technology, _ecosystem, _nodereward, _boxreward, _nodereceive)
}

// SetRate is a paid mutator transaction binding the contract method 0xfaf4b24b.
//
// Solidity: function setRate(uint256 _ratetechnology, uint256 _rateecosystem, uint256 _ratenodereward, uint256 _rateboxreward, uint256 _rateupdate) returns()
func (_Mycontract *MycontractTransactor) SetRate(opts *bind.TransactOpts, _ratetechnology *big.Int, _rateecosystem *big.Int, _ratenodereward *big.Int, _rateboxreward *big.Int, _rateupdate *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setRate", _ratetechnology, _rateecosystem, _ratenodereward, _rateboxreward, _rateupdate)
}

// SetRate is a paid mutator transaction binding the contract method 0xfaf4b24b.
//
// Solidity: function setRate(uint256 _ratetechnology, uint256 _rateecosystem, uint256 _ratenodereward, uint256 _rateboxreward, uint256 _rateupdate) returns()
func (_Mycontract *MycontractSession) SetRate(_ratetechnology *big.Int, _rateecosystem *big.Int, _ratenodereward *big.Int, _rateboxreward *big.Int, _rateupdate *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetRate(&_Mycontract.TransactOpts, _ratetechnology, _rateecosystem, _ratenodereward, _rateboxreward, _rateupdate)
}

// SetRate is a paid mutator transaction binding the contract method 0xfaf4b24b.
//
// Solidity: function setRate(uint256 _ratetechnology, uint256 _rateecosystem, uint256 _ratenodereward, uint256 _rateboxreward, uint256 _rateupdate) returns()
func (_Mycontract *MycontractTransactorSession) SetRate(_ratetechnology *big.Int, _rateecosystem *big.Int, _ratenodereward *big.Int, _rateboxreward *big.Int, _rateupdate *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetRate(&_Mycontract.TransactOpts, _ratetechnology, _rateecosystem, _ratenodereward, _rateboxreward, _rateupdate)
}

// SetRigisternum is a paid mutator transaction binding the contract method 0x15162570.
//
// Solidity: function setRigisternum(uint256 _amount, uint256 _nodeamount) returns()
func (_Mycontract *MycontractTransactor) SetRigisternum(opts *bind.TransactOpts, _amount *big.Int, _nodeamount *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setRigisternum", _amount, _nodeamount)
}

// SetRigisternum is a paid mutator transaction binding the contract method 0x15162570.
//
// Solidity: function setRigisternum(uint256 _amount, uint256 _nodeamount) returns()
func (_Mycontract *MycontractSession) SetRigisternum(_amount *big.Int, _nodeamount *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetRigisternum(&_Mycontract.TransactOpts, _amount, _nodeamount)
}

// SetRigisternum is a paid mutator transaction binding the contract method 0x15162570.
//
// Solidity: function setRigisternum(uint256 _amount, uint256 _nodeamount) returns()
func (_Mycontract *MycontractTransactorSession) SetRigisternum(_amount *big.Int, _nodeamount *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetRigisternum(&_Mycontract.TransactOpts, _amount, _nodeamount)
}

// SetWithAddress is a paid mutator transaction binding the contract method 0xeeafd8b3.
//
// Solidity: function setWithAddress(address _to) returns()
func (_Mycontract *MycontractTransactor) SetWithAddress(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setWithAddress", _to)
}

// SetWithAddress is a paid mutator transaction binding the contract method 0xeeafd8b3.
//
// Solidity: function setWithAddress(address _to) returns()
func (_Mycontract *MycontractSession) SetWithAddress(_to common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.SetWithAddress(&_Mycontract.TransactOpts, _to)
}

// SetWithAddress is a paid mutator transaction binding the contract method 0xeeafd8b3.
//
// Solidity: function setWithAddress(address _to) returns()
func (_Mycontract *MycontractTransactorSession) SetWithAddress(_to common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.SetWithAddress(&_Mycontract.TransactOpts, _to)
}

// SeverFromtowith is a paid mutator transaction binding the contract method 0x349f4323.
//
// Solidity: function severFromtowith(address _from, address _to, uint256 _amount) returns()
func (_Mycontract *MycontractTransactor) SeverFromtowith(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "severFromtowith", _from, _to, _amount)
}

// SeverFromtowith is a paid mutator transaction binding the contract method 0x349f4323.
//
// Solidity: function severFromtowith(address _from, address _to, uint256 _amount) returns()
func (_Mycontract *MycontractSession) SeverFromtowith(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SeverFromtowith(&_Mycontract.TransactOpts, _from, _to, _amount)
}

// SeverFromtowith is a paid mutator transaction binding the contract method 0x349f4323.
//
// Solidity: function severFromtowith(address _from, address _to, uint256 _amount) returns()
func (_Mycontract *MycontractTransactorSession) SeverFromtowith(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SeverFromtowith(&_Mycontract.TransactOpts, _from, _to, _amount)
}

// SeverRigister is a paid mutator transaction binding the contract method 0x2b9e8299.
//
// Solidity: function severRigister(address _address, address _invite) returns(bool)
func (_Mycontract *MycontractTransactor) SeverRigister(opts *bind.TransactOpts, _address common.Address, _invite common.Address) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "severRigister", _address, _invite)
}

// SeverRigister is a paid mutator transaction binding the contract method 0x2b9e8299.
//
// Solidity: function severRigister(address _address, address _invite) returns(bool)
func (_Mycontract *MycontractSession) SeverRigister(_address common.Address, _invite common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.SeverRigister(&_Mycontract.TransactOpts, _address, _invite)
}

// SeverRigister is a paid mutator transaction binding the contract method 0x2b9e8299.
//
// Solidity: function severRigister(address _address, address _invite) returns(bool)
func (_Mycontract *MycontractTransactorSession) SeverRigister(_address common.Address, _invite common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.SeverRigister(&_Mycontract.TransactOpts, _address, _invite)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mycontract *MycontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mycontract *MycontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferOwnership(&_Mycontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mycontract *MycontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferOwnership(&_Mycontract.TransactOpts, newOwner)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x8adaa8f6.
//
// Solidity: function updateNode() returns(bool)
func (_Mycontract *MycontractTransactor) UpdateNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "updateNode")
}

// UpdateNode is a paid mutator transaction binding the contract method 0x8adaa8f6.
//
// Solidity: function updateNode() returns(bool)
func (_Mycontract *MycontractSession) UpdateNode() (*types.Transaction, error) {
	return _Mycontract.Contract.UpdateNode(&_Mycontract.TransactOpts)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x8adaa8f6.
//
// Solidity: function updateNode() returns(bool)
func (_Mycontract *MycontractTransactorSession) UpdateNode() (*types.Transaction, error) {
	return _Mycontract.Contract.UpdateNode(&_Mycontract.TransactOpts)
}

// UpdateUserInfo is a paid mutator transaction binding the contract method 0xf140fe6b.
//
// Solidity: function updateUserInfo(address _oldAddr, address _newAddr, address _invite, uint8 _vip, uint8 _node, uint256 _nodereward, uint256 _boxreward) returns()
func (_Mycontract *MycontractTransactor) UpdateUserInfo(opts *bind.TransactOpts, _oldAddr common.Address, _newAddr common.Address, _invite common.Address, _vip uint8, _node uint8, _nodereward *big.Int, _boxreward *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "updateUserInfo", _oldAddr, _newAddr, _invite, _vip, _node, _nodereward, _boxreward)
}

// UpdateUserInfo is a paid mutator transaction binding the contract method 0xf140fe6b.
//
// Solidity: function updateUserInfo(address _oldAddr, address _newAddr, address _invite, uint8 _vip, uint8 _node, uint256 _nodereward, uint256 _boxreward) returns()
func (_Mycontract *MycontractSession) UpdateUserInfo(_oldAddr common.Address, _newAddr common.Address, _invite common.Address, _vip uint8, _node uint8, _nodereward *big.Int, _boxreward *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.UpdateUserInfo(&_Mycontract.TransactOpts, _oldAddr, _newAddr, _invite, _vip, _node, _nodereward, _boxreward)
}

// UpdateUserInfo is a paid mutator transaction binding the contract method 0xf140fe6b.
//
// Solidity: function updateUserInfo(address _oldAddr, address _newAddr, address _invite, uint8 _vip, uint8 _node, uint256 _nodereward, uint256 _boxreward) returns()
func (_Mycontract *MycontractTransactorSession) UpdateUserInfo(_oldAddr common.Address, _newAddr common.Address, _invite common.Address, _vip uint8, _node uint8, _nodereward *big.Int, _boxreward *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.UpdateUserInfo(&_Mycontract.TransactOpts, _oldAddr, _newAddr, _invite, _vip, _node, _nodereward, _boxreward)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _words) returns()
func (_Mycontract *MycontractTransactor) WithdrawAll(opts *bind.TransactOpts, _words *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "withdrawAll", _words)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _words) returns()
func (_Mycontract *MycontractSession) WithdrawAll(_words *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawAll(&_Mycontract.TransactOpts, _words)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _words) returns()
func (_Mycontract *MycontractTransactorSession) WithdrawAll(_words *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawAll(&_Mycontract.TransactOpts, _words)
}

// WithdrawBnb is a paid mutator transaction binding the contract method 0xca109946.
//
// Solidity: function withdrawBnb(address _to) returns()
func (_Mycontract *MycontractTransactor) WithdrawBnb(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "withdrawBnb", _to)
}

// WithdrawBnb is a paid mutator transaction binding the contract method 0xca109946.
//
// Solidity: function withdrawBnb(address _to) returns()
func (_Mycontract *MycontractSession) WithdrawBnb(_to common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawBnb(&_Mycontract.TransactOpts, _to)
}

// WithdrawBnb is a paid mutator transaction binding the contract method 0xca109946.
//
// Solidity: function withdrawBnb(address _to) returns()
func (_Mycontract *MycontractTransactorSession) WithdrawBnb(_to common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawBnb(&_Mycontract.TransactOpts, _to)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x3a0967cd.
//
// Solidity: function withdrawFrom(address _from, uint256 _amount, uint256 _words) returns()
func (_Mycontract *MycontractTransactor) WithdrawFrom(opts *bind.TransactOpts, _from common.Address, _amount *big.Int, _words *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "withdrawFrom", _from, _amount, _words)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x3a0967cd.
//
// Solidity: function withdrawFrom(address _from, uint256 _amount, uint256 _words) returns()
func (_Mycontract *MycontractSession) WithdrawFrom(_from common.Address, _amount *big.Int, _words *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawFrom(&_Mycontract.TransactOpts, _from, _amount, _words)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x3a0967cd.
//
// Solidity: function withdrawFrom(address _from, uint256 _amount, uint256 _words) returns()
func (_Mycontract *MycontractTransactorSession) WithdrawFrom(_from common.Address, _amount *big.Int, _words *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawFrom(&_Mycontract.TransactOpts, _from, _amount, _words)
}

// WithdrawUsdt is a paid mutator transaction binding the contract method 0x6865b8e7.
//
// Solidity: function withdrawUsdt() returns()
func (_Mycontract *MycontractTransactor) WithdrawUsdt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "withdrawUsdt")
}

// WithdrawUsdt is a paid mutator transaction binding the contract method 0x6865b8e7.
//
// Solidity: function withdrawUsdt() returns()
func (_Mycontract *MycontractSession) WithdrawUsdt() (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawUsdt(&_Mycontract.TransactOpts)
}

// WithdrawUsdt is a paid mutator transaction binding the contract method 0x6865b8e7.
//
// Solidity: function withdrawUsdt() returns()
func (_Mycontract *MycontractTransactorSession) WithdrawUsdt() (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawUsdt(&_Mycontract.TransactOpts)
}

// MycontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mycontract contract.
type MycontractOwnershipTransferredIterator struct {
	Event *MycontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MycontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractOwnershipTransferred)
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
		it.Event = new(MycontractOwnershipTransferred)
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
func (it *MycontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractOwnershipTransferred represents a OwnershipTransferred event raised by the Mycontract contract.
type MycontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mycontract *MycontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MycontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MycontractOwnershipTransferredIterator{contract: _Mycontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mycontract *MycontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MycontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractOwnershipTransferred)
				if err := _Mycontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mycontract *MycontractFilterer) ParseOwnershipTransferred(log types.Log) (*MycontractOwnershipTransferred, error) {
	event := new(MycontractOwnershipTransferred)
	if err := _Mycontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractTransferAndActionEventIterator is returned from FilterTransferAndActionEvent and is used to iterate over the raw logs and unpacked data for TransferAndActionEvent events raised by the Mycontract contract.
type MycontractTransferAndActionEventIterator struct {
	Event *MycontractTransferAndActionEvent // Event containing the contract specifics and raw log

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
func (it *MycontractTransferAndActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractTransferAndActionEvent)
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
		it.Event = new(MycontractTransferAndActionEvent)
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
func (it *MycontractTransferAndActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractTransferAndActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractTransferAndActionEvent represents a TransferAndActionEvent event raised by the Mycontract contract.
type MycontractTransferAndActionEvent struct {
	From       common.Address
	To         common.Address
	Amount     *big.Int
	ActionType uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferAndActionEvent is a free log retrieval operation binding the contract event 0xf6ccaf88266dc17759b964283c91904cb88ef919b3efe4a3e590f83a545e8f94.
//
// Solidity: event TransferAndActionEvent(address indexed from, address indexed to, uint256 amount, uint8 actionType)
func (_Mycontract *MycontractFilterer) FilterTransferAndActionEvent(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MycontractTransferAndActionEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "TransferAndActionEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MycontractTransferAndActionEventIterator{contract: _Mycontract.contract, event: "TransferAndActionEvent", logs: logs, sub: sub}, nil
}

// WatchTransferAndActionEvent is a free log subscription operation binding the contract event 0xf6ccaf88266dc17759b964283c91904cb88ef919b3efe4a3e590f83a545e8f94.
//
// Solidity: event TransferAndActionEvent(address indexed from, address indexed to, uint256 amount, uint8 actionType)
func (_Mycontract *MycontractFilterer) WatchTransferAndActionEvent(opts *bind.WatchOpts, sink chan<- *MycontractTransferAndActionEvent, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "TransferAndActionEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractTransferAndActionEvent)
				if err := _Mycontract.contract.UnpackLog(event, "TransferAndActionEvent", log); err != nil {
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

// ParseTransferAndActionEvent is a log parse operation binding the contract event 0xf6ccaf88266dc17759b964283c91904cb88ef919b3efe4a3e590f83a545e8f94.
//
// Solidity: event TransferAndActionEvent(address indexed from, address indexed to, uint256 amount, uint8 actionType)
func (_Mycontract *MycontractFilterer) ParseTransferAndActionEvent(log types.Log) (*MycontractTransferAndActionEvent, error) {
	event := new(MycontractTransferAndActionEvent)
	if err := _Mycontract.contract.UnpackLog(event, "TransferAndActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractWithdrawnAllIterator is returned from FilterWithdrawnAll and is used to iterate over the raw logs and unpacked data for WithdrawnAll events raised by the Mycontract contract.
type MycontractWithdrawnAllIterator struct {
	Event *MycontractWithdrawnAll // Event containing the contract specifics and raw log

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
func (it *MycontractWithdrawnAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractWithdrawnAll)
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
		it.Event = new(MycontractWithdrawnAll)
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
func (it *MycontractWithdrawnAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractWithdrawnAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractWithdrawnAll represents a WithdrawnAll event raised by the Mycontract contract.
type MycontractWithdrawnAll struct {
	To          common.Address
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAll is a free log retrieval operation binding the contract event 0x9acea1e63773c1b36db05fe8c8a86c1c1324ddd077d7a891e981b51d70491906.
//
// Solidity: event WithdrawnAll(address indexed to, uint256 totalAmount)
func (_Mycontract *MycontractFilterer) FilterWithdrawnAll(opts *bind.FilterOpts, to []common.Address) (*MycontractWithdrawnAllIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "WithdrawnAll", toRule)
	if err != nil {
		return nil, err
	}
	return &MycontractWithdrawnAllIterator{contract: _Mycontract.contract, event: "WithdrawnAll", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAll is a free log subscription operation binding the contract event 0x9acea1e63773c1b36db05fe8c8a86c1c1324ddd077d7a891e981b51d70491906.
//
// Solidity: event WithdrawnAll(address indexed to, uint256 totalAmount)
func (_Mycontract *MycontractFilterer) WatchWithdrawnAll(opts *bind.WatchOpts, sink chan<- *MycontractWithdrawnAll, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "WithdrawnAll", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractWithdrawnAll)
				if err := _Mycontract.contract.UnpackLog(event, "WithdrawnAll", log); err != nil {
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

// ParseWithdrawnAll is a log parse operation binding the contract event 0x9acea1e63773c1b36db05fe8c8a86c1c1324ddd077d7a891e981b51d70491906.
//
// Solidity: event WithdrawnAll(address indexed to, uint256 totalAmount)
func (_Mycontract *MycontractFilterer) ParseWithdrawnAll(log types.Log) (*MycontractWithdrawnAll, error) {
	event := new(MycontractWithdrawnAll)
	if err := _Mycontract.contract.UnpackLog(event, "WithdrawnAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
