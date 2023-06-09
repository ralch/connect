// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	connect "github.com/bufbuild/connect-go"
	"github.com/ralch/connect/internal/proto/connect/runtime/v1/runtimev1connect"
)

type FakeHealthServiceHandler struct {
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
	WatchStub        func(context.Context, *connect.Request[HealthCheckRequest], *connect.ServerStream[HealthCheckResponse]) error
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 context.Context
		arg2 *connect.Request[HealthCheckRequest]
		arg3 *connect.ServerStream[HealthCheckResponse]
	}
	watchReturns struct {
		result1 error
	}
	watchReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHealthServiceHandler) Check(arg1 context.Context, arg2 *connect.Request[HealthCheckRequest]) (*connect.Response[HealthCheckResponse], error) {
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

func (fake *FakeHealthServiceHandler) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeHealthServiceHandler) CheckCalls(stub func(context.Context, *connect.Request[HealthCheckRequest]) (*connect.Response[HealthCheckResponse], error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeHealthServiceHandler) CheckArgsForCall(i int) (context.Context, *connect.Request[HealthCheckRequest]) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHealthServiceHandler) CheckReturns(result1 *connect.Response[HealthCheckResponse], result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 *connect.Response[HealthCheckResponse]
		result2 error
	}{result1, result2}
}

func (fake *FakeHealthServiceHandler) CheckReturnsOnCall(i int, result1 *connect.Response[HealthCheckResponse], result2 error) {
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

func (fake *FakeHealthServiceHandler) Watch(arg1 context.Context, arg2 *connect.Request[HealthCheckRequest], arg3 *connect.ServerStream[HealthCheckResponse]) error {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 context.Context
		arg2 *connect.Request[HealthCheckRequest]
		arg3 *connect.ServerStream[HealthCheckResponse]
	}{arg1, arg2, arg3})
	stub := fake.WatchStub
	fakeReturns := fake.watchReturns
	fake.recordInvocation("Watch", []interface{}{arg1, arg2, arg3})
	fake.watchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHealthServiceHandler) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *FakeHealthServiceHandler) WatchCalls(stub func(context.Context, *connect.Request[HealthCheckRequest], *connect.ServerStream[HealthCheckResponse]) error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *FakeHealthServiceHandler) WatchArgsForCall(i int) (context.Context, *connect.Request[HealthCheckRequest], *connect.ServerStream[HealthCheckResponse]) {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeHealthServiceHandler) WatchReturns(result1 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHealthServiceHandler) WatchReturnsOnCall(i int, result1 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHealthServiceHandler) Invocations() map[string][][]interface{} {
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

func (fake *FakeHealthServiceHandler) recordInvocation(key string, args []interface{}) {
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

var _ runtimev1connect.HealthServiceHandler = new(FakeHealthServiceHandler)
