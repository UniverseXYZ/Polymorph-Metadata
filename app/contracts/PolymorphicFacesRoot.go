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

// PolymorphicFacesRootParams is an auto generated low-level Go binding around an user-defined struct.
type PolymorphicFacesRootParams struct {
	Name                  string
	Symbol                string
	BaseURI               string
	DaoAddress            common.Address
	RoyaltyFee            *big.Int
	BaseGenomeChangePrice *big.Int
	MaxSupply             *big.Int
	RandomizeGenomePrice  *big.Int
	ArweaveAssetsJSON     string
	PolymorphV2Address    common.Address
}

const PolymorphicFacesRootABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_daoAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_royaltyFee\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"_baseGenomeChangePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_randomizeGenomePrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_arweaveAssetsJSON\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_polymorphV2Address\",\"type\":\"address\"}],\"internalType\":\"structPolymorphicFacesRoot.Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"arweaveAssetsJSON\",\"type\":\"string\"}],\"name\":\"ArweaveAssetsJSONChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGenomeChange\",\"type\":\"uint256\"}],\"name\":\"BaseGenomeChangePriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ConsumerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"newDefaultRoyalty\",\"type\":\"uint96\"}],\"name\":\"DefaultRoyaltyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPolyV2Address\",\"type\":\"address\"}],\"name\":\"PolyV2AddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRandomizeGenomePriceChange\",\"type\":\"uint256\"}],\"name\":\"RandomizeGenomePriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumPolymorphicFaces.FacesEventType\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"TokenMorphed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arweaveAssetsJSON\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseGenomeChangePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGenomeChangePrice\",\"type\":\"uint256\"}],\"name\":\"changeBaseGenomeChangePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consumer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"changeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRandomizeGenomePrice\",\"type\":\"uint256\"}],\"name\":\"changeRandomizeGenomePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"consumerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"geneOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"genomeChanges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"genomeChnages\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isNotVirgin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genePosition\",\"type\":\"uint256\"}],\"name\":\"morphGene\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polymorphV2Contract\",\"outputs\":[{\"internalType\":\"contractPolymorphRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"priceForGenomeChange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"randomizeGenome\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomizeGenomePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_arweaveAssetsJSON\",\"type\":\"string\"}],\"name\":\"setArweaveAssetsJSON\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyFee\",\"type\":\"uint96\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPolyV2Address\",\"type\":\"address\"}],\"name\":\"setPolyV2Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"whitelistBridgeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistTunnelAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVirgin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"genomeChangesCount\",\"type\":\"uint256\"}],\"name\":\"wormholeUpdateGene\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PolymorphicFacesRoot is an auto generated Go binding around an Ethereum contract.
type PolymorphicFacesRoot struct {
	PolymorphicFacesRootCaller     // Read-only binding to the contract
	PolymorphicFacesRootTransactor // Write-only binding to the contract
	PolymorphicFacesRootFilterer   // Log filterer for contract events
}

// PolymorphicFacesRootCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolymorphicFacesRootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphicFacesRootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolymorphicFacesRootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphicFacesRootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolymorphicFacesRootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphicFacesRootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolymorphicFacesRootSession struct {
	Contract     *PolymorphicFacesRoot // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PolymorphicFacesRootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolymorphicFacesRootCallerSession struct {
	Contract *PolymorphicFacesRootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PolymorphicFacesRootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolymorphicFacesRootTransactorSession struct {
	Contract     *PolymorphicFacesRootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PolymorphicFacesRootRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolymorphicFacesRootRaw struct {
	Contract *PolymorphicFacesRoot // Generic contract binding to access the raw methods on
}

// PolymorphicFacesRootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolymorphicFacesRootCallerRaw struct {
	Contract *PolymorphicFacesRootCaller // Generic read-only contract binding to access the raw methods on
}

// PolymorphicFacesRootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolymorphicFacesRootTransactorRaw struct {
	Contract *PolymorphicFacesRootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolymorphicFacesRoot creates a new instance of PolymorphicFacesRoot, bound to a specific deployed contract.
func NewPolymorphicFacesRoot(address common.Address, backend bind.ContractBackend) (*PolymorphicFacesRoot, error) {
	contract, err := bindPolymorphicFacesRoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRoot{PolymorphicFacesRootCaller: PolymorphicFacesRootCaller{contract: contract}, PolymorphicFacesRootTransactor: PolymorphicFacesRootTransactor{contract: contract}, PolymorphicFacesRootFilterer: PolymorphicFacesRootFilterer{contract: contract}}, nil
}

// NewPolymorphicFacesRootCaller creates a new read-only instance of PolymorphicFacesRoot, bound to a specific deployed contract.
func NewPolymorphicFacesRootCaller(address common.Address, caller bind.ContractCaller) (*PolymorphicFacesRootCaller, error) {
	contract, err := bindPolymorphicFacesRoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootCaller{contract: contract}, nil
}

// NewPolymorphicFacesRootTransactor creates a new write-only instance of PolymorphicFacesRoot, bound to a specific deployed contract.
func NewPolymorphicFacesRootTransactor(address common.Address, transactor bind.ContractTransactor) (*PolymorphicFacesRootTransactor, error) {
	contract, err := bindPolymorphicFacesRoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootTransactor{contract: contract}, nil
}

// NewPolymorphicFacesRootFilterer creates a new log filterer instance of PolymorphicFacesRoot, bound to a specific deployed contract.
func NewPolymorphicFacesRootFilterer(address common.Address, filterer bind.ContractFilterer) (*PolymorphicFacesRootFilterer, error) {
	contract, err := bindPolymorphicFacesRoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootFilterer{contract: contract}, nil
}

