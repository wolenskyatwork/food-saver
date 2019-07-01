package store

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"database/sql"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"github.com/wolenskyatwork/food-saver/dao"
	"testing"
)

type StoreSuite struct {
	suite.Suite
	store *DBStore
}

func (s *StoreSuite) SetupSuite() {
	// TODO read this from the config
	connString := "dbname=life_logger_test user=life_logger_app sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	s.store = &DBStore{DB: db}
}

func (s *StoreSuite) SetupTest() {
	createGooseDatabase(s.store.DB)
}

func (s *StoreSuite) TearDownSuite() {
	s.store.DB.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestGetActivityNames() {
	Convey("when known data exists in db", s.T(), func() {
		_, err := s.store.DB.Exec("INSERT INTO activity_name (name) VALUES ('medicine'),('workout'),('climbing'),('biking'),('archery'),('spartan');")
		if err != nil {
			s.T().Fatal(err)
		}

		_, err = s.store.DB.Exec("INSERT INTO app_user (username) VALUES ('sailorflares'), ('closgmr');")
		if err != nil {
			s.T().Fatal(err)
		}

		_, err = s.store.DB.Exec("INSERT INTO user_activity (activity_id, app_user_id) VALUES (1,1), (2,1), (3,1), (4,2);")
		if err != nil {
			s.T().Fatal(err)
		}

		actual, err := s.store.GetActivityNames("1")
		if err != nil {
			s.T().Fatal(err)
		}

		expected := []*dao.ActivityName{
			{Name: "medicine", Id: "1" },
			{Name: "workout", Id: "2" },
			{Name: "climbing", Id: "3" },
		}

		So(actual, ShouldHaveLength, len(expected))
		So(actual, ShouldResemble, expected)

		actual, err = s.store.GetActivityNames("2")
		if err != nil {
			s.T().Fatal(err)
		}

		expected = []*dao.ActivityName{
			{Name: "biking", Id: "4" },
		}

		So(actual, ShouldHaveLength, len(expected))
		So(actual, ShouldResemble, expected)
	})
}

func (s *StoreSuite) TestGetActivities() {
	Convey("when the data exists", s.T(), func() {
		_, err := s.store.DB.Exec("INSERT INTO activity_name (name) VALUES ('medicine'),('workout'),('climbing'),('biking'),('archery'),('spartan');")
		if err != nil {
			s.T().Fatal(err)
		}

		_, err = s.store.DB.Exec("INSERT INTO app_user (username) VALUES ('sailorflares'), ('closgmr');")
		if err != nil {
			s.T().Fatal(err)
		}

		_, err = s.store.DB.Exec("INSERT INTO user_activity (activity_id, app_user_id) VALUES (1,1), (2,1), (3,1), (4,2);")
		if err != nil {
			s.T().Fatal(err)
		}

		_, err = s.store.DB.Exec("INSERT INTO activity (activity_id, app_user_id, datetime_completed) VALUES (1,1,'2019-4-12'), (2,1,'2018-4-12'), (3,1,'2017-4-12'), (4,2,'2016-4-12');")
		if err != nil {
			s.T().Fatal(err)
		}

		actual, err := s.store.GetActivities("1")

		if err != nil {
			s.T().Fatal(err)
		}

		expected := []*dao.Activity{
			{Name: "medicine", Id: "1", UserId: "1", DateCompleted: "2019-04-12T00:00:00Z"},
			{Name: "workout", Id: "2", UserId: "1", DateCompleted: "2018-04-12T00:00:00Z" },
			{Name: "climbing", Id: "3", UserId: "1", DateCompleted: "2017-04-12T00:00:00Z" },
		}

		So(actual, ShouldHaveLength, len(expected))
		So(actual, ShouldResemble, expected)
	})
}

func (s *StoreSuite) TestCreateActivity() {
	Convey("When the data exists", s.T(), func() {
		_, err := s.store.DB.Exec("INSERT INTO activity_name (name) VALUES ('medicine'),('workout'),('climbing'),('biking'),('archery'),('spartan');")
		if err != nil {
			So(err, ShouldBeNil)
		}

		_, err = s.store.DB.Exec("INSERT INTO app_user (username) VALUES ('sailorflares'), ('closgmr');")
		if err != nil {
			So(err, ShouldBeNil)
		}

		s.store.CreateActivity(dao.Activity{
			Id: "2",
			UserId: "1",
			DateCompleted: "2019-1-2",
		})

		rows, err := s.store.DB.Query(`SELECT activity_id, app_user_id, datetime_completed FROM activity;`)
		if err != nil {
			So(err, ShouldBeNil)
		}

		activities := []*dao.Activity{}
		for rows.Next() {
			activity := &dao.Activity{}
			if err := rows.Scan(&activity.Id, &activity.UserId, &activity.DateCompleted); err != nil {
				So(err, ShouldBeNil)
			}

			activities = append(activities, activity)
		}

		So(activities, ShouldHaveLength, 1)
		So(activities[0], ShouldResemble, &dao.Activity{
			Id: "2",
			UserId: "1",
			DateCompleted: "2019-01-02T00:00:00Z",
		})
	})
}

func (s *StoreSuite) TestCreateActivityMalformedData() {
	Convey("When the activity is missing required fields", s.T(), func() {
		err := s.store.CreateActivity(dao.Activity{
			Id: "2",
			UserId: "1",
		})

		So(err, ShouldNotBeNil)
	})
}

func (s *StoreSuite) TestCreateActivityName() {
	Convey("When createactivityname is called", s.T(), func() {
		s.store.CreateActivityName(&dao.ActivityName{
			Name: "test name",
		})

		res, err := s.store.DB.Query(`SELECT COUNT(*) FROM activity_name WHERE name='test name';`)
		if err != nil {
			So(err, ShouldBeNil)
		}

		var count int
		for res.Next() {
			err := res.Scan(&count)
			if err != nil {
				So(err, ShouldBeNil)
			}
		}

		So(count, ShouldEqual, 1)
	})
}

func (s *StoreSuite) TestGetWeights() {
	Convey("When known weights exist", s.T(), func() {
		_, err := s.store.DB.Exec("INSERT INTO app_user (username) VALUES ('sailorflares'), ('closgmr');")
		if err != nil {
			So(err, ShouldBeNil)
		}

		_, err = s.store.DB.Exec("INSERT INTO weight (app_user_id, weight, weight_date) VALUES ('1', 143.7, '2019-06-23');")
		if err != nil {
			So(err, ShouldBeNil)
		}

		_, err = s.store.DB.Exec("INSERT INTO weight (app_user_id, weight, weight_date) VALUES ('1', 142.7, '2019-06-22');")
		if err != nil {
			So(err, ShouldBeNil)
		}

		_, err = s.store.DB.Exec("INSERT INTO weight (app_user_id, weight, weight_date) VALUES ('1', 141.7, '2019-06-21');")
		if err != nil {
			So(err, ShouldBeNil)
		}

		w0 := dao.Weight{
			Id: "1",
			UserId: "1",
			Weight: 143.7,
			WeightDate: "2019-06-23T00:00:00Z",
		}

		w1 := dao.Weight{
			Id: "2",
			UserId: "1",
			Weight: 142.7,
			WeightDate: "2019-06-22T00:00:00Z",
		}

		w2 := dao.Weight{
			Id: "3",
			UserId: "1",
			Weight: 141.7,
			WeightDate: "2019-06-21T00:00:00Z",
		}

		Convey("then Service returns correctly mapped weights", func() {
			weights, err := s.store.GetWeights("1")
			So(err, ShouldBeNil)
			So(weights, ShouldHaveLength, 3)
			So(weights[0], ShouldResemble, &w0)
			So(weights[1], ShouldResemble, &w1)
			So(weights[2], ShouldResemble, &w2)
		})
	})
}

func (s *StoreSuite) TestEmptyGetWeights() {
	Convey("When no weights exist for user", s.T(), func() {
		_, err := s.store.DB.Exec("INSERT INTO app_user (username) VALUES ('sailorflares'), ('closgmr');")
		if err != nil {
			So(err, ShouldBeNil)
		}

		Convey("then Service returns correctly mapped weights", func() {
			weights, err := s.store.GetWeights("1")
			So(err, ShouldBeNil)
			So(weights, ShouldHaveLength, 0)
		})
	})
}

func (s *StoreSuite) TestCreateWeight() {
	_, err := s.store.DB.Exec("INSERT INTO app_user (username) VALUES ('sailorflares'), ('closgmr');")
	if err != nil {
		So(err, ShouldBeNil)
	}

	Convey("When calling createWeight function", s.T(), func() {
		weight := dao.Weight{
			UserId: "1",
			Weight: 142.1,
			WeightDate: "2019-06-22",
		}
		err := s.store.CreateWeight(weight)
		So(err, ShouldBeNil)

		Convey("then the row should have been created in the db", func() {
			res, err := s.store.DB.Query(`SELECT id, app_user_id, weight, weight_date FROM weight WHERE app_user_id = 1;`)
			if err != nil {
				So(err, ShouldBeNil)
			}

			var weight dao.Weight
			for res.Next() {
				err := res.Scan(&weight.Id, &weight.UserId, &weight.Weight, &weight.WeightDate)
				if err != nil {
					So(err, ShouldBeNil)
				}
			}

			So(weight, ShouldResemble, dao.Weight{
				Id: "1",
				UserId: "1",
				Weight: 142.1,
				WeightDate: "2019-06-22T00:00:00Z",
			})
		})
	})
}

func createGooseDatabase(db *sql.DB) {
	_, err := db.Exec("drop schema public cascade; create schema public;")
	if err != nil {
		panic(err)
	}

	conf, err := goose.NewDBConf("../db", "testing", "")
	if err != nil {
		panic(err)
	}

	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		panic(err)
	}

	if err := goose.RunMigrationsOnDb(conf, conf.MigrationsDir, target, db); err != nil {
		panic(err)
	}
}
