// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	connect "github.com/bufbuild/connect-go"
	"github.com/ralch/connect/internal/proto/connect/runtime/v1/runtimev1connect"
)

type FakeMessageServiceClient struct {
	PushMessageStub        func(context.Context, *connect.Request[PushMessageRequest]) (*connect.Response[PushMessageResponse], error)
	pushMessageMutex       sync.RWMutex
	pushMessageArgsForCall []struct {
		arg1 context.Context
		arg2 *connect.Request[PushMessageRequest]
	}
	pushMessageReturns struct {
		result1 *connect.Response[PushMessageResponse]
		result2 error
	}
	pushMessageReturnsOnCall map[int]struct {
		result1 *connect.Response[PushMessageResponse]
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMessageServiceClient) PushMessage(arg1 context.Context, arg2 *connect.Request[PushMessageRequest]) (*connect.Response[PushMessageResponse], error) {
	fake.pushMessageMutex.Lock()
	ret, specificReturn := fake.pushMessageReturnsOnCall[len(fake.pushMessageArgsForCall)]
	fake.pushMessageArgsForCall = append(fake.pushMessageArgsForCall, struct {
		arg1 context.Context
		arg2 *connect.Request[PushMessageRequest]
	}{arg1, arg2})
	stub := fake.PushMessageStub
	fakeReturns := fake.pushMessageReturns
	fake.recordInvocation("PushMessage", []interface{}{arg1, arg2})
	fake.pushMessageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMessageServiceClient) PushMessageCallCount() int {
	fake.pushMessageMutex.RLock()
	defer fake.pushMessageMutex.RUnlock()
	return len(fake.pushMessageArgsForCall)
}

func (fake *FakeMessageServiceClient) PushMessageCalls(stub func(context.Context, *connect.Request[PushMessageRequest]) (*connect.Response[PushMessageResponse], error)) {
	fake.pushMessageMutex.Lock()
	defer fake.pushMessageMutex.Unlock()
	fake.PushMessageStub = stub
}

func (fake *FakeMessageServiceClient) PushMessageArgsForCall(i int) (context.Context, *connect.Request[PushMessageRequest]) {
	fake.pushMessageMutex.RLock()
	defer fake.pushMessageMutex.RUnlock()
	argsForCall := fake.pushMessageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMessageServiceClient) PushMessageReturns(result1 *connect.Response[PushMessageResponse], result2 error) {
	fake.pushMessageMutex.Lock()
	defer fake.pushMessageMutex.Unlock()
	fake.PushMessageStub = nil
	fake.pushMessageReturns = struct {
		result1 *connect.Response[PushMessageResponse]
		result2 error
	}{result1, result2}
}

func (fake *FakeMessageServiceClient) PushMessageReturnsOnCall(i int, result1 *connect.Response[PushMessageResponse], result2 error) {
	fake.pushMessageMutex.Lock()
	defer fake.pushMessageMutex.Unlock()
	fake.PushMessageStub = nil
	if fake.pushMessageReturnsOnCall == nil {
		fake.pushMessageReturnsOnCall = make(map[int]struct {
			result1 *connect.Response[PushMessageResponse]
			result2 error
		})
	}
	fake.pushMessageReturnsOnCall[i] = struct {
		result1 *connect.Response[PushMessageResponse]
		result2 error
	}{result1, result2}
}

func (fake *FakeMessageServiceClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pushMessageMutex.RLock()
	defer fake.pushMessageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMessageServiceClient) recordInvocation(key string, args []interface{}) {
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

var _ runtimev1connect.MessageServiceClient = new(FakeMessageServiceClient)
