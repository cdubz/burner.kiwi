// Code generated by mockery v1.0.0. DO NOT EDIT.

package testEmail

import data "github.com/haydenwoodhead/burner.kiwi/data"

import mock "github.com/stretchr/testify/mock"
import mux "github.com/gorilla/mux"

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// DeleteExipredRoutes provides a mock function with given fields:
func (p *Provider) DeleteExipredRoutes() error {
	ret := p.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterRoute provides a mock function with given fields: i
func (p *Provider) RegisterRoute(i data.Inbox) (string, error) {
	ret := p.Called(i)

	var r0 string
	if rf, ok := ret.Get(0).(func(data.Inbox) string); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(data.Inbox) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: websiteAddr, db, r, isBlacklisted
func (p *Provider) Start(websiteAddr string, db data.Database, r *mux.Router, isBlacklisted func(string) bool) error {
	ret := p.Called(websiteAddr, db, r, isBlacklisted)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, data.Database, *mux.Router, func(string) bool) error); ok {
		r0 = rf(websiteAddr, db, r, isBlacklisted)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (p *Provider) Stop() error {
	ret := p.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
