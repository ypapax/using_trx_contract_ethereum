// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// TronTokenABI is the input ABI used to generate the binding from.
const TronTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addressFounder\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TronTokenFuncSigs maps the 4-byte function signature to its string representation.
var TronTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"c47f0027": "setName(string)",
	"be9a6555": "start()",
	"07da68f5": "stop()",
	"75f12b21": "stopped()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// TronTokenBin is the compiled bytecode used for deploying new contracts.
var TronTokenBin = "0x60c0604052600660808190527f54726f6e6978000000000000000000000000000000000000000000000000000060a090815261003e916000919061013c565b506040805180820190915260038082527f545258000000000000000000000000000000000000000000000000000000000060209092019182526100839160019161013c565b506006600281905560006005558054600160a860020a03191690553480156100aa57600080fd5b50604051602080610abc83398101604081815291516006805461010060a860020a031916336101000217905567016345785d8a00006005819055600160a060020a03821660008181526003602090815286822084905592855294519294909390927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506101d7565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061017d57805160ff19168380011785556101aa565b828001600101855582156101aa579182015b828111156101aa57825182559160200191906001019061018f565b506101b69291506101ba565b5090565b6101d491905b808211156101b657600081556001016101c0565b90565b6108d6806101e66000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100d457806307da68f51461015e578063095ea7b31461017557806318160ddd146101ad57806323b872dd146101d4578063313ce567146101fe57806342966c681461021357806370a082311461022b57806375f12b211461024c57806395d89b4114610261578063a9059cbb14610276578063be9a65551461029a578063c47f0027146102af578063dd62ed3e14610308575b600080fd5b3480156100e057600080fd5b506100e961032f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561012357818101518382015260200161010b565b50505050905090810190601f1680156101505780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016a57600080fd5b506101736103bd565b005b34801561018157600080fd5b50610199600160a060020a03600435166024356103e5565b604080519115158252519081900360200190f35b3480156101b957600080fd5b506101c261049e565b60408051918252519081900360200190f35b3480156101e057600080fd5b50610199600160a060020a03600435811690602435166044356104a4565b34801561020a57600080fd5b506101c26105c1565b34801561021f57600080fd5b506101736004356105c7565b34801561023757600080fd5b506101c2600160a060020a036004351661065e565b34801561025857600080fd5b50610199610670565b34801561026d57600080fd5b506100e9610679565b34801561028257600080fd5b50610199600160a060020a03600435166024356106d3565b3480156102a657600080fd5b5061017361079d565b3480156102bb57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101739436949293602493928401919081908401838280828437509497506107c29650505050505050565b34801561031457600080fd5b506101c2600160a060020a03600435811690602435166107f2565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103b55780601f1061038a576101008083540402835291602001916103b5565b820191906000526020600020905b81548152906001019060200180831161039857829003601f168201915b505050505081565b6006546101009004600160a060020a031633146103d657fe5b6006805460ff19166001179055565b60065460009060ff16156103f557fe5b3315156103fe57fe5b81158061042c5750336000908152600460209081526040808320600160a060020a0387168452909152902054155b151561043757600080fd5b336000818152600460209081526040808320600160a060020a03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60055481565b60065460009060ff16156104b457fe5b3315156104bd57fe5b600160a060020a0384166000908152600360205260409020548211156104e257600080fd5b600160a060020a038316600090815260036020526040902054828101101561050957600080fd5b600160a060020a038416600090815260046020908152604080832033845290915290205482111561053957600080fd5b600160a060020a03808416600081815260036020908152604080832080548801905593881680835284832080548890039055600482528483203384528252918490208054879003905583518681529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060019392505050565b60025481565b336000908152600360205260409020548111156105e357600080fd5b336000818152600360209081526040808320805486900390558280527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff805486019055805185815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350565b60036020526000908152604090205481565b60065460ff1681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103b55780601f1061038a576101008083540402835291602001916103b5565b60065460009060ff16156106e357fe5b3315156106ec57fe5b3360009081526003602052604090205482111561070857600080fd5b600160a060020a038316600090815260036020526040902054828101101561072f57600080fd5b33600081815260036020908152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350600192915050565b6006546101009004600160a060020a031633146107b657fe5b6006805460ff19169055565b6006546101009004600160a060020a031633146107db57fe5b80516107ee90600090602084019061080f565b5050565b600460209081526000928352604080842090915290825290205481565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061085057805160ff191683800117855561087d565b8280016001018555821561087d579182015b8281111561087d578251825591602001919060010190610862565b5061088992915061088d565b5090565b6108a791905b808211156108895760008155600101610893565b905600a165627a7a723058201d6c2fbf20d9e14fb11ae69e0c8447a171fa17c901dbd2a4ef263a96d8c7f1110029"

