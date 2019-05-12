package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
req, err := http.NewRequest("GET", "", nil)
if err != nil {
t.Fatal(err)
}

recorder := httptest.NewRecorder()
hf := http.HandlerFunc(GetHealthcheckHandler)

hf.ServeHTTP(recorder, req)

if status := recorder.Code; status != http.StatusOK {
t.Errorf("handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
}

expected := `Hello World!`
actual := recorder.Body.String()

if actual != expected {
t.Errorf("handler returned unexpected body: got %v wanted %v", actual, expected)
}
}
