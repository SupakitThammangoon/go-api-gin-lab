// วิธี run ใช้คำสั่ง go run . ได้เลยมันจะรันเป็นทั้ง package
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/SupakitThammangoon/go-api-gin-lab/config"
	"github.com/SupakitThammangoon/go-api-gin-lab/handlers"
	"github.com/SupakitThammangoon/go-api-gin-lab/repositories"
	"github.com/SupakitThammangoon/go-api-gin-lab/services"
)

func main() {
	db := config.InitDB()

	repo := &repositories.StudentRepository{DB: db}
	service := &services.StudentService{Repo: repo}
	handler := &handlers.StudentHandler{Service: service}

	r := gin.Default()

	r.GET("/students", handler.GetStudents)
	r.GET("/students/:id", handler.GetStudentByID)
	r.POST("/students", handler.CreateStudent)
	r.PUT("/students/:id", handler.UpdateStudent)
	r.DELETE("/students/:id", handler.DeleteStudent)

	r.Run(":8080")
}
