package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/wolenskyatwork/food-saver/dao"
	"testing"
)

type StoreSuite struct {
	suite.Suite
	store *DBStore
	db *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	connString := "dbname=life_logger_test sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &DBStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	_, err := s.db.Query("DELETE FROM activity_name")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateActivity() {
	s.store.CreateActivity(&dao.Activity{
		Name: "test name",
	})

	res, err := s.db.Query(`SELECT COUNT(*) FROM activity_name WHERE name='test name';`)
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