// bindPolymorphicFacesRoot binds a generic wrapper to an already deployed contract.
func bindPolymorphicFacesRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolymorphicFacesRootABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolymorphicFacesRoot *PolymorphicFacesRootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolymorphicFacesRoot.Contract.PolymorphicFacesRootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolymorphicFacesRoot *PolymorphicFacesRootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.PolymorphicFacesRootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolymorphicFacesRoot *PolymorphicFacesRootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.PolymorphicFacesRootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolymorphicFacesRoot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PolymorphicFacesRoot.Contract.DEFAULTADMINROLE(&_PolymorphicFacesRoot.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PolymorphicFacesRoot.Contract.DEFAULTADMINROLE(&_PolymorphicFacesRoot.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) MINTERROLE() ([32]byte, error) {
	return _PolymorphicFacesRoot.Contract.MINTERROLE(&_PolymorphicFacesRoot.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) MINTERROLE() ([32]byte, error) {
	return _PolymorphicFacesRoot.Contract.MINTERROLE(&_PolymorphicFacesRoot.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) PAUSERROLE() ([32]byte, error) {
	return _PolymorphicFacesRoot.Contract.PAUSERROLE(&_PolymorphicFacesRoot.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) PAUSERROLE() ([32]byte, error) {
	return _PolymorphicFacesRoot.Contract.PAUSERROLE(&_PolymorphicFacesRoot.CallOpts)
}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) ArweaveAssetsJSON(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "arweaveAssetsJSON")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) ArweaveAssetsJSON() (string, error) {
	return _PolymorphicFacesRoot.Contract.ArweaveAssetsJSON(&_PolymorphicFacesRoot.CallOpts)
}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) ArweaveAssetsJSON() (string, error) {
	return _PolymorphicFacesRoot.Contract.ArweaveAssetsJSON(&_PolymorphicFacesRoot.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.BalanceOf(&_PolymorphicFacesRoot.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.BalanceOf(&_PolymorphicFacesRoot.CallOpts, owner)
}

// BaseGenomeChangePrice is a free data retrieval call binding the contract method 0xce14617d.
//
// Solidity: function baseGenomeChangePrice() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) BaseGenomeChangePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "baseGenomeChangePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseGenomeChangePrice is a free data retrieval call binding the contract method 0xce14617d.
//
// Solidity: function baseGenomeChangePrice() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) BaseGenomeChangePrice() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.BaseGenomeChangePrice(&_PolymorphicFacesRoot.CallOpts)
}

// BaseGenomeChangePrice is a free data retrieval call binding the contract method 0xce14617d.
//
// Solidity: function baseGenomeChangePrice() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) BaseGenomeChangePrice() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.BaseGenomeChangePrice(&_PolymorphicFacesRoot.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) BaseURI() (string, error) {
	return _PolymorphicFacesRoot.Contract.BaseURI(&_PolymorphicFacesRoot.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) BaseURI() (string, error) {
	return _PolymorphicFacesRoot.Contract.BaseURI(&_PolymorphicFacesRoot.CallOpts)
}

// ConsumerOf is a free data retrieval call binding the contract method 0xe5892331.
//
// Solidity: function consumerOf(uint256 _tokenId) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) ConsumerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "consumerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConsumerOf is a free data retrieval call binding the contract method 0xe5892331.
//
// Solidity: function consumerOf(uint256 _tokenId) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) ConsumerOf(_tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.ConsumerOf(&_PolymorphicFacesRoot.CallOpts, _tokenId)
}

// ConsumerOf is a free data retrieval call binding the contract method 0xe5892331.
//
// Solidity: function consumerOf(uint256 _tokenId) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) ConsumerOf(_tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.ConsumerOf(&_PolymorphicFacesRoot.CallOpts, _tokenId)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "daoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) DaoAddress() (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.DaoAddress(&_PolymorphicFacesRoot.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) DaoAddress() (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.DaoAddress(&_PolymorphicFacesRoot.CallOpts)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) GeneOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "geneOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.GeneOf(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.GeneOf(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// GenomeChanges is a free data retrieval call binding the contract method 0x23c8d07a.
//
// Solidity: function genomeChanges(uint256 tokenId) view returns(uint256 genomeChnages)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) GenomeChanges(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "genomeChanges", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenomeChanges is a free data retrieval call binding the contract method 0x23c8d07a.
//
// Solidity: function genomeChanges(uint256 tokenId) view returns(uint256 genomeChnages)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) GenomeChanges(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.GenomeChanges(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// GenomeChanges is a free data retrieval call binding the contract method 0x23c8d07a.
//
// Solidity: function genomeChanges(uint256 tokenId) view returns(uint256 genomeChnages)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) GenomeChanges(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.GenomeChanges(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.GetApproved(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.GetApproved(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PolymorphicFacesRoot.Contract.GetRoleAdmin(&_PolymorphicFacesRoot.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PolymorphicFacesRoot.Contract.GetRoleAdmin(&_PolymorphicFacesRoot.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.GetRoleMember(&_PolymorphicFacesRoot.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.GetRoleMember(&_PolymorphicFacesRoot.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.GetRoleMemberCount(&_PolymorphicFacesRoot.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.GetRoleMemberCount(&_PolymorphicFacesRoot.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PolymorphicFacesRoot.Contract.HasRole(&_PolymorphicFacesRoot.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PolymorphicFacesRoot.Contract.HasRole(&_PolymorphicFacesRoot.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PolymorphicFacesRoot.Contract.IsApprovedForAll(&_PolymorphicFacesRoot.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PolymorphicFacesRoot.Contract.IsApprovedForAll(&_PolymorphicFacesRoot.CallOpts, owner, operator)
}

// IsNotVirgin is a free data retrieval call binding the contract method 0x4df77416.
//
// Solidity: function isNotVirgin(uint256 ) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) IsNotVirgin(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "isNotVirgin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNotVirgin is a free data retrieval call binding the contract method 0x4df77416.
//
// Solidity: function isNotVirgin(uint256 ) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) IsNotVirgin(arg0 *big.Int) (bool, error) {
	return _PolymorphicFacesRoot.Contract.IsNotVirgin(&_PolymorphicFacesRoot.CallOpts, arg0)
}

// IsNotVirgin is a free data retrieval call binding the contract method 0x4df77416.
//
// Solidity: function isNotVirgin(uint256 ) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) IsNotVirgin(arg0 *big.Int) (bool, error) {
	return _PolymorphicFacesRoot.Contract.IsNotVirgin(&_PolymorphicFacesRoot.CallOpts, arg0)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) LastTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "lastTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) LastTokenId() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.LastTokenId(&_PolymorphicFacesRoot.CallOpts)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) LastTokenId() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.LastTokenId(&_PolymorphicFacesRoot.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) MaxSupply() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.MaxSupply(&_PolymorphicFacesRoot.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) MaxSupply() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.MaxSupply(&_PolymorphicFacesRoot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) Name() (string, error) {
	return _PolymorphicFacesRoot.Contract.Name(&_PolymorphicFacesRoot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) Name() (string, error) {
	return _PolymorphicFacesRoot.Contract.Name(&_PolymorphicFacesRoot.CallOpts)
}

// NumClaimed is a free data retrieval call binding the contract method 0x074cba6b.
//
// Solidity: function numClaimed(address ) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) NumClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "numClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumClaimed is a free data retrieval call binding the contract method 0x074cba6b.
//
// Solidity: function numClaimed(address ) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) NumClaimed(arg0 common.Address) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.NumClaimed(&_PolymorphicFacesRoot.CallOpts, arg0)
}

