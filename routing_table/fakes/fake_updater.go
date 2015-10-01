// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/cf-tcp-router/routing_table"
	"github.com/cloudfoundry-incubator/routing-api"
)

type FakeUpdater struct {
	HandleEventStub        func(event routing_api.TcpEvent) error
	handleEventMutex       sync.RWMutex
	handleEventArgsForCall []struct {
		event routing_api.TcpEvent
	}
	handleEventReturns struct {
		result1 error
	}
}

func (fake *FakeUpdater) HandleEvent(event routing_api.TcpEvent) error {
	fake.handleEventMutex.Lock()
	fake.handleEventArgsForCall = append(fake.handleEventArgsForCall, struct {
		event routing_api.TcpEvent
	}{event})
	fake.handleEventMutex.Unlock()
	if fake.HandleEventStub != nil {
		return fake.HandleEventStub(event)
	} else {
		return fake.handleEventReturns.result1
	}
}

func (fake *FakeUpdater) HandleEventCallCount() int {
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	return len(fake.handleEventArgsForCall)
}

func (fake *FakeUpdater) HandleEventArgsForCall(i int) routing_api.TcpEvent {
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	return fake.handleEventArgsForCall[i].event
}

func (fake *FakeUpdater) HandleEventReturns(result1 error) {
	fake.HandleEventStub = nil
	fake.handleEventReturns = struct {
		result1 error
	}{result1}
}

var _ routing_table.Updater = new(FakeUpdater)
