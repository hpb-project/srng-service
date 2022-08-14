// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// Commit is an auto generated low-level Go binding around an user-defined struct.
type Commit struct {
	Author        common.Address
	Commit        [32]byte
	Block         *big.Int
	Seed          [32]byte
	Revealed      bool
	VerifiedBlock *big.Int
	Consumer      common.Address
	Subsender     common.Address
	SubBlock      *big.Int
	Substatus     uint8
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"CommitHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"commiter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"RandomConsumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"commiter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"RevealSeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"commiter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Subscribe\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"commiter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"UnSubscribe\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"commiter\",\"type\":\"address\"}],\"name\":\"getCommiterValidCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"getConsumerConsumedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"getRandom\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"commiter\",\"type\":\"address\"}],\"name\":\"getUserCommitsList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"commit\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"revealed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"verifiedBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subsender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"substatus\",\"type\":\"uint8\"}],\"internalType\":\"structCommit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"getUserSubscribed\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"commit\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"revealed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"verifiedBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subsender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"substatus\",\"type\":\"uint8\"}],\"internalType\":\"structCommit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"commiter\",\"type\":\"address\"}],\"name\":\"getUserUnverifiedList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"commit\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"revealed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"verifiedBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subsender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"substatus\",\"type\":\"uint8\"}],\"internalType\":\"structCommit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"}],\"name\":\"requestRandom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_deposit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_store\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_commitReveal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stat\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_internalstore\",\"type\":\"address\"}],\"name\":\"setting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"unsubscribeRandom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// GetCommiterValidCount is a free data retrieval call binding the contract method 0x2b68d9c1.
//
// Solidity: function getCommiterValidCount(address commiter) view returns(uint256)
func (_Oracle *OracleCaller) GetCommiterValidCount(opts *bind.CallOpts, commiter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getCommiterValidCount", commiter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommiterValidCount is a free data retrieval call binding the contract method 0x2b68d9c1.
//
// Solidity: function getCommiterValidCount(address commiter) view returns(uint256)
func (_Oracle *OracleSession) GetCommiterValidCount(commiter common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetCommiterValidCount(&_Oracle.CallOpts, commiter)
}

// GetCommiterValidCount is a free data retrieval call binding the contract method 0x2b68d9c1.
//
// Solidity: function getCommiterValidCount(address commiter) view returns(uint256)
func (_Oracle *OracleCallerSession) GetCommiterValidCount(commiter common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetCommiterValidCount(&_Oracle.CallOpts, commiter)
}

// GetConsumerConsumedCount is a free data retrieval call binding the contract method 0x7d7171b6.
//
// Solidity: function getConsumerConsumedCount(address consumer) view returns(uint256)
func (_Oracle *OracleCaller) GetConsumerConsumedCount(opts *bind.CallOpts, consumer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getConsumerConsumedCount", consumer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsumerConsumedCount is a free data retrieval call binding the contract method 0x7d7171b6.
//
// Solidity: function getConsumerConsumedCount(address consumer) view returns(uint256)
func (_Oracle *OracleSession) GetConsumerConsumedCount(consumer common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetConsumerConsumedCount(&_Oracle.CallOpts, consumer)
}

// GetConsumerConsumedCount is a free data retrieval call binding the contract method 0x7d7171b6.
//
// Solidity: function getConsumerConsumedCount(address consumer) view returns(uint256)
func (_Oracle *OracleCallerSession) GetConsumerConsumedCount(consumer common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetConsumerConsumedCount(&_Oracle.CallOpts, consumer)
}

// GetHash is a free data retrieval call binding the contract method 0x3cf5040a.
//
// Solidity: function getHash(bytes32 seed) view returns(bytes32)
func (_Oracle *OracleCaller) GetHash(opts *bind.CallOpts, seed [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getHash", seed)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x3cf5040a.
//
// Solidity: function getHash(bytes32 seed) view returns(bytes32)
func (_Oracle *OracleSession) GetHash(seed [32]byte) ([32]byte, error) {
	return _Oracle.Contract.GetHash(&_Oracle.CallOpts, seed)
}

// GetHash is a free data retrieval call binding the contract method 0x3cf5040a.
//
// Solidity: function getHash(bytes32 seed) view returns(bytes32)
func (_Oracle *OracleCallerSession) GetHash(seed [32]byte) ([32]byte, error) {
	return _Oracle.Contract.GetHash(&_Oracle.CallOpts, seed)
}

// GetRandom is a free data retrieval call binding the contract method 0x5b6b90d2.
//
// Solidity: function getRandom(bytes32 commit, bytes signature) view returns(bytes32)
func (_Oracle *OracleCaller) GetRandom(opts *bind.CallOpts, commit [32]byte, signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRandom", commit, signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRandom is a free data retrieval call binding the contract method 0x5b6b90d2.
//
// Solidity: function getRandom(bytes32 commit, bytes signature) view returns(bytes32)
func (_Oracle *OracleSession) GetRandom(commit [32]byte, signature []byte) ([32]byte, error) {
	return _Oracle.Contract.GetRandom(&_Oracle.CallOpts, commit, signature)
}

// GetRandom is a free data retrieval call binding the contract method 0x5b6b90d2.
//
// Solidity: function getRandom(bytes32 commit, bytes signature) view returns(bytes32)
func (_Oracle *OracleCallerSession) GetRandom(commit [32]byte, signature []byte) ([32]byte, error) {
	return _Oracle.Contract.GetRandom(&_Oracle.CallOpts, commit, signature)
}

// GetTotalStat is a free data retrieval call binding the contract method 0x0256d5f8.
//
// Solidity: function getTotalStat() view returns(uint256, uint256, uint256)
func (_Oracle *OracleCaller) GetTotalStat(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getTotalStat")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetTotalStat is a free data retrieval call binding the contract method 0x0256d5f8.
//
// Solidity: function getTotalStat() view returns(uint256, uint256, uint256)
func (_Oracle *OracleSession) GetTotalStat() (*big.Int, *big.Int, *big.Int, error) {
	return _Oracle.Contract.GetTotalStat(&_Oracle.CallOpts)
}

// GetTotalStat is a free data retrieval call binding the contract method 0x0256d5f8.
//
// Solidity: function getTotalStat() view returns(uint256, uint256, uint256)
func (_Oracle *OracleCallerSession) GetTotalStat() (*big.Int, *big.Int, *big.Int, error) {
	return _Oracle.Contract.GetTotalStat(&_Oracle.CallOpts)
}

// GetUserCommitsList is a free data retrieval call binding the contract method 0x23344446.
//
// Solidity: function getUserCommitsList(address commiter) view returns((address,bytes32,uint256,bytes32,bool,uint256,address,address,uint256,uint8)[])
func (_Oracle *OracleCaller) GetUserCommitsList(opts *bind.CallOpts, commiter common.Address) ([]Commit, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getUserCommitsList", commiter)

	if err != nil {
		return *new([]Commit), err
	}

	out0 := *abi.ConvertType(out[0], new([]Commit)).(*[]Commit)

	return out0, err

}

// GetUserCommitsList is a free data retrieval call binding the contract method 0x23344446.
//
// Solidity: function getUserCommitsList(address commiter) view returns((address,bytes32,uint256,bytes32,bool,uint256,address,address,uint256,uint8)[])
func (_Oracle *OracleSession) GetUserCommitsList(commiter common.Address) ([]Commit, error) {
	return _Oracle.Contract.GetUserCommitsList(&_Oracle.CallOpts, commiter)
}

// GetUserCommitsList is a free data retrieval call binding the contract method 0x23344446.
//
// Solidity: function getUserCommitsList(address commiter) view returns((address,bytes32,uint256,bytes32,bool,uint256,address,address,uint256,uint8)[])
func (_Oracle *OracleCallerSession) GetUserCommitsList(commiter common.Address) ([]Commit, error) {
	return _Oracle.Contract.GetUserCommitsList(&_Oracle.CallOpts, commiter)
}

// GetUserSubscribed is a free data retrieval call binding the contract method 0xd4b612a2.
//
// Solidity: function getUserSubscribed(address consumer) view returns((address,bytes32,uint256,bytes32,bool,uint256,address,address,uint256,uint8)[])
func (_Oracle *OracleCaller) GetUserSubscribed(opts *bind.CallOpts, consumer common.Address) ([]Commit, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getUserSubscribed", consumer)

	if err != nil {
		return *new([]Commit), err
	}

	out0 := *abi.ConvertType(out[0], new([]Commit)).(*[]Commit)

	return out0, err

}

// GetUserSubscribed is a free data retrieval call binding the contract method 0xd4b612a2.
//
// Solidity: function getUserSubscribed(address consumer) view returns((address,bytes32,uint256,bytes32,bool,uint256,address,address,uint256,uint8)[])
func (_Oracle *OracleSession) GetUserSubscribed(consumer common.Address) ([]Commit, error) {
	return _Oracle.Contract.GetUserSubscribed(&_Oracle.CallOpts, consumer)
}

// GetUserSubscribed is a free data retrieval call binding the contract method 0xd4b612a2.
//
// Solidity: function getUserSubscribed(address consumer) view returns((address,bytes32,uint256,bytes32,bool,uint256,address,address,uint256,uint8)[])
func (_Oracle *OracleCallerSession) GetUserSubscribed(consumer common.Address) ([]Commit, error) {
	return _Oracle.Contract.GetUserSubscribed(&_Oracle.CallOpts, consumer)
}

// GetUserUnverifiedList is a free data retrieval call binding the contract method 0x82129d8e.
//
// Solidity: function getUserUnverifiedList(address commiter) view returns((address,bytes32,uint256,bytes32,bool,uint256,address,address,uint256,uint8)[])
func (_Oracle *OracleCaller) GetUserUnverifiedList(opts *bind.CallOpts, commiter common.Address) ([]Commit, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getUserUnverifiedList", commiter)

	if err != nil {
		return *new([]Commit), err
	}

	out0 := *abi.ConvertType(out[0], new([]Commit)).(*[]Commit)

	return out0, err

}

// GetUserUnverifiedList is a free data retrieval call binding the contract method 0x82129d8e.
//
// Solidity: function getUserUnverifiedList(address commiter) view returns((address,bytes32,uint256,bytes32,bool,uint256,address,address,uint256,uint8)[])
func (_Oracle *OracleSession) GetUserUnverifiedList(commiter common.Address) ([]Commit, error) {
	return _Oracle.Contract.GetUserUnverifiedList(&_Oracle.CallOpts, commiter)
}

// GetUserUnverifiedList is a free data retrieval call binding the contract method 0x82129d8e.
//
// Solidity: function getUserUnverifiedList(address commiter) view returns((address,bytes32,uint256,bytes32,bool,uint256,address,address,uint256,uint8)[])
func (_Oracle *OracleCallerSession) GetUserUnverifiedList(commiter common.Address) ([]Commit, error) {
	return _Oracle.Contract.GetUserUnverifiedList(&_Oracle.CallOpts, commiter)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCallerSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_Oracle *OracleTransactor) AddAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addAdmin", _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_Oracle *OracleSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddAdmin(&_Oracle.TransactOpts, _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_Oracle *OracleTransactorSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddAdmin(&_Oracle.TransactOpts, _admin)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 hash) returns()
func (_Oracle *OracleTransactor) Commit(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "commit", hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 hash) returns()
func (_Oracle *OracleSession) Commit(hash [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.Commit(&_Oracle.TransactOpts, hash)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 hash) returns()
func (_Oracle *OracleTransactorSession) Commit(hash [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.Commit(&_Oracle.TransactOpts, hash)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_Oracle *OracleTransactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_Oracle *OracleSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveAdmin(&_Oracle.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_Oracle *OracleTransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveAdmin(&_Oracle.TransactOpts, _admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceOwnership(&_Oracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oracle *OracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceOwnership(&_Oracle.TransactOpts)
}

// RequestRandom is a paid mutator transaction binding the contract method 0x951ab42e.
//
// Solidity: function requestRandom(address user, address consumer, bytes32 token) returns(bool)
func (_Oracle *OracleTransactor) RequestRandom(opts *bind.TransactOpts, user common.Address, consumer common.Address, token [32]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "requestRandom", user, consumer, token)
}

// RequestRandom is a paid mutator transaction binding the contract method 0x951ab42e.
//
// Solidity: function requestRandom(address user, address consumer, bytes32 token) returns(bool)
func (_Oracle *OracleSession) RequestRandom(user common.Address, consumer common.Address, token [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.RequestRandom(&_Oracle.TransactOpts, user, consumer, token)
}

// RequestRandom is a paid mutator transaction binding the contract method 0x951ab42e.
//
// Solidity: function requestRandom(address user, address consumer, bytes32 token) returns(bool)
func (_Oracle *OracleTransactorSession) RequestRandom(user common.Address, consumer common.Address, token [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.RequestRandom(&_Oracle.TransactOpts, user, consumer, token)
}

// Reveal is a paid mutator transaction binding the contract method 0xfc334e8c.
//
// Solidity: function reveal(bytes32 hash, bytes32 seed) returns()
func (_Oracle *OracleTransactor) Reveal(opts *bind.TransactOpts, hash [32]byte, seed [32]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "reveal", hash, seed)
}

// Reveal is a paid mutator transaction binding the contract method 0xfc334e8c.
//
// Solidity: function reveal(bytes32 hash, bytes32 seed) returns()
func (_Oracle *OracleSession) Reveal(hash [32]byte, seed [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.Reveal(&_Oracle.TransactOpts, hash, seed)
}

// Reveal is a paid mutator transaction binding the contract method 0xfc334e8c.
//
// Solidity: function reveal(bytes32 hash, bytes32 seed) returns()
func (_Oracle *OracleTransactorSession) Reveal(hash [32]byte, seed [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.Reveal(&_Oracle.TransactOpts, hash, seed)
}

// Setting is a paid mutator transaction binding the contract method 0x500cbf45.
//
// Solidity: function setting(address _token, address _config, address _deposit, address _store, address _commitReveal, address _stat, address _internalstore) returns()
func (_Oracle *OracleTransactor) Setting(opts *bind.TransactOpts, _token common.Address, _config common.Address, _deposit common.Address, _store common.Address, _commitReveal common.Address, _stat common.Address, _internalstore common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setting", _token, _config, _deposit, _store, _commitReveal, _stat, _internalstore)
}

// Setting is a paid mutator transaction binding the contract method 0x500cbf45.
//
// Solidity: function setting(address _token, address _config, address _deposit, address _store, address _commitReveal, address _stat, address _internalstore) returns()
func (_Oracle *OracleSession) Setting(_token common.Address, _config common.Address, _deposit common.Address, _store common.Address, _commitReveal common.Address, _stat common.Address, _internalstore common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Setting(&_Oracle.TransactOpts, _token, _config, _deposit, _store, _commitReveal, _stat, _internalstore)
}

// Setting is a paid mutator transaction binding the contract method 0x500cbf45.
//
// Solidity: function setting(address _token, address _config, address _deposit, address _store, address _commitReveal, address _stat, address _internalstore) returns()
func (_Oracle *OracleTransactorSession) Setting(_token common.Address, _config common.Address, _deposit common.Address, _store common.Address, _commitReveal common.Address, _stat common.Address, _internalstore common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Setting(&_Oracle.TransactOpts, _token, _config, _deposit, _store, _commitReveal, _stat, _internalstore)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oracle *OracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, newOwner)
}

// UnsubscribeRandom is a paid mutator transaction binding the contract method 0xc12a3802.
//
// Solidity: function unsubscribeRandom(address consumer, bytes32 hash) returns()
func (_Oracle *OracleTransactor) UnsubscribeRandom(opts *bind.TransactOpts, consumer common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "unsubscribeRandom", consumer, hash)
}

// UnsubscribeRandom is a paid mutator transaction binding the contract method 0xc12a3802.
//
// Solidity: function unsubscribeRandom(address consumer, bytes32 hash) returns()
func (_Oracle *OracleSession) UnsubscribeRandom(consumer common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.UnsubscribeRandom(&_Oracle.TransactOpts, consumer, hash)
}

// UnsubscribeRandom is a paid mutator transaction binding the contract method 0xc12a3802.
//
// Solidity: function unsubscribeRandom(address consumer, bytes32 hash) returns()
func (_Oracle *OracleTransactorSession) UnsubscribeRandom(consumer common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.UnsubscribeRandom(&_Oracle.TransactOpts, consumer, hash)
}

// OracleCommitHashIterator is returned from FilterCommitHash and is used to iterate over the raw logs and unpacked data for CommitHash events raised by the Oracle contract.
type OracleCommitHashIterator struct {
	Event *OracleCommitHash // Event containing the contract specifics and raw log

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
func (it *OracleCommitHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleCommitHash)
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
		it.Event = new(OracleCommitHash)
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
func (it *OracleCommitHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleCommitHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleCommitHash represents a CommitHash event raised by the Oracle contract.
type OracleCommitHash struct {
	Sender common.Address
	Hash   [32]byte
	Block  *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCommitHash is a free log retrieval operation binding the contract event 0x09b2b06005aa246accb0fa2d5147d18a48b2580ebb471e3a669968562e0737e1.
//
// Solidity: event CommitHash(address sender, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) FilterCommitHash(opts *bind.FilterOpts) (*OracleCommitHashIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "CommitHash")
	if err != nil {
		return nil, err
	}
	return &OracleCommitHashIterator{contract: _Oracle.contract, event: "CommitHash", logs: logs, sub: sub}, nil
}

// WatchCommitHash is a free log subscription operation binding the contract event 0x09b2b06005aa246accb0fa2d5147d18a48b2580ebb471e3a669968562e0737e1.
//
// Solidity: event CommitHash(address sender, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) WatchCommitHash(opts *bind.WatchOpts, sink chan<- *OracleCommitHash) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "CommitHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleCommitHash)
				if err := _Oracle.contract.UnpackLog(event, "CommitHash", log); err != nil {
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

// ParseCommitHash is a log parse operation binding the contract event 0x09b2b06005aa246accb0fa2d5147d18a48b2580ebb471e3a669968562e0737e1.
//
// Solidity: event CommitHash(address sender, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) ParseCommitHash(log types.Log) (*OracleCommitHash, error) {
	event := new(OracleCommitHash)
	if err := _Oracle.contract.UnpackLog(event, "CommitHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oracle contract.
type OracleOwnershipTransferredIterator struct {
	Event *OracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleOwnershipTransferred)
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
		it.Event = new(OracleOwnershipTransferred)
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
func (it *OracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleOwnershipTransferred represents a OwnershipTransferred event raised by the Oracle contract.
type OracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleOwnershipTransferredIterator{contract: _Oracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleOwnershipTransferred)
				if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseOwnershipTransferred(log types.Log) (*OracleOwnershipTransferred, error) {
	event := new(OracleOwnershipTransferred)
	if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRandomConsumedIterator is returned from FilterRandomConsumed and is used to iterate over the raw logs and unpacked data for RandomConsumed events raised by the Oracle contract.
type OracleRandomConsumedIterator struct {
	Event *OracleRandomConsumed // Event containing the contract specifics and raw log

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
func (it *OracleRandomConsumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRandomConsumed)
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
		it.Event = new(OracleRandomConsumed)
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
func (it *OracleRandomConsumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRandomConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRandomConsumed represents a RandomConsumed event raised by the Oracle contract.
type OracleRandomConsumed struct {
	Commiter common.Address
	Consumer common.Address
	Hash     [32]byte
	Block    *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRandomConsumed is a free log retrieval operation binding the contract event 0xa6b50d46fd68ec1507309bfc0107c38fba251b3e06ad8a36416b1fb74e529bfe.
//
// Solidity: event RandomConsumed(address commiter, address consumer, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) FilterRandomConsumed(opts *bind.FilterOpts) (*OracleRandomConsumedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "RandomConsumed")
	if err != nil {
		return nil, err
	}
	return &OracleRandomConsumedIterator{contract: _Oracle.contract, event: "RandomConsumed", logs: logs, sub: sub}, nil
}

// WatchRandomConsumed is a free log subscription operation binding the contract event 0xa6b50d46fd68ec1507309bfc0107c38fba251b3e06ad8a36416b1fb74e529bfe.
//
// Solidity: event RandomConsumed(address commiter, address consumer, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) WatchRandomConsumed(opts *bind.WatchOpts, sink chan<- *OracleRandomConsumed) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "RandomConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRandomConsumed)
				if err := _Oracle.contract.UnpackLog(event, "RandomConsumed", log); err != nil {
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

// ParseRandomConsumed is a log parse operation binding the contract event 0xa6b50d46fd68ec1507309bfc0107c38fba251b3e06ad8a36416b1fb74e529bfe.
//
// Solidity: event RandomConsumed(address commiter, address consumer, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) ParseRandomConsumed(log types.Log) (*OracleRandomConsumed, error) {
	event := new(OracleRandomConsumed)
	if err := _Oracle.contract.UnpackLog(event, "RandomConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRevealSeedIterator is returned from FilterRevealSeed and is used to iterate over the raw logs and unpacked data for RevealSeed events raised by the Oracle contract.
type OracleRevealSeedIterator struct {
	Event *OracleRevealSeed // Event containing the contract specifics and raw log

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
func (it *OracleRevealSeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRevealSeed)
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
		it.Event = new(OracleRevealSeed)
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
func (it *OracleRevealSeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRevealSeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRevealSeed represents a RevealSeed event raised by the Oracle contract.
type OracleRevealSeed struct {
	Commiter common.Address
	Hash     [32]byte
	Seed     [32]byte
	Block    *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRevealSeed is a free log retrieval operation binding the contract event 0x0f55b32212e9b88d2bf9643f51936750a283fab160176b01c340af4ec1e85572.
//
// Solidity: event RevealSeed(address commiter, bytes32 hash, bytes32 seed, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) FilterRevealSeed(opts *bind.FilterOpts) (*OracleRevealSeedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "RevealSeed")
	if err != nil {
		return nil, err
	}
	return &OracleRevealSeedIterator{contract: _Oracle.contract, event: "RevealSeed", logs: logs, sub: sub}, nil
}

// WatchRevealSeed is a free log subscription operation binding the contract event 0x0f55b32212e9b88d2bf9643f51936750a283fab160176b01c340af4ec1e85572.
//
// Solidity: event RevealSeed(address commiter, bytes32 hash, bytes32 seed, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) WatchRevealSeed(opts *bind.WatchOpts, sink chan<- *OracleRevealSeed) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "RevealSeed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRevealSeed)
				if err := _Oracle.contract.UnpackLog(event, "RevealSeed", log); err != nil {
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

// ParseRevealSeed is a log parse operation binding the contract event 0x0f55b32212e9b88d2bf9643f51936750a283fab160176b01c340af4ec1e85572.
//
// Solidity: event RevealSeed(address commiter, bytes32 hash, bytes32 seed, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) ParseRevealSeed(log types.Log) (*OracleRevealSeed, error) {
	event := new(OracleRevealSeed)
	if err := _Oracle.contract.UnpackLog(event, "RevealSeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleSubscribeIterator is returned from FilterSubscribe and is used to iterate over the raw logs and unpacked data for Subscribe events raised by the Oracle contract.
type OracleSubscribeIterator struct {
	Event *OracleSubscribe // Event containing the contract specifics and raw log

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
func (it *OracleSubscribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleSubscribe)
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
		it.Event = new(OracleSubscribe)
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
func (it *OracleSubscribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleSubscribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleSubscribe represents a Subscribe event raised by the Oracle contract.
type OracleSubscribe struct {
	Consumer common.Address
	Commiter common.Address
	Hash     [32]byte
	Block    *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubscribe is a free log retrieval operation binding the contract event 0xb4d80af06c39f4ab03973a4e14cb9b17ab4213fa0d7bd904a5c7d6bf5d5716d7.
//
// Solidity: event Subscribe(address consumer, address commiter, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) FilterSubscribe(opts *bind.FilterOpts) (*OracleSubscribeIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Subscribe")
	if err != nil {
		return nil, err
	}
	return &OracleSubscribeIterator{contract: _Oracle.contract, event: "Subscribe", logs: logs, sub: sub}, nil
}

// WatchSubscribe is a free log subscription operation binding the contract event 0xb4d80af06c39f4ab03973a4e14cb9b17ab4213fa0d7bd904a5c7d6bf5d5716d7.
//
// Solidity: event Subscribe(address consumer, address commiter, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) WatchSubscribe(opts *bind.WatchOpts, sink chan<- *OracleSubscribe) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Subscribe")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleSubscribe)
				if err := _Oracle.contract.UnpackLog(event, "Subscribe", log); err != nil {
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

// ParseSubscribe is a log parse operation binding the contract event 0xb4d80af06c39f4ab03973a4e14cb9b17ab4213fa0d7bd904a5c7d6bf5d5716d7.
//
// Solidity: event Subscribe(address consumer, address commiter, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) ParseSubscribe(log types.Log) (*OracleSubscribe, error) {
	event := new(OracleSubscribe)
	if err := _Oracle.contract.UnpackLog(event, "Subscribe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleUnSubscribeIterator is returned from FilterUnSubscribe and is used to iterate over the raw logs and unpacked data for UnSubscribe events raised by the Oracle contract.
type OracleUnSubscribeIterator struct {
	Event *OracleUnSubscribe // Event containing the contract specifics and raw log

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
func (it *OracleUnSubscribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleUnSubscribe)
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
		it.Event = new(OracleUnSubscribe)
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
func (it *OracleUnSubscribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleUnSubscribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleUnSubscribe represents a UnSubscribe event raised by the Oracle contract.
type OracleUnSubscribe struct {
	Consumer common.Address
	Commiter common.Address
	Hash     [32]byte
	Block    *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnSubscribe is a free log retrieval operation binding the contract event 0x575cd5bf167e8f74f63163ab979775fe588116fc382a6cd1cc3b8bf76e07eacc.
//
// Solidity: event UnSubscribe(address consumer, address commiter, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) FilterUnSubscribe(opts *bind.FilterOpts) (*OracleUnSubscribeIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "UnSubscribe")
	if err != nil {
		return nil, err
	}
	return &OracleUnSubscribeIterator{contract: _Oracle.contract, event: "UnSubscribe", logs: logs, sub: sub}, nil
}

// WatchUnSubscribe is a free log subscription operation binding the contract event 0x575cd5bf167e8f74f63163ab979775fe588116fc382a6cd1cc3b8bf76e07eacc.
//
// Solidity: event UnSubscribe(address consumer, address commiter, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) WatchUnSubscribe(opts *bind.WatchOpts, sink chan<- *OracleUnSubscribe) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "UnSubscribe")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleUnSubscribe)
				if err := _Oracle.contract.UnpackLog(event, "UnSubscribe", log); err != nil {
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

// ParseUnSubscribe is a log parse operation binding the contract event 0x575cd5bf167e8f74f63163ab979775fe588116fc382a6cd1cc3b8bf76e07eacc.
//
// Solidity: event UnSubscribe(address consumer, address commiter, bytes32 hash, uint256 block, uint256 time)
func (_Oracle *OracleFilterer) ParseUnSubscribe(log types.Log) (*OracleUnSubscribe, error) {
	event := new(OracleUnSubscribe)
	if err := _Oracle.contract.UnpackLog(event, "UnSubscribe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
