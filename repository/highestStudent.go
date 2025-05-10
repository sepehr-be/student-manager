package repository

import (
	"student-app/entity"
)

func (r *StudenRepository) HighestStudent() entity.Student {

	if len(r.Data) == 0 {
		return entity.Student{}
	}
	var highestStudent entity.Student
	for _,student := range r.Data{
		if student.Grade >highestStudent.Grade {
			highestStudent = student
		}
	}
	return highestStudent
}