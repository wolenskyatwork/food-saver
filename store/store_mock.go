package store

import (
	"github.com/stretchr/testify/mock"
	"github.com/wolenskyatwork/food-saver/dao"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateActivity(activity *dao.Activity) error {
	res := m.Called(activity)
	return res.Error(0)
}

func (m *MockStore) GetActivities() ([]*dao.Activity, error) {
	res := m.Called()
	return res.Get(0).([]*dao.Activity), res.Error(1)
}

