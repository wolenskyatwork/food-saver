package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetActivitiesHandler(t *testing.T) {
	mockStore := InitMockStore()
	mockStore.On("GetActivities").Return([]*Activity{
		{Name: "fake activity"},
	}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(GetActivityHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Activity{ "fake activity" }
	a := []Activity{}
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
