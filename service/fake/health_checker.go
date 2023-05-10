// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/ralch/connect/service"
)

type FakeHealthChecker struct {
	CheckStub        func(context.Context, *grpchealth.CheckRequest) (*grpchealth.CheckResponse, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 context.Context
		arg2 *grpchealth.CheckRequest
	}
	checkReturns struct {
		result1 *grpchealth.CheckResponse
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 *grpchealth.CheckResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHealthChecker) Check(arg1 context.Context, arg2 *grpchealth.CheckRequest) (*grpchealth.CheckResponse, error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 context.Context
		arg2 *grpchealth.CheckRequest
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

func (fake *FakeHealthChecker) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeHealthChecker) CheckCalls(stub func(context.Context, *grpchealth.CheckRequest) (*grpchealth.CheckResponse, error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeHealthChecker) CheckArgsForCall(i int) (context.Context, *grpchealth.CheckRequest) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHealthChecker) CheckReturns(result1 *grpchealth.CheckResponse, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 *grpchealth.CheckResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeHealthChecker) CheckReturnsOnCall(i int, result1 *grpchealth.CheckResponse, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 *grpchealth.CheckResponse
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 *grpchealth.CheckResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeHealthChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHealthChecker) recordInvocation(key string, args []interface{}) {
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

var _ service.HealthChecker = new(FakeHealthChecker)