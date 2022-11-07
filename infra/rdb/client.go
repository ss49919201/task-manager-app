package rdb

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", getDataSourceName())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDBXorm() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", getDataSourceName())
	if err != nil {
		return nil, err
	}

	// TODO: Configure connection pool
	logger := log.NewSimpleLogger(os.Stdout)
	logger.ShowSQL(true)
	engine.SetLogger(logger)

	return engine, nil
}

func getDataSourceName() string {
	cnf := mysql.Config{
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "database",
		Collation: "utf8mb4_general_ci",
		Loc:       time.UTC,
		User:      os.Getenv("DATABASE_USER"),
		Passwd:    os.Getenv("DATABASE_PASSWORD"),
		// Addr:   os.Getenv("DATABASE_HOST"),
		// DBName: os.Getenv("DATABASE_NAME"),
	}
	fmt.Println(cnf.FormatDSN())
	return cnf.FormatDSN()
}
