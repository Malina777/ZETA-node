// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	core "github.com/ethereum/go-ethereum/core"

	evmtypes "github.com/evmos/ethermint/x/evm/types"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"

	vm "github.com/ethereum/go-ethereum/core/vm"
)

// FungibleEVMKeeper is an autogenerated mock type for the FungibleEVMKeeper type
type FungibleEVMKeeper struct {
	mock.Mock
}

// ApplyMessage provides a mock function with given fields: ctx, msg, tracer, commit
func (_m *FungibleEVMKeeper) ApplyMessage(ctx types.Context, msg core.Message, tracer vm.EVMLogger, commit bool) (*evmtypes.MsgEthereumTxResponse, error) {
	ret := _m.Called(ctx, msg, tracer, commit)

	var r0 *evmtypes.MsgEthereumTxResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, core.Message, vm.EVMLogger, bool) (*evmtypes.MsgEthereumTxResponse, error)); ok {
		return rf(ctx, msg, tracer, commit)
	}
	if rf, ok := ret.Get(0).(func(types.Context, core.Message, vm.EVMLogger, bool) *evmtypes.MsgEthereumTxResponse); ok {
		r0 = rf(ctx, msg, tracer, commit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*evmtypes.MsgEthereumTxResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, core.Message, vm.EVMLogger, bool) error); ok {
		r1 = rf(ctx, msg, tracer, commit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainID provides a mock function with given fields:
func (_m *FungibleEVMKeeper) ChainID() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// EstimateGas provides a mock function with given fields: c, req
func (_m *FungibleEVMKeeper) EstimateGas(c context.Context, req *evmtypes.EthCallRequest) (*evmtypes.EstimateGasResponse, error) {
	ret := _m.Called(c, req)

	var r0 *evmtypes.EstimateGasResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *evmtypes.EthCallRequest) (*evmtypes.EstimateGasResponse, error)); ok {
		return rf(c, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *evmtypes.EthCallRequest) *evmtypes.EstimateGasResponse); ok {
		r0 = rf(c, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*evmtypes.EstimateGasResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *evmtypes.EthCallRequest) error); ok {
		r1 = rf(c, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockBloomTransient provides a mock function with given fields: ctx
func (_m *FungibleEVMKeeper) GetBlockBloomTransient(ctx types.Context) *big.Int {
	ret := _m.Called(ctx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(types.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetLogSizeTransient provides a mock function with given fields: ctx
func (_m *FungibleEVMKeeper) GetLogSizeTransient(ctx types.Context) uint64 {
	ret := _m.Called(ctx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(types.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// SetBlockBloomTransient provides a mock function with given fields: ctx, bloom
func (_m *FungibleEVMKeeper) SetBlockBloomTransient(ctx types.Context, bloom *big.Int) {
	_m.Called(ctx, bloom)
}

// SetLogSizeTransient provides a mock function with given fields: ctx, logSize
func (_m *FungibleEVMKeeper) SetLogSizeTransient(ctx types.Context, logSize uint64) {
	_m.Called(ctx, logSize)
}

// WithChainID provides a mock function with given fields: ctx
func (_m *FungibleEVMKeeper) WithChainID(ctx types.Context) {
	_m.Called(ctx)
}

// NewFungibleEVMKeeper creates a new instance of FungibleEVMKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFungibleEVMKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *FungibleEVMKeeper {
	mock := &FungibleEVMKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
