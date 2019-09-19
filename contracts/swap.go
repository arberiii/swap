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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SwapABI is the input ABI used to generate the binding from.
const SwapABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_FTSE\",\"type\":\"uint256\"},{\"name\":\"_LIBOR\",\"type\":\"uint256\"},{\"name\":\"_IPFShash\",\"type\":\"bytes32\"}],\"name\":\"main\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"address_A_\",\"type\":\"address\"},{\"name\":\"address_B_\",\"type\":\"address\"},{\"name\":\"partyAname\",\"type\":\"string\"},{\"name\":\"partyBname\",\"type\":\"string\"},{\"name\":\"swap_amount\",\"type\":\"uint256\"},{\"name\":\"LIBOR_amount\",\"type\":\"uint256\"},{\"name\":\"FTSE_amount\",\"type\":\"uint256\"},{\"name\":\"jurisdiction_\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SwapBin is the compiled bytecode used for deploying new contracts.
const SwapBin = `0x608060405234801561001057600080fd5b5060405161094b38038061094b8339810160409081528151602083015191830151606080850151608086015160a087015160c088015160e089015196989586019693860195929491939092019060009081908061006b610356565b61007361037f565b61007b610356565b61008361037f565b8f97508e96508d95508c94508b8460000181815250508a84602001818152505060b4846060018181525050878360000190600160a060020a03169081600160a060020a0316815250508583602001819052506000836040018181525050888360600181905250838360a0018190525060008051602061092b833981519152600102836080019060001916908160001916815250508b8260000181815250508982604001818152505060b4826060018181525050868160000190600160a060020a03169081600160a060020a0316815250508481602001819052506000816040018181525050888160600181905250818160a0018190525060008051602061092b83398151915260010281608001906000191690816000191681525050826000808a600160a060020a0316600160a060020a0316815260200190815260200160002060008201518160000160006101000a815481600160a060020a030219169083600160a060020a03160217905550602082015181600101908051906020019061020d9291906103c8565b5060408201516002820155606082015180516102339160038401916020909101906103c8565b506080820151600482015560a0909101518051600583015560208082015160068401556040808301516007850155606090920151600890930192909255600160a060020a0389811660009081528084529190912083518154600160a060020a031916921691909117815582820151805184936102b69260018501929101906103c8565b5060408201516002820155606082015180516102dc9160038401916020909101906103c8565b506080820151600482015560a09091015180516005830155602081015160068301556040810151600783015560600151600890910155505060018054600160a060020a03978816600160a060020a0319918216179091556002805496909716951694909417909455506104639a5050505050505050505050565b608060405190810160405280600081526020016000815260200160008152602001600081525090565b610120604051908101604052806000600160a060020a03168152602001606081526020016000815260200160608152602001600080191681526020016103c3610356565b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061040957805160ff1916838001178555610436565b82800160010185558215610436579182015b8281111561043657825182559160200191906001019061041b565b50610442929150610446565b5090565b61046091905b80821115610442576000815560010161044c565b90565b6104b9806104726000396000f30060806040526004361061004a5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041662113e08811461004f57806365c046051461007d575b600080fd5b34801561005b57600080fd5b506100646100af565b6040805192835260208301919091528051918290030190f35b34801561008957600080fd5b5061009b6004356024356044356100e1565b604080519115158252519081900360200190f35b600154600160a060020a0390811660009081526020819052604080822060029081015481549094168352912001549091565b60006100eb61027f565b15156100f657600080fd5b60028054600160a060020a031660009081526020819052604090819020905160039091018054909282918491600019600183161561010002019091160480156101765780601f10610154576101008083540402835291820191610176565b820191906000526020600020905b815481529060010190602001808311610162575b50506040805191829003822060018054600160a060020a0316600090815260208190529290922060030180549195509350829184916002918116156101000260001901160480156101fe5780601f106101dc5761010080835404028352918201916101fe565b820191906000526020600020905b8154815290600101906020018083116101ea575b505091505060405180910390206000191614151561021b57600080fd5b600154600160a060020a0316600090815260208190526040902060040154821461024457600080fd5b600254600160a060020a0316600090815260208190526040902060040154821461026d57600080fd5b61027784846102b5565b509392505050565b60018054600160a060020a0390811660009081526020819052604080822060029081018390558054909316825281209091015590565b600254600160a060020a03166000908152602081905260408120600701548190819081908190819081908911156103e457600254600160a060020a0390811660009081526020819052604080822060070154600154909316825290206005810154600890910154918b0397509550935061016861271089870204850204925061271085870204915082821115610386575060018054600160a060020a039081166000908152602081905260408082206002908101805488880390810190915594549093168252902001549650610481565b828210156103ce575060028054600160a060020a0390811660009081526020819052604080822084018054868803908101909155845490931682529020909101549650610481565b828214156103df5760009650610481565b610481565b600254600160a060020a0390811660009081526020819052604080822060070154600154909316825290206005810154600890910154918b90039750955093506127108587020461016885612710888c02040281151561044057fe5b60028054600160a060020a03908116600090815260208190526040808220840180549690950496909601948501909355815416825292902090910154975092505b505050505050929150505600a165627a7a72305820b88c171bf81a4c2e56bc8aa8e45e2a6242bb98da85a7f18d506bd559a0a755f70029861ce807bbdcba4de666e2552ed170ea92adc4f35da9e97723e3a6ee374458d2`

