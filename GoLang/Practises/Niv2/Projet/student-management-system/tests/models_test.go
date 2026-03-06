package tests

import (
	"time"
	"testing"
	"github.com/Atinkene/tech-learning-lab/GoLang/Practises/Niv2/student-management-system/internal/models"
)

var e = &models.Student{
		LastName :       "Bassene",
		FirstName :    "Massina",
		Email :     "massina@esp.sn",
		Phone : "+221771234567",
		DateOfBirth : time.Date(2000, 5, 10, 0, 0, 0, 0, time.UTC),
		Level :    "M2",
		Specialization :"GI",
		RegistrationDate : time.Now(),
		Active :     true,
	}

func TestEtudiantValidation(t *testing.T) {

	err := e.Validate()

	if err != nil {
		t.Errorf("Validation should success but error: %v", err)
	}
}

func TestAgeCalcul (t *testing.T) {
	e.Age()

}



func TestPrintInfos (t *testing.T) {
	e.PrintInfo()
}