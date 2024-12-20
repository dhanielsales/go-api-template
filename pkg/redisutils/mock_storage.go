// Code generated by MockGen. DO NOT EDIT.
// Source: ./storage.go
//
// Generated by this command:
//
//	mockgen -source ./storage.go -destination ./mock_storage.go -package redisutils
//

// Package redisutils is a generated GoMock package.
package redisutils

import (
	context "context"
	reflect "reflect"
	time "time"

	redis "github.com/redis/go-redis/v9"
	gomock "go.uber.org/mock/gomock"
)

// MockRedisClient is a mock of RedisClient interface.
type MockRedisClient struct {
	ctrl     *gomock.Controller
	recorder *MockRedisClientMockRecorder
	isgomock struct{}
}

// MockRedisClientMockRecorder is the mock recorder for MockRedisClient.
type MockRedisClientMockRecorder struct {
	mock *MockRedisClient
}

// NewMockRedisClient creates a new mock instance.
func NewMockRedisClient(ctrl *gomock.Controller) *MockRedisClient {
	mock := &MockRedisClient{ctrl: ctrl}
	mock.recorder = &MockRedisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisClient) EXPECT() *MockRedisClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRedisClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRedisClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedisClient)(nil).Close))
}

// Del mocks base method.
func (m *MockRedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockRedisClientMockRecorder) Del(ctx any, keys ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedisClient)(nil).Del), varargs...)
}

// Get mocks base method.
func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockRedisClientMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisClient)(nil).Get), ctx, key)
}

// Ping mocks base method.
func (m *MockRedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockRedisClientMockRecorder) Ping(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRedisClient)(nil).Ping), ctx)
}

// Scan mocks base method.
func (m *MockRedisClient) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", ctx, cursor, match, count)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockRedisClientMockRecorder) Scan(ctx, cursor, match, count any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRedisClient)(nil).Scan), ctx, cursor, match, count)
}

// Set mocks base method.
func (m *MockRedisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisClientMockRecorder) Set(ctx, key, value, expiration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisClient)(nil).Set), ctx, key, value, expiration)
}

// TxPipelined mocks base method.
func (m *MockRedisClient) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPipelined", ctx, fn)
	ret0, _ := ret[0].([]redis.Cmder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxPipelined indicates an expected call of TxPipelined.
func (mr *MockRedisClientMockRecorder) TxPipelined(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPipelined", reflect.TypeOf((*MockRedisClient)(nil).TxPipelined), ctx, fn)
}

// Watch mocks base method.
func (m *MockRedisClient) Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, fn}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch.
func (mr *MockRedisClientMockRecorder) Watch(ctx, fn any, keys ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, fn}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockRedisClient)(nil).Watch), varargs...)
}
