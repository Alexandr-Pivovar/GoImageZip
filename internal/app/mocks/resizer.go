// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "GoImageZip/internal/domain"
import mock "github.com/stretchr/testify/mock"

// Resizer is an autogenerated mock type for the Resizer type
type Resizer struct {
	mock.Mock
}

// Do provides a mock function with given fields: _a0
func (_m *Resizer) Do(_a0 domain.Image) (domain.Image, error) {
	ret := _m.Called(_a0)

	var r0 domain.Image
	if rf, ok := ret.Get(0).(func(domain.Image) domain.Image); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.Image)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Image) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}