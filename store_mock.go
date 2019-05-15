package main

import (
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateActivity(activity *Activity) error {
	rets := m.Called(activity)
	return rets.Error(0)
}

func (m *MockStore) GetActivities() ([]*Activity, error) {
	rets := m.Called()
	return rets.Get(0).([]*Activity), rets.Error(1)
}

func InitMockStore() *MockStore {
	s := new(MockStore)
	store = s
	return s
}
