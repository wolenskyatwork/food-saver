package controller

import (
	"encoding/json"
	"github.com/wolenskyatwork/food-saver/dao"
	"github.com/wolenskyatwork/food-saver/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetActivityNamesHandler(t *testing.T) {
	mockStore := store.MockStore{}

	mockStore.On("GetActivityNames", "").Return([]*dao.ActivityName{
		{ Name: "fake activity", Id: "1" },
	}, nil)

	activityController := ActivityNameController{
		Service: mockStore,
	}

	req, err := http.NewRequest("GET", "/user/1/activityName", nil)
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
