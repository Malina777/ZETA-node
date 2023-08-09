// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zevmswap

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

// Context is an auto generated low-level Go binding around an user-defined struct.
type Context struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// ZEVMSwapAppMetaData contains all meta data concerning the ZEVMSwapApp contract.
var ZEVMSwapAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router02_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"systemContract_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LowAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"decodeMemo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetZRC20\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"encodeMemo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structContext\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router02\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001333380380620013338339818101604052810190620000379190620000c4565b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b8152505050506200015e565b600081519050620000be8162000144565b92915050565b60008060408385031215620000de57620000dd6200013f565b5b6000620000ee85828601620000ad565b92505060206200010185828601620000ad565b9150509250929050565b600062000118826200011f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200014f816200010b565b81146200015b57600080fd5b50565b60805160601c60a05160601c6111946200019f600039600081816101b201526101fa0152600081816101d60152818161039d015261043101526111946000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063a06ea8bc1461005c578063bb88b7691461008d578063bd00c9c4146100ab578063de43156e146100c9578063df73044e146100e5575b600080fd5b61007660048036038101906100719190610a30565b610115565b604051610084929190610c8d565b60405180910390f35b6100956101b0565b6040516100a29190610c72565b60405180910390f35b6100b36101d4565b6040516100c09190610c72565b60405180910390f35b6100e360048036038101906100de9190610a7d565b6101f8565b005b6100ff60048036038101906100fa919061091a565b61076f565b60405161010c9190610ce6565b60405180910390f35b600060608060008585905090506000868660009060149261013893929190610e4e565b906101439190610f31565b60601c90508686601490809261015b93929190610e4e565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505092508083945094505050509250929050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461027d576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000606061028b8484610115565b80925081935050506060600267ffffffffffffffff8111156102b0576102af611097565b5b6040519080825280602002602001820160405280156102de5781602001602082028036833780820191505090505b50905086816000815181106102f6576102f5611068565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828160018151811061034557610344611068565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f0000000000000000000000000000000000000000000000000000000000000000886040518363ffffffff1660e01b81526004016103da929190610cbd565b602060405180830381600087803b1580156103f457600080fd5b505af1158015610408573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042c9190610a03565b5060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398860008530680100000000000000006040518663ffffffff1660e01b815260040161049a959493929190610d38565b600060405180830381600087803b1580156104b457600080fd5b505af11580156104c8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104f191906109ba565b90506000808573ffffffffffffffffffffffffffffffffffffffff1663d9eeebed6040518163ffffffff1660e01b8152600401604080518083038186803b15801561053b57600080fd5b505afa15801561054f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610573919061097a565b915091508173ffffffffffffffffffffffffffffffffffffffff1663095ea7b387836040518363ffffffff1660e01b81526004016105b2929190610cbd565b602060405180830381600087803b1580156105cc57600080fd5b505af11580156105e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106049190610a03565b508573ffffffffffffffffffffffffffffffffffffffff1663095ea7b3878560018151811061063657610635611068565b5b60200260200101516040518363ffffffff1660e01b815260040161065b929190610cbd565b602060405180830381600087803b15801561067557600080fd5b505af1158015610689573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ad9190610a03565b508573ffffffffffffffffffffffffffffffffffffffff1663c70126268683866001815181106106e0576106df611068565b5b60200260200101516106f29190610e89565b6040518363ffffffff1660e01b815260040161070f929190610d08565b602060405180830381600087803b15801561072957600080fd5b505af115801561073d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107619190610a03565b505050505050505050505050565b606083838360405160200161078693929190610c48565b60405160208183030381529060405290509392505050565b60006107b16107ac84610db7565b610d92565b905080838252602082019050828560208602820111156107d4576107d36110df565b5b60005b8581101561080457816107ea8882610905565b8452602084019350602083019250506001810190506107d7565b5050509392505050565b60008135905061081d81611119565b92915050565b60008151905061083281611119565b92915050565b600082601f83011261084d5761084c6110cb565b5b815161085d84826020860161079e565b91505092915050565b60008151905061087581611130565b92915050565b60008083601f840112610891576108906110cb565b5b8235905067ffffffffffffffff8111156108ae576108ad6110c6565b5b6020830191508360018202830111156108ca576108c96110df565b5b9250929050565b6000606082840312156108e7576108e66110d0565b5b81905092915050565b6000813590506108ff81611147565b92915050565b60008151905061091481611147565b92915050565b600080600060408486031215610933576109326110e9565b5b60006109418682870161080e565b935050602084013567ffffffffffffffff811115610962576109616110e4565b5b61096e8682870161087b565b92509250509250925092565b60008060408385031215610991576109906110e9565b5b600061099f85828601610823565b92505060206109b085828601610905565b9150509250929050565b6000602082840312156109d0576109cf6110e9565b5b600082015167ffffffffffffffff8111156109ee576109ed6110e4565b5b6109fa84828501610838565b91505092915050565b600060208284031215610a1957610a186110e9565b5b6000610a2784828501610866565b91505092915050565b60008060208385031215610a4757610a466110e9565b5b600083013567ffffffffffffffff811115610a6557610a646110e4565b5b610a718582860161087b565b92509250509250929050565b600080600080600060808688031215610a9957610a986110e9565b5b600086013567ffffffffffffffff811115610ab757610ab66110e4565b5b610ac3888289016108d1565b9550506020610ad48882890161080e565b9450506040610ae5888289016108f0565b935050606086013567ffffffffffffffff811115610b0657610b056110e4565b5b610b128882890161087b565b92509250509295509295909350565b6000610b2d8383610b39565b60208301905092915050565b610b4281610ebd565b82525050565b610b5181610ebd565b82525050565b610b68610b6382610ebd565b611015565b82525050565b6000610b7982610df3565b610b838185610e21565b9350610b8e83610de3565b8060005b83811015610bbf578151610ba68882610b21565b9750610bb183610e14565b925050600181019050610b92565b5085935050505092915050565b6000610bd88385610e43565b9350610be5838584610fa2565b82840190509392505050565b6000610bfc82610e09565b610c068185610e32565b9350610c16818560208601610fb1565b610c1f816110ee565b840191505092915050565b610c3381610f90565b82525050565b610c4281610f27565b82525050565b6000610c548286610b57565b601482019150610c65828486610bcc565b9150819050949350505050565b6000602082019050610c876000830184610b48565b92915050565b6000604082019050610ca26000830185610b48565b8181036020830152610cb48184610bf1565b90509392505050565b6000604082019050610cd26000830185610b48565b610cdf6020830184610c39565b9392505050565b60006020820190508181036000830152610d008184610bf1565b905092915050565b60006040820190508181036000830152610d228185610bf1565b9050610d316020830184610c39565b9392505050565b600060a082019050610d4d6000830188610c39565b610d5a6020830187610c2a565b8181036040830152610d6c8186610b6e565b9050610d7b6060830185610b48565b610d886080830184610c39565b9695505050505050565b6000610d9c610dad565b9050610da88282610fe4565b919050565b6000604051905090565b600067ffffffffffffffff821115610dd257610dd1611097565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600082905092915050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60008085851115610e6257610e616110da565b5b83861115610e7357610e726110d5565b5b6001850283019150848603905094509492505050565b6000610e9482610f27565b9150610e9f83610f27565b925082821015610eb257610eb1611039565b5b828203905092915050565b6000610ec882610f07565b9050919050565b60008115159050919050565b60007fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610f3d8383610dfe565b82610f488135610edb565b92506014821015610f8857610f837fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008360140360080261110c565b831692505b505092915050565b6000610f9b82610f27565b9050919050565b82818337600083830152505050565b60005b83811015610fcf578082015181840152602081019050610fb4565b83811115610fde576000848401525b50505050565b610fed826110ee565b810181811067ffffffffffffffff8211171561100c5761100b611097565b5b80604052505050565b600061102082611027565b9050919050565b6000611032826110ff565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b600082821b905092915050565b61112281610ebd565b811461112d57600080fd5b50565b61113981610ecf565b811461114457600080fd5b50565b61115081610f27565b811461115b57600080fd5b5056fea264697066735822122077546c2f4741f7d79b7a2e01a3100f966ababfb11046c7872e4f310304d04ce564736f6c63430008070033",
}

