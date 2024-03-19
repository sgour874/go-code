package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"sql-student/src/db"
	"sql-student/src/student"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Apis struct {
	conn db.Conn
	lg   *log.Logger
}

func GetApis(conn db.Conn, lg *log.Logger) Apis {
	return Apis{
		conn: conn,
		lg:   lg,
	}
}

func (a *Apis) GetAllStudents(c *gin.Context) {

	stud, err := student.GetStudents(a.conn, a.lg)
	if err != nil {
		a.lg.Printf("Error: Can't get student detail %v", err)
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, stud)

}

func (a *Apis) GetStudentById(c *gin.Context) {

	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		a.lg.Printf("Error: in parsing the string to integer")
		c.JSON(http.StatusBadRequest, nil)
	}

	stud, err := student.GetStudentById(a.conn, a.lg, id)
	if err != nil {
		a.lg.Printf("Error: Can't get student detail %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, stud)

}

func (a *Apis) UpdateStudentById(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		a.lg.Printf("Error: in parsing the string to integer")
		c.JSON(http.StatusBadRequest, nil)
	}

	stud := student.Student{}

	if err := json.NewDecoder(c.Request.Body).Decode(&stud); err != nil {
		log.Printf("Error: Unable to decode")
	}

	if err := student.UpdateById(a.conn, a.lg, id, stud); err != nil {
		log.Printf("Error: Unable to update : id: %d data: %v", id, stud)
	}

	c.JSON(http.StatusOK, stud)

}
func (a *Apis) InsertStudent(c *gin.Context) {

	stud := student.Student{}

	if err := json.NewDecoder(c.Request.Body).Decode(&stud); err != nil {
		log.Printf("Error : Unable to Create Student")
	}

	if err := student.InsterStudent(a.conn, a.lg, stud); err != nil {
		log.Printf("Error : Unable to Insert Data : %v", stud)
	}
	c.JSON(http.StatusOK, stud)
}

func (a *Apis) DeleteStudentById(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		log.Printf("Error => in parsing the string to integer")
		c.JSON(http.StatusBadRequest, nil)
	}
	err = student.DeleteById(a.conn, a.lg, id)
	if err != nil {
		a.lg.Printf("Error: Can't get student detail %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, id)

}
