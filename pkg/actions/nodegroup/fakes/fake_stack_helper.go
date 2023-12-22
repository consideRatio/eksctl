// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/weaveworks/eksctl/pkg/actions/nodegroup"
	"github.com/weaveworks/eksctl/pkg/cfn/manager"
	"github.com/weaveworks/eksctl/pkg/utils/tasks"
)

type FakeStackHelper struct {
	ListNodeGroupStacksWithStatusesStub        func(context.Context) ([]manager.NodeGroupStack, error)
	listNodeGroupStacksWithStatusesMutex       sync.RWMutex
	listNodeGroupStacksWithStatusesArgsForCall []struct {
		arg1 context.Context
	}
	listNodeGroupStacksWithStatusesReturns struct {
		result1 []manager.NodeGroupStack
		result2 error
	}
	listNodeGroupStacksWithStatusesReturnsOnCall map[int]struct {
		result1 []manager.NodeGroupStack
		result2 error
	}
	NewTaskToDeleteUnownedNodeGroupStub        func(context.Context, string, string, manager.NodeGroupDeleter, *manager.DeleteWaitCondition) tasks.Task
	newTaskToDeleteUnownedNodeGroupMutex       sync.RWMutex
	newTaskToDeleteUnownedNodeGroupArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 manager.NodeGroupDeleter
		arg5 *manager.DeleteWaitCondition
	}
	newTaskToDeleteUnownedNodeGroupReturns struct {
		result1 tasks.Task
	}
	newTaskToDeleteUnownedNodeGroupReturnsOnCall map[int]struct {
		result1 tasks.Task
	}
	NewTasksToDeleteNodeGroupsStub        func([]manager.NodeGroupStack, func(_ string) bool, bool, func(chan error, string) error) (*tasks.TaskTree, error)
	newTasksToDeleteNodeGroupsMutex       sync.RWMutex
	newTasksToDeleteNodeGroupsArgsForCall []struct {
		arg1 []manager.NodeGroupStack
		arg2 func(_ string) bool
		arg3 bool
		arg4 func(chan error, string) error
	}
	newTasksToDeleteNodeGroupsReturns struct {
		result1 *tasks.TaskTree
		result2 error
	}
	newTasksToDeleteNodeGroupsReturnsOnCall map[int]struct {
		result1 *tasks.TaskTree
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStackHelper) ListNodeGroupStacksWithStatuses(arg1 context.Context) ([]manager.NodeGroupStack, error) {
	fake.listNodeGroupStacksWithStatusesMutex.Lock()
	ret, specificReturn := fake.listNodeGroupStacksWithStatusesReturnsOnCall[len(fake.listNodeGroupStacksWithStatusesArgsForCall)]
	fake.listNodeGroupStacksWithStatusesArgsForCall = append(fake.listNodeGroupStacksWithStatusesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListNodeGroupStacksWithStatusesStub
	fakeReturns := fake.listNodeGroupStacksWithStatusesReturns
	fake.recordInvocation("ListNodeGroupStacksWithStatuses", []interface{}{arg1})
	fake.listNodeGroupStacksWithStatusesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackHelper) ListNodeGroupStacksWithStatusesCallCount() int {
	fake.listNodeGroupStacksWithStatusesMutex.RLock()
	defer fake.listNodeGroupStacksWithStatusesMutex.RUnlock()
	return len(fake.listNodeGroupStacksWithStatusesArgsForCall)
}

func (fake *FakeStackHelper) ListNodeGroupStacksWithStatusesCalls(stub func(context.Context) ([]manager.NodeGroupStack, error)) {
	fake.listNodeGroupStacksWithStatusesMutex.Lock()
	defer fake.listNodeGroupStacksWithStatusesMutex.Unlock()
	fake.ListNodeGroupStacksWithStatusesStub = stub
}

func (fake *FakeStackHelper) ListNodeGroupStacksWithStatusesArgsForCall(i int) context.Context {
	fake.listNodeGroupStacksWithStatusesMutex.RLock()
	defer fake.listNodeGroupStacksWithStatusesMutex.RUnlock()
	argsForCall := fake.listNodeGroupStacksWithStatusesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStackHelper) ListNodeGroupStacksWithStatusesReturns(result1 []manager.NodeGroupStack, result2 error) {
	fake.listNodeGroupStacksWithStatusesMutex.Lock()
	defer fake.listNodeGroupStacksWithStatusesMutex.Unlock()
	fake.ListNodeGroupStacksWithStatusesStub = nil
	fake.listNodeGroupStacksWithStatusesReturns = struct {
		result1 []manager.NodeGroupStack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackHelper) ListNodeGroupStacksWithStatusesReturnsOnCall(i int, result1 []manager.NodeGroupStack, result2 error) {
	fake.listNodeGroupStacksWithStatusesMutex.Lock()
	defer fake.listNodeGroupStacksWithStatusesMutex.Unlock()
	fake.ListNodeGroupStacksWithStatusesStub = nil
	if fake.listNodeGroupStacksWithStatusesReturnsOnCall == nil {
		fake.listNodeGroupStacksWithStatusesReturnsOnCall = make(map[int]struct {
			result1 []manager.NodeGroupStack
			result2 error
		})
	}
	fake.listNodeGroupStacksWithStatusesReturnsOnCall[i] = struct {
		result1 []manager.NodeGroupStack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackHelper) NewTaskToDeleteUnownedNodeGroup(arg1 context.Context, arg2 string, arg3 string, arg4 manager.NodeGroupDeleter, arg5 *manager.DeleteWaitCondition) tasks.Task {
	fake.newTaskToDeleteUnownedNodeGroupMutex.Lock()
	ret, specificReturn := fake.newTaskToDeleteUnownedNodeGroupReturnsOnCall[len(fake.newTaskToDeleteUnownedNodeGroupArgsForCall)]
	fake.newTaskToDeleteUnownedNodeGroupArgsForCall = append(fake.newTaskToDeleteUnownedNodeGroupArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 manager.NodeGroupDeleter
		arg5 *manager.DeleteWaitCondition
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.NewTaskToDeleteUnownedNodeGroupStub
	fakeReturns := fake.newTaskToDeleteUnownedNodeGroupReturns
	fake.recordInvocation("NewTaskToDeleteUnownedNodeGroup", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.newTaskToDeleteUnownedNodeGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStackHelper) NewTaskToDeleteUnownedNodeGroupCallCount() int {
	fake.newTaskToDeleteUnownedNodeGroupMutex.RLock()
	defer fake.newTaskToDeleteUnownedNodeGroupMutex.RUnlock()
	return len(fake.newTaskToDeleteUnownedNodeGroupArgsForCall)
}

func (fake *FakeStackHelper) NewTaskToDeleteUnownedNodeGroupCalls(stub func(context.Context, string, string, manager.NodeGroupDeleter, *manager.DeleteWaitCondition) tasks.Task) {
	fake.newTaskToDeleteUnownedNodeGroupMutex.Lock()
	defer fake.newTaskToDeleteUnownedNodeGroupMutex.Unlock()
	fake.NewTaskToDeleteUnownedNodeGroupStub = stub
}

func (fake *FakeStackHelper) NewTaskToDeleteUnownedNodeGroupArgsForCall(i int) (context.Context, string, string, manager.NodeGroupDeleter, *manager.DeleteWaitCondition) {
	fake.newTaskToDeleteUnownedNodeGroupMutex.RLock()
	defer fake.newTaskToDeleteUnownedNodeGroupMutex.RUnlock()
	argsForCall := fake.newTaskToDeleteUnownedNodeGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeStackHelper) NewTaskToDeleteUnownedNodeGroupReturns(result1 tasks.Task) {
	fake.newTaskToDeleteUnownedNodeGroupMutex.Lock()
	defer fake.newTaskToDeleteUnownedNodeGroupMutex.Unlock()
	fake.NewTaskToDeleteUnownedNodeGroupStub = nil
	fake.newTaskToDeleteUnownedNodeGroupReturns = struct {
		result1 tasks.Task
	}{result1}
}

func (fake *FakeStackHelper) NewTaskToDeleteUnownedNodeGroupReturnsOnCall(i int, result1 tasks.Task) {
	fake.newTaskToDeleteUnownedNodeGroupMutex.Lock()
	defer fake.newTaskToDeleteUnownedNodeGroupMutex.Unlock()
	fake.NewTaskToDeleteUnownedNodeGroupStub = nil
	if fake.newTaskToDeleteUnownedNodeGroupReturnsOnCall == nil {
		fake.newTaskToDeleteUnownedNodeGroupReturnsOnCall = make(map[int]struct {
			result1 tasks.Task
		})
	}
	fake.newTaskToDeleteUnownedNodeGroupReturnsOnCall[i] = struct {
		result1 tasks.Task
	}{result1}
}

func (fake *FakeStackHelper) NewTasksToDeleteNodeGroups(arg1 []manager.NodeGroupStack, arg2 func(_ string) bool, arg3 bool, arg4 func(chan error, string) error) (*tasks.TaskTree, error) {
	var arg1Copy []manager.NodeGroupStack
	if arg1 != nil {
		arg1Copy = make([]manager.NodeGroupStack, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.newTasksToDeleteNodeGroupsMutex.Lock()
	ret, specificReturn := fake.newTasksToDeleteNodeGroupsReturnsOnCall[len(fake.newTasksToDeleteNodeGroupsArgsForCall)]
	fake.newTasksToDeleteNodeGroupsArgsForCall = append(fake.newTasksToDeleteNodeGroupsArgsForCall, struct {
		arg1 []manager.NodeGroupStack
		arg2 func(_ string) bool
		arg3 bool
		arg4 func(chan error, string) error
	}{arg1Copy, arg2, arg3, arg4})
	stub := fake.NewTasksToDeleteNodeGroupsStub
	fakeReturns := fake.newTasksToDeleteNodeGroupsReturns
	fake.recordInvocation("NewTasksToDeleteNodeGroups", []interface{}{arg1Copy, arg2, arg3, arg4})
	fake.newTasksToDeleteNodeGroupsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackHelper) NewTasksToDeleteNodeGroupsCallCount() int {
	fake.newTasksToDeleteNodeGroupsMutex.RLock()
	defer fake.newTasksToDeleteNodeGroupsMutex.RUnlock()
	return len(fake.newTasksToDeleteNodeGroupsArgsForCall)
}

func (fake *FakeStackHelper) NewTasksToDeleteNodeGroupsCalls(stub func([]manager.NodeGroupStack, func(_ string) bool, bool, func(chan error, string) error) (*tasks.TaskTree, error)) {
	fake.newTasksToDeleteNodeGroupsMutex.Lock()
	defer fake.newTasksToDeleteNodeGroupsMutex.Unlock()
	fake.NewTasksToDeleteNodeGroupsStub = stub
}

func (fake *FakeStackHelper) NewTasksToDeleteNodeGroupsArgsForCall(i int) ([]manager.NodeGroupStack, func(_ string) bool, bool, func(chan error, string) error) {
	fake.newTasksToDeleteNodeGroupsMutex.RLock()
	defer fake.newTasksToDeleteNodeGroupsMutex.RUnlock()
	argsForCall := fake.newTasksToDeleteNodeGroupsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeStackHelper) NewTasksToDeleteNodeGroupsReturns(result1 *tasks.TaskTree, result2 error) {
	fake.newTasksToDeleteNodeGroupsMutex.Lock()
	defer fake.newTasksToDeleteNodeGroupsMutex.Unlock()
	fake.NewTasksToDeleteNodeGroupsStub = nil
	fake.newTasksToDeleteNodeGroupsReturns = struct {
		result1 *tasks.TaskTree
		result2 error
	}{result1, result2}
}

func (fake *FakeStackHelper) NewTasksToDeleteNodeGroupsReturnsOnCall(i int, result1 *tasks.TaskTree, result2 error) {
	fake.newTasksToDeleteNodeGroupsMutex.Lock()
	defer fake.newTasksToDeleteNodeGroupsMutex.Unlock()
	fake.NewTasksToDeleteNodeGroupsStub = nil
	if fake.newTasksToDeleteNodeGroupsReturnsOnCall == nil {
		fake.newTasksToDeleteNodeGroupsReturnsOnCall = make(map[int]struct {
			result1 *tasks.TaskTree
			result2 error
		})
	}
	fake.newTasksToDeleteNodeGroupsReturnsOnCall[i] = struct {
		result1 *tasks.TaskTree
		result2 error
	}{result1, result2}
}

func (fake *FakeStackHelper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listNodeGroupStacksWithStatusesMutex.RLock()
	defer fake.listNodeGroupStacksWithStatusesMutex.RUnlock()
	fake.newTaskToDeleteUnownedNodeGroupMutex.RLock()
	defer fake.newTaskToDeleteUnownedNodeGroupMutex.RUnlock()
	fake.newTasksToDeleteNodeGroupsMutex.RLock()
	defer fake.newTasksToDeleteNodeGroupsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStackHelper) recordInvocation(key string, args []interface{}) {
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

var _ nodegroup.StackHelper = new(FakeStackHelper)
