// Code generated by counterfeiter. DO NOT EDIT.
package helpersfakes

import (
	"context"
	"database/sql"
	"sync"

	"code.cloudfoundry.org/bbs/db/sqldb/helpers"
)

type FakeTx struct {
	CommitStub        func() error
	commitMutex       sync.RWMutex
	commitArgsForCall []struct {
	}
	commitReturns struct {
		result1 error
	}
	commitReturnsOnCall map[int]struct {
		result1 error
	}
	ExecContextStub        func(context.Context, string, ...interface{}) (sql.Result, error)
	execContextMutex       sync.RWMutex
	execContextArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []interface{}
	}
	execContextReturns struct {
		result1 sql.Result
		result2 error
	}
	execContextReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	PrepareContextStub        func(context.Context, string) (*sql.Stmt, error)
	prepareContextMutex       sync.RWMutex
	prepareContextArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	prepareContextReturns struct {
		result1 *sql.Stmt
		result2 error
	}
	prepareContextReturnsOnCall map[int]struct {
		result1 *sql.Stmt
		result2 error
	}
	QueryContextStub        func(context.Context, string, ...interface{}) (*sql.Rows, error)
	queryContextMutex       sync.RWMutex
	queryContextArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []interface{}
	}
	queryContextReturns struct {
		result1 *sql.Rows
		result2 error
	}
	queryContextReturnsOnCall map[int]struct {
		result1 *sql.Rows
		result2 error
	}
	QueryRowContextStub        func(context.Context, string, ...interface{}) helpers.RowScanner
	queryRowContextMutex       sync.RWMutex
	queryRowContextArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []interface{}
	}
	queryRowContextReturns struct {
		result1 helpers.RowScanner
	}
	queryRowContextReturnsOnCall map[int]struct {
		result1 helpers.RowScanner
	}
	RollbackStub        func() error
	rollbackMutex       sync.RWMutex
	rollbackArgsForCall []struct {
	}
	rollbackReturns struct {
		result1 error
	}
	rollbackReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTx) Commit() error {
	fake.commitMutex.Lock()
	ret, specificReturn := fake.commitReturnsOnCall[len(fake.commitArgsForCall)]
	fake.commitArgsForCall = append(fake.commitArgsForCall, struct {
	}{})
	stub := fake.CommitStub
	fakeReturns := fake.commitReturns
	fake.recordInvocation("Commit", []interface{}{})
	fake.commitMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTx) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *FakeTx) CommitCalls(stub func() error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = stub
}

func (fake *FakeTx) CommitReturns(result1 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) CommitReturnsOnCall(i int, result1 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	if fake.commitReturnsOnCall == nil {
		fake.commitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.commitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) ExecContext(arg1 context.Context, arg2 string, arg3 ...interface{}) (sql.Result, error) {
	fake.execContextMutex.Lock()
	ret, specificReturn := fake.execContextReturnsOnCall[len(fake.execContextArgsForCall)]
	fake.execContextArgsForCall = append(fake.execContextArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []interface{}
	}{arg1, arg2, arg3})
	stub := fake.ExecContextStub
	fakeReturns := fake.execContextReturns
	fake.recordInvocation("ExecContext", []interface{}{arg1, arg2, arg3})
	fake.execContextMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTx) ExecContextCallCount() int {
	fake.execContextMutex.RLock()
	defer fake.execContextMutex.RUnlock()
	return len(fake.execContextArgsForCall)
}

func (fake *FakeTx) ExecContextCalls(stub func(context.Context, string, ...interface{}) (sql.Result, error)) {
	fake.execContextMutex.Lock()
	defer fake.execContextMutex.Unlock()
	fake.ExecContextStub = stub
}

