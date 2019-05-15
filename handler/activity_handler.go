package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/wolenskyatwork/food-saver/store"
)

type ActivityHandler struct {
	Service store.Store
}

func (h ActivityHandler) GetActivityHandler(w http.ResponseWriter, r *http.Request) {
	activities, err := h.Service.GetActivities()
	activityListBytes, err := json.Marshal(activities)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(activityListBytes)
}
