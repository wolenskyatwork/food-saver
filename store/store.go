package store

import (
	"database/sql"
	"github.com/wolenskyatwork/food-saver/handler"
)

type Store interface {
	CreateActivity(activity *handler.Activity) error
	GetActivities() ([]*handler.Activity, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateActivity(activity *handler.Activity) error {
	_, err := store.db.Query("INSERT INTO activity(name) VALUES ($1);", activity.Name)
	return err
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
