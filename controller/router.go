package controller

import (
	"github.com/gorilla/mux"
	"github.com/wolenskyatwork/food-saver/store"
)

func NewRouter(service store.Store) *mux.Router{
	router := mux.NewRouter()

	activityNameController := ActivityNameController{ Service: service }
	activityController := ActivityController{ Service: service }

	router.HandleFunc("/healthcheck", GetHealthcheckController).Methods("GET")
	router.HandleFunc("/activityName", activityNameController.Index).Methods("GET")
	router.HandleFunc("/activity", activityController.Index).Methods("GET")

	return router
}
