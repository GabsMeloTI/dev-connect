package db_postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"treads/infra/database"
)

func ConnDB(config *database.Config, appRunMigration bool) *sql.DB {
	var db *sql.DB
	driver := config.Driver
	dsn := config.Driver + "://" + config.User + ":" + config.Password + "@" +
		config.Host + ":" + config.Port + "/" + config.Database + config.SSLMode
	db, err := sql.Open(driver, dsn)
	if err != nil {
		errConnection(config.Environment, err)
	}

	if appRunMigration {
		if errM := runMigrations(db, config.Environment); errM != nil {
			errConnection(config.Environment, errM)
		}
	}

	return db
}

func errConnection(environment string, err error) {
	panic("failed to connect " + environment + " postgres database:" + err.Error())
}
