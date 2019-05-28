package dao

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type PostgresSQLSuite struct {
	suite.Suite
	db 				*sql.DB
}

func (p *PostgresSQLSuite) SetupSuite() {
	// TODO read this from the config
	connString := "host=localhost port=5432 user=postgres dbname=mims sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	createTestDatabase(db)
}

func createTestDatabase(db *sql.DB) {
	_, err := db.Exec("drop schema public cascade; create schema public;")
	if err != nil {
		panic(err)
	}

	conf, err := goose.NewDBConf("./", "testing", "")
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
