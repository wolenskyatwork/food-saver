package store

import (
	"database/sql"
	"github.com/wolenskyatwork/food-saver/dao"
)

type Store interface {
	CreateActivity(activity *dao.Activity) error
	GetActivities() ([]*dao.Activity, error)
}

type DBStore struct {
	DB *sql.DB
}

func (store *DBStore) SaveActivityWIP() error {

}

func (store *DBStore) CreateActivity(activity *dao.Activity) error {
	_, err := store.DB.Query("INSERT INTO activity_name (name) VALUES ($1);", activity.Name)
	return err
}

func (store *DBStore) GetActivities() ([]*dao.Activity, error) {
	rows, err := store.DB.Query("SELECT name FROM activity_name;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	activities := []*dao.Activity{}
	for rows.Next() {
		activity := &dao.Activity{}
		if err := rows.Scan(&activity.Name); err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}

	return activities, nil
}
