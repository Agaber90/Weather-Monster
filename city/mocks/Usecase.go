package mocks

import (
	"Weather-Monster/models"
	"context"

	mock "github.com/stretchr/testify/mock"
)

// CityUsecase is an autogenerated mock type for the Usecase type
type CityUsecase struct {
	mock.Mock
}

// CreateCity provides a mock function with given fields
func (m *CityUsecase) CreateCity(c context.Context, cty *models.Cities) error {
	ret := m.Called(c, cty)

	var err error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Cities) error); ok {
		err = rf(c, cty)
	} else {
		err = ret.Error(0)
	}

	return err
}

// UpdateCity provides a mock function with given fields: c, cty
func (m *CityUsecase) UpdateCity(c context.Context, cty *models.Cities) error {
	ret := m.Called(c, cty)

	var err error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Cities) error); ok {
		err = rf(c, cty)
	} else {
		err = ret.Error(0)
	}

	return err
}

//DeleteCity provides a mock function with given fields: c, id
func (m *CityUsecase) DeleteCity(c context.Context, id int64) error {
	ret := m.Called(c, id)

	var err error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		err = rf(c, id)
	} else {
		err = ret.Error(0)
	}

	return err
}
