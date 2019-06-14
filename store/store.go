package store

import (
	"database/sql"
	"github.com/wolenskyatwork/food-saver/dao"
	"github.com/wolenskyatwork/food-saver/user"
)

type Store interface {
	CreateActivityName(activity *dao.ActivityName) error
	GetActivityNames(userId string) ([]*dao.ActivityName, error)
	CreateActivity(dao.Activity) error
	GetActivities() ([]*dao.Activity, error)
}

type DBStore struct {
	DB *sql.DB
}

func NewDBStore(db *sql.DB) Store {
	return &DBStore{DB: db}
}

func (store *DBStore) CreateActivity(activity dao.Activity) error {
	_, err := store.DB.Exec("INSERT INTO activity (activity_name, app_user_id, datetime_completed) VALUES ($1, $2, $3)",
		activity.Id, user.GetUserId(), activity.DateCompleted)
	return err
}

func (store *DBStore) GetActivities() ([]*dao.Activity, error) {
	return []*dao.Activity{}, nil
}

func (store *DBStore) CreateActivityName(activity *dao.ActivityName) error {
	_, err := store.DB.Exec("INSERT INTO activity_name (name) VALUES ($1);", activity.Name)
	return err
}

func (store *DBStore) GetActivityNames(userId string) ([]*dao.ActivityName, error) {
	rows, err := store.DB.Query("SELECT an.id, an.name FROM user_activity ua INNER JOIN activity_name an ON an.id = ua.activity_id WHERE ua.app_user_id = $1;", userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	activityNames := []*dao.ActivityName{}
	for rows.Next() {
		activityName := &dao.ActivityName{}
		if err := rows.Scan(&activityName.Id, &activityName.Name); err != nil {
			return nil, err
		}

		activityNames = append(activityNames, activityName)
	}

	return activityNames, nil
}
