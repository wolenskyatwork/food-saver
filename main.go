package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/wolenskyatwork/food-saver/handler"
	"github.com/wolenskyatwork/food-saver/store"
	"net/http"
)

func newRouter(activityHandler handler.ActivityHandler) *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/healthcheck", handler.GetHealthcheckHandler).Methods("GET")
	router.HandleFunc("/activity", activityHandler.GetActivityHandler).Methods("GET")

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
	activityHandler := handler.ActivityHandler{Service: &dbService}

	http.ListenAndServe(":8081", newRouter(activityHandler))
}


