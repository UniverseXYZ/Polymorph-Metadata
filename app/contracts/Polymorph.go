// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PolymorphABI is the input ABI used to generate the binding from.
const PolymorphABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSlope\",\"type\":\"uint256\"}],\"name\":\"changeSlope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"geneOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenNumber\",\"type\":\"uint256\"}],\"name\":\"priceFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genePosition\",\"type\":\"uint256\"}],\"name\":\"morphGene\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"randomizeGenome\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"priceForGenomeChange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGenomeChangePrice\",\"type\":\"uint256\"}],\"name\":\"changeBaseGenomeChangePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Polymorph is an auto generated Go binding around an Ethereum contract.
type Polymorph struct {
	PolymorphCaller     // Read-only binding to the contract
	PolymorphTransactor // Write-only binding to the contract
	PolymorphFilterer   // Log filterer for contract events
}

// PolymorphCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolymorphCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolymorphTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolymorphFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolymorphSession struct {
	Contract     *Polymorph        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolymorphCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolymorphCallerSession struct {
	Contract *PolymorphCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PolymorphTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolymorphTransactorSession struct {
	Contract     *PolymorphTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PolymorphRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolymorphRaw struct {
	Contract *Polymorph // Generic contract binding to access the raw methods on
}

// PolymorphCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolymorphCallerRaw struct {
	Contract *PolymorphCaller // Generic read-only contract binding to access the raw methods on
}

// PolymorphTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolymorphTransactorRaw struct {
	Contract *PolymorphTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolymorph creates a new instance of Polymorph, bound to a specific deployed contract.
func NewPolymorph(address common.Address, backend bind.ContractBackend) (*Polymorph, error) {
	contract, err := bindPolymorph(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polymorph{PolymorphCaller: PolymorphCaller{contract: contract}, PolymorphTransactor: PolymorphTransactor{contract: contract}, PolymorphFilterer: PolymorphFilterer{contract: contract}}, nil
}

// NewPolymorphCaller creates a new read-only instance of Polymorph, bound to a specific deployed contract.
func NewPolymorphCaller(address common.Address, caller bind.ContractCaller) (*PolymorphCaller, error) {
	contract, err := bindPolymorph(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolymorphCaller{contract: contract}, nil
}

// NewPolymorphTransactor creates a new write-only instance of Polymorph, bound to a specific deployed contract.
func NewPolymorphTransactor(address common.Address, transactor bind.ContractTransactor) (*PolymorphTransactor, error) {
	contract, err := bindPolymorph(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolymorphTransactor{contract: contract}, nil
}

// NewPolymorphFilterer creates a new log filterer instance of Polymorph, bound to a specific deployed contract.
func NewPolymorphFilterer(address common.Address, filterer bind.ContractFilterer) (*PolymorphFilterer, error) {
	contract, err := bindPolymorph(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolymorphFilterer{contract: contract}, nil
}

// bindPolymorph binds a generic wrapper to an already deployed contract.
func bindPolymorph(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolymorphABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polymorph *PolymorphRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polymorph.Contract.PolymorphCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polymorph *PolymorphRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polymorph.Contract.PolymorphTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polymorph *PolymorphRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polymorph.Contract.PolymorphTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polymorph *PolymorphCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polymorph.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polymorph *PolymorphTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polymorph.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polymorph *PolymorphTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polymorph.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Polymorph *PolymorphCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Polymorph.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Polymorph *PolymorphSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Polymorph.Contract.BalanceOf(&_Polymorph.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Polymorph *PolymorphCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Polymorph.Contract.BalanceOf(&_Polymorph.CallOpts, owner)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_Polymorph *PolymorphCaller) GeneOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Polymorph.contract.Call(opts, &out, "geneOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_Polymorph *PolymorphSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _Polymorph.Contract.GeneOf(&_Polymorph.CallOpts, tokenId)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_Polymorph *PolymorphCallerSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _Polymorph.Contract.GeneOf(&_Polymorph.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_Polymorph *PolymorphCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Polymorph.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_Polymorph *PolymorphSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Polymorph.Contract.GetApproved(&_Polymorph.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_Polymorph *PolymorphCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Polymorph.Contract.GetApproved(&_Polymorph.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Polymorph *PolymorphCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Polymorph.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Polymorph *PolymorphSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Polymorph.Contract.IsApprovedForAll(&_Polymorph.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Polymorph *PolymorphCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Polymorph.Contract.IsApprovedForAll(&_Polymorph.CallOpts, owner, operator)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_Polymorph *PolymorphCaller) LastTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polymorph.contract.Call(opts, &out, "lastTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_Polymorph *PolymorphSession) LastTokenId() (*big.Int, error) {
	return _Polymorph.Contract.LastTokenId(&_Polymorph.CallOpts)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_Polymorph *PolymorphCallerSession) LastTokenId() (*big.Int, error) {
	return _Polymorph.Contract.LastTokenId(&_Polymorph.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Polymorph *PolymorphCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Polymorph.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Polymorph *PolymorphSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Polymorph.Contract.OwnerOf(&_Polymorph.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Polymorph *PolymorphCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Polymorph.Contract.OwnerOf(&_Polymorph.CallOpts, tokenId)
}

// PriceFor is a free data retrieval call binding the contract method 0x8d5555f2.
//
// Solidity: function priceFor(uint256 tokenNumber) view returns(uint256 price)
func (_Polymorph *PolymorphCaller) PriceFor(opts *bind.CallOpts, tokenNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Polymorph.contract.Call(opts, &out, "priceFor", tokenNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceFor is a free data retrieval call binding the contract method 0x8d5555f2.
//
// Solidity: function priceFor(uint256 tokenNumber) view returns(uint256 price)
func (_Polymorph *PolymorphSession) PriceFor(tokenNumber *big.Int) (*big.Int, error) {
	return _Polymorph.Contract.PriceFor(&_Polymorph.CallOpts, tokenNumber)
}

// PriceFor is a free data retrieval call binding the contract method 0x8d5555f2.
//
// Solidity: function priceFor(uint256 tokenNumber) view returns(uint256 price)
func (_Polymorph *PolymorphCallerSession) PriceFor(tokenNumber *big.Int) (*big.Int, error) {
	return _Polymorph.Contract.PriceFor(&_Polymorph.CallOpts, tokenNumber)
}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_Polymorph *PolymorphCaller) PriceForGenomeChange(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Polymorph.contract.Call(opts, &out, "priceForGenomeChange", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_Polymorph *PolymorphSession) PriceForGenomeChange(tokenId *big.Int) (*big.Int, error) {
	return _Polymorph.Contract.PriceForGenomeChange(&_Polymorph.CallOpts, tokenId)
}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_Polymorph *PolymorphCallerSession) PriceForGenomeChange(tokenId *big.Int) (*big.Int, error) {
	return _Polymorph.Contract.PriceForGenomeChange(&_Polymorph.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Polymorph *PolymorphCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Polymorph.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Polymorph *PolymorphSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Polymorph.Contract.SupportsInterface(&_Polymorph.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Polymorph *PolymorphCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Polymorph.Contract.SupportsInterface(&_Polymorph.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Polymorph *PolymorphTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Polymorph *PolymorphSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.Approve(&_Polymorph.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Polymorph *PolymorphTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.Approve(&_Polymorph.TransactOpts, to, tokenId)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_Polymorph *PolymorphTransactor) ChangeBaseGenomeChangePrice(opts *bind.TransactOpts, newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "changeBaseGenomeChangePrice", newGenomeChangePrice)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_Polymorph *PolymorphSession) ChangeBaseGenomeChangePrice(newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.ChangeBaseGenomeChangePrice(&_Polymorph.TransactOpts, newGenomeChangePrice)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_Polymorph *PolymorphTransactorSession) ChangeBaseGenomeChangePrice(newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.ChangeBaseGenomeChangePrice(&_Polymorph.TransactOpts, newGenomeChangePrice)
}

// ChangeSlope is a paid mutator transaction binding the contract method 0xf1cd9565.
//
// Solidity: function changeSlope(uint256 newSlope) returns()
func (_Polymorph *PolymorphTransactor) ChangeSlope(opts *bind.TransactOpts, newSlope *big.Int) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "changeSlope", newSlope)
}

// ChangeSlope is a paid mutator transaction binding the contract method 0xf1cd9565.
//
// Solidity: function changeSlope(uint256 newSlope) returns()
func (_Polymorph *PolymorphSession) ChangeSlope(newSlope *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.ChangeSlope(&_Polymorph.TransactOpts, newSlope)
}

// ChangeSlope is a paid mutator transaction binding the contract method 0xf1cd9565.
//
// Solidity: function changeSlope(uint256 newSlope) returns()
func (_Polymorph *PolymorphTransactorSession) ChangeSlope(newSlope *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.ChangeSlope(&_Polymorph.TransactOpts, newSlope)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Polymorph *PolymorphTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Polymorph *PolymorphSession) Mint() (*types.Transaction, error) {
	return _Polymorph.Contract.Mint(&_Polymorph.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Polymorph *PolymorphTransactorSession) Mint() (*types.Transaction, error) {
	return _Polymorph.Contract.Mint(&_Polymorph.TransactOpts)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_Polymorph *PolymorphTransactor) MorphGene(opts *bind.TransactOpts, tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "morphGene", tokenId, genePosition)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_Polymorph *PolymorphSession) MorphGene(tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.MorphGene(&_Polymorph.TransactOpts, tokenId, genePosition)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_Polymorph *PolymorphTransactorSession) MorphGene(tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.MorphGene(&_Polymorph.TransactOpts, tokenId, genePosition)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_Polymorph *PolymorphTransactor) RandomizeGenome(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "randomizeGenome", tokenId)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_Polymorph *PolymorphSession) RandomizeGenome(tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.RandomizeGenome(&_Polymorph.TransactOpts, tokenId)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_Polymorph *PolymorphTransactorSession) RandomizeGenome(tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.RandomizeGenome(&_Polymorph.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Polymorph *PolymorphTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Polymorph *PolymorphSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.SafeTransferFrom(&_Polymorph.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Polymorph *PolymorphTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.SafeTransferFrom(&_Polymorph.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Polymorph *PolymorphTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Polymorph *PolymorphSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Polymorph.Contract.SafeTransferFrom0(&_Polymorph.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Polymorph *PolymorphTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Polymorph.Contract.SafeTransferFrom0(&_Polymorph.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_Polymorph *PolymorphTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_Polymorph *PolymorphSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Polymorph.Contract.SetApprovalForAll(&_Polymorph.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_Polymorph *PolymorphTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Polymorph.Contract.SetApprovalForAll(&_Polymorph.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Polymorph *PolymorphTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Polymorph *PolymorphSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.TransferFrom(&_Polymorph.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Polymorph *PolymorphTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Polymorph.Contract.TransferFrom(&_Polymorph.TransactOpts, from, to, tokenId)
}

// PolymorphApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Polymorph contract.
type PolymorphApprovalIterator struct {
	Event *PolymorphApproval // Event containing the contract specifics and raw log

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
func (it *PolymorphApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphApproval)
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
		it.Event = new(PolymorphApproval)
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
func (it *PolymorphApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphApproval represents a Approval event raised by the Polymorph contract.
type PolymorphApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Polymorph *PolymorphFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PolymorphApprovalIterator, error) {

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

	logs, sub, err := _Polymorph.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphApprovalIterator{contract: _Polymorph.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Polymorph *PolymorphFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PolymorphApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Polymorph.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphApproval)
				if err := _Polymorph.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Polymorph *PolymorphFilterer) ParseApproval(log types.Log) (*PolymorphApproval, error) {
	event := new(PolymorphApproval)
	if err := _Polymorph.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PolymorphApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Polymorph contract.
type PolymorphApprovalForAllIterator struct {
	Event *PolymorphApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PolymorphApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphApprovalForAll)
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
		it.Event = new(PolymorphApprovalForAll)
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
func (it *PolymorphApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphApprovalForAll represents a ApprovalForAll event raised by the Polymorph contract.
type PolymorphApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Polymorph *PolymorphFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PolymorphApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Polymorph.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphApprovalForAllIterator{contract: _Polymorph.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Polymorph *PolymorphFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PolymorphApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Polymorph.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphApprovalForAll)
				if err := _Polymorph.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Polymorph *PolymorphFilterer) ParseApprovalForAll(log types.Log) (*PolymorphApprovalForAll, error) {
	event := new(PolymorphApprovalForAll)
	if err := _Polymorph.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PolymorphTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Polymorph contract.
type PolymorphTransferIterator struct {
	Event *PolymorphTransfer // Event containing the contract specifics and raw log

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
func (it *PolymorphTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphTransfer)
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
		it.Event = new(PolymorphTransfer)
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
func (it *PolymorphTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphTransfer represents a Transfer event raised by the Polymorph contract.
type PolymorphTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Polymorph *PolymorphFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PolymorphTransferIterator, error) {

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

	logs, sub, err := _Polymorph.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphTransferIterator{contract: _Polymorph.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Polymorph *PolymorphFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PolymorphTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Polymorph.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphTransfer)
				if err := _Polymorph.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Polymorph *PolymorphFilterer) ParseTransfer(log types.Log) (*PolymorphTransfer, error) {
	event := new(PolymorphTransfer)
	if err := _Polymorph.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
