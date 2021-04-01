// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "github.com/carRub/academy-go-q12021/model"
	mock "github.com/stretchr/testify/mock"
)

// CharacterService is an autogenerated mock type for the CharacterService type
type CharacterService struct {
	mock.Mock
}

// GetCharacterByID provides a mock function with given fields: id
func (_m *CharacterService) GetCharacterByID(id int) (*model.Character, error) {
	ret := _m.Called(id)

	var r0 *model.Character
	if rf, ok := ret.Get(0).(func(int) *model.Character); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Character)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCharacters provides a mock function with given fields:
func (_m *CharacterService) GetCharacters() ([]model.Character, error) {
	ret := _m.Called()

	var r0 []model.Character
	if rf, ok := ret.Get(0).(func() []model.Character); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Character)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertExternalCharacter provides a mock function with given fields: id
func (_m *CharacterService) InsertExternalCharacter(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
