/*
 * Student struct and related methods
 * This struct represents a student in the system, with fields for ID, registration number, name, email, phone, date of birth, level, specialization, registration date, and active status.
 * The methods include validation for the student's data, calculating age, checking if the student is an adult, activating/deactivating the student, and printing the student's information.
 */

package models

import (
	"fmt"
	"regexp"
	"time"
)


/*
 * Student represents a student in the formation management system, with details such as ID, registration number, name, email, phone, date of birth, level, specialization, registration date, and active status.
 */

type Student struct {
	ID                 string    // Format: ESP2024001
	RegistrationNumber string    // Unique, format: M2024001
	LastName           string    // Mandatory, min 2 characters
	FirstName          string    // Mandatory, min 2 characters
	Email              string    // Mandatory, valid format (xxx@esp.sn)
	Phone              string    // Format: +221XXXXXXXXX
	DateOfBirth        time.Time // Date of birth
	Level              Level    // L1, L2, L3, M1, M2
	Specialization     string    // GI, RT, GM, GC, etc.
	RegistrationDate   time.Time // Date of registration
	Active             bool      // Active/inactive status
}


/*
 * Validate checks if the student has valid values for its fields. It ensures that the last name and first name are not too short, the email format is valid, and the phone format is valid.
 */

func (s *Student) Validate() error {
	if len(s.LastName) < 2 {
		return fmt.Errorf("last name must be at least 2 characters")
	}
	if len(s.FirstName) < 2 {
		return fmt.Errorf("first name must be at least 2 characters")
	}
	if !isValidEmail(s.Email) {
		return fmt.Errorf("invalid email format")
	}
	if !isValidPhone(s.Phone) {
		return fmt.Errorf("invalid phone format")
	}
	return nil
}


/*
 * isValidEmail checks if the email has a valid format.
 */
 
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@esp\.sn$`)
	return re.MatchString(email)
}


/*
 * isValidPhone checks if the phone number has a valid format.
 */

func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\+221\d{9}$`)
	return re.MatchString(phone)
}


/*
 * Age calculates the age of the student based on their date of birth.
 */
func (s * Student) Age() int {
	now := time.Now()
	age := now.Year() - s.DateOfBirth.Year()
	if now.YearDay() < s.DateOfBirth.YearDay() {
		age--
	}
	return age
}


/*
 * IsAdult checks if the student is an adult (18 years or older).
 */

func (s *Student) IsAdult() bool {
	return s.Age() >= 18
}


/*
 * Activate activates the student.
 */

func (s *Student) Activate() {
	s.Active = true
}


/* * Deactivate deactivates the student.
 */

func (s *Student) Deactivate() {
	s.Active = false
}


/* * PrintInfo displays the student's information in a readable format. It prints the student's ID, registration number, name, email, phone, date of birth, level, specialization, registration date, and active status.
 */
 
func (s *Student) PrintInfo() {
	fmt.Printf("ID: %s\n", s.ID)
	fmt.Printf("Registration Number: %s\n", s.RegistrationNumber)
	fmt.Printf("Name: %s %s\n", s.FirstName, s.LastName)
	fmt.Printf("Email: %s\n", s.Email)
	fmt.Printf("Phone: %s\n", s.Phone)
	fmt.Printf("Date of Birth: %s\n", s.DateOfBirth.Format("2006-01-02"))
	fmt.Printf("Level: %s\n", s.Level)
	fmt.Printf("Specialization: %s\n", s.Specialization)
	fmt.Printf("Registration Date: %s\n", s.RegistrationDate.Format("2006-01-02"))
	fmt.Printf("Active: %t\n", s.Active)
}