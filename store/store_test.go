package store

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"database/sql"
	. "github.com/smartystreets/goconvey/convey"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/wolenskyatwork/food-saver/dao"
	"reflect"
	"testing"
)

type StoreSuite struct {
	suite.Suite
	store *DBStore
}

func (s *StoreSuite) SetupSuite() {
	// TODO read this from the config
	connString := "dbname=life_logger_test sslmode=disable"
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
	_, err := s.store.DB.Exec("INSERT INTO activity_name (name) VALUES ('medicine'),('workout'),('climbing'),('biking'),('archery'),('spartan');")
	if err != nil {
		s.T().Fatal(err)
	}

	-, err := s.store.DB.Exec("INSERT INTO app_user (username) VALUES ('sailorflares'), ('closgmr');")
	if err != nil {
		s.T().Fatal(err)
	}

	_, err = s.store.DB.Exec("INSERT INTO user_activity (activity_id, app_user_id) VALUES (1,1), (2,1), (3,1), (4,2);")
	if err != nil {
		s.T().Fatal(err)
	}

	got, err := s.store.GetActivityNames("1")
	if err != nil {
		s.T().Fatal(err)
	}

	wanted := []*dao.ActivityName{
		{Name: "medicine", Id: "1" },
		{Name: "workout", Id: "2" },
		{Name: "climbing", Id: "3" },
	}

	if len(got) != len(wanted) {
		s.T().Errorf("incorrect count, wanted %d, got %d", len(wanted), len(got))
	}

	if !reflect.DeepEqual(wanted, got) {
		So(wanted, ShouldResemble, got)
		// s.T().Errorf("Slices do not match, wanted %s, got %s", wanted, got)
	}

	got, err = s.store.GetActivityNames("2")
	if err != nil {
		s.T().Fatal(err)
	}

	wanted = []*dao.ActivityName{
		{Name: "biking", Id: "4" },
	}

	if len(got) != len(wanted) {
		s.T().Errorf("incorrect count, wanted %d, got %d", len(wanted), len(got))
	}

	if !reflect.DeepEqual(wanted, got) {
		s.T().Errorf("Slices do not match, wanted %s, got %s", wanted, got)
	}
}

func (s *StoreSuite) TestCreateActivityName() {
	s.store.CreateActivityName(&dao.ActivityName{
		Name: "test name",
	})

	res, err := s.store.DB.Query(`SELECT COUNT(*) FROM activity_name WHERE name='test name';`)
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
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
