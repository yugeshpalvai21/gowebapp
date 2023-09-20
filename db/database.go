package db

import (
	"database/sql"
	"fmt"
	"gowebapp/config" // import database config details

	_ "github.com/lib/pq" // postgrres adaptor module
)

// intiate dabase connection
func InitDB() (*sql.DB, error) {
	dbConfig := config.GetDatabaseConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	// initiate database connection through postgres adaptor
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	// make connection with database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// return database instance
	return db, nil
}
