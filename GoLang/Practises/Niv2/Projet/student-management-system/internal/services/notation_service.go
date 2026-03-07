/*
 *
 */


 package services

 import (
    "sync"
    "github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/models"
	"github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/storage"

 )

 type NotationService struct {
    storage storage.Storage
    mu      sync.RWMutex
}


func (s *NotationService) AddGrade(grade *models.Grade) error {
    return nil
}

func (s *NotationService) CalculateCourseResults(courseID string) (/*[]*models.CourseResult,*/ error) {
    return nil
}

func (s *NotationService) GenerateReportCard(studentID string, semester int) (/**ReportCard,*/ error) {
    return nil
}

type ReportCard struct {
    Student      *models.Student
    Semester      int
    Results     []*models.CourseResult
    genAverage    float64
    TotalCredits  int
    Rank          int 
    Mention       string
}