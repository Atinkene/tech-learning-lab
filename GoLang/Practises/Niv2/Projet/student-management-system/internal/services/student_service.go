/*
 * StudentService provides methods to manage students, including creating new students and searching for students by name. It uses a storage interface to persist student data and a cache for faster retrieval of frequently accessed data.
 * The CreateStudent method generates a unique ID and matricule for each new student, validates the student data, saves it to storage, and caches the result. The SearchStudentByName method first checks the cache for results before querying the storage.
 */

package services

import (
	"sync"
	"time"
	"github.com/massina/gestion-formation/internal/models"
	"github.com/massina/gestion-formation/internal/storage"
)


/*
 * Cache is a simple in-memory cache implementation with TTL (Time To Live) functionality. It allows storing key-value pairs with an expiration time, and provides methods to set, get, delete, and clear cache entries. Additionally, it includes a cleanup mechanism to remove expired entries at regular intervals.
 */
/*
 * Storage is an interface that defines methods for managing students, courses, registrations, and grades. It includes methods for saving, retrieving, and deleting data related to these entities. The Storage interface abstracts the underlying data storage mechanism, allowing for flexibility in how data is stored (e.g., in-memory, JSON files, databases).
 */

/*
 * StudentService provides methods to manage students, including creating new students and searching for students by name. It uses a storage interface to persist student data and a cache for faster retrieval of frequently accessed data.
 * The CreateStudent method generates a unique ID and matricule for each new student, validates the student data, saves it to storage, and caches the result. The SearchStudentByName method first checks the cache for results before querying the storage.
 */

type StudentService struct {
	storage Storage
	cache   *Cache
	mu	  sync.RWMutex
}


/*
 * NewStudentService creates a new instance of StudentService with the provided storage and cache. It initializes the service with the necessary dependencies for managing student data.
 */

func NewStudentService(storage Storage, cache *Cache) *StudentService {
	return &StudentService{
		storage: storage,
		cache:   cache,
	}
}


/*
 * genererID generates a unique ID for a student in the format "ESP2024001". It uses the current time in nanoseconds to ensure uniqueness.
 * genererMatricule generates a unique matricule for a student in the format "M2024001". Similar to genererID, it uses the current time in nanoseconds for uniqueness.
 */

func genererID() string {
	// Generate a unique ID (ex: ESP2024001)
	return fmt.Sprintf("ESP%d", time.Now().UnixNano())
}


/*
 * genererMatricule generates a unique matricule for a student in the format "M2024001". Similar to genererID, it uses the current time in nanoseconds for uniqueness.
 */

func genererMatricule() string {	// Generate a unique matricule (ex: M2024001)
	return fmt.Sprintf("M%d", time.Now().UnixNano())
}


/*
 * CreateStudent creates a new student in the system. It generates a unique ID and matricule for the student, sets the registration date and active status, validates the student data, saves it to storage, and caches the result.
 * The method ensures that the student data is valid before saving, and it handles concurrency with a mutex to protect access to the storage and cache.
 */

func (s *StudentService) CreateStudent(student *models.Student) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	student.ID = genererID()
	student.RegistrationNumber = genererMatricule()
	student.RegistrationDate = time.Now()
	student.Active = true

	if err := student.Validate(); err != nil {
		return err
	}

	if err := s.storage.SaveStudent(student); err != nil {
		return err
	}

	s.cache.Set(student.ID, student)
	return nil
}


/*
 * SearchStudentByName searches for students by their name. It first checks the cache for results, and if not found, it queries the storage. The results are then cached for future requests.
 * The method uses a mutex to protect access to the cache and storage during the search operation.
 */

func (s *StudentService) SearchStudentByName(name string) ([]*models.Student, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cached, found := s.cache.Get("search:" + name)
	if found {
		return cached.([]*models.Student), nil
	}

	students, err := s.storage.GetStudent(name)
	if err != nil {
		return nil, err
	}

	s.cache.Set("search:" + name, students)
	return students, nil
}


/*
 * GetStudentByLevel 
 */
 

func (s *StudentService) GetStudentByLevel(level string) ([]*models.Student, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cached, found := s.cache.Get("level:" + level)
	if found {
		return cached.([]*models.Student), nil
	}

	students, err := s.storage.GetAllStudents()
	studentsByLevel []*models.Student

	for _, student := range students {
		if strings.EqualFold(student.Level, level) {
			studentsByLevel := student
		}
	}

	s.cache.Set("level" + level, studentsByLevel)

	return studentsByLevel, nil
}


/*
 * GetStudentBySpecialization
 */

func (s *StudentService) GetStudentBySpecialization(specialization Specialization) ([]*models.Student, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cached, found := s.cache.Get("Specialization" + specialization)
	if found {
		return cached.([]*models.Student), nil
	}

	students, err := s.storage.GetAllStudents()
	studentsBySpecialization []*models.Student

	for _, student := range students {
		if strings.EqualFold(student.Specialization, specialization) {
			studentsBySpecialization := student
		}
	}

	s.cache.Set("level" + level, studentsBySpecialization)

	return studentsBySpecialization, nil

}


// func (s *EtudiantService) StatsByLevel() map[string]int      {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()

// 	students, err := s.storage.GetAllStudent()
// 	return students, nil

// }
// func (s *EtudiantService) StatsBySpecialization() map[string]int {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()

// 	students, err := s.storage.GetAllStudent()
// 	return students, nil

// }
// func (s *EtudiantService) ActiveStudentsStats() int {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()

// 	students, err := s.storage.GetAllStudent()
// 	return students, nil

// }