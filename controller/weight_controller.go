package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wolenskyatwork/food-saver/store"
	"net/http"
)

type WeightController struct {
	Service store.Store
}

func NewWeightController(service store.Store) Controller {
	return WeightController{ Service: service }
}

func (wc WeightController) Index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weights, err := wc.Service.GetWeights(vars["userId"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(weights)
	}
}

func (wc WeightController) Create(w http.ResponseWriter, r *http.Request) {

}
