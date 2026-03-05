/*
 *
 */

 package services


type CourseService struct {
	storage Storage
	cache   *Cache
	mu	  sync.RWMutex
}


/*
 * NewCourseService creates a new instance of CourseService with the provided storage and cache. It initializes the service with the necessary dependencies for managing course data.
 */

func NewCourseService(storage Storage, cache *Cache) *CourseService {
	return &CourseService{
		storage: storage,
		cache:   cache,
	}
}


func (c *CourseService) CreateCourse (course *models.Cours) error
func (c *CourseService) GetCoursesByLevel (level string) ([]*models.Cours, error)
func (c *CourseService) GetAvailableCourses () ([]*models.Cours, error)
func (c *CourseService) GetRegistrationsStats () map[string]int
