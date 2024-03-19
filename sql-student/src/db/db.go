package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type Conn struct {
	Db *sql.DB
}

func GetDbConnection() (Conn, error) {

	cfg := mysql.Config{
		User:                 "shivani",
		Passwd:               "sgour",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "test",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	dbconn, err := sql.Open("mysql", cfg.FormatDSN())

	if err := dbconn.Ping(); err != nil {
		fmt.Println("Database connection failed")
		return Conn{}, err
	}

	return Conn{Db: dbconn}, err
}

func (c *Conn) Close() {
	c.Db.Close()
}