// DeployTronToken deploys a new Ethereum contract, binding an instance of TronToken to it.
func DeployTronToken(auth *bind.TransactOpts, backend bind.ContractBackend, _addressFounder common.Address) (common.Address, *types.Transaction, *TronToken, error) {
	parsed, err := abi.JSON(strings.NewReader(TronTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TronTokenBin), backend, _addressFounder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TronToken{TronTokenCaller: TronTokenCaller{contract: contract}, TronTokenTransactor: TronTokenTransactor{contract: contract}, TronTokenFilterer: TronTokenFilterer{contract: contract}}, nil
}

// TronToken is an auto generated Go binding around an Ethereum contract.
type TronToken struct {
	TronTokenCaller     // Read-only binding to the contract
	TronTokenTransactor // Write-only binding to the contract
	TronTokenFilterer   // Log filterer for contract events
}

// TronTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TronTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TronTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TronTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TronTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TronTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TronTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TronTokenSession struct {
	Contract     *TronToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TronTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TronTokenCallerSession struct {
	Contract *TronTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TronTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TronTokenTransactorSession struct {
	Contract     *TronTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TronTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TronTokenRaw struct {
	Contract *TronToken // Generic contract binding to access the raw methods on
}

// TronTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TronTokenCallerRaw struct {
	Contract *TronTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TronTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TronTokenTransactorRaw struct {
	Contract *TronTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTronToken creates a new instance of TronToken, bound to a specific deployed contract.
func NewTronToken(address common.Address, backend bind.ContractBackend) (*TronToken, error) {
	contract, err := bindTronToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TronToken{TronTokenCaller: TronTokenCaller{contract: contract}, TronTokenTransactor: TronTokenTransactor{contract: contract}, TronTokenFilterer: TronTokenFilterer{contract: contract}}, nil
}

// NewTronTokenCaller creates a new read-only instance of TronToken, bound to a specific deployed contract.
func NewTronTokenCaller(address common.Address, caller bind.ContractCaller) (*TronTokenCaller, error) {
	contract, err := bindTronToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TronTokenCaller{contract: contract}, nil
}

// NewTronTokenTransactor creates a new write-only instance of TronToken, bound to a specific deployed contract.
func NewTronTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TronTokenTransactor, error) {
	contract, err := bindTronToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TronTokenTransactor{contract: contract}, nil
}

// NewTronTokenFilterer creates a new log filterer instance of TronToken, bound to a specific deployed contract.
func NewTronTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TronTokenFilterer, error) {
	contract, err := bindTronToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TronTokenFilterer{contract: contract}, nil
}

// bindTronToken binds a generic wrapper to an already deployed contract.
func bindTronToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TronTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TronToken *TronTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TronToken.Contract.TronTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TronToken *TronTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TronToken.Contract.TronTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TronToken *TronTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TronToken.Contract.TronTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TronToken *TronTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TronToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TronToken *TronTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TronToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TronToken *TronTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TronToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_TronToken *TronTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_TronToken *TronTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TronToken.Contract.Allowance(&_TronToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_TronToken *TronTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TronToken.Contract.Allowance(&_TronToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_TronToken *TronTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_TronToken *TronTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TronToken.Contract.BalanceOf(&_TronToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_TronToken *TronTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TronToken.Contract.BalanceOf(&_TronToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_TronToken *TronTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_TronToken *TronTokenSession) Decimals() (*big.Int, error) {
	return _TronToken.Contract.Decimals(&_TronToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_TronToken *TronTokenCallerSession) Decimals() (*big.Int, error) {
	return _TronToken.Contract.Decimals(&_TronToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TronToken *TronTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TronToken *TronTokenSession) Name() (string, error) {
	return _TronToken.Contract.Name(&_TronToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TronToken *TronTokenCallerSession) Name() (string, error) {
	return _TronToken.Contract.Name(&_TronToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_TronToken *TronTokenCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_TronToken *TronTokenSession) Stopped() (bool, error) {
	return _TronToken.Contract.Stopped(&_TronToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_TronToken *TronTokenCallerSession) Stopped() (bool, error) {
	return _TronToken.Contract.Stopped(&_TronToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TronToken *TronTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TronToken *TronTokenSession) Symbol() (string, error) {
	return _TronToken.Contract.Symbol(&_TronToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TronToken *TronTokenCallerSession) Symbol() (string, error) {
	return _TronToken.Contract.Symbol(&_TronToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TronToken *TronTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TronToken *TronTokenSession) TotalSupply() (*big.Int, error) {
	return _TronToken.Contract.TotalSupply(&_TronToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TronToken *TronTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TronToken.Contract.TotalSupply(&_TronToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_TronToken *TronTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_TronToken *TronTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Approve(&_TronToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_TronToken *TronTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Approve(&_TronToken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_TronToken *TronTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_TronToken *TronTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Burn(&_TronToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_TronToken *TronTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Burn(&_TronToken.TransactOpts, _value)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_TronToken *TronTokenTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_TronToken *TronTokenSession) SetName(_name string) (*types.Transaction, error) {
	return _TronToken.Contract.SetName(&_TronToken.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_TronToken *TronTokenTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _TronToken.Contract.SetName(&_TronToken.TransactOpts, _name)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_TronToken *TronTokenTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_TronToken *TronTokenSession) Start() (*types.Transaction, error) {
	return _TronToken.Contract.Start(&_TronToken.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_TronToken *TronTokenTransactorSession) Start() (*types.Transaction, error) {
	return _TronToken.Contract.Start(&_TronToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_TronToken *TronTokenTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_TronToken *TronTokenSession) Stop() (*types.Transaction, error) {
	return _TronToken.Contract.Stop(&_TronToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_TronToken *TronTokenTransactorSession) Stop() (*types.Transaction, error) {
	return _TronToken.Contract.Stop(&_TronToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_TronToken *TronTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_TronToken *TronTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Transfer(&_TronToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_TronToken *TronTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Transfer(&_TronToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_TronToken *TronTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_TronToken *TronTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.TransferFrom(&_TronToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_TronToken *TronTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.TransferFrom(&_TronToken.TransactOpts, _from, _to, _value)
}

// TronTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TronToken contract.
type TronTokenApprovalIterator struct {
	Event *TronTokenApproval // Event containing the contract specifics and raw log

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
func (it *TronTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TronTokenApproval)
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
		it.Event = new(TronTokenApproval)
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
func (it *TronTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TronTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TronTokenApproval represents a Approval event raised by the TronToken contract.
type TronTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TronToken *TronTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TronTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TronToken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TronTokenApprovalIterator{contract: _TronToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TronToken *TronTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TronTokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TronToken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TronTokenApproval)
				if err := _TronToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TronToken *TronTokenFilterer) ParseApproval(log types.Log) (*TronTokenApproval, error) {
	event := new(TronTokenApproval)
	if err := _TronToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TronTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TronToken contract.
type TronTokenTransferIterator struct {
	Event *TronTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TronTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TronTokenTransfer)
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
		it.Event = new(TronTokenTransfer)
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
func (it *TronTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TronTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TronTokenTransfer represents a Transfer event raised by the TronToken contract.
type TronTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TronToken *TronTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TronTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TronToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TronTokenTransferIterator{contract: _TronToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TronToken *TronTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TronTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TronToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TronTokenTransfer)
				if err := _TronToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TronToken *TronTokenFilterer) ParseTransfer(log types.Log) (*TronTokenTransfer, error) {
	event := new(TronTokenTransfer)
	if err := _TronToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
