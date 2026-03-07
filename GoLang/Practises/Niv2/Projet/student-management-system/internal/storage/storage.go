/**
 * This file defines the Storage interface for managing students, courses, registrations, and grades.
 * It includes methods for saving, retrieving, and deleting data related to these entities.
 * The Storage interface abstracts the underlying data storage mechanism, allowing for flexibility
 * in how data is stored (e.g., in-memory, JSON files, databases).
 */

package storage

import 	"github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/models"


type Storage interface {
	SaveStudent(student *models.Student) error
	GetStudent(id string) (*models.Student, error)
	GetAllStudents() ([]*models.Student, error)
	DeleteStudent(id string) error

	SaveCourse(course *models.Course) error
	GetCourse(id string) (*models.Course, error)
	GetAllCourses() ([]*models.Course, error)
	DeleteCourse(id string) error

	SaveRegistration(registration *models.Registration) error
	GetRegistration(id string) (*models.Registration, error)
	GetRegistrationsByStudent(studentID string) ([]*models.Registration, error)
	GetRegistrationsByCourse(courseID string) ([]*models.Registration, error)

	SaveGrade(grade *models.Grade) error
	GetGradesByRegistration(registrationID string) ([]*models.Grade, error)
}

