// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/mevansam/cf-cli-api/cfapi"
)

type FakeCfSessionProvider struct {
	NewCfSessionStub        func(apiEndPoint string, userName string, password string, orgName string, spaceName string, sslDisabled bool, logger *cfapi.Logger) (cfSession cfapi.CfSession, err error)
	newCfSessionMutex       sync.RWMutex
	newCfSessionArgsForCall []struct {
		apiEndPoint string
		userName    string
		password    string
		orgName     string
		spaceName   string
		sslDisabled bool
		logger      *cfapi.Logger
	}
	newCfSessionReturns struct {
		result1 cfapi.CfSession
		result2 error
	}
	NewCfSessionFromFilepathStub        func(configPath string, sslDisabled bool, logger *cfapi.Logger) (cfSession cfapi.CfSession, err error)
	newCfSessionFromFilepathMutex       sync.RWMutex
	newCfSessionFromFilepathArgsForCall []struct {
		configPath  string
		sslDisabled bool
		logger      *cfapi.Logger
	}
	newCfSessionFromFilepathReturns struct {
		result1 cfapi.CfSession
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCfSessionProvider) NewCfSession(apiEndPoint string, userName string, password string, orgName string, spaceName string, sslDisabled bool, logger *cfapi.Logger) (cfSession cfapi.CfSession, err error) {
	fake.newCfSessionMutex.Lock()
	fake.newCfSessionArgsForCall = append(fake.newCfSessionArgsForCall, struct {
		apiEndPoint string
		userName    string
		password    string
		orgName     string
		spaceName   string
		sslDisabled bool
		logger      *cfapi.Logger
	}{apiEndPoint, userName, password, orgName, spaceName, sslDisabled, logger})
	fake.recordInvocation("NewCfSession", []interface{}{apiEndPoint, userName, password, orgName, spaceName, sslDisabled, logger})
	fake.newCfSessionMutex.Unlock()
	if fake.NewCfSessionStub != nil {
		return fake.NewCfSessionStub(apiEndPoint, userName, password, orgName, spaceName, sslDisabled, logger)
	} else {
		return fake.newCfSessionReturns.result1, fake.newCfSessionReturns.result2
	}
}

func (fake *FakeCfSessionProvider) NewCfSessionCallCount() int {
	fake.newCfSessionMutex.RLock()
	defer fake.newCfSessionMutex.RUnlock()
	return len(fake.newCfSessionArgsForCall)
}

func (fake *FakeCfSessionProvider) NewCfSessionArgsForCall(i int) (string, string, string, string, string, bool, *cfapi.Logger) {
	fake.newCfSessionMutex.RLock()
	defer fake.newCfSessionMutex.RUnlock()
	return fake.newCfSessionArgsForCall[i].apiEndPoint, fake.newCfSessionArgsForCall[i].userName, fake.newCfSessionArgsForCall[i].password, fake.newCfSessionArgsForCall[i].orgName, fake.newCfSessionArgsForCall[i].spaceName, fake.newCfSessionArgsForCall[i].sslDisabled, fake.newCfSessionArgsForCall[i].logger
}

func (fake *FakeCfSessionProvider) NewCfSessionReturns(result1 cfapi.CfSession, result2 error) {
	fake.NewCfSessionStub = nil
	fake.newCfSessionReturns = struct {
		result1 cfapi.CfSession
		result2 error
	}{result1, result2}
}

func (fake *FakeCfSessionProvider) NewCfSessionFromFilepath(configPath string, sslDisabled bool, logger *cfapi.Logger) (cfSession cfapi.CfSession, err error) {
	fake.newCfSessionFromFilepathMutex.Lock()
	fake.newCfSessionFromFilepathArgsForCall = append(fake.newCfSessionFromFilepathArgsForCall, struct {
		configPath  string
		sslDisabled bool
		logger      *cfapi.Logger
	}{configPath, sslDisabled, logger})
	fake.recordInvocation("NewCfSessionFromFilepath", []interface{}{configPath, sslDisabled, logger})
	fake.newCfSessionFromFilepathMutex.Unlock()
	if fake.NewCfSessionFromFilepathStub != nil {
		return fake.NewCfSessionFromFilepathStub(configPath, sslDisabled, logger)
	} else {
		return fake.newCfSessionFromFilepathReturns.result1, fake.newCfSessionFromFilepathReturns.result2
	}
}

func (fake *FakeCfSessionProvider) NewCfSessionFromFilepathCallCount() int {
	fake.newCfSessionFromFilepathMutex.RLock()
	defer fake.newCfSessionFromFilepathMutex.RUnlock()
	return len(fake.newCfSessionFromFilepathArgsForCall)
}

func (fake *FakeCfSessionProvider) NewCfSessionFromFilepathArgsForCall(i int) (string, bool, *cfapi.Logger) {
	fake.newCfSessionFromFilepathMutex.RLock()
	defer fake.newCfSessionFromFilepathMutex.RUnlock()
	return fake.newCfSessionFromFilepathArgsForCall[i].configPath, fake.newCfSessionFromFilepathArgsForCall[i].sslDisabled, fake.newCfSessionFromFilepathArgsForCall[i].logger
}

func (fake *FakeCfSessionProvider) NewCfSessionFromFilepathReturns(result1 cfapi.CfSession, result2 error) {
	fake.NewCfSessionFromFilepathStub = nil
	fake.newCfSessionFromFilepathReturns = struct {
		result1 cfapi.CfSession
		result2 error
	}{result1, result2}
}

func (fake *FakeCfSessionProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newCfSessionMutex.RLock()
	defer fake.newCfSessionMutex.RUnlock()
	fake.newCfSessionFromFilepathMutex.RLock()
	defer fake.newCfSessionFromFilepathMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCfSessionProvider) recordInvocation(key string, args []interface{}) {
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

var _ cfapi.CfSessionProvider = new(FakeCfSessionProvider)
