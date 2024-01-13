package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/tomazcx/screen-seat-api/config"
)

var dbConn *sql.DB 

func GetDbConnection() *sql.DB{
	return dbConn
}

func ConnectToDB(conf *config.Cfg) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName)
	dbConn, _ = sql.Open("postgres", connStr)
	err := dbConn.Ping()

	if err != nil {
		log.Fatalf("Failed to ping to the database: %v", err)
		return nil, err
	}

	return dbConn, nil
}
