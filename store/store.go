package store

import (
	"database/sql"
	"github.com/wolenskyatwork/food-saver/dao"
)

type Store interface {
	CreateActivityName(activity *dao.ActivityName) error
	GetActivityNames() ([]*dao.ActivityName, error)
}

type DBStore struct {
	DB *sql.DB
}

func (store DBStore) CreateActivityName(activity *dao.ActivityName) error {
	_, err := store.DB.Query("INSERT INTO activity_name (name) VALUES ($1);", activity.Name)
	return err
}

func (store DBStore) GetActivityNames() ([]*dao.ActivityName, error) {
	rows, err := store.DB.Query("SELECT name FROM activity_name;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	activityNames := []*dao.ActivityName{}
	for rows.Next() {
		activityName := &dao.ActivityName{}
		if err := rows.Scan(&activityName.Name); err != nil {
			return nil, err
		}

		activityNames = append(activityNames, activityName)
	}

	return activityNames, nil
}
