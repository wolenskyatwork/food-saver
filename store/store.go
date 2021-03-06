package store

import (
	"database/sql"
	"errors"
	"github.com/wolenskyatwork/food-saver/dao"
	"github.com/wolenskyatwork/food-saver/user"
)

type Store interface {
	CreateActivityName(activity *dao.ActivityName) error
	GetActivityNames(userId string) ([]*dao.ActivityName, error)
	CreateActivity(dao.Activity) error
	GetActivities(userId string) ([]*dao.Activity, error)
	GetWeights(userId string) ([]*dao.Weight, error)
	CreateWeight(dao.Weight) error
}

type DBStore struct {
	DB *sql.DB
}

func NewDBStore(db *sql.DB) Store {
	return &DBStore{DB: db}
}

func (store *DBStore) CreateWeight(weight dao.Weight) error {
	_, err := store.DB.Exec("INSERT INTO weight (app_user_id, weight, weight_date) VALUES ($1, $2, $3)",
		weight.UserId, weight.Weight, weight.WeightDate)
	return err
}

func (store *DBStore) GetWeights(userId string) ([]*dao.Weight, error) {
 rows, err := store.DB.Query("SELECT id, app_user_id, weight, weight_date FROM weight WHERE app_user_id = $1", userId)
 if err != nil {
 	return nil, err
 }
 defer rows.Close()

 var weights = make([]*dao.Weight, 0)
 for rows.Next() {
	 weight := &dao.Weight{}
	 if err := rows.Scan(&weight.Id, &weight.UserId, &weight.Weight, &weight.WeightDate); err != nil {
		 return nil, err
	 }

	 weights = append(weights, weight)
 }

 return weights, nil
}

func (store *DBStore) CreateActivity(activity dao.Activity) error {
	if activity.Id == "" || activity.DateCompleted == "" || activity.UserId == "" {
		return errors.New("Missing id or datecompleted on activity")
	} else {
		_, err := store.DB.Exec("INSERT INTO activity (activity_id, app_user_id, datetime_completed) VALUES ($1, $2, $3)",
			activity.Id, user.GetUserId(), activity.DateCompleted)
		return err
	}
}

func (store *DBStore) GetActivities(userId string) ([]*dao.Activity, error) {
	rows, err := store.DB.Query("SELECT a.activity_id, a.app_user_id, a.datetime_completed, an.name FROM activity a LEFT JOIN activity_name an ON a.activity_id = an.id WHERE a.app_user_id = $1;", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*dao.Activity
	for rows.Next() {
		activity := &dao.Activity{}
		if err := rows.Scan(&activity.Id, &activity.UserId, &activity.DateCompleted, &activity.Name); err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}
	return activities, nil
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
