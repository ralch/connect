// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	connect "github.com/bufbuild/connect-go"
	"github.com/ralch/connect/internal/proto/connect/runtime/v1/runtimev1connect"
)

type FakeHealthServiceClient struct {
	CheckStub        func(context.Context, *connect.Request[HealthCheckRequest]) (*connect.Response[HealthCheckResponse], error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 context.Context
		arg2 *connect.Request[HealthCheckRequest]
	}
	checkReturns struct {
		result1 *connect.Response[HealthCheckResponse]
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 *connect.Response[HealthCheckResponse]
		result2 error
	}
	WatchStub        func(context.Context, *connect.Request[HealthCheckRequest]) (*connect.ServerStreamForClient[HealthCheckResponse], error)
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 context.Context
		arg2 *connect.Request[HealthCheckRequest]
	}
	watchReturns struct {
		result1 *connect.ServerStreamForClient[HealthCheckResponse]
		result2 error
	}
	watchReturnsOnCall map[int]struct {
		result1 *connect.ServerStreamForClient[HealthCheckResponse]
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHealthServiceClient) Check(arg1 context.Context, arg2 *connect.Request[HealthCheckRequest]) (*connect.Response[HealthCheckResponse], error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 context.Context
		arg2 *connect.Request[HealthCheckRequest]
	}{arg1, arg2})
	stub := fake.CheckStub
	fakeReturns := fake.checkReturns
	fake.recordInvocation("Check", []interface{}{arg1, arg2})
	fake.checkMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHealthServiceClient) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeHealthServiceClient) CheckCalls(stub func(context.Context, *connect.Request[HealthCheckRequest]) (*connect.Response[HealthCheckResponse], error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeHealthServiceClient) CheckArgsForCall(i int) (context.Context, *connect.Request[HealthCheckRequest]) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHealthServiceClient) CheckReturns(result1 *connect.Response[HealthCheckResponse], result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 *connect.Response[HealthCheckResponse]
		result2 error
	}{result1, result2}
}

func (fake *FakeHealthServiceClient) CheckReturnsOnCall(i int, result1 *connect.Response[HealthCheckResponse], result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 *connect.Response[HealthCheckResponse]
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 *connect.Response[HealthCheckResponse]
		result2 error
	}{result1, result2}
}

func (fake *FakeHealthServiceClient) Watch(arg1 context.Context, arg2 *connect.Request[HealthCheckRequest]) (*connect.ServerStreamForClient[HealthCheckResponse], error) {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 context.Context
		arg2 *connect.Request[HealthCheckRequest]
	}{arg1, arg2})
	stub := fake.WatchStub
	fakeReturns := fake.watchReturns
	fake.recordInvocation("Watch", []interface{}{arg1, arg2})
	fake.watchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHealthServiceClient) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *FakeHealthServiceClient) WatchCalls(stub func(context.Context, *connect.Request[HealthCheckRequest]) (*connect.ServerStreamForClient[HealthCheckResponse], error)) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *FakeHealthServiceClient) WatchArgsForCall(i int) (context.Context, *connect.Request[HealthCheckRequest]) {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHealthServiceClient) WatchReturns(result1 *connect.ServerStreamForClient[HealthCheckResponse], result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 *connect.ServerStreamForClient[HealthCheckResponse]
		result2 error
	}{result1, result2}
}

func (fake *FakeHealthServiceClient) WatchReturnsOnCall(i int, result1 *connect.ServerStreamForClient[HealthCheckResponse], result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 *connect.ServerStreamForClient[HealthCheckResponse]
			result2 error
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 *connect.ServerStreamForClient[HealthCheckResponse]
		result2 error
	}{result1, result2}
}

func (fake *FakeHealthServiceClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHealthServiceClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ runtimev1connect.HealthServiceClient = new(FakeHealthServiceClient)