// DeploySwap deploys a new Ethereum contract, binding an instance of Swap to it.
func DeploySwap(auth *bind.TransactOpts, backend bind.ContractBackend, address_A_ common.Address, address_B_ common.Address, partyAname string, partyBname string, swap_amount *big.Int, LIBOR_amount *big.Int, FTSE_amount *big.Int, jurisdiction_ string) (common.Address, *types.Transaction, *Swap, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwapBin), backend, address_A_, address_B_, partyAname, partyBname, swap_amount, LIBOR_amount, FTSE_amount, jurisdiction_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// GetBalances is a free data retrieval call binding the contract method 0x00113e08.
//
// Solidity: function getBalances() constant returns(uint256, uint256)
func (_Swap *SwapCaller) GetBalances(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Swap.contract.Call(opts, out, "getBalances")
	return *ret0, *ret1, err
}

// GetBalances is a free data retrieval call binding the contract method 0x00113e08.
//
// Solidity: function getBalances() constant returns(uint256, uint256)
func (_Swap *SwapSession) GetBalances() (*big.Int, *big.Int, error) {
	return _Swap.Contract.GetBalances(&_Swap.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x00113e08.
//
// Solidity: function getBalances() constant returns(uint256, uint256)
func (_Swap *SwapCallerSession) GetBalances() (*big.Int, *big.Int, error) {
	return _Swap.Contract.GetBalances(&_Swap.CallOpts)
}

// Main is a paid mutator transaction binding the contract method 0x65c04605.
//
// Solidity: function main(_FTSE uint256, _LIBOR uint256, _IPFShash bytes32) returns(bool)
func (_Swap *SwapTransactor) Main(opts *bind.TransactOpts, _FTSE *big.Int, _LIBOR *big.Int, _IPFShash [32]byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "main", _FTSE, _LIBOR, _IPFShash)
}

// Main is a paid mutator transaction binding the contract method 0x65c04605.
//
// Solidity: function main(_FTSE uint256, _LIBOR uint256, _IPFShash bytes32) returns(bool)
func (_Swap *SwapSession) Main(_FTSE *big.Int, _LIBOR *big.Int, _IPFShash [32]byte) (*types.Transaction, error) {
	return _Swap.Contract.Main(&_Swap.TransactOpts, _FTSE, _LIBOR, _IPFShash)
}

// Main is a paid mutator transaction binding the contract method 0x65c04605.
//
// Solidity: function main(_FTSE uint256, _LIBOR uint256, _IPFShash bytes32) returns(bool)
func (_Swap *SwapTransactorSession) Main(_FTSE *big.Int, _LIBOR *big.Int, _IPFShash [32]byte) (*types.Transaction, error) {
	return _Swap.Contract.Main(&_Swap.TransactOpts, _FTSE, _LIBOR, _IPFShash)
}
