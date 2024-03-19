package student

import (
	"log"
	"sql-student/src/db"
)

type Student struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Std        int    `json:"std"`
	Gender     string `json:"gender"`
}

func GetStudents(dbConn db.Conn, lg *log.Logger) ([]Student, error) {
	var sts []Student
	query := `
		SELECT
			id,
			first_name,
			last_name,
			std,
			gender
		FROM
			STUDENT
	`

	result, err := dbConn.Db.Query(query)
	if err != nil {
		lg.Printf("Error: %v", err)
	}

	for result.Next() {
		st := Student{}
		result.Scan(&st.Id, &st.First_name, &st.Last_name, &st.Std, &st.Gender)
		if err != nil {
			lg.Printf("Error: %v", err)
		}

		sts = append(sts, st)

	}
	return sts, err

}

func GetStudentById(dbConn db.Conn, lg *log.Logger, id int) (Student, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			std,
			gender
		FROM
			STUDENT
		WHERE
			id = ?
	`
	st := Student{}

	result := dbConn.Db.QueryRow(query, id)

	var err error
	if err = result.Scan(
		&st.Id,
		&st.First_name,
		&st.Last_name,
		&st.Std,
		&st.Gender,
	); err != nil {
		lg.Printf("Error: %v", err)
	}

	return st, err

}
