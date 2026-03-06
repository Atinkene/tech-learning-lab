/*
 *
 */

 package services

 import (
	"sync"
	"github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/models"
	"github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/storage"
 )

type CourseService struct {
	storage storage.Storage
	cache   *storage.Cache
	mu	  sync.RWMutex
}


/*
 * NewCourseService creates a new instance of CourseService with the provided storage and cache. It initializes the service with the necessary dependencies for managing course data.
 */

func NewCourseService(storage storage.Storage, cache *storage.Cache) *CourseService {
	return &CourseService{
		storage: storage,
		cache:   cache,
	}
}


func (c *CourseService) CreateCourse (course *models.Course) error
func (c *CourseService) GetCoursesByLevel (level string) ([]*models.Course, error)
func (c *CourseService) GetAvailableCourses () ([]*models.Course, error)
func (c *CourseService) GetRegistrationsStats () map[string]int
