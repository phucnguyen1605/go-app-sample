// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Token is an autogenerated mock type for the Token type
type Token struct {
	mock.Mock
}

// Generate provides a mock function with given fields: claims
func (_m *Token) Generate(claims map[string]interface{}) (string, error) {
	ret := _m.Called(claims)

	var r0 string
	if rf, ok := ret.Get(0).(func(map[string]interface{}) string); ok {
		r0 = rf(claims)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: _a0
func (_m *Token) Validate(_a0 string) (map[string]interface{}, error) {
	ret := _m.Called(_a0)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}