// NumClaimed is a free data retrieval call binding the contract method 0x074cba6b.
//
// Solidity: function numClaimed(address ) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) NumClaimed(arg0 common.Address) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.NumClaimed(&_PolymorphicFacesRoot.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.OwnerOf(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.OwnerOf(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) Paused() (bool, error) {
	return _PolymorphicFacesRoot.Contract.Paused(&_PolymorphicFacesRoot.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) Paused() (bool, error) {
	return _PolymorphicFacesRoot.Contract.Paused(&_PolymorphicFacesRoot.CallOpts)
}

// PolymorphV2Contract is a free data retrieval call binding the contract method 0x016fa7d6.
//
// Solidity: function polymorphV2Contract() view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) PolymorphV2Contract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "polymorphV2Contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolymorphV2Contract is a free data retrieval call binding the contract method 0x016fa7d6.
//
// Solidity: function polymorphV2Contract() view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) PolymorphV2Contract() (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.PolymorphV2Contract(&_PolymorphicFacesRoot.CallOpts)
}

// PolymorphV2Contract is a free data retrieval call binding the contract method 0x016fa7d6.
//
// Solidity: function polymorphV2Contract() view returns(address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) PolymorphV2Contract() (common.Address, error) {
	return _PolymorphicFacesRoot.Contract.PolymorphV2Contract(&_PolymorphicFacesRoot.CallOpts)
}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) PriceForGenomeChange(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "priceForGenomeChange", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) PriceForGenomeChange(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.PriceForGenomeChange(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) PriceForGenomeChange(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.PriceForGenomeChange(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// RandomizeGenomePrice is a free data retrieval call binding the contract method 0xec9c074c.
//
// Solidity: function randomizeGenomePrice() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) RandomizeGenomePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "randomizeGenomePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomizeGenomePrice is a free data retrieval call binding the contract method 0xec9c074c.
//
// Solidity: function randomizeGenomePrice() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) RandomizeGenomePrice() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.RandomizeGenomePrice(&_PolymorphicFacesRoot.CallOpts)
}

// RandomizeGenomePrice is a free data retrieval call binding the contract method 0xec9c074c.
//
// Solidity: function randomizeGenomePrice() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) RandomizeGenomePrice() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.RandomizeGenomePrice(&_PolymorphicFacesRoot.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _PolymorphicFacesRoot.Contract.RoyaltyInfo(&_PolymorphicFacesRoot.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _PolymorphicFacesRoot.Contract.RoyaltyInfo(&_PolymorphicFacesRoot.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PolymorphicFacesRoot.Contract.SupportsInterface(&_PolymorphicFacesRoot.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PolymorphicFacesRoot.Contract.SupportsInterface(&_PolymorphicFacesRoot.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) Symbol() (string, error) {
	return _PolymorphicFacesRoot.Contract.Symbol(&_PolymorphicFacesRoot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) Symbol() (string, error) {
	return _PolymorphicFacesRoot.Contract.Symbol(&_PolymorphicFacesRoot.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.TokenByIndex(&_PolymorphicFacesRoot.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.TokenByIndex(&_PolymorphicFacesRoot.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.TokenOfOwnerByIndex(&_PolymorphicFacesRoot.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.TokenOfOwnerByIndex(&_PolymorphicFacesRoot.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PolymorphicFacesRoot.Contract.TokenURI(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PolymorphicFacesRoot.Contract.TokenURI(&_PolymorphicFacesRoot.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) TotalSupply() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.TotalSupply(&_PolymorphicFacesRoot.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) TotalSupply() (*big.Int, error) {
	return _PolymorphicFacesRoot.Contract.TotalSupply(&_PolymorphicFacesRoot.CallOpts)
}

// WhitelistTunnelAddresses is a free data retrieval call binding the contract method 0xcccb6d0d.
//
// Solidity: function whitelistTunnelAddresses(address ) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCaller) WhitelistTunnelAddresses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PolymorphicFacesRoot.contract.Call(opts, &out, "whitelistTunnelAddresses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistTunnelAddresses is a free data retrieval call binding the contract method 0xcccb6d0d.
//
// Solidity: function whitelistTunnelAddresses(address ) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) WhitelistTunnelAddresses(arg0 common.Address) (bool, error) {
	return _PolymorphicFacesRoot.Contract.WhitelistTunnelAddresses(&_PolymorphicFacesRoot.CallOpts, arg0)
}

// WhitelistTunnelAddresses is a free data retrieval call binding the contract method 0xcccb6d0d.
//
// Solidity: function whitelistTunnelAddresses(address ) view returns(bool)
func (_PolymorphicFacesRoot *PolymorphicFacesRootCallerSession) WhitelistTunnelAddresses(arg0 common.Address) (bool, error) {
	return _PolymorphicFacesRoot.Contract.WhitelistTunnelAddresses(&_PolymorphicFacesRoot.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Approve(&_PolymorphicFacesRoot.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Approve(&_PolymorphicFacesRoot.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Burn(&_PolymorphicFacesRoot.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Burn(&_PolymorphicFacesRoot.TransactOpts, tokenId)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) ChangeBaseGenomeChangePrice(opts *bind.TransactOpts, newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "changeBaseGenomeChangePrice", newGenomeChangePrice)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) ChangeBaseGenomeChangePrice(newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.ChangeBaseGenomeChangePrice(&_PolymorphicFacesRoot.TransactOpts, newGenomeChangePrice)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) ChangeBaseGenomeChangePrice(newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.ChangeBaseGenomeChangePrice(&_PolymorphicFacesRoot.TransactOpts, newGenomeChangePrice)
}

// ChangeConsumer is a paid mutator transaction binding the contract method 0x70b5aecb.
//
// Solidity: function changeConsumer(address _consumer, uint256 _tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) ChangeConsumer(opts *bind.TransactOpts, _consumer common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "changeConsumer", _consumer, _tokenId)
}

// ChangeConsumer is a paid mutator transaction binding the contract method 0x70b5aecb.
//
// Solidity: function changeConsumer(address _consumer, uint256 _tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) ChangeConsumer(_consumer common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.ChangeConsumer(&_PolymorphicFacesRoot.TransactOpts, _consumer, _tokenId)
}

// ChangeConsumer is a paid mutator transaction binding the contract method 0x70b5aecb.
//
// Solidity: function changeConsumer(address _consumer, uint256 _tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) ChangeConsumer(_consumer common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.ChangeConsumer(&_PolymorphicFacesRoot.TransactOpts, _consumer, _tokenId)
}

// ChangeRandomizeGenomePrice is a paid mutator transaction binding the contract method 0x98c5c078.
//
// Solidity: function changeRandomizeGenomePrice(uint256 newRandomizeGenomePrice) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) ChangeRandomizeGenomePrice(opts *bind.TransactOpts, newRandomizeGenomePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "changeRandomizeGenomePrice", newRandomizeGenomePrice)
}

// ChangeRandomizeGenomePrice is a paid mutator transaction binding the contract method 0x98c5c078.
//
// Solidity: function changeRandomizeGenomePrice(uint256 newRandomizeGenomePrice) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) ChangeRandomizeGenomePrice(newRandomizeGenomePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.ChangeRandomizeGenomePrice(&_PolymorphicFacesRoot.TransactOpts, newRandomizeGenomePrice)
}

// ChangeRandomizeGenomePrice is a paid mutator transaction binding the contract method 0x98c5c078.
//
// Solidity: function changeRandomizeGenomePrice(uint256 newRandomizeGenomePrice) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) ChangeRandomizeGenomePrice(newRandomizeGenomePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.ChangeRandomizeGenomePrice(&_PolymorphicFacesRoot.TransactOpts, newRandomizeGenomePrice)
}

// DaoMint is a paid mutator transaction binding the contract method 0x593518aa.
//
// Solidity: function daoMint() returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) DaoMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "daoMint")
}

// DaoMint is a paid mutator transaction binding the contract method 0x593518aa.
//
// Solidity: function daoMint() returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) DaoMint() (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.DaoMint(&_PolymorphicFacesRoot.TransactOpts)
}

// DaoMint is a paid mutator transaction binding the contract method 0x593518aa.
//
// Solidity: function daoMint() returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) DaoMint() (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.DaoMint(&_PolymorphicFacesRoot.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.GrantRole(&_PolymorphicFacesRoot.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.GrantRole(&_PolymorphicFacesRoot.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) Mint(to common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Mint(&_PolymorphicFacesRoot.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Mint(&_PolymorphicFacesRoot.TransactOpts, to)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) Mint0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "mint0", amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Mint0(&_PolymorphicFacesRoot.TransactOpts, amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Mint0(&_PolymorphicFacesRoot.TransactOpts, amount)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) MorphGene(opts *bind.TransactOpts, tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "morphGene", tokenId, genePosition)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) MorphGene(tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.MorphGene(&_PolymorphicFacesRoot.TransactOpts, tokenId, genePosition)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) MorphGene(tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.MorphGene(&_PolymorphicFacesRoot.TransactOpts, tokenId, genePosition)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) Pause() (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Pause(&_PolymorphicFacesRoot.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) Pause() (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Pause(&_PolymorphicFacesRoot.TransactOpts)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) RandomizeGenome(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "randomizeGenome", tokenId)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) RandomizeGenome(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.RandomizeGenome(&_PolymorphicFacesRoot.TransactOpts, tokenId)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) RandomizeGenome(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.RandomizeGenome(&_PolymorphicFacesRoot.TransactOpts, tokenId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.RenounceRole(&_PolymorphicFacesRoot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.RenounceRole(&_PolymorphicFacesRoot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.RevokeRole(&_PolymorphicFacesRoot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.RevokeRole(&_PolymorphicFacesRoot.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SafeTransferFrom(&_PolymorphicFacesRoot.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SafeTransferFrom(&_PolymorphicFacesRoot.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SafeTransferFrom0(&_PolymorphicFacesRoot.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SafeTransferFrom0(&_PolymorphicFacesRoot.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetApprovalForAll(&_PolymorphicFacesRoot.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetApprovalForAll(&_PolymorphicFacesRoot.TransactOpts, operator, approved)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) SetArweaveAssetsJSON(opts *bind.TransactOpts, _arweaveAssetsJSON string) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "setArweaveAssetsJSON", _arweaveAssetsJSON)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) SetArweaveAssetsJSON(_arweaveAssetsJSON string) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetArweaveAssetsJSON(&_PolymorphicFacesRoot.TransactOpts, _arweaveAssetsJSON)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) SetArweaveAssetsJSON(_arweaveAssetsJSON string) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetArweaveAssetsJSON(&_PolymorphicFacesRoot.TransactOpts, _arweaveAssetsJSON)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetBaseURI(&_PolymorphicFacesRoot.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetBaseURI(&_PolymorphicFacesRoot.TransactOpts, _baseURI)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 royaltyFee) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) SetDefaultRoyalty(opts *bind.TransactOpts, receiver common.Address, royaltyFee *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "setDefaultRoyalty", receiver, royaltyFee)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 royaltyFee) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) SetDefaultRoyalty(receiver common.Address, royaltyFee *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetDefaultRoyalty(&_PolymorphicFacesRoot.TransactOpts, receiver, royaltyFee)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 royaltyFee) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) SetDefaultRoyalty(receiver common.Address, royaltyFee *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetDefaultRoyalty(&_PolymorphicFacesRoot.TransactOpts, receiver, royaltyFee)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) SetMaxSupply(opts *bind.TransactOpts, _maxSupply *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "setMaxSupply", _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetMaxSupply(&_PolymorphicFacesRoot.TransactOpts, _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetMaxSupply(&_PolymorphicFacesRoot.TransactOpts, _maxSupply)
}

// SetPolyV2Address is a paid mutator transaction binding the contract method 0x274ea5f1.
//
// Solidity: function setPolyV2Address(address newPolyV2Address) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) SetPolyV2Address(opts *bind.TransactOpts, newPolyV2Address common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "setPolyV2Address", newPolyV2Address)
}

// SetPolyV2Address is a paid mutator transaction binding the contract method 0x274ea5f1.
//
// Solidity: function setPolyV2Address(address newPolyV2Address) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) SetPolyV2Address(newPolyV2Address common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetPolyV2Address(&_PolymorphicFacesRoot.TransactOpts, newPolyV2Address)
}

