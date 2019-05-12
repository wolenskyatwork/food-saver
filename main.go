package main

import (
	"github.com/gorilla/mux"
	"github.com/wolenskyatwork/food-saver/handler"
	"net/http"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/healthcheck", handler.GetHealthcheckHandler).Methods("GET")
	router.HandleFunc("/activity", handler.GetActivityHandler).Methods("GET")

	return router
}

func main() {
	router := newRouter()
	http.ListenAndServe(":8081", router)
}
