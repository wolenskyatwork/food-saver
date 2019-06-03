package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/wolenskyatwork/food-saver/controller"
	"github.com/wolenskyatwork/food-saver/store"
	"net/http"
)

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

	http.ListenAndServe(":8081", controller.NewRouter(dbService))
}


