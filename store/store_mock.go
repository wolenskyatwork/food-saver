package store

import (
	"github.com/stretchr/testify/mock"
	"github.com/wolenskyatwork/food-saver/dao"
)

type MockStore struct {
	mock.Mock
}

//func NewMockStore() {
//	return MockStore{}
//}

func (m *MockStore) CreateActivity(activity dao.Activity) error {
	res := m.Called(activity)
	return res.Error(0)
}

func (m *MockStore) GetActivities() ([]*dao.Activity, error){
	res := m.Called()
	return res.Get(0).([]*dao.Activity), res.Error(1)
}

func (m *MockStore) GetActivityNames(userId string) ([]*dao.ActivityName, error){
	// testing environment not recognizing mux.Vars(r), must call this function with "" as userId
	res := m.Called(userId)
	return res.Get(0).([]*dao.ActivityName), res.Error(1)
}

func (m *MockStore) CreateActivityName(activity *dao.ActivityName) error {
	res := m.Called(activity)
	return res.Error(0)
}