// ZEVMSwapAppABI is the input ABI used to generate the binding from.
// Deprecated: Use ZEVMSwapAppMetaData.ABI instead.
var ZEVMSwapAppABI = ZEVMSwapAppMetaData.ABI

// ZEVMSwapAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZEVMSwapAppMetaData.Bin instead.
var ZEVMSwapAppBin = ZEVMSwapAppMetaData.Bin

// DeployZEVMSwapApp deploys a new Ethereum contract, binding an instance of ZEVMSwapApp to it.
func DeployZEVMSwapApp(auth *bind.TransactOpts, backend bind.ContractBackend, router02_ common.Address, systemContract_ common.Address) (common.Address, *types.Transaction, *ZEVMSwapApp, error) {
	parsed, err := ZEVMSwapAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZEVMSwapAppBin), backend, router02_, systemContract_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZEVMSwapApp{ZEVMSwapAppCaller: ZEVMSwapAppCaller{contract: contract}, ZEVMSwapAppTransactor: ZEVMSwapAppTransactor{contract: contract}, ZEVMSwapAppFilterer: ZEVMSwapAppFilterer{contract: contract}}, nil
}

// ZEVMSwapApp is an auto generated Go binding around an Ethereum contract.
type ZEVMSwapApp struct {
	ZEVMSwapAppCaller     // Read-only binding to the contract
	ZEVMSwapAppTransactor // Write-only binding to the contract
	ZEVMSwapAppFilterer   // Log filterer for contract events
}

// ZEVMSwapAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZEVMSwapAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZEVMSwapAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZEVMSwapAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZEVMSwapAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZEVMSwapAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZEVMSwapAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZEVMSwapAppSession struct {
	Contract     *ZEVMSwapApp      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZEVMSwapAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZEVMSwapAppCallerSession struct {
	Contract *ZEVMSwapAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZEVMSwapAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZEVMSwapAppTransactorSession struct {
	Contract     *ZEVMSwapAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZEVMSwapAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZEVMSwapAppRaw struct {
	Contract *ZEVMSwapApp // Generic contract binding to access the raw methods on
}

// ZEVMSwapAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZEVMSwapAppCallerRaw struct {
	Contract *ZEVMSwapAppCaller // Generic read-only contract binding to access the raw methods on
}

// ZEVMSwapAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZEVMSwapAppTransactorRaw struct {
	Contract *ZEVMSwapAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZEVMSwapApp creates a new instance of ZEVMSwapApp, bound to a specific deployed contract.
func NewZEVMSwapApp(address common.Address, backend bind.ContractBackend) (*ZEVMSwapApp, error) {
	contract, err := bindZEVMSwapApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZEVMSwapApp{ZEVMSwapAppCaller: ZEVMSwapAppCaller{contract: contract}, ZEVMSwapAppTransactor: ZEVMSwapAppTransactor{contract: contract}, ZEVMSwapAppFilterer: ZEVMSwapAppFilterer{contract: contract}}, nil
}

// NewZEVMSwapAppCaller creates a new read-only instance of ZEVMSwapApp, bound to a specific deployed contract.
func NewZEVMSwapAppCaller(address common.Address, caller bind.ContractCaller) (*ZEVMSwapAppCaller, error) {
	contract, err := bindZEVMSwapApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZEVMSwapAppCaller{contract: contract}, nil
}

// NewZEVMSwapAppTransactor creates a new write-only instance of ZEVMSwapApp, bound to a specific deployed contract.
func NewZEVMSwapAppTransactor(address common.Address, transactor bind.ContractTransactor) (*ZEVMSwapAppTransactor, error) {
	contract, err := bindZEVMSwapApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZEVMSwapAppTransactor{contract: contract}, nil
}

// NewZEVMSwapAppFilterer creates a new log filterer instance of ZEVMSwapApp, bound to a specific deployed contract.
func NewZEVMSwapAppFilterer(address common.Address, filterer bind.ContractFilterer) (*ZEVMSwapAppFilterer, error) {
	contract, err := bindZEVMSwapApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZEVMSwapAppFilterer{contract: contract}, nil
}

// bindZEVMSwapApp binds a generic wrapper to an already deployed contract.
func bindZEVMSwapApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZEVMSwapAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZEVMSwapApp *ZEVMSwapAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZEVMSwapApp.Contract.ZEVMSwapAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZEVMSwapApp *ZEVMSwapAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.ZEVMSwapAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZEVMSwapApp *ZEVMSwapAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.ZEVMSwapAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZEVMSwapApp *ZEVMSwapAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZEVMSwapApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZEVMSwapApp *ZEVMSwapAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZEVMSwapApp *ZEVMSwapAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.contract.Transact(opts, method, params...)
}

