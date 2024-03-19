package db

import (
	"Database/sql"
	"fmt"
)

type Conne struct {
	Db *sql.DB
}

func Databaseconn() (Conne, error) {
	cfg := mysql.config{
		User:                "shivani",
		Passwd:              "sgour",
		Net:                 "tcp",
		Addr:                "localhost:3306",
		Database:            "test",
		AllowNativePassword: ture,
	}
	// Get a database handle.
	dbconn, err := sql.Open("mysql", cfg.FormatDSN())

	if err := dbconn.Ping(); err != nil {
		fmt.Println("Database connection failed")
		return Conne{}, err
	}

	return Conne{Db: dbconn}, err
}

func (c *Conne) Close() {
	c.Db.Close()
}
