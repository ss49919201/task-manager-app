package infra

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", getDataSourceName())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDataSourceName() string {
	cnf := mysql.Config{
		Net:                  "tcp",
		Addr:                 os.Getenv("DATABASE_HOST"),
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		DBName:               os.Getenv("DATABASE_NAME"),
		Loc:                  time.UTC,
		Collation:            "utf8mb4_general_ci",
		AllowNativePasswords: true,
	}
	return cnf.FormatDSN()
}
