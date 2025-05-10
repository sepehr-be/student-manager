package studentservice

import (
	"student-app/dto"
	"student-app/entity"
)


type Repository interface {
	NewStudent(dto.NewStudentRequest) (dto.NewStudentResponse,error)
	AddCourse(c string,stdID int) error
	CalculatGradeAverage() float64
	HighestStudent() entity.Student
	LowesStudent() entity.Student
	GetAll() []entity.Student
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}