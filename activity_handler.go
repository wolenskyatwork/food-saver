package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Activity struct {
	Name string `json:"name"`
}

func GetActivityHandler(w http.ResponseWriter, r *http.Request) {
	activities, err := store.GetActivities()
	activityListBytes, err := json.Marshal(activities)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(activityListBytes)
}
