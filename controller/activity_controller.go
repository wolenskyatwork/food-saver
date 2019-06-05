package controller
//
//import (
//	"encoding/json"
//	"github.com/wolenskyatwork/food-saver/store"
//	"net/http"
//)
//
//type ActivityController struct {
//	Service store.Store
//}
//
//func NewActivityController(service store.Store) Controller {
//	return ActivityController{ Service: service }
//}
//
//func (h ActivityController) Index(w http.ResponseWriter, r *http.Request) {
//	activities, err := h.Service.GetActivities()
//
//	if err != nil {
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusInternalServerError)
//		json.NewEncoder(w).Encode(err.Error())
//	} else {
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(activities)
//	}
//}
