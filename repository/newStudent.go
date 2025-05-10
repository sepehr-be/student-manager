package repository

import (
	"student-app/dto"
	"student-app/entity"
)

func (r *StudenRepository) NewStudent(req dto.NewStudentRequest) (dto.NewStudentResponse, error) {
	id := r.getNextIDFromMap()
	student := entity.Student{
		ID:     id,
		Name:   req.Name,
		Age:    req.Age,
		Grade:  req.Grade,
		Course: req.Course,
	}
	r.Data[id] = student
	return dto.NewStudentResponse{ID: id},nil
}
