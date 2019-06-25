package controller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	activities, err := ac.Service.GetActivities(vars["userId"])

	activityDates := make(map[string][]*dao.Activity)

	for _, activity := range activities {
		if v, ok := activityDates[activity.DateCompleted]; ok {
			activityDates[activity.DateCompleted] = append(v, activity)
		} else {
			activityDates[activity.DateCompleted] = []*dao.Activity{activity}
		}
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(activityDates)
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
		return
	} else if activity.Id == "" || activity.DateCompleted == "" || activity.UserId == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("missing required info to create activity"))
		return
	}

	err = ac.Service.CreateActivity(activity)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}