func (fake *FakeTx) ExecContextArgsForCall(i int) (context.Context, string, []interface{}) {
	fake.execContextMutex.RLock()
	defer fake.execContextMutex.RUnlock()
	argsForCall := fake.execContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTx) ExecContextReturns(result1 sql.Result, result2 error) {
	fake.execContextMutex.Lock()
	defer fake.execContextMutex.Unlock()
	fake.ExecContextStub = nil
	fake.execContextReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeTx) ExecContextReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.execContextMutex.Lock()
	defer fake.execContextMutex.Unlock()
	fake.ExecContextStub = nil
	if fake.execContextReturnsOnCall == nil {
		fake.execContextReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.execContextReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeTx) PrepareContext(arg1 context.Context, arg2 string) (*sql.Stmt, error) {
	fake.prepareContextMutex.Lock()
	ret, specificReturn := fake.prepareContextReturnsOnCall[len(fake.prepareContextArgsForCall)]
	fake.prepareContextArgsForCall = append(fake.prepareContextArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.PrepareContextStub
	fakeReturns := fake.prepareContextReturns
	fake.recordInvocation("PrepareContext", []interface{}{arg1, arg2})
	fake.prepareContextMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTx) PrepareContextCallCount() int {
	fake.prepareContextMutex.RLock()
	defer fake.prepareContextMutex.RUnlock()
	return len(fake.prepareContextArgsForCall)
}

func (fake *FakeTx) PrepareContextCalls(stub func(context.Context, string) (*sql.Stmt, error)) {
	fake.prepareContextMutex.Lock()
	defer fake.prepareContextMutex.Unlock()
	fake.PrepareContextStub = stub
}

func (fake *FakeTx) PrepareContextArgsForCall(i int) (context.Context, string) {
	fake.prepareContextMutex.RLock()
	defer fake.prepareContextMutex.RUnlock()
	argsForCall := fake.prepareContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTx) PrepareContextReturns(result1 *sql.Stmt, result2 error) {
	fake.prepareContextMutex.Lock()
	defer fake.prepareContextMutex.Unlock()
	fake.PrepareContextStub = nil
	fake.prepareContextReturns = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeTx) PrepareContextReturnsOnCall(i int, result1 *sql.Stmt, result2 error) {
	fake.prepareContextMutex.Lock()
	defer fake.prepareContextMutex.Unlock()
	fake.PrepareContextStub = nil
	if fake.prepareContextReturnsOnCall == nil {
		fake.prepareContextReturnsOnCall = make(map[int]struct {
			result1 *sql.Stmt
			result2 error
		})
	}
	fake.prepareContextReturnsOnCall[i] = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeTx) QueryContext(arg1 context.Context, arg2 string, arg3 ...interface{}) (*sql.Rows, error) {
	fake.queryContextMutex.Lock()
	ret, specificReturn := fake.queryContextReturnsOnCall[len(fake.queryContextArgsForCall)]
	fake.queryContextArgsForCall = append(fake.queryContextArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []interface{}
	}{arg1, arg2, arg3})
	stub := fake.QueryContextStub
	fakeReturns := fake.queryContextReturns
	fake.recordInvocation("QueryContext", []interface{}{arg1, arg2, arg3})
	fake.queryContextMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTx) QueryContextCallCount() int {
	fake.queryContextMutex.RLock()
	defer fake.queryContextMutex.RUnlock()
	return len(fake.queryContextArgsForCall)
}

func (fake *FakeTx) QueryContextCalls(stub func(context.Context, string, ...interface{}) (*sql.Rows, error)) {
	fake.queryContextMutex.Lock()
	defer fake.queryContextMutex.Unlock()
	fake.QueryContextStub = stub
}

