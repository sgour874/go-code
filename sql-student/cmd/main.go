package main

import (
	"log"
	"os"
	"sql-student/src/apis"
	"sql-student/src/db"

	"github.com/gin-gonic/gin"
)

func setupAPI(api *apis.Apis) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/student/:id", api.GetStudentById)
	router.GET("/student", api.GetAllStudents)
	router.PUT("/student/:id", api.UpdateStudentById)
	router.POST("/student", api.InsertStudent)
	router.DELETE("/student/:id", api.DeleteStudentById)
	router.Run("localhost:8080")

}

func main() {
	conn, _ := db.GetDbConnection()
	file, _ := os.OpenFile("build/student.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	lg := log.New(file, "", log.Lmicroseconds)

	api := apis.GetApis(conn, lg)
	setupAPI(&api)
}
