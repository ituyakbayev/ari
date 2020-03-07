// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/CyCoreSystems/ari/v5"
	mock "github.com/stretchr/testify/mock"
)

// Player is an autogenerated mock type for the Player type
type Player struct {
	mock.Mock
}

// Play provides a mock function with given fields: _a0, _a1
func (_m *Player) Play(_a0 string, _a1 string) (*ari.PlaybackHandle, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ari.PlaybackHandle
	if rf, ok := ret.Get(0).(func(string, string) *ari.PlaybackHandle); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.PlaybackHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StagePlay provides a mock function with given fields: _a0, _a1
func (_m *Player) StagePlay(_a0 string, _a1 string) (*ari.PlaybackHandle, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ari.PlaybackHandle
	if rf, ok := ret.Get(0).(func(string, string) *ari.PlaybackHandle); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.PlaybackHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: n
func (_m *Player) Subscribe(n ...string) ari.Subscription {
	_va := make([]interface{}, len(n))
	for _i := range n {
		_va[_i] = n[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ari.Subscription
	if rf, ok := ret.Get(0).(func(...string) ari.Subscription); ok {
		r0 = rf(n...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Subscription)
		}
	}

	return r0
}
