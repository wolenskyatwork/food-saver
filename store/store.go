package store

import (
	"database/sql"
	"github.com/wolenskyatwork/food-saver/handler"
)

type Store interface {
	GetActivities() ([]*handler.Activity, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) GetActivities() ([]*handler.Activity, error) {
	rows, err := store.db.Query("SELECT name FROM activity;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	activities := []*handler.Activity{}
	for rows.Next() {
		activity := &handler.Activity{}

		if err := rows.Scan(&activity.Name); err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}
	return activities, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
