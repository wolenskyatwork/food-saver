package store

import (
	"github.com/stretchr/testify/mock"
	"github.com/wolenskyatwork/food-saver/dao"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateActivityName(activity *dao.ActivityName) error {
	res := m.Called(activity)
	return res.Error(0)
}

func (m *MockStore) GetActivityNames() ([]*dao.ActivityName, error) {
	res := m.Called()
	return res.Get(0).([]*dao.ActivityName), res.Error(1)
}

func (m *MockStore) GetActivities() ([]*dao.Activity, error) {
	//TODO fix this
	res := m.Called()
	return res.Get(0).([]*dao.Activity), res.Error(1)
}

