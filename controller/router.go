package controller

import (
	"github.com/gorilla/mux"
	"github.com/wolenskyatwork/food-saver/store"
)

func NewRouter(service store.Store) *mux.Router{
	router := mux.NewRouter()

	activityNameController := ActivityNameController{ Service: service }
	activityController := NewActivityController(service)
	weightController := NewWeightController(service)

	router.HandleFunc("/healthcheck", GetHealthcheckController).Methods("GET")
	router.HandleFunc("/user/{userId}/activityName", activityNameController.Index).Methods("GET")
	router.HandleFunc("/user/{userId}/activity", activityController.Index).Methods("GET")
	router.HandleFunc("/user/{userId}/activity", activityController.Create).Methods("POST")
	router.HandleFunc("/user/{userId}/weight", weightController.Index).Methods("GET")
	router.HandleFunc("/user/{userId}/weight", weightController.Create).Methods("POST")

	return router
}