// SetPolyV2Address is a paid mutator transaction binding the contract method 0x274ea5f1.
//
// Solidity: function setPolyV2Address(address newPolyV2Address) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) SetPolyV2Address(newPolyV2Address common.Address) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.SetPolyV2Address(&_PolymorphicFacesRoot.TransactOpts, newPolyV2Address)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.TransferFrom(&_PolymorphicFacesRoot.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.TransferFrom(&_PolymorphicFacesRoot.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) Unpause() (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Unpause(&_PolymorphicFacesRoot.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) Unpause() (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.Unpause(&_PolymorphicFacesRoot.TransactOpts)
}

// WhitelistBridgeAddress is a paid mutator transaction binding the contract method 0xab39a3c8.
//
// Solidity: function whitelistBridgeAddress(address bridgeAddress, bool status) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) WhitelistBridgeAddress(opts *bind.TransactOpts, bridgeAddress common.Address, status bool) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "whitelistBridgeAddress", bridgeAddress, status)
}

// WhitelistBridgeAddress is a paid mutator transaction binding the contract method 0xab39a3c8.
//
// Solidity: function whitelistBridgeAddress(address bridgeAddress, bool status) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) WhitelistBridgeAddress(bridgeAddress common.Address, status bool) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.WhitelistBridgeAddress(&_PolymorphicFacesRoot.TransactOpts, bridgeAddress, status)
}

