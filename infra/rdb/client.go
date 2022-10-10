package rdb

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func NewDB() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", getDataSourceName())
	if err != nil {
		return nil, err
	}
	// TODO: Configure connection pool
	return engine, nil
}

func getDataSourceName() string {
	// Set data source name
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}
