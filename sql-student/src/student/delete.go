package student

import (
	"log"
	"sql-student/src/db"
)

func DeleteById(dbConn db.Conn, lg *log.Logger, id int) error {
	query := `
		DELETE FROM STUDENT
		WHERE id = ?
	`
	_, err := dbConn.Db.Exec(query, id)
	if err != nil {
		lg.Printf("Error: %v\n", err)

	}
	return err
}
