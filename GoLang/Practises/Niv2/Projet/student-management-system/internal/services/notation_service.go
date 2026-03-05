/*
 *
 */


 package services


 type NotationService struct {
    storage Storage
    mu      sync.RWMutex
}


func (s *NotationService) AddGrade(grade *models.Grade) error

func (s *NotationService) CalculateCourseResults(courseID string) ([]*models.CourseResult, error)

func (s *NotationService) GenerateReportCard(studentID string, semester int) (*ReportCard, error)

type ReportCard struct {
    Student      *models.Student
    Semester      int
    Results     []*models.ResultatCours
    genAverage    float64
    TotalCredits  int
    Rank          int 
    Mention       string
}