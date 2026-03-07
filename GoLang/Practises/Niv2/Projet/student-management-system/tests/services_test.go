package tests

import (
    "errors"
    "testing"
    "time"

    "github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/models"
    "github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/storage"
    "github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/services"
)

// MockStorage is a mock implementation of the Storage interface for testing
type MockStorage struct {
    students []*models.Student
    err      error
}

func (m *MockStorage) GetAllStudents() ([]*models.Student, error) {
    return m.students, m.err
}

func (m *MockStorage) SaveStudent(student *models.Student) error {
    return nil
}

// corrected signature to match storage.Storage
func (m *MockStorage) GetStudent(id string) (*models.Student, error) {
    return nil, nil
}

func (m *MockStorage) DeleteStudent(id string) error {
    return nil
}

func (m *MockStorage) SaveCourse(course *models.Course) error {
    return nil
}

func (m *MockStorage) GetCourse(id string) (*models.Course, error) {
    return nil, nil
}

func (m *MockStorage) GetAllCourses() ([]*models.Course, error) {
    return nil, nil
}

func (m *MockStorage) DeleteCourse(id string) error {
    return nil
}

func (m *MockStorage) SaveRegistration(registration *models.Registration) error {
    return nil
}

func (m *MockStorage) GetRegistration(id string) (*models.Registration, error) {
    return nil, nil
}

func (m *MockStorage) GetRegistrationsByStudent(studentID string) ([]*models.Registration, error) {
    return nil, nil
}

func (m *MockStorage) GetRegistrationsByCourse(courseID string) ([]*models.Registration, error) {
    return nil, nil
}

func (m *MockStorage) SaveGrade(grade *models.Grade) error {
    return nil
}

func (m *MockStorage) GetGradesByRegistration(registrationID string) ([]*models.Grade, error) {
    return nil, nil
}

// … reste des tests inchangé …

// Test: StatsByLevel returns correct count for each level
func TestStatsByLevel_CorrectCounts(t *testing.T) {
    mockStorage := &MockStorage{
        students: []*models.Student{
            {ID: "ESP1", Level: models.L1},
            {ID: "ESP2", Level: models.L1},
            {ID: "ESP3", Level: models.L2},
            {ID: "ESP4", Level: models.L3},
            {ID: "ESP5", Level: models.L1},
        },
    }

    cache := storage.NewCache(1 * time.Hour)
    service := services.NewStudentService(mockStorage, cache)

    result, err := service.StatsByLevel()

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    expectedStats := map[string]int{
        "L1": 3,
        "L2": 1,
        "L3": 1,
    }

    if len(result) != len(expectedStats) {
        t.Errorf("Expected %d levels, got %d", len(expectedStats), len(result))
    }

    for level, expectedCount := range expectedStats {
        actualCount, exists := result[level]
        if !exists {
            t.Errorf("Expected level %s to exist in results", level)
        }
        if actualCount != expectedCount {
            t.Errorf("For level %s: expected count %d, got %d", level, expectedCount, actualCount)
        }
    }
}

// Test: StatsByLevel returns empty map for no students
func TestStatsByLevel_NoStudents(t *testing.T) {
    mockStorage := &MockStorage{
        students: []*models.Student{},
    }

    cache := storage.NewCache(1 * time.Hour)
    service := services.NewStudentService(mockStorage, cache)

    result, err := service.StatsByLevel()

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if len(result) != 0 {
        t.Errorf("Expected empty map, got %d entries", len(result))
    }
}

// Test: StatsByLevel handles storage error
func TestStatsByLevel_StorageError(t *testing.T) {
    mockStorage := &MockStorage{
        err: errors.New("database connection failed"),
    }

    cache := storage.NewCache(1 * time.Hour)
    service := services.NewStudentService(mockStorage, cache)

    result, err := service.StatsByLevel()

    if err == nil {
        t.Fatal("Expected an error, got nil")
    }

    if result != nil {
        t.Errorf("Expected nil result, got %v", result)
    }

    if err.Error() != "database connection failed" {
        t.Errorf("Expected specific error message, got %v", err)
    }
}

// Test: StatsByLevel handles all level types
func TestStatsByLevel_AllLevelTypes(t *testing.T) {
    mockStorage := &MockStorage{
        students: []*models.Student{
            {ID: "ESP1", Level: models.L1},
            {ID: "ESP2", Level: models.L2},
            {ID: "ESP3", Level: models.L3},
            {ID: "ESP4", Level: models.M1},
            {ID: "ESP5", Level: models.M2},
        },
    }

    cache := storage.NewCache(1 * time.Hour)
    service := services.NewStudentService(mockStorage, cache)

    result, err := service.StatsByLevel()
	
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    levels := []string{"L1", "L2", "L3", "M1", "M2"}
    for _, level := range levels {
        if _, exists := result[level]; !exists {
            t.Errorf("Expected level %s to be in results", level)
        }
        if result[level] != 1 {
            t.Errorf("Expected count 1 for level %s, got %d", level, result[level])
        }
    }
}

// Test: StatsByLevel handles large dataset
func TestStatsByLevel_LargeDataset(t *testing.T) {
    var students []*models.Student
    for i := 0; i < 1000; i++ {
        level := models.L1
        if i%5 == 0 {
            level = models.L2
        } else if i%7 == 0 {
            level = models.L3
        }
        students = append(students, &models.Student{
            ID:    "ESP" + string(rune(i)),
            Level: level,
        })
    }

    mockStorage := &MockStorage{
        students: students,
    }

    cache := storage.NewCache(1 * time.Hour)
    service := services.NewStudentService(mockStorage, cache)

    result, err := service.StatsByLevel()

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    totalCount := 0
    for _, count := range result {
        totalCount += count
    }

    if totalCount != len(students) {
        t.Errorf("Expected total count %d, got %d", len(students), totalCount)
    }
}

// Test: StatsByLevel with single student
func TestStatsByLevel_SingleStudent(t *testing.T) {
    mockStorage := &MockStorage{
        students: []*models.Student{
            {ID: "ESP1", Level: models.L1},
        },
    }

    cache := storage.NewCache(1 * time.Hour)
    service := services.NewStudentService(mockStorage, cache)

    result, err := service.StatsByLevel()

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if result["L1"] != 1 {
        t.Errorf("Expected count 1 for L1, got %d", result["L1"])
    }
}

// Test: StatsByLevel with only M2 students
func TestStatsByLevel_OnlyMastersDegree(t *testing.T) {
    mockStorage := &MockStorage{
        students: []*models.Student{
            {ID: "ESP1", Level: models.M2},
            {ID: "ESP2", Level: models.M2},
            {ID: "ESP3", Level: models.M1},
        },
    }

    cache := storage.NewCache(1 * time.Hour)
    service := services.NewStudentService(mockStorage, cache)

    result, err := service.StatsByLevel()

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if result["M2"] != 2 {
        t.Errorf("Expected count 2 for M2, got %d", result["M2"])
    }
    if result["M1"] != 1 {
        t.Errorf("Expected count 1 for M1, got %d", result["M1"])
    }
}