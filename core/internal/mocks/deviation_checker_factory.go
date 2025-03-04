// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	fluxmonitor "github.com/smartcontractkit/chainlink/core/services/fluxmonitor"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	orm "github.com/smartcontractkit/chainlink/core/store/orm"
)

// DeviationCheckerFactory is an autogenerated mock type for the DeviationCheckerFactory type
type DeviationCheckerFactory struct {
	mock.Mock
}

// New provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *DeviationCheckerFactory) New(_a0 models.Initiator, _a1 fluxmonitor.RunManager, _a2 *orm.ORM, _a3 models.Duration) (fluxmonitor.DeviationChecker, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 fluxmonitor.DeviationChecker
	if rf, ok := ret.Get(0).(func(models.Initiator, fluxmonitor.RunManager, *orm.ORM, models.Duration) fluxmonitor.DeviationChecker); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fluxmonitor.DeviationChecker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Initiator, fluxmonitor.RunManager, *orm.ORM, models.Duration) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
