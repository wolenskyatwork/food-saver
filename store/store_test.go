package store

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"github.com/wolenskyatwork/food-saver/handler"
	"testing"
	_ "github.com/lib/pq"
)

type StoreSuite struct {
	suite.Suite
	store *dbStore
	db *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	connString := "dbname=life_logger_test sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	_, err := s.db.Query("DELETE FROM activity")
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
	s.store.CreateActivity(&handler.Activity{
		Name: "test name",
	})

	res, err := s.db.Query(`SELECT COUNT(*) FROM activities WHERE name='test name';`)
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