// WhitelistBridgeAddress is a paid mutator transaction binding the contract method 0xab39a3c8.
//
// Solidity: function whitelistBridgeAddress(address bridgeAddress, bool status) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) WhitelistBridgeAddress(bridgeAddress common.Address, status bool) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.WhitelistBridgeAddress(&_PolymorphicFacesRoot.TransactOpts, bridgeAddress, status)
}

// WormholeUpdateGene is a paid mutator transaction binding the contract method 0x6a1c03dc.
//
// Solidity: function wormholeUpdateGene(uint256 tokenId, uint256 gene, bool isVirgin, uint256 genomeChangesCount) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactor) WormholeUpdateGene(opts *bind.TransactOpts, tokenId *big.Int, gene *big.Int, isVirgin bool, genomeChangesCount *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.contract.Transact(opts, "wormholeUpdateGene", tokenId, gene, isVirgin, genomeChangesCount)
}

// WormholeUpdateGene is a paid mutator transaction binding the contract method 0x6a1c03dc.
//
// Solidity: function wormholeUpdateGene(uint256 tokenId, uint256 gene, bool isVirgin, uint256 genomeChangesCount) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootSession) WormholeUpdateGene(tokenId *big.Int, gene *big.Int, isVirgin bool, genomeChangesCount *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.WormholeUpdateGene(&_PolymorphicFacesRoot.TransactOpts, tokenId, gene, isVirgin, genomeChangesCount)
}

// WormholeUpdateGene is a paid mutator transaction binding the contract method 0x6a1c03dc.
//
// Solidity: function wormholeUpdateGene(uint256 tokenId, uint256 gene, bool isVirgin, uint256 genomeChangesCount) returns()
func (_PolymorphicFacesRoot *PolymorphicFacesRootTransactorSession) WormholeUpdateGene(tokenId *big.Int, gene *big.Int, isVirgin bool, genomeChangesCount *big.Int) (*types.Transaction, error) {
	return _PolymorphicFacesRoot.Contract.WormholeUpdateGene(&_PolymorphicFacesRoot.TransactOpts, tokenId, gene, isVirgin, genomeChangesCount)
}

// PolymorphicFacesRootApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootApprovalIterator struct {
	Event *PolymorphicFacesRootApproval // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootApproval)
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
		it.Event = new(PolymorphicFacesRootApproval)
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
func (it *PolymorphicFacesRootApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootApproval represents a Approval event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PolymorphicFacesRootApprovalIterator, error) {

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

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootApprovalIterator{contract: _PolymorphicFacesRoot.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootApproval)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseApproval(log types.Log) (*PolymorphicFacesRootApproval, error) {
	event := new(PolymorphicFacesRootApproval)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootApprovalForAllIterator struct {
	Event *PolymorphicFacesRootApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootApprovalForAll)
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
		it.Event = new(PolymorphicFacesRootApprovalForAll)
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
func (it *PolymorphicFacesRootApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootApprovalForAll represents a ApprovalForAll event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PolymorphicFacesRootApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootApprovalForAllIterator{contract: _PolymorphicFacesRoot.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootApprovalForAll)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseApprovalForAll(log types.Log) (*PolymorphicFacesRootApprovalForAll, error) {
	event := new(PolymorphicFacesRootApprovalForAll)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootArweaveAssetsJSONChangedIterator is returned from FilterArweaveAssetsJSONChanged and is used to iterate over the raw logs and unpacked data for ArweaveAssetsJSONChanged events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootArweaveAssetsJSONChangedIterator struct {
	Event *PolymorphicFacesRootArweaveAssetsJSONChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootArweaveAssetsJSONChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootArweaveAssetsJSONChanged)
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
		it.Event = new(PolymorphicFacesRootArweaveAssetsJSONChanged)
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
func (it *PolymorphicFacesRootArweaveAssetsJSONChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootArweaveAssetsJSONChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootArweaveAssetsJSONChanged represents a ArweaveAssetsJSONChanged event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootArweaveAssetsJSONChanged struct {
	ArweaveAssetsJSON string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterArweaveAssetsJSONChanged is a free log retrieval operation binding the contract event 0x4a826ca029d05af64e411551e15f7ee1e70af0b9bc43a31154ace86a863397b4.
//
// Solidity: event ArweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterArweaveAssetsJSONChanged(opts *bind.FilterOpts) (*PolymorphicFacesRootArweaveAssetsJSONChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "ArweaveAssetsJSONChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootArweaveAssetsJSONChangedIterator{contract: _PolymorphicFacesRoot.contract, event: "ArweaveAssetsJSONChanged", logs: logs, sub: sub}, nil
}

// WatchArweaveAssetsJSONChanged is a free log subscription operation binding the contract event 0x4a826ca029d05af64e411551e15f7ee1e70af0b9bc43a31154ace86a863397b4.
//
// Solidity: event ArweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchArweaveAssetsJSONChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootArweaveAssetsJSONChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "ArweaveAssetsJSONChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootArweaveAssetsJSONChanged)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "ArweaveAssetsJSONChanged", log); err != nil {
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

// ParseArweaveAssetsJSONChanged is a log parse operation binding the contract event 0x4a826ca029d05af64e411551e15f7ee1e70af0b9bc43a31154ace86a863397b4.
//
// Solidity: event ArweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseArweaveAssetsJSONChanged(log types.Log) (*PolymorphicFacesRootArweaveAssetsJSONChanged, error) {
	event := new(PolymorphicFacesRootArweaveAssetsJSONChanged)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "ArweaveAssetsJSONChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootBaseGenomeChangePriceChangedIterator is returned from FilterBaseGenomeChangePriceChanged and is used to iterate over the raw logs and unpacked data for BaseGenomeChangePriceChanged events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootBaseGenomeChangePriceChangedIterator struct {
	Event *PolymorphicFacesRootBaseGenomeChangePriceChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootBaseGenomeChangePriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootBaseGenomeChangePriceChanged)
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
		it.Event = new(PolymorphicFacesRootBaseGenomeChangePriceChanged)
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
func (it *PolymorphicFacesRootBaseGenomeChangePriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootBaseGenomeChangePriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootBaseGenomeChangePriceChanged represents a BaseGenomeChangePriceChanged event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootBaseGenomeChangePriceChanged struct {
	NewGenomeChange *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBaseGenomeChangePriceChanged is a free log retrieval operation binding the contract event 0xb1d78271daba9a366098d40b64d642a1399cabaa22c5234bacc87e92cef82ae6.
//
// Solidity: event BaseGenomeChangePriceChanged(uint256 newGenomeChange)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterBaseGenomeChangePriceChanged(opts *bind.FilterOpts) (*PolymorphicFacesRootBaseGenomeChangePriceChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "BaseGenomeChangePriceChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootBaseGenomeChangePriceChangedIterator{contract: _PolymorphicFacesRoot.contract, event: "BaseGenomeChangePriceChanged", logs: logs, sub: sub}, nil
}

// WatchBaseGenomeChangePriceChanged is a free log subscription operation binding the contract event 0xb1d78271daba9a366098d40b64d642a1399cabaa22c5234bacc87e92cef82ae6.
//
// Solidity: event BaseGenomeChangePriceChanged(uint256 newGenomeChange)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchBaseGenomeChangePriceChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootBaseGenomeChangePriceChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "BaseGenomeChangePriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootBaseGenomeChangePriceChanged)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "BaseGenomeChangePriceChanged", log); err != nil {
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

// ParseBaseGenomeChangePriceChanged is a log parse operation binding the contract event 0xb1d78271daba9a366098d40b64d642a1399cabaa22c5234bacc87e92cef82ae6.
//
// Solidity: event BaseGenomeChangePriceChanged(uint256 newGenomeChange)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseBaseGenomeChangePriceChanged(log types.Log) (*PolymorphicFacesRootBaseGenomeChangePriceChanged, error) {
	event := new(PolymorphicFacesRootBaseGenomeChangePriceChanged)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "BaseGenomeChangePriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootBaseURIChangedIterator is returned from FilterBaseURIChanged and is used to iterate over the raw logs and unpacked data for BaseURIChanged events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootBaseURIChangedIterator struct {
	Event *PolymorphicFacesRootBaseURIChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootBaseURIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootBaseURIChanged)
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
		it.Event = new(PolymorphicFacesRootBaseURIChanged)
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
func (it *PolymorphicFacesRootBaseURIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootBaseURIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootBaseURIChanged represents a BaseURIChanged event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootBaseURIChanged struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseURIChanged is a free log retrieval operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterBaseURIChanged(opts *bind.FilterOpts) (*PolymorphicFacesRootBaseURIChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootBaseURIChangedIterator{contract: _PolymorphicFacesRoot.contract, event: "BaseURIChanged", logs: logs, sub: sub}, nil
}

// WatchBaseURIChanged is a free log subscription operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchBaseURIChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootBaseURIChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootBaseURIChanged)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
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

// ParseBaseURIChanged is a log parse operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseBaseURIChanged(log types.Log) (*PolymorphicFacesRootBaseURIChanged, error) {
	event := new(PolymorphicFacesRootBaseURIChanged)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootConsumerChangedIterator is returned from FilterConsumerChanged and is used to iterate over the raw logs and unpacked data for ConsumerChanged events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootConsumerChangedIterator struct {
	Event *PolymorphicFacesRootConsumerChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootConsumerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootConsumerChanged)
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
		it.Event = new(PolymorphicFacesRootConsumerChanged)
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
func (it *PolymorphicFacesRootConsumerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootConsumerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootConsumerChanged represents a ConsumerChanged event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootConsumerChanged struct {
	Owner    common.Address
	Consumer common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsumerChanged is a free log retrieval operation binding the contract event 0x42ef856c2602f37ce625d252830bed486c5c8e9a4de8aa36cc3d15f304eb662b.
//
// Solidity: event ConsumerChanged(address indexed owner, address indexed consumer, uint256 indexed tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterConsumerChanged(opts *bind.FilterOpts, owner []common.Address, consumer []common.Address, tokenId []*big.Int) (*PolymorphicFacesRootConsumerChangedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "ConsumerChanged", ownerRule, consumerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootConsumerChangedIterator{contract: _PolymorphicFacesRoot.contract, event: "ConsumerChanged", logs: logs, sub: sub}, nil
}

// WatchConsumerChanged is a free log subscription operation binding the contract event 0x42ef856c2602f37ce625d252830bed486c5c8e9a4de8aa36cc3d15f304eb662b.
//
// Solidity: event ConsumerChanged(address indexed owner, address indexed consumer, uint256 indexed tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchConsumerChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootConsumerChanged, owner []common.Address, consumer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "ConsumerChanged", ownerRule, consumerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootConsumerChanged)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "ConsumerChanged", log); err != nil {
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

// ParseConsumerChanged is a log parse operation binding the contract event 0x42ef856c2602f37ce625d252830bed486c5c8e9a4de8aa36cc3d15f304eb662b.
//
// Solidity: event ConsumerChanged(address indexed owner, address indexed consumer, uint256 indexed tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseConsumerChanged(log types.Log) (*PolymorphicFacesRootConsumerChanged, error) {
	event := new(PolymorphicFacesRootConsumerChanged)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "ConsumerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootDefaultRoyaltyChangedIterator is returned from FilterDefaultRoyaltyChanged and is used to iterate over the raw logs and unpacked data for DefaultRoyaltyChanged events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootDefaultRoyaltyChangedIterator struct {
	Event *PolymorphicFacesRootDefaultRoyaltyChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootDefaultRoyaltyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootDefaultRoyaltyChanged)
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
		it.Event = new(PolymorphicFacesRootDefaultRoyaltyChanged)
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
func (it *PolymorphicFacesRootDefaultRoyaltyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootDefaultRoyaltyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootDefaultRoyaltyChanged represents a DefaultRoyaltyChanged event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootDefaultRoyaltyChanged struct {
	NewReceiver       common.Address
	NewDefaultRoyalty *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyaltyChanged is a free log retrieval operation binding the contract event 0xe5ed39918c4170e24337471011e1ccdeb5e4a433f53fae4eb2ad73e03cd21bda.
//
// Solidity: event DefaultRoyaltyChanged(address newReceiver, uint96 newDefaultRoyalty)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterDefaultRoyaltyChanged(opts *bind.FilterOpts) (*PolymorphicFacesRootDefaultRoyaltyChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "DefaultRoyaltyChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootDefaultRoyaltyChangedIterator{contract: _PolymorphicFacesRoot.contract, event: "DefaultRoyaltyChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyaltyChanged is a free log subscription operation binding the contract event 0xe5ed39918c4170e24337471011e1ccdeb5e4a433f53fae4eb2ad73e03cd21bda.
//
// Solidity: event DefaultRoyaltyChanged(address newReceiver, uint96 newDefaultRoyalty)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchDefaultRoyaltyChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootDefaultRoyaltyChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "DefaultRoyaltyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootDefaultRoyaltyChanged)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "DefaultRoyaltyChanged", log); err != nil {
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

// ParseDefaultRoyaltyChanged is a log parse operation binding the contract event 0xe5ed39918c4170e24337471011e1ccdeb5e4a433f53fae4eb2ad73e03cd21bda.
//
// Solidity: event DefaultRoyaltyChanged(address newReceiver, uint96 newDefaultRoyalty)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseDefaultRoyaltyChanged(log types.Log) (*PolymorphicFacesRootDefaultRoyaltyChanged, error) {
	event := new(PolymorphicFacesRootDefaultRoyaltyChanged)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "DefaultRoyaltyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootMaxSupplyChangedIterator is returned from FilterMaxSupplyChanged and is used to iterate over the raw logs and unpacked data for MaxSupplyChanged events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootMaxSupplyChangedIterator struct {
	Event *PolymorphicFacesRootMaxSupplyChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootMaxSupplyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootMaxSupplyChanged)
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
		it.Event = new(PolymorphicFacesRootMaxSupplyChanged)
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
func (it *PolymorphicFacesRootMaxSupplyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootMaxSupplyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootMaxSupplyChanged represents a MaxSupplyChanged event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootMaxSupplyChanged struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyChanged is a free log retrieval operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterMaxSupplyChanged(opts *bind.FilterOpts) (*PolymorphicFacesRootMaxSupplyChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootMaxSupplyChangedIterator{contract: _PolymorphicFacesRoot.contract, event: "MaxSupplyChanged", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyChanged is a free log subscription operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchMaxSupplyChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootMaxSupplyChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootMaxSupplyChanged)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
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

// ParseMaxSupplyChanged is a log parse operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseMaxSupplyChanged(log types.Log) (*PolymorphicFacesRootMaxSupplyChanged, error) {
	event := new(PolymorphicFacesRootMaxSupplyChanged)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootPausedIterator struct {
	Event *PolymorphicFacesRootPaused // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootPaused)
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
		it.Event = new(PolymorphicFacesRootPaused)
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
func (it *PolymorphicFacesRootPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootPaused represents a Paused event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterPaused(opts *bind.FilterOpts) (*PolymorphicFacesRootPausedIterator, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootPausedIterator{contract: _PolymorphicFacesRoot.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootPaused) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootPaused)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address account)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParsePaused(log types.Log) (*PolymorphicFacesRootPaused, error) {
	event := new(PolymorphicFacesRootPaused)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootPolyV2AddressChangedIterator is returned from FilterPolyV2AddressChanged and is used to iterate over the raw logs and unpacked data for PolyV2AddressChanged events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootPolyV2AddressChangedIterator struct {
	Event *PolymorphicFacesRootPolyV2AddressChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootPolyV2AddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootPolyV2AddressChanged)
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
		it.Event = new(PolymorphicFacesRootPolyV2AddressChanged)
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
func (it *PolymorphicFacesRootPolyV2AddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootPolyV2AddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootPolyV2AddressChanged represents a PolyV2AddressChanged event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootPolyV2AddressChanged struct {
	NewPolyV2Address common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPolyV2AddressChanged is a free log retrieval operation binding the contract event 0x5ec2919217d5d2381c3a8e75d708902dfe1bc0a6e3b2829ea35289a0cea7877e.
//
// Solidity: event PolyV2AddressChanged(address newPolyV2Address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterPolyV2AddressChanged(opts *bind.FilterOpts) (*PolymorphicFacesRootPolyV2AddressChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "PolyV2AddressChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootPolyV2AddressChangedIterator{contract: _PolymorphicFacesRoot.contract, event: "PolyV2AddressChanged", logs: logs, sub: sub}, nil
}

// WatchPolyV2AddressChanged is a free log subscription operation binding the contract event 0x5ec2919217d5d2381c3a8e75d708902dfe1bc0a6e3b2829ea35289a0cea7877e.
//
// Solidity: event PolyV2AddressChanged(address newPolyV2Address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchPolyV2AddressChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootPolyV2AddressChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "PolyV2AddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootPolyV2AddressChanged)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "PolyV2AddressChanged", log); err != nil {
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

// ParsePolyV2AddressChanged is a log parse operation binding the contract event 0x5ec2919217d5d2381c3a8e75d708902dfe1bc0a6e3b2829ea35289a0cea7877e.
//
// Solidity: event PolyV2AddressChanged(address newPolyV2Address)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParsePolyV2AddressChanged(log types.Log) (*PolymorphicFacesRootPolyV2AddressChanged, error) {
	event := new(PolymorphicFacesRootPolyV2AddressChanged)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "PolyV2AddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootRandomizeGenomePriceChangedIterator is returned from FilterRandomizeGenomePriceChanged and is used to iterate over the raw logs and unpacked data for RandomizeGenomePriceChanged events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootRandomizeGenomePriceChangedIterator struct {
	Event *PolymorphicFacesRootRandomizeGenomePriceChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootRandomizeGenomePriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootRandomizeGenomePriceChanged)
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
		it.Event = new(PolymorphicFacesRootRandomizeGenomePriceChanged)
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
func (it *PolymorphicFacesRootRandomizeGenomePriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootRandomizeGenomePriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootRandomizeGenomePriceChanged represents a RandomizeGenomePriceChanged event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootRandomizeGenomePriceChanged struct {
	NewRandomizeGenomePriceChange *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterRandomizeGenomePriceChanged is a free log retrieval operation binding the contract event 0xff4da8d01e7184cc8c9d6c57d64b336b1de6d676b6215408967bd071c8da7e3d.
//
// Solidity: event RandomizeGenomePriceChanged(uint256 newRandomizeGenomePriceChange)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterRandomizeGenomePriceChanged(opts *bind.FilterOpts) (*PolymorphicFacesRootRandomizeGenomePriceChangedIterator, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "RandomizeGenomePriceChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootRandomizeGenomePriceChangedIterator{contract: _PolymorphicFacesRoot.contract, event: "RandomizeGenomePriceChanged", logs: logs, sub: sub}, nil
}

// WatchRandomizeGenomePriceChanged is a free log subscription operation binding the contract event 0xff4da8d01e7184cc8c9d6c57d64b336b1de6d676b6215408967bd071c8da7e3d.
//
// Solidity: event RandomizeGenomePriceChanged(uint256 newRandomizeGenomePriceChange)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchRandomizeGenomePriceChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootRandomizeGenomePriceChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "RandomizeGenomePriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootRandomizeGenomePriceChanged)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "RandomizeGenomePriceChanged", log); err != nil {
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

// ParseRandomizeGenomePriceChanged is a log parse operation binding the contract event 0xff4da8d01e7184cc8c9d6c57d64b336b1de6d676b6215408967bd071c8da7e3d.
//
// Solidity: event RandomizeGenomePriceChanged(uint256 newRandomizeGenomePriceChange)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseRandomizeGenomePriceChanged(log types.Log) (*PolymorphicFacesRootRandomizeGenomePriceChanged, error) {
	event := new(PolymorphicFacesRootRandomizeGenomePriceChanged)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "RandomizeGenomePriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootRoleAdminChangedIterator struct {
	Event *PolymorphicFacesRootRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootRoleAdminChanged)
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
		it.Event = new(PolymorphicFacesRootRoleAdminChanged)
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
func (it *PolymorphicFacesRootRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootRoleAdminChanged represents a RoleAdminChanged event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolymorphicFacesRootRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootRoleAdminChangedIterator{contract: _PolymorphicFacesRoot.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootRoleAdminChanged)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseRoleAdminChanged(log types.Log) (*PolymorphicFacesRootRoleAdminChanged, error) {
	event := new(PolymorphicFacesRootRoleAdminChanged)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootRoleGrantedIterator struct {
	Event *PolymorphicFacesRootRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootRoleGranted)
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
		it.Event = new(PolymorphicFacesRootRoleGranted)
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
func (it *PolymorphicFacesRootRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootRoleGranted represents a RoleGranted event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolymorphicFacesRootRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootRoleGrantedIterator{contract: _PolymorphicFacesRoot.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootRoleGranted)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseRoleGranted(log types.Log) (*PolymorphicFacesRootRoleGranted, error) {
	event := new(PolymorphicFacesRootRoleGranted)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootRoleRevokedIterator struct {
	Event *PolymorphicFacesRootRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootRoleRevoked)
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
		it.Event = new(PolymorphicFacesRootRoleRevoked)
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
func (it *PolymorphicFacesRootRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootRoleRevoked represents a RoleRevoked event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolymorphicFacesRootRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootRoleRevokedIterator{contract: _PolymorphicFacesRoot.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootRoleRevoked)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseRoleRevoked(log types.Log) (*PolymorphicFacesRootRoleRevoked, error) {
	event := new(PolymorphicFacesRootRoleRevoked)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootTokenMintedIterator struct {
	Event *PolymorphicFacesRootTokenMinted // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootTokenMinted)
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
		it.Event = new(PolymorphicFacesRootTokenMinted)
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
func (it *PolymorphicFacesRootTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootTokenMinted represents a TokenMinted event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootTokenMinted struct {
	TokenId *big.Int
	NewGene *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterTokenMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*PolymorphicFacesRootTokenMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootTokenMintedIterator{contract: _PolymorphicFacesRoot.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootTokenMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootTokenMinted)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseTokenMinted(log types.Log) (*PolymorphicFacesRootTokenMinted, error) {
	event := new(PolymorphicFacesRootTokenMinted)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootTokenMorphedIterator is returned from FilterTokenMorphed and is used to iterate over the raw logs and unpacked data for TokenMorphed events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootTokenMorphedIterator struct {
	Event *PolymorphicFacesRootTokenMorphed // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootTokenMorphedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootTokenMorphed)
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
		it.Event = new(PolymorphicFacesRootTokenMorphed)
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
func (it *PolymorphicFacesRootTokenMorphedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootTokenMorphedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootTokenMorphed represents a TokenMorphed event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootTokenMorphed struct {
	TokenId   *big.Int
	OldGene   *big.Int
	NewGene   *big.Int
	Price     *big.Int
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenMorphed is a free log retrieval operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterTokenMorphed(opts *bind.FilterOpts, tokenId []*big.Int) (*PolymorphicFacesRootTokenMorphedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootTokenMorphedIterator{contract: _PolymorphicFacesRoot.contract, event: "TokenMorphed", logs: logs, sub: sub}, nil
}

// WatchTokenMorphed is a free log subscription operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchTokenMorphed(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootTokenMorphed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootTokenMorphed)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
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

// ParseTokenMorphed is a log parse operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseTokenMorphed(log types.Log) (*PolymorphicFacesRootTokenMorphed, error) {
	event := new(PolymorphicFacesRootTokenMorphed)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootTransferIterator struct {
	Event *PolymorphicFacesRootTransfer // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootTransfer)
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
		it.Event = new(PolymorphicFacesRootTransfer)
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
func (it *PolymorphicFacesRootTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootTransfer represents a Transfer event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PolymorphicFacesRootTransferIterator, error) {

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

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootTransferIterator{contract: _PolymorphicFacesRoot.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootTransfer)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseTransfer(log types.Log) (*PolymorphicFacesRootTransfer, error) {
	event := new(PolymorphicFacesRootTransfer)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphicFacesRootUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootUnpausedIterator struct {
	Event *PolymorphicFacesRootUnpaused // Event containing the contract specifics and raw log

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
func (it *PolymorphicFacesRootUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphicFacesRootUnpaused)
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
		it.Event = new(PolymorphicFacesRootUnpaused)
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
func (it *PolymorphicFacesRootUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphicFacesRootUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphicFacesRootUnpaused represents a Unpaused event raised by the PolymorphicFacesRoot contract.
type PolymorphicFacesRootUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PolymorphicFacesRootUnpausedIterator, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PolymorphicFacesRootUnpausedIterator{contract: _PolymorphicFacesRoot.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PolymorphicFacesRootUnpaused) (event.Subscription, error) {

	logs, sub, err := _PolymorphicFacesRoot.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphicFacesRootUnpaused)
				if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address account)
func (_PolymorphicFacesRoot *PolymorphicFacesRootFilterer) ParseUnpaused(log types.Log) (*PolymorphicFacesRootUnpaused, error) {
	event := new(PolymorphicFacesRootUnpaused)
	if err := _PolymorphicFacesRoot.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
