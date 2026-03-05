/*
 * File: registration.go
 * Description: This file defines the Registration struct and its associated methods for managing course registrations.
 * The Registration struct includes fields for student ID, course ID, registration date, status, academic year, and semester.
 * Methods include validation of registration data, status management (abandonment), and utility functions to check if a registration is active and to calculate its duration.
 */

package models

import (
	"fmt"
	"time"
)


/*
 * Registration represents a course registration in the formation management system, with details such as student ID, course ID, registration date, status, academic year, and semester.
 */

type Registration struct {
	ID           string    // Auto-generated
	StudentID    string    // Reference to Student
	CourseID     string    // Reference to Course
	RegistrationDate time.Time // Date of registration
	Status       enum["pending", "approved", "rejected"]    // "en_cours", "validé", "abandonné"
	Year         int       // Academic year (2024)
	Semester     enum[1, 2]       // 1 or 2
}	


/* * Registration methods
 */


/*
 * Validate checks if the registration has valid values for its fields. It ensures that the student ID and course ID are not empty, the academic year is reasonable, and the semester is either 1 or 2.
 */

func (r *Registration) Validate() error {
	if r.StudentID == "" {
		return fmt.Errorf("student ID is required")
	}
	if r.CourseID == "" {
		return fmt.Errorf("course ID is required")
	}
	if r.Year < 2000 || r.Year > time.Now().Year() {
		return fmt.Errorf("invalid academic year")
	}
	if r.Semester != 1 && r.Semester != 2 {
		return fmt.Errorf("semester must be 1 or 2")
	}
	return nil
}


/*
 * validate_Status checks if the registration status is valid. It ensures that the status is one of the allowed values ("pending", "approved", "rejected").
 */
func (r *Registration) validate_Status() error {
	if r.Status != "pending" && r.Status != "approved" && r.Status != "rejected" {
		return fmt.Errorf("invalid status")
	}
	return nil
}


/*
 * Abandon sets the registration status to "rejected", indicating that the student has abandoned the course.
 */

func (r *Registration) Abandon() error {
	r.Status = "rejected"
	return nil
}


/*
 * IsActive checks if the registration is active, meaning its status is either "pending" or "approved".
 */
 
func (r *Registration) IsActive() bool {
	return r.Status == "pending" || r.Status == "approved"
}


/*
 * Duration calculates the duration of the registration from the registration date to the current time.
 */
 
func (r *Registration) Duration() time.Duration {
	return time.Since(r.RegistrationDate)
}
