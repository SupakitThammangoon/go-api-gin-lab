// Layer (Handler) -> Handle HTTP request/response
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SupakitThammangoon/go-api-gin-lab/models"
	"github.com/SupakitThammangoon/go-api-gin-lab/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

// GetAll
func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

// GetByID
func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.Service.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// POST
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add check validation: ID, Name, GPA
	if student.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must not be empty"})
		return
	}
	if student.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name must not be empty"})
		return
	}
	if student.GPA < 0.00 || student.GPA > 4.00 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gpa must be between 0.00 and 4.00"})
		return
	}

	if err := h.Service.CreateStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// PUT
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// check validation: Name, GPA
	if student.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name must not be empty"})
		return
	}
	if student.GPA < 0.00 || student.GPA > 4.00 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gpa must be between 0.00 and 4.00"})
		return
	}

	// check จาก return 0 ที่ repositories (RowsAffected)
	affected, err := h.Service.UpdateStudent(id, student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// return value after updated
	updated, err := h.Service.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DELETE
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	// check จาก return 0 ที่ repositories (RowsAffected)
	affected, err := h.Service.DeleteStudent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
