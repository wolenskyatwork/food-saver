package controller

import (
	"github.com/wolenskyatwork/food-saver/store"
	"net/http"
)

type TrainingController struct {
	Service store.Store
}

func NewTrainingController(service store.Store) Controller {
	return TrainingController{ Service: service }
}

func (tc TrainingController) Index(w http.ResponseWriter, r *http.Request) {

}
