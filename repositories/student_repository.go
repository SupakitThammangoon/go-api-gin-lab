// Layer (Repositories) -> Database access
package repositories

import (
	"database/sql"

	"github.com/SupakitThammangoon/go-api-gin-lab/models"
)

type StudentRepository struct {
	DB *sql.DB
}

// GetAll
func (r *StudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.DB.Query("SELECT id, name, major, gpa FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		rows.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
		students = append(students, s)
	}
	return students, nil
}

// GetByID
func (r *StudentRepository) GetByID(id string) (*models.Student, error) {
	row := r.DB.QueryRow(
		"SELECT id, name, major, gpa FROM students WHERE id = ?",
		id,
	)

	var s models.Student
	err := row.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// POST
func (r *StudentRepository) Create(s models.Student) error {
	_, err := r.DB.Exec(
		"INSERT INTO students (id, name, major, gpa) VALUES (?, ?, ?, ?)",
		s.Id, s.Name, s.Major, s.GPA,
	)
	return err
}

// PUT
func (r *StudentRepository) UpdateByID(id string, s models.Student) (int64, error) {
	res, err := r.DB.Exec(
		"UPDATE students SET name = ?, major = ?, gpa = ? WHERE id = ?",
		s.Name, s.Major, s.GPA, id,
	)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// DELETE
func (r *StudentRepository) DeleteByID(id string) (int64, error) {
	res, err := r.DB.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
