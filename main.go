package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/wolenskyatwork/food-saver/controller"
	"github.com/wolenskyatwork/food-saver/store"
	"net/http"
)

func newRouter(service *store.DBStore) *mux.Router{
	router := mux.NewRouter()

	activityHandler := controller.ActivityController{ Service: service }


	router.HandleFunc("/healthcheck", controller.GetHealthcheckController).Methods("GET")
	router.HandleFunc("/activity", activityHandler.Index).Methods("GET")

	return router
}

func main() {
	connString := "dbname=life_logger user=life_logger_app sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	dbService := store.DBStore{DB: db}

	http.ListenAndServe(":8081", newRouter(&dbService))
}


