package arimocks

import mock "github.com/stretchr/testify/mock"

// EndpointEvent is an autogenerated mock type for the EndpointEvent type
type EndpointEvent struct {
	mock.Mock
}

// GetEndpointIDs provides a mock function with given fields:
func (_m *EndpointEvent) GetEndpointIDs() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}