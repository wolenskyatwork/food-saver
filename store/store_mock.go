package store

import (
	"github.com/stretchr/testify/mock"
	"github.com/wolenskyatwork/food-saver/dao"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateActivity(activity *dao.Activity) error {
	rets := m.Called(activity)
	return rets.Error(0)
}

func (m *MockStore) GetActivities() ([]*dao.Activity, error) {
	rets := m.Called()
	return rets.Get(0).([]*dao.Activity), rets.Error(1)
}
//
//func InitMockStore() *MockStore {
//	s := new(MockStore)
//	store = s
//	return s
//}
