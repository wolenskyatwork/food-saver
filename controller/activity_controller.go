package controller

import (
	"encoding/json"
	"github.com/wolenskyatwork/food-saver/dao"
	"github.com/wolenskyatwork/food-saver/store"
	"net/http"
)

type ActivityController struct {
	Service store.Store
}

func NewActivityController(service store.Store) Controller {
	return ActivityController{ Service: service }
}

func (ac ActivityController) Index(w http.ResponseWriter, r *http.Request) {
	activities, err := ac.Service.GetActivities()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(activities)
	}
}

func (ac ActivityController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var activity dao.Activity
	err := decoder.Decode(&activity)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // I guess?
		json.NewEncoder(w).Encode(err.Error())
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}
