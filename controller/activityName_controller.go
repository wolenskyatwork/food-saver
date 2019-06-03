package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/wolenskyatwork/food-saver/store"
)

type ActivityNameController struct {
	Service store.Store
}

func (h ActivityNameController) Index(w http.ResponseWriter, r *http.Request) {
	activities, err := h.Service.GetActivityNames()
	activityListBytes, err := json.Marshal(activities)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(activityListBytes)
}

// TODO make sure name is correct

func (h ActivityNameController) Create(w http.ResponseWriter, r *http.Request) {

}

