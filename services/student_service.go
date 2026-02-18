// Layer (Service) -> Business logic
package services

import (
	"github.com/SupakitThammangoon/go-api-gin-lab/models"
	"github.com/SupakitThammangoon/go-api-gin-lab/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

// GetAll
func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

// GetByID
func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

// POST
func (s *StudentService) CreateStudent(student models.Student) error {
	return s.Repo.Create(student)
}

// PUT
func (s *StudentService) UpdateStudent(id string, student models.Student) (int64, error) {
	return s.Repo.UpdateByID(id, student)
}

// DELETE
func (s *StudentService) DeleteStudent(id string) (int64, error) {
	return s.Repo.DeleteByID(id)
}