// DecodeMemo is a free data retrieval call binding the contract method 0xa06ea8bc.
//
// Solidity: function decodeMemo(bytes data) pure returns(address, bytes)
func (_ZEVMSwapApp *ZEVMSwapAppCaller) DecodeMemo(opts *bind.CallOpts, data []byte) (common.Address, []byte, error) {
	var out []interface{}
	err := _ZEVMSwapApp.contract.Call(opts, &out, "decodeMemo", data)

	if err != nil {
		return *new(common.Address), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// DecodeMemo is a free data retrieval call binding the contract method 0xa06ea8bc.
//
// Solidity: function decodeMemo(bytes data) pure returns(address, bytes)
func (_ZEVMSwapApp *ZEVMSwapAppSession) DecodeMemo(data []byte) (common.Address, []byte, error) {
	return _ZEVMSwapApp.Contract.DecodeMemo(&_ZEVMSwapApp.CallOpts, data)
}

// DecodeMemo is a free data retrieval call binding the contract method 0xa06ea8bc.
//
// Solidity: function decodeMemo(bytes data) pure returns(address, bytes)
func (_ZEVMSwapApp *ZEVMSwapAppCallerSession) DecodeMemo(data []byte) (common.Address, []byte, error) {
	return _ZEVMSwapApp.Contract.DecodeMemo(&_ZEVMSwapApp.CallOpts, data)
}

// EncodeMemo is a free data retrieval call binding the contract method 0xdf73044e.
//
// Solidity: function encodeMemo(address targetZRC20, bytes recipient) pure returns(bytes)
func (_ZEVMSwapApp *ZEVMSwapAppCaller) EncodeMemo(opts *bind.CallOpts, targetZRC20 common.Address, recipient []byte) ([]byte, error) {
	var out []interface{}
	err := _ZEVMSwapApp.contract.Call(opts, &out, "encodeMemo", targetZRC20, recipient)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeMemo is a free data retrieval call binding the contract method 0xdf73044e.
//
// Solidity: function encodeMemo(address targetZRC20, bytes recipient) pure returns(bytes)
func (_ZEVMSwapApp *ZEVMSwapAppSession) EncodeMemo(targetZRC20 common.Address, recipient []byte) ([]byte, error) {
	return _ZEVMSwapApp.Contract.EncodeMemo(&_ZEVMSwapApp.CallOpts, targetZRC20, recipient)
}

// EncodeMemo is a free data retrieval call binding the contract method 0xdf73044e.
//
// Solidity: function encodeMemo(address targetZRC20, bytes recipient) pure returns(bytes)
func (_ZEVMSwapApp *ZEVMSwapAppCallerSession) EncodeMemo(targetZRC20 common.Address, recipient []byte) ([]byte, error) {
	return _ZEVMSwapApp.Contract.EncodeMemo(&_ZEVMSwapApp.CallOpts, targetZRC20, recipient)
}

// Router02 is a free data retrieval call binding the contract method 0xbd00c9c4.
//
// Solidity: function router02() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppCaller) Router02(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZEVMSwapApp.contract.Call(opts, &out, "router02")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router02 is a free data retrieval call binding the contract method 0xbd00c9c4.
//
// Solidity: function router02() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppSession) Router02() (common.Address, error) {
	return _ZEVMSwapApp.Contract.Router02(&_ZEVMSwapApp.CallOpts)
}

// Router02 is a free data retrieval call binding the contract method 0xbd00c9c4.
//
// Solidity: function router02() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppCallerSession) Router02() (common.Address, error) {
	return _ZEVMSwapApp.Contract.Router02(&_ZEVMSwapApp.CallOpts)
}

// SystemContract is a free data retrieval call binding the contract method 0xbb88b769.
//
// Solidity: function systemContract() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppCaller) SystemContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZEVMSwapApp.contract.Call(opts, &out, "systemContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemContract is a free data retrieval call binding the contract method 0xbb88b769.
//
// Solidity: function systemContract() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppSession) SystemContract() (common.Address, error) {
	return _ZEVMSwapApp.Contract.SystemContract(&_ZEVMSwapApp.CallOpts)
}

// SystemContract is a free data retrieval call binding the contract method 0xbb88b769.
//
// Solidity: function systemContract() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppCallerSession) SystemContract() (common.Address, error) {
	return _ZEVMSwapApp.Contract.SystemContract(&_ZEVMSwapApp.CallOpts)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) , address zrc20, uint256 amount, bytes message) returns()
func (_ZEVMSwapApp *ZEVMSwapAppTransactor) OnCrossChainCall(opts *bind.TransactOpts, arg0 Context, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ZEVMSwapApp.contract.Transact(opts, "onCrossChainCall", arg0, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) , address zrc20, uint256 amount, bytes message) returns()
func (_ZEVMSwapApp *ZEVMSwapAppSession) OnCrossChainCall(arg0 Context, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.OnCrossChainCall(&_ZEVMSwapApp.TransactOpts, arg0, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) , address zrc20, uint256 amount, bytes message) returns()
func (_ZEVMSwapApp *ZEVMSwapAppTransactorSession) OnCrossChainCall(arg0 Context, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.OnCrossChainCall(&_ZEVMSwapApp.TransactOpts, arg0, zrc20, amount, message)
}
