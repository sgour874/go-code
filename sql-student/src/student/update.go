package student

import (
	"log"
	"sql-student/src/db"
)

func UpdateById(dbConn db.Conn, lg *log.Logger, id int, stud Student) error {
	query := `
		Update  STUDENT
		SET id = ?,
		first_name = ?,
		last_name = ?,
		std = ?,
		gender = ?
		WHERE id = ?
	`
	_, err := dbConn.Db.Exec(query, stud.Id, stud.First_name, stud.Last_name, stud.Std, stud.Gender, id)
	if err != nil {
		lg.Printf("Error: %v\n", err)
	}

	return err
}
