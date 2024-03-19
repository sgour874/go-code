package student

import (
	"log"
	"sql-student/src/db"
)

func InsterStudent(dbConn db.Conn, lg *log.Logger, stud Student) error {
	query := `
		INSERT INTO 
		STUDENT (
			id,
			first_name,
			last_name,
			std,
			gender
		) VALUES (
			?,
			?,
			?,
			?,
			?
		)	
	`
	_, err := dbConn.Db.Exec(query, stud.Id, stud.First_name, stud.Last_name, stud.Std, stud.Gender)
	if err != nil {
		lg.Printf("Error: %v\n", err)
	}

	return err
}
