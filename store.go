package main

import (
	"database/sql"
)

type Store interface {
	CreateActivity(activity *Activity) error
	GetActivities() ([]*Activity, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateActivity(activity *Activity) error {
	_, err := store.db.Query("INSERT INTO activity_name (name) VALUES ($1);", activity.Name)
	return err
}

func (store *dbStore) GetActivities() ([]*Activity, error) {
	rows, err := store.db.Query("SELECT name FROM activity_name;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	activities := []*Activity{}
	for rows.Next() {
		activity := &Activity{}

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
