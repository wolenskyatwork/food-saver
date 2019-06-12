package controller

import (
	"encoding/json"
	"github.com/wolenskyatwork/food-saver/dao"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m MockStore) CreateActivity(activity dao.Activity) error {
	res := m.Called(activity)
	return res.Error(0)
}

func (m MockStore) GetActivities() ([]*dao.Activity, error){
	res := m.Called()
	return res.Get(0).([]*dao.Activity), res.Error(1)
}

func (m MockStore) GetActivityNames(userId string) ([]*dao.ActivityName, error){
	res := m.Called()
	return res.Get(0).([]*dao.ActivityName), res.Error(1)
}

func (m MockStore) CreateActivityName(activity *dao.ActivityName) error {
	res := m.Called(activity)
	return res.Error(0)
}

func TestGetActivityNamesHandler(t *testing.T) {
	mockStore := new(MockStore)
	activityController := ActivityNameController{
		Service: mockStore,
	}

	mockStore.On("GetActivityNames").Return([]*dao.ActivityName{
		{ Name: "fake activity", Id: "1" },
	}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(activityController.Index)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := dao.ActivityName{ Name: "fake activity", Id: "1" }
	a := []dao.ActivityName{}
	err = json.NewDecoder(recorder.Body).Decode(&a)

	if err != nil {
		t.Fatal(err)
	}

	actual := a[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	mockStore.AssertExpectations(t)
}
