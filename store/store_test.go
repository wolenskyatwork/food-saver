package store

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	store *dbStore
	db *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	connString := "dbname=db_test sslmode=disable"

}
