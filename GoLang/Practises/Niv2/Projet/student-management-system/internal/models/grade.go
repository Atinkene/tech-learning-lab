/*
 * This file defines the Grade and CourseResult models, along with methods for validation and average calculation.
 * The Grade struct represents a single grade for a course registration, while the CourseResult struct aggregates
 * grades for a course and determines the average and status of the student in that course.
 */

package models

import (
	"fmt"
	"time"
)


/*
 * Grade represents a single grade for a course registration, including the type of evaluation, value, coefficient, and date.
 */

type Grade struct {
	ID            string    // Auto-generated
	InscriptionID string    // Reference to Inscription
	TypeEval      enum["CA", "PW", "Exam", "Project"]    // Evaluation type
	Valeur        float64   // Grade out of 20
	Coefficient   float64   // Weight of the grade
	DateEval      time.Time // Date of evaluation
	Commentaire   string    // Optional
}


/*
* CourseResult aggregates grades for a course and determines the average and status of the student in that course.
*/

type CourseResult struct {
	InscriptionID string
	CourseCode    string
	Grades        []Grade
	Average       float64
	Status        enum["Pass", "Fail", "Retake"]
}


/*
* Grade methods
*/


/*
 * Validate checks if the grade has valid values for its fields. It ensures that the evaluation type is one of the allowed types, the grade value is between 0 and 20,
 * and the coefficient is positive.
 */

func (g *Grade) Validate() error {
	if g.TypeEval != "CA" && g.TypeEval != "PW" && g.TypeEval != "Exam" && g.TypeEval != "Project" {
		return fmt.Errorf("invalid evaluation type")
	}
	if g.Valeur < 0 || g.Valeur > 20 {
		return fmt.Errorf("grade must be between 0 and 20")
	}
	if g.Coefficient <= 0 {
		return fmt.Errorf("coefficient must be positive")
	}
	return nil
}


/*
 * CalculateAverage computes the weighted average of the grades for a course result. It sums the products of grade values and their coefficients, and divides by the total coefficient.
 * If there are no grades or the total coefficient is zero, it returns an average of 0.
 */

func (cr *CourseResult) CalculateAverage(grades []Grade) float64 {
	var total float64
	var totalCoeff float64
	for _, grade := range grades {
		total += grade.Valeur * grade.Coefficient
		totalCoeff += grade.Coefficient
	}
	if totalCoeff == 0 {
		return 0
	}
	return total / totalCoeff
}

/*
 * IsValid checks if the course result is valid based on the average grade. A course result is considered valid if the average is 10 or above.
 */

func (cr *CourseResult) IsValid() bool {
	return cr.Average >= 10
}

/*
 * IsEliminating checks if the course result is eliminating based on the average grade. A course result is considered eliminating if the average is below 8.
 */

func (cr *CourseResult) IsEliminating() bool {
	return cr.Average < 8
}


/*
 * Mention returns a string representing the mention (honors) based on the average grade. It categorizes the average into different mentions such as "summa cum laude", "magna cum laude", "With honours", "Pass", or "Fail".
 */

func (cr *CourseResult) Mention() string {
	if cr.Average >= 16 {
		return "summa cum laude"
	}
	if cr.Average >= 14 {
		return "magna cum laude"
	}
	if cr.Average >= 12 {
		return "With honours"
	}
	if cr.Average >= 10 {
		return "Pass"
	}
	return "Fail"
}