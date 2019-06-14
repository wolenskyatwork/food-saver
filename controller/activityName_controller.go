package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/wolenskyatwork/food-saver/store"
)

type ActivityNameController struct {
	Service store.Store
}

func (h ActivityNameController) Index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activities, err := h.Service.GetActivityNames(vars["userId"])

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(activities)
}

// TODO make sure name is correct

func (h ActivityNameController) Create(w http.ResponseWriter, r *http.Request) {

}

