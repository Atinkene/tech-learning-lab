/**
 * JSONStorage is a simple file-based storage implementation for managing students, courses, registrations, and grades.
 * It uses JSON files to persist data and provides methods for CRUD operations on each entity.
 * Concurrency is handled with a RWMutex to ensure thread safety during read/write operations.
 */

package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/models"
)

type JSONStorage struct {
	dataDir      string
	students    map[string]*models.Student
	courses        map[string]*models.Course
	registrations map[string]*models.Registration
	grades        map[string]*models.Grade
	mu           sync.RWMutex // Concurrency protection
}

/**
 * Constructor and data loading/saving methods
 */


func NewJSONStorage(dataDir string) (*JSONStorage, error) {
	storage := &JSONStorage{
		dataDir:      dataDir,
		students:    make(map[string]*models.Student),
		courses:        make(map[string]*models.Course),
		registrations: make(map[string]*models.Registration),
		grades:        make(map[string]*models.Grade),
	}
	if err := storage.LoadData(); err != nil {
		return nil, err
	}
	return storage, nil
}

func (s *JSONStorage) LoadData() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	load := func(path string, target interface{}) error {
		file, err := os.Open(path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return nil // No data file, start with empty data
			}
			return err
		}
		defer file.Close()

		return json.NewDecoder(file).Decode(target)
	}
	if err := load(filepath.Join(s.dataDir, "students.json"), &s.students); err != nil {
		return err	
	}
	if err := load(filepath.Join(s.dataDir, "courses.json"), &s.courses); err != nil {
		return err
	}
	if err := load(filepath.Join(s.dataDir, "registrations.json"), &s.registrations); err != nil {
		return err
	}
	if err := load(filepath.Join(s.dataDir, "grades.json"), &s.grades); err != nil {
		return err
	}
	return nil
}



func (s *JSONStorage) SaveData() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	save := func(path string, data interface{}) error {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		
		return json.NewEncoder(file).Encode(data)
	}
	if err := save(filepath.Join(s.dataDir, "students.json"), s.students); err != nil {
		return err
	}
	if err := save(filepath.Join(s.dataDir, "courses.json"), s.courses); err != nil {
		return err
	}
	if err := save(filepath.Join(s.dataDir, "registrations.json"), s.registrations); err != nil {
		return err
	}
	if err := save(filepath.Join(s.dataDir, "grades.json"), s.grades); err != nil {
		return err
	}
	return nil
}

/**
 * Student methods
 */

func (s *JSONStorage) SaveStudent(student *models.Student) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.students[student.ID] = student
	return s.SaveData()
}

func (s *JSONStorage) GetStudent(id string) (*models.Student, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	student, exists := s.students[id]
	if !exists {
		return nil, fmt.Errorf("student not found")
	}
	return student, nil
}

func (s *JSONStorage) GetAllStudents() ([]*models.Student, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	students := make([]*models.Student, 0, len(s.students))
	for _, student := range s.students {
		students = append(students, student)
	}
	return students, nil
}

func (s *JSONStorage) DeleteStudent(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.students, id)
	return s.SaveData()
}


/**
 * Course methods
 */


func (s *JSONStorage) SaveCourse(course *models.Course) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.courses[course.ID] = course
	return s.SaveData()
}

func (s *JSONStorage) GetCourse(id string) (*models.Course, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	course, exists := s.courses[id]
	if !exists {
		return nil, fmt.Errorf("course not found")
	}
	return course, nil
}

func (s *JSONStorage) GetAllCourses() ([]*models.Course, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	courses := make([]*models.Course, 0, len(s.courses))
	for _, course := range s.courses {
		courses = append(courses, course)
	}
	return courses, nil
}

func (s *JSONStorage) DeleteCourse(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.courses, id)
	return s.SaveData()
}	


func (s *JSONStorage) SaveRegistration(registration *models.Registration) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.registrations[registration.ID] = registration
	return s.SaveData()
}

func (s *JSONStorage) GetRegistration(id string) (*models.Registration, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	registration, exists := s.registrations[id]
	if !exists {
		return nil, fmt.Errorf("registration not found")
	}
	return registration, nil
}

func (s *JSONStorage) GetRegistrationsByStudent(studentID string) ([]*models.Registration, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var registrations []*models.Registration
	for _, reg := range s.registrations {
		if reg.StudentID == studentID {
			registrations = append(registrations, reg)
		}
	}
	return registrations, nil
}

func (s *JSONStorage) GetRegistrationsByCourse(courseID string) ([]*models.Registration, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var registrations []*models.Registration
	for _, reg := range s.registrations {
		if reg.CourseID == courseID {
			registrations = append(registrations, reg)
		}
	}
	return registrations, nil
}


/**
 * Grade methods
 */


func (s *JSONStorage) SaveGrade(grade *models.Grade) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.grades[grade.ID] = grade
	return s.SaveData()
}

func (s *JSONStorage) GetGradesByRegistration(registrationID string) ([]*models.Grade, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var grades []*models.Grade
	for _, grade := range s.grades {
		if grade.RegistrationID == registrationID {
			grades = append(grades, grade)
		}
	}
	return grades, nil
}