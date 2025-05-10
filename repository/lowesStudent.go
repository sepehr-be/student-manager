package repository

import "student-app/entity"


func (r *StudenRepository) LowesStudent() entity.Student {
	if len(r.Data) == 0 {
		return entity.Student{}
	}
	var lowestStudent entity.Student
	for _,student := range r.Data{
		if student.Grade >lowestStudent.Grade {
			lowestStudent = student
		}
	}
	 return lowestStudent
}