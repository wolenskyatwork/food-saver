package main

import (
	"github.com/wolenskyatwork/food-saver/controller"
	"github.com/wolenskyatwork/food-saver/store"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckRouter(t *testing.T) {
	mockStore := new(store.MockStore)

	r := controller.NewRouter(mockStore)
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/healthcheck")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "Hello World!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	mockStore := new(store.MockStore)

	r := controller.NewRouter(mockStore)
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL + "/healthcheck", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}