func (fake *FakeTx) QueryContextArgsForCall(i int) (context.Context, string, []interface{}) {
	fake.queryContextMutex.RLock()
	defer fake.queryContextMutex.RUnlock()
	argsForCall := fake.queryContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTx) QueryContextReturns(result1 *sql.Rows, result2 error) {
	fake.queryContextMutex.Lock()
	defer fake.queryContextMutex.Unlock()
	fake.QueryContextStub = nil
	fake.queryContextReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeTx) QueryContextReturnsOnCall(i int, result1 *sql.Rows, result2 error) {
	fake.queryContextMutex.Lock()
	defer fake.queryContextMutex.Unlock()
	fake.QueryContextStub = nil
	if fake.queryContextReturnsOnCall == nil {
		fake.queryContextReturnsOnCall = make(map[int]struct {
			result1 *sql.Rows
			result2 error
		})
	}
	fake.queryContextReturnsOnCall[i] = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeTx) QueryRowContext(arg1 context.Context, arg2 string, arg3 ...interface{}) helpers.RowScanner {
	fake.queryRowContextMutex.Lock()
	ret, specificReturn := fake.queryRowContextReturnsOnCall[len(fake.queryRowContextArgsForCall)]
	fake.queryRowContextArgsForCall = append(fake.queryRowContextArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []interface{}
	}{arg1, arg2, arg3})
	stub := fake.QueryRowContextStub
	fakeReturns := fake.queryRowContextReturns
	fake.recordInvocation("QueryRowContext", []interface{}{arg1, arg2, arg3})
	fake.queryRowContextMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTx) QueryRowContextCallCount() int {
	fake.queryRowContextMutex.RLock()
	defer fake.queryRowContextMutex.RUnlock()
	return len(fake.queryRowContextArgsForCall)
}

func (fake *FakeTx) QueryRowContextCalls(stub func(context.Context, string, ...interface{}) helpers.RowScanner) {
	fake.queryRowContextMutex.Lock()
	defer fake.queryRowContextMutex.Unlock()
	fake.QueryRowContextStub = stub
}

func (fake *FakeTx) QueryRowContextArgsForCall(i int) (context.Context, string, []interface{}) {
	fake.queryRowContextMutex.RLock()
	defer fake.queryRowContextMutex.RUnlock()
	argsForCall := fake.queryRowContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTx) QueryRowContextReturns(result1 helpers.RowScanner) {
	fake.queryRowContextMutex.Lock()
	defer fake.queryRowContextMutex.Unlock()
	fake.QueryRowContextStub = nil
	fake.queryRowContextReturns = struct {
		result1 helpers.RowScanner
	}{result1}
}

func (fake *FakeTx) QueryRowContextReturnsOnCall(i int, result1 helpers.RowScanner) {
	fake.queryRowContextMutex.Lock()
	defer fake.queryRowContextMutex.Unlock()
	fake.QueryRowContextStub = nil
	if fake.queryRowContextReturnsOnCall == nil {
		fake.queryRowContextReturnsOnCall = make(map[int]struct {
			result1 helpers.RowScanner
		})
	}
	fake.queryRowContextReturnsOnCall[i] = struct {
		result1 helpers.RowScanner
	}{result1}
}

func (fake *FakeTx) Rollback() error {
	fake.rollbackMutex.Lock()
	ret, specificReturn := fake.rollbackReturnsOnCall[len(fake.rollbackArgsForCall)]
	fake.rollbackArgsForCall = append(fake.rollbackArgsForCall, struct {
	}{})
	stub := fake.RollbackStub
	fakeReturns := fake.rollbackReturns
	fake.recordInvocation("Rollback", []interface{}{})
	fake.rollbackMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTx) RollbackCallCount() int {
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	return len(fake.rollbackArgsForCall)
}

func (fake *FakeTx) RollbackCalls(stub func() error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = stub
}

func (fake *FakeTx) RollbackReturns(result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	fake.rollbackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) RollbackReturnsOnCall(i int, result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	if fake.rollbackReturnsOnCall == nil {
		fake.rollbackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rollbackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	fake.execContextMutex.RLock()
	defer fake.execContextMutex.RUnlock()
	fake.prepareContextMutex.RLock()
	defer fake.prepareContextMutex.RUnlock()
	fake.queryContextMutex.RLock()
	defer fake.queryContextMutex.RUnlock()
	fake.queryRowContextMutex.RLock()
	defer fake.queryRowContextMutex.RUnlock()
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTx) recordInvocation(key string, args []interface{}) {
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

var _ helpers.Tx = new(FakeTx)