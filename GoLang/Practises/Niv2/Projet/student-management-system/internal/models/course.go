/*
 * Course struct and methods
 * This struct represents a course in the formation management system. It includes fields for course details such as ID, code, title, description, credits, semester, level, professor, maximum number of students, and current enrollment.
 * The methods associated with the Course struct provide functionality for validating course data, enrolling and unenrolling students, checking if the course is full, calculating the folding rate (enrollment percentage), and printing course information.
 */

package models

import (
	"fmt"
)


/*
 * Course represents a course in the formation management system, with details such as ID, code, title, description, credits, semester, level, professor, maximum number of students, and current enrollment.
 */

type Course struct {
	ID          string    // Format: C2024001
	Code        string    // Unique, format: CS101
	Title       string    // Mandatory, min 5 characters
	Description string    // Optional, max 500 characters
	Credits     enum["1", "2", "3", "4", "5"]       // Number of credits (1-5)
	Semester    int       // 1 or 2
	Level       enum["L1", "L2", "L3", "M1", "M2"]    // L1, L2, L3, M1, M2
	Professor   string    // Name of the professor
	MaxStudents int       // Maximum number of students (20-100)
	Enrolled    int       // Current number of enrolled students
}


/*
 * Course methods
 */


 /*
 * Validate checks if the course has valid values for its fields. It ensures that the title is at least 5 characters long, credits are between 1 and 5, semester is either 1 or 2, level is one of the allowed levels, and maximum number of students is between 20 and 100.
 */

func (c *Course) Validate() error {
	if len(c.Title) < 5 {
		return fmt.Errorf("title must be at least 5 characters")
	}
	if c.Credits < 1 || c.Credits > 5 {
		return fmt.Errorf("credits must be between 1 and 5")
	}
	if c.Semester != 1 && c.Semester != 2 {
		return fmt.Errorf("semester must be 1 or 2")
	}
	if c.Level != "L1" && c.Level != "L2" && c.Level != "L3" && c.Level != "M1" && c.Level != "M2" {
		return fmt.Errorf("level must be L1, L2, L3, M1, or M2")
	}
	if c.MaxStudents < 20 || c.MaxStudents > 100 {
		return fmt.Errorf("max students must be between 20 and 100")
	}
	return nil
}


/*
 * EnrollStudent attempts to enroll a student in the course. It checks if the course is full before enrolling.
 */

func (c *Course) EnrollStudent() error {
	if c.Enrolled >= c.MaxStudents {
		return fmt.Errorf("course is full")
	}
	c.Enrolled++
	return nil
}


/*
 * UnenrollStudent attempts to unenroll a student from the course. It checks if there are students enrolled before unenrolling.
 */

func (c *Course) UnenrollStudent() error {
	if c.Enrolled <= 0 {
		return fmt.Errorf("no students to unenroll")
	}
	c.Enrolled--
	return nil
}


/*
 * IsFull checks if the course has reached its maximum enrollment. It returns true if the number of enrolled students is greater than or equal to the maximum number of students, and false otherwise.
 */

func (c *Course) IsFull() bool {
	return c.Enrolled >= c.MaxStudents
}


/*
 * foldingRate calculates the enrollment percentage for the course. It returns the percentage of enrolled students relative to the maximum number of students. If the maximum number of students is zero, it returns 0 to avoid division by zero.
 */

func (c *Course) foldingRate() float64 {
	if c.MaxStudents == 0 {
		return 0
	}
	return (float64(c.Enrolled) / float64(c.MaxStudents)) * 100
}

/*
 * PrintInfo displays the course information in a readable format. It prints the course ID, code, title, description, credits, semester, level, professor, maximum number of students, and current enrollment.
 */

func (c *Course) PrintInfo() {
	fmt.Printf("ID: %s\n", c.ID)
	fmt.Printf("Code: %s\n", c.Code)
	fmt.Printf("Title: %s\n", c.Title)
	fmt.Printf("Description: %s\n", c.Description)
	fmt.Printf("Credits: %s\n", c.Credits)
	fmt.Printf("Semester: %d\n", c.Semester)
	fmt.Printf("Level: %s\n", c.Level)
	fmt.Printf("Professor: %s\n", c.Professor)
	fmt.Printf("Max Students: %d\n", c.MaxStudents)
	fmt.Printf("Enrolled: %d\n", c.Enrolled)
}