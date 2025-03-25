// Code generated by MockGen. DO NOT EDIT.
// Source: x/auth/ante/expected_keepers.go
//
// Generated by this command:
//
//	mockgen -source=x/auth/ante/expected_keepers.go -package testutil -destination x/auth/ante/testutil/expected_keepers_mocks.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"
	time "time"

	address "cosmossdk.io/core/address"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
	isgomock struct{}
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// AddressCodec mocks base method.
func (m *MockAccountKeeper) AddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// AddressCodec indicates an expected call of AddressCodec.
func (mr *MockAccountKeeperMockRecorder) AddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressCodec", reflect.TypeOf((*MockAccountKeeper)(nil).AddressCodec))
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx context.Context, addr types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(moduleName string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", moduleName)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(moduleName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), moduleName)
}

// GetParams mocks base method.
func (m *MockAccountKeeper) GetParams(ctx context.Context) types0.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types0.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockAccountKeeperMockRecorder) GetParams(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockAccountKeeper)(nil).GetParams), ctx)
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(ctx context.Context, acc types.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(ctx, acc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), ctx, acc)
}

// MockUnorderedNonceManager is a mock of UnorderedNonceManager interface.
type MockUnorderedNonceManager struct {
	ctrl     *gomock.Controller
	recorder *MockUnorderedNonceManagerMockRecorder
	isgomock struct{}
}

// MockUnorderedNonceManagerMockRecorder is the mock recorder for MockUnorderedNonceManager.
type MockUnorderedNonceManagerMockRecorder struct {
	mock *MockUnorderedNonceManager
}

// NewMockUnorderedNonceManager creates a new mock instance.
func NewMockUnorderedNonceManager(ctrl *gomock.Controller) *MockUnorderedNonceManager {
	mock := &MockUnorderedNonceManager{ctrl: ctrl}
	mock.recorder = &MockUnorderedNonceManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnorderedNonceManager) EXPECT() *MockUnorderedNonceManagerMockRecorder {
	return m.recorder
}

// RemoveExpiredUnorderedNonces mocks base method.
func (m *MockUnorderedNonceManager) RemoveExpiredUnorderedNonces(ctx types.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveExpiredUnorderedNonces", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveExpiredUnorderedNonces indicates an expected call of RemoveExpiredUnorderedNonces.
func (mr *MockUnorderedNonceManagerMockRecorder) RemoveExpiredUnorderedNonces(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveExpiredUnorderedNonces", reflect.TypeOf((*MockUnorderedNonceManager)(nil).RemoveExpiredUnorderedNonces), ctx)
}

// TryAddUnorderedNonce mocks base method.
func (m *MockUnorderedNonceManager) TryAddUnorderedNonce(ctx types.Context, sender []byte, timestamp time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryAddUnorderedNonce", ctx, sender, timestamp)
	ret0, _ := ret[0].(error)
	return ret0
}

// TryAddUnorderedNonce indicates an expected call of TryAddUnorderedNonce.
func (mr *MockUnorderedNonceManagerMockRecorder) TryAddUnorderedNonce(ctx, sender, timestamp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryAddUnorderedNonce", reflect.TypeOf((*MockUnorderedNonceManager)(nil).TryAddUnorderedNonce), ctx, sender, timestamp)
}

// MockFeegrantKeeper is a mock of FeegrantKeeper interface.
type MockFeegrantKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockFeegrantKeeperMockRecorder
	isgomock struct{}
}

// MockFeegrantKeeperMockRecorder is the mock recorder for MockFeegrantKeeper.
type MockFeegrantKeeperMockRecorder struct {
	mock *MockFeegrantKeeper
}

// NewMockFeegrantKeeper creates a new mock instance.
func NewMockFeegrantKeeper(ctrl *gomock.Controller) *MockFeegrantKeeper {
	mock := &MockFeegrantKeeper{ctrl: ctrl}
	mock.recorder = &MockFeegrantKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeegrantKeeper) EXPECT() *MockFeegrantKeeperMockRecorder {
	return m.recorder
}

// UseGrantedFees mocks base method.
func (m *MockFeegrantKeeper) UseGrantedFees(ctx context.Context, granter, grantee types.AccAddress, fee types.Coins, msgs []types.Msg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseGrantedFees", ctx, granter, grantee, fee, msgs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UseGrantedFees indicates an expected call of UseGrantedFees.
func (mr *MockFeegrantKeeperMockRecorder) UseGrantedFees(ctx, granter, grantee, fee, msgs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseGrantedFees", reflect.TypeOf((*MockFeegrantKeeper)(nil).UseGrantedFees), ctx, granter, grantee, fee, msgs)
}
