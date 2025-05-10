package repository

import (
	"fmt"
	"strings"
)

func (r *StudenRepository) AddCourse(c string, stdID int) error {
	student, ok := r.Data[stdID]
	if !ok {
		return fmt.Errorf("user is not exist")
	}
	for _, course := range student.Course {
		if strings.EqualFold(course, c) {
			return fmt.Errorf("the entered course already exists.")
		}
	}
	student.Course = append(student.Course, c)
	r.Data[stdID] = student

	return nil
}
