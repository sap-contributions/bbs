// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/bbs"
	"github.com/cloudfoundry-incubator/bbs/events"
	"github.com/cloudfoundry-incubator/bbs/models"
)

type FakeClient struct {
	DomainsStub        func() ([]string, error)
	domainsMutex       sync.RWMutex
	domainsArgsForCall []struct{}
	domainsReturns struct {
		result1 []string
		result2 error
	}
	UpsertDomainStub        func(domain string, ttl time.Duration) error
	upsertDomainMutex       sync.RWMutex
	upsertDomainArgsForCall []struct {
		domain string
		ttl    time.Duration
	}
	upsertDomainReturns struct {
		result1 error
	}
	ActualLRPGroupsStub        func(models.ActualLRPFilter) ([]*models.ActualLRPGroup, error)
	actualLRPGroupsMutex       sync.RWMutex
	actualLRPGroupsArgsForCall []struct {
		arg1 models.ActualLRPFilter
	}
	actualLRPGroupsReturns struct {
		result1 []*models.ActualLRPGroup
		result2 error
	}
	ActualLRPGroupsByProcessGuidStub        func(processGuid string) ([]*models.ActualLRPGroup, error)
	actualLRPGroupsByProcessGuidMutex       sync.RWMutex
	actualLRPGroupsByProcessGuidArgsForCall []struct {
		processGuid string
	}
	actualLRPGroupsByProcessGuidReturns struct {
		result1 []*models.ActualLRPGroup
		result2 error
	}
	ActualLRPGroupByProcessGuidAndIndexStub        func(processGuid string, index int) (*models.ActualLRPGroup, error)
	actualLRPGroupByProcessGuidAndIndexMutex       sync.RWMutex
	actualLRPGroupByProcessGuidAndIndexArgsForCall []struct {
		processGuid string
		index       int
	}
	actualLRPGroupByProcessGuidAndIndexReturns struct {
		result1 *models.ActualLRPGroup
		result2 error
	}
	ClaimActualLRPStub        func(processGuid string, index int, instanceKey *models.ActualLRPInstanceKey) (*models.ActualLRP, error)
	claimActualLRPMutex       sync.RWMutex
	claimActualLRPArgsForCall []struct {
		processGuid string
		index       int
		instanceKey *models.ActualLRPInstanceKey
	}
	claimActualLRPReturns struct {
		result1 *models.ActualLRP
		result2 error
	}
	StartActualLRPStub        func(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, netInfo *models.ActualLRPNetInfo) (*models.ActualLRP, error)
	startActualLRPMutex       sync.RWMutex
	startActualLRPArgsForCall []struct {
		key         *models.ActualLRPKey
		instanceKey *models.ActualLRPInstanceKey
		netInfo     *models.ActualLRPNetInfo
	}
	startActualLRPReturns struct {
		result1 *models.ActualLRP
		result2 error
	}
	CrashActualLRPStub        func(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, errorMessage string) error
	crashActualLRPMutex       sync.RWMutex
	crashActualLRPArgsForCall []struct {
		key          *models.ActualLRPKey
		instanceKey  *models.ActualLRPInstanceKey
		errorMessage string
	}
	crashActualLRPReturns struct {
		result1 error
	}
	FailActualLRPStub        func(key *models.ActualLRPKey, errorMessage string) error
	failActualLRPMutex       sync.RWMutex
	failActualLRPArgsForCall []struct {
		key          *models.ActualLRPKey
		errorMessage string
	}
	failActualLRPReturns struct {
		result1 error
	}
	RemoveActualLRPStub        func(processGuid string, index int) error
	removeActualLRPMutex       sync.RWMutex
	removeActualLRPArgsForCall []struct {
		processGuid string
		index       int
	}
	removeActualLRPReturns struct {
		result1 error
	}
	DesiredLRPsStub        func(models.DesiredLRPFilter) ([]*models.DesiredLRP, error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct {
		arg1 models.DesiredLRPFilter
	}
	desiredLRPsReturns struct {
		result1 []*models.DesiredLRP
		result2 error
	}
	DesiredLRPByProcessGuidStub        func(processGuid string) (*models.DesiredLRP, error)
	desiredLRPByProcessGuidMutex       sync.RWMutex
	desiredLRPByProcessGuidArgsForCall []struct {
		processGuid string
	}
	desiredLRPByProcessGuidReturns struct {
		result1 *models.DesiredLRP
		result2 error
	}
	TasksStub        func() ([]*models.Task, error)
	tasksMutex       sync.RWMutex
	tasksArgsForCall []struct{}
	tasksReturns struct {
		result1 []*models.Task
		result2 error
	}
	TasksByDomainStub        func(domain string) ([]*models.Task, error)
	tasksByDomainMutex       sync.RWMutex
	tasksByDomainArgsForCall []struct {
		domain string
	}
	tasksByDomainReturns struct {
		result1 []*models.Task
		result2 error
	}
	TasksByCellIDStub        func(cellId string) ([]*models.Task, error)
	tasksByCellIDMutex       sync.RWMutex
	tasksByCellIDArgsForCall []struct {
		cellId string
	}
	tasksByCellIDReturns struct {
		result1 []*models.Task
		result2 error
	}
	TaskByGuidStub        func(guid string) (*models.Task, error)
	taskByGuidMutex       sync.RWMutex
	taskByGuidArgsForCall []struct {
		guid string
	}
	taskByGuidReturns struct {
		result1 *models.Task
		result2 error
	}
	SubscribeToEventsStub        func() (events.EventSource, error)
	subscribeToEventsMutex       sync.RWMutex
	subscribeToEventsArgsForCall []struct{}
	subscribeToEventsReturns struct {
		result1 events.EventSource
		result2 error
	}
}

func (fake *FakeClient) Domains() ([]string, error) {
	fake.domainsMutex.Lock()
	fake.domainsArgsForCall = append(fake.domainsArgsForCall, struct{}{})
	fake.domainsMutex.Unlock()
	if fake.DomainsStub != nil {
		return fake.DomainsStub()
	} else {
		return fake.domainsReturns.result1, fake.domainsReturns.result2
	}
}

func (fake *FakeClient) DomainsCallCount() int {
	fake.domainsMutex.RLock()
	defer fake.domainsMutex.RUnlock()
	return len(fake.domainsArgsForCall)
}

func (fake *FakeClient) DomainsReturns(result1 []string, result2 error) {
	fake.DomainsStub = nil
	fake.domainsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpsertDomain(domain string, ttl time.Duration) error {
	fake.upsertDomainMutex.Lock()
	fake.upsertDomainArgsForCall = append(fake.upsertDomainArgsForCall, struct {
		domain string
		ttl    time.Duration
	}{domain, ttl})
	fake.upsertDomainMutex.Unlock()
	if fake.UpsertDomainStub != nil {
		return fake.UpsertDomainStub(domain, ttl)
	} else {
		return fake.upsertDomainReturns.result1
	}
}

func (fake *FakeClient) UpsertDomainCallCount() int {
	fake.upsertDomainMutex.RLock()
	defer fake.upsertDomainMutex.RUnlock()
	return len(fake.upsertDomainArgsForCall)
}

func (fake *FakeClient) UpsertDomainArgsForCall(i int) (string, time.Duration) {
	fake.upsertDomainMutex.RLock()
	defer fake.upsertDomainMutex.RUnlock()
	return fake.upsertDomainArgsForCall[i].domain, fake.upsertDomainArgsForCall[i].ttl
}

func (fake *FakeClient) UpsertDomainReturns(result1 error) {
	fake.UpsertDomainStub = nil
	fake.upsertDomainReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) ActualLRPGroups(arg1 models.ActualLRPFilter) ([]*models.ActualLRPGroup, error) {
	fake.actualLRPGroupsMutex.Lock()
	fake.actualLRPGroupsArgsForCall = append(fake.actualLRPGroupsArgsForCall, struct {
		arg1 models.ActualLRPFilter
	}{arg1})
	fake.actualLRPGroupsMutex.Unlock()
	if fake.ActualLRPGroupsStub != nil {
		return fake.ActualLRPGroupsStub(arg1)
	} else {
		return fake.actualLRPGroupsReturns.result1, fake.actualLRPGroupsReturns.result2
	}
}

func (fake *FakeClient) ActualLRPGroupsCallCount() int {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return len(fake.actualLRPGroupsArgsForCall)
}

func (fake *FakeClient) ActualLRPGroupsArgsForCall(i int) models.ActualLRPFilter {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return fake.actualLRPGroupsArgsForCall[i].arg1
}

func (fake *FakeClient) ActualLRPGroupsReturns(result1 []*models.ActualLRPGroup, result2 error) {
	fake.ActualLRPGroupsStub = nil
	fake.actualLRPGroupsReturns = struct {
		result1 []*models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ActualLRPGroupsByProcessGuid(processGuid string) ([]*models.ActualLRPGroup, error) {
	fake.actualLRPGroupsByProcessGuidMutex.Lock()
	fake.actualLRPGroupsByProcessGuidArgsForCall = append(fake.actualLRPGroupsByProcessGuidArgsForCall, struct {
		processGuid string
	}{processGuid})
	fake.actualLRPGroupsByProcessGuidMutex.Unlock()
	if fake.ActualLRPGroupsByProcessGuidStub != nil {
		return fake.ActualLRPGroupsByProcessGuidStub(processGuid)
	} else {
		return fake.actualLRPGroupsByProcessGuidReturns.result1, fake.actualLRPGroupsByProcessGuidReturns.result2
	}
}

func (fake *FakeClient) ActualLRPGroupsByProcessGuidCallCount() int {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return len(fake.actualLRPGroupsByProcessGuidArgsForCall)
}

func (fake *FakeClient) ActualLRPGroupsByProcessGuidArgsForCall(i int) string {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return fake.actualLRPGroupsByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeClient) ActualLRPGroupsByProcessGuidReturns(result1 []*models.ActualLRPGroup, result2 error) {
	fake.ActualLRPGroupsByProcessGuidStub = nil
	fake.actualLRPGroupsByProcessGuidReturns = struct {
		result1 []*models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ActualLRPGroupByProcessGuidAndIndex(processGuid string, index int) (*models.ActualLRPGroup, error) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Lock()
	fake.actualLRPGroupByProcessGuidAndIndexArgsForCall = append(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall, struct {
		processGuid string
		index       int
	}{processGuid, index})
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Unlock()
	if fake.ActualLRPGroupByProcessGuidAndIndexStub != nil {
		return fake.ActualLRPGroupByProcessGuidAndIndexStub(processGuid, index)
	} else {
		return fake.actualLRPGroupByProcessGuidAndIndexReturns.result1, fake.actualLRPGroupByProcessGuidAndIndexReturns.result2
	}
}

func (fake *FakeClient) ActualLRPGroupByProcessGuidAndIndexCallCount() int {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return len(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall)
}

func (fake *FakeClient) ActualLRPGroupByProcessGuidAndIndexArgsForCall(i int) (string, int) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].processGuid, fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].index
}

func (fake *FakeClient) ActualLRPGroupByProcessGuidAndIndexReturns(result1 *models.ActualLRPGroup, result2 error) {
	fake.ActualLRPGroupByProcessGuidAndIndexStub = nil
	fake.actualLRPGroupByProcessGuidAndIndexReturns = struct {
		result1 *models.ActualLRPGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ClaimActualLRP(processGuid string, index int, instanceKey *models.ActualLRPInstanceKey) (*models.ActualLRP, error) {
	fake.claimActualLRPMutex.Lock()
	fake.claimActualLRPArgsForCall = append(fake.claimActualLRPArgsForCall, struct {
		processGuid string
		index       int
		instanceKey *models.ActualLRPInstanceKey
	}{processGuid, index, instanceKey})
	fake.claimActualLRPMutex.Unlock()
	if fake.ClaimActualLRPStub != nil {
		return fake.ClaimActualLRPStub(processGuid, index, instanceKey)
	} else {
		return fake.claimActualLRPReturns.result1, fake.claimActualLRPReturns.result2
	}
}

func (fake *FakeClient) ClaimActualLRPCallCount() int {
	fake.claimActualLRPMutex.RLock()
	defer fake.claimActualLRPMutex.RUnlock()
	return len(fake.claimActualLRPArgsForCall)
}

func (fake *FakeClient) ClaimActualLRPArgsForCall(i int) (string, int, *models.ActualLRPInstanceKey) {
	fake.claimActualLRPMutex.RLock()
	defer fake.claimActualLRPMutex.RUnlock()
	return fake.claimActualLRPArgsForCall[i].processGuid, fake.claimActualLRPArgsForCall[i].index, fake.claimActualLRPArgsForCall[i].instanceKey
}

func (fake *FakeClient) ClaimActualLRPReturns(result1 *models.ActualLRP, result2 error) {
	fake.ClaimActualLRPStub = nil
	fake.claimActualLRPReturns = struct {
		result1 *models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) StartActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, netInfo *models.ActualLRPNetInfo) (*models.ActualLRP, error) {
	fake.startActualLRPMutex.Lock()
	fake.startActualLRPArgsForCall = append(fake.startActualLRPArgsForCall, struct {
		key         *models.ActualLRPKey
		instanceKey *models.ActualLRPInstanceKey
		netInfo     *models.ActualLRPNetInfo
	}{key, instanceKey, netInfo})
	fake.startActualLRPMutex.Unlock()
	if fake.StartActualLRPStub != nil {
		return fake.StartActualLRPStub(key, instanceKey, netInfo)
	} else {
		return fake.startActualLRPReturns.result1, fake.startActualLRPReturns.result2
	}
}

func (fake *FakeClient) StartActualLRPCallCount() int {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	return len(fake.startActualLRPArgsForCall)
}

func (fake *FakeClient) StartActualLRPArgsForCall(i int) (*models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo) {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	return fake.startActualLRPArgsForCall[i].key, fake.startActualLRPArgsForCall[i].instanceKey, fake.startActualLRPArgsForCall[i].netInfo
}

func (fake *FakeClient) StartActualLRPReturns(result1 *models.ActualLRP, result2 error) {
	fake.StartActualLRPStub = nil
	fake.startActualLRPReturns = struct {
		result1 *models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CrashActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, errorMessage string) error {
	fake.crashActualLRPMutex.Lock()
	fake.crashActualLRPArgsForCall = append(fake.crashActualLRPArgsForCall, struct {
		key          *models.ActualLRPKey
		instanceKey  *models.ActualLRPInstanceKey
		errorMessage string
	}{key, instanceKey, errorMessage})
	fake.crashActualLRPMutex.Unlock()
	if fake.CrashActualLRPStub != nil {
		return fake.CrashActualLRPStub(key, instanceKey, errorMessage)
	} else {
		return fake.crashActualLRPReturns.result1
	}
}

func (fake *FakeClient) CrashActualLRPCallCount() int {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	return len(fake.crashActualLRPArgsForCall)
}

func (fake *FakeClient) CrashActualLRPArgsForCall(i int) (*models.ActualLRPKey, *models.ActualLRPInstanceKey, string) {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	return fake.crashActualLRPArgsForCall[i].key, fake.crashActualLRPArgsForCall[i].instanceKey, fake.crashActualLRPArgsForCall[i].errorMessage
}

func (fake *FakeClient) CrashActualLRPReturns(result1 error) {
	fake.CrashActualLRPStub = nil
	fake.crashActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) FailActualLRP(key *models.ActualLRPKey, errorMessage string) error {
	fake.failActualLRPMutex.Lock()
	fake.failActualLRPArgsForCall = append(fake.failActualLRPArgsForCall, struct {
		key          *models.ActualLRPKey
		errorMessage string
	}{key, errorMessage})
	fake.failActualLRPMutex.Unlock()
	if fake.FailActualLRPStub != nil {
		return fake.FailActualLRPStub(key, errorMessage)
	} else {
		return fake.failActualLRPReturns.result1
	}
}

func (fake *FakeClient) FailActualLRPCallCount() int {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	return len(fake.failActualLRPArgsForCall)
}

func (fake *FakeClient) FailActualLRPArgsForCall(i int) (*models.ActualLRPKey, string) {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	return fake.failActualLRPArgsForCall[i].key, fake.failActualLRPArgsForCall[i].errorMessage
}

func (fake *FakeClient) FailActualLRPReturns(result1 error) {
	fake.FailActualLRPStub = nil
	fake.failActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) RemoveActualLRP(processGuid string, index int) error {
	fake.removeActualLRPMutex.Lock()
	fake.removeActualLRPArgsForCall = append(fake.removeActualLRPArgsForCall, struct {
		processGuid string
		index       int
	}{processGuid, index})
	fake.removeActualLRPMutex.Unlock()
	if fake.RemoveActualLRPStub != nil {
		return fake.RemoveActualLRPStub(processGuid, index)
	} else {
		return fake.removeActualLRPReturns.result1
	}
}

func (fake *FakeClient) RemoveActualLRPCallCount() int {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return len(fake.removeActualLRPArgsForCall)
}

func (fake *FakeClient) RemoveActualLRPArgsForCall(i int) (string, int) {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return fake.removeActualLRPArgsForCall[i].processGuid, fake.removeActualLRPArgsForCall[i].index
}

func (fake *FakeClient) RemoveActualLRPReturns(result1 error) {
	fake.RemoveActualLRPStub = nil
	fake.removeActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DesiredLRPs(arg1 models.DesiredLRPFilter) ([]*models.DesiredLRP, error) {
	fake.desiredLRPsMutex.Lock()
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct {
		arg1 models.DesiredLRPFilter
	}{arg1})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub(arg1)
	} else {
		return fake.desiredLRPsReturns.result1, fake.desiredLRPsReturns.result2
	}
}

func (fake *FakeClient) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeClient) DesiredLRPsArgsForCall(i int) models.DesiredLRPFilter {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return fake.desiredLRPsArgsForCall[i].arg1
}

func (fake *FakeClient) DesiredLRPsReturns(result1 []*models.DesiredLRP, result2 error) {
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 []*models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DesiredLRPByProcessGuid(processGuid string) (*models.DesiredLRP, error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	fake.desiredLRPByProcessGuidArgsForCall = append(fake.desiredLRPByProcessGuidArgsForCall, struct {
		processGuid string
	}{processGuid})
	fake.desiredLRPByProcessGuidMutex.Unlock()
	if fake.DesiredLRPByProcessGuidStub != nil {
		return fake.DesiredLRPByProcessGuidStub(processGuid)
	} else {
		return fake.desiredLRPByProcessGuidReturns.result1, fake.desiredLRPByProcessGuidReturns.result2
	}
}

func (fake *FakeClient) DesiredLRPByProcessGuidCallCount() int {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.desiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeClient) DesiredLRPByProcessGuidArgsForCall(i int) string {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return fake.desiredLRPByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeClient) DesiredLRPByProcessGuidReturns(result1 *models.DesiredLRP, result2 error) {
	fake.DesiredLRPByProcessGuidStub = nil
	fake.desiredLRPByProcessGuidReturns = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Tasks() ([]*models.Task, error) {
	fake.tasksMutex.Lock()
	fake.tasksArgsForCall = append(fake.tasksArgsForCall, struct{}{})
	fake.tasksMutex.Unlock()
	if fake.TasksStub != nil {
		return fake.TasksStub()
	} else {
		return fake.tasksReturns.result1, fake.tasksReturns.result2
	}
}

func (fake *FakeClient) TasksCallCount() int {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return len(fake.tasksArgsForCall)
}

func (fake *FakeClient) TasksReturns(result1 []*models.Task, result2 error) {
	fake.TasksStub = nil
	fake.tasksReturns = struct {
		result1 []*models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) TasksByDomain(domain string) ([]*models.Task, error) {
	fake.tasksByDomainMutex.Lock()
	fake.tasksByDomainArgsForCall = append(fake.tasksByDomainArgsForCall, struct {
		domain string
	}{domain})
	fake.tasksByDomainMutex.Unlock()
	if fake.TasksByDomainStub != nil {
		return fake.TasksByDomainStub(domain)
	} else {
		return fake.tasksByDomainReturns.result1, fake.tasksByDomainReturns.result2
	}
}

func (fake *FakeClient) TasksByDomainCallCount() int {
	fake.tasksByDomainMutex.RLock()
	defer fake.tasksByDomainMutex.RUnlock()
	return len(fake.tasksByDomainArgsForCall)
}

func (fake *FakeClient) TasksByDomainArgsForCall(i int) string {
	fake.tasksByDomainMutex.RLock()
	defer fake.tasksByDomainMutex.RUnlock()
	return fake.tasksByDomainArgsForCall[i].domain
}

func (fake *FakeClient) TasksByDomainReturns(result1 []*models.Task, result2 error) {
	fake.TasksByDomainStub = nil
	fake.tasksByDomainReturns = struct {
		result1 []*models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) TasksByCellID(cellId string) ([]*models.Task, error) {
	fake.tasksByCellIDMutex.Lock()
	fake.tasksByCellIDArgsForCall = append(fake.tasksByCellIDArgsForCall, struct {
		cellId string
	}{cellId})
	fake.tasksByCellIDMutex.Unlock()
	if fake.TasksByCellIDStub != nil {
		return fake.TasksByCellIDStub(cellId)
	} else {
		return fake.tasksByCellIDReturns.result1, fake.tasksByCellIDReturns.result2
	}
}

func (fake *FakeClient) TasksByCellIDCallCount() int {
	fake.tasksByCellIDMutex.RLock()
	defer fake.tasksByCellIDMutex.RUnlock()
	return len(fake.tasksByCellIDArgsForCall)
}

func (fake *FakeClient) TasksByCellIDArgsForCall(i int) string {
	fake.tasksByCellIDMutex.RLock()
	defer fake.tasksByCellIDMutex.RUnlock()
	return fake.tasksByCellIDArgsForCall[i].cellId
}

func (fake *FakeClient) TasksByCellIDReturns(result1 []*models.Task, result2 error) {
	fake.TasksByCellIDStub = nil
	fake.tasksByCellIDReturns = struct {
		result1 []*models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) TaskByGuid(guid string) (*models.Task, error) {
	fake.taskByGuidMutex.Lock()
	fake.taskByGuidArgsForCall = append(fake.taskByGuidArgsForCall, struct {
		guid string
	}{guid})
	fake.taskByGuidMutex.Unlock()
	if fake.TaskByGuidStub != nil {
		return fake.TaskByGuidStub(guid)
	} else {
		return fake.taskByGuidReturns.result1, fake.taskByGuidReturns.result2
	}
}

func (fake *FakeClient) TaskByGuidCallCount() int {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return len(fake.taskByGuidArgsForCall)
}

func (fake *FakeClient) TaskByGuidArgsForCall(i int) string {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return fake.taskByGuidArgsForCall[i].guid
}

func (fake *FakeClient) TaskByGuidReturns(result1 *models.Task, result2 error) {
	fake.TaskByGuidStub = nil
	fake.taskByGuidReturns = struct {
		result1 *models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SubscribeToEvents() (events.EventSource, error) {
	fake.subscribeToEventsMutex.Lock()
	fake.subscribeToEventsArgsForCall = append(fake.subscribeToEventsArgsForCall, struct{}{})
	fake.subscribeToEventsMutex.Unlock()
	if fake.SubscribeToEventsStub != nil {
		return fake.SubscribeToEventsStub()
	} else {
		return fake.subscribeToEventsReturns.result1, fake.subscribeToEventsReturns.result2
	}
}

func (fake *FakeClient) SubscribeToEventsCallCount() int {
	fake.subscribeToEventsMutex.RLock()
	defer fake.subscribeToEventsMutex.RUnlock()
	return len(fake.subscribeToEventsArgsForCall)
}

func (fake *FakeClient) SubscribeToEventsReturns(result1 events.EventSource, result2 error) {
	fake.SubscribeToEventsStub = nil
	fake.subscribeToEventsReturns = struct {
		result1 events.EventSource
		result2 error
	}{result1, result2}
}

var _ bbs.Client = new(FakeClient)
