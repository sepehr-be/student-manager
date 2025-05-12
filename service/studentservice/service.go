package studentservice

import (
	"fmt"
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

func (s *Service) NewStudent(req dto.NewStudentRequest) (dto.NewStudentResponse, error) {
	ID, err := s.repo.NewStudent(req)
	if err != nil {
		return dto.NewStudentResponse{}, fmt.Errorf("error :%v", err)
	}
	return dto.NewStudentResponse{ID: ID.ID}, nil
}

func (s Service) GetAll() []entity.Student {
	studenList := s.repo.GetAll()

	return studenList
}

func (s *Service) AddCourse(c string,stdID int) error {
	err := s.repo.AddCourse(c,stdID)
	if err != nil {
		return fmt.Errorf("error: %v",err)
	}

	return nil
}

func (s Service) CalculatGradeAverage() float64 {
	average := s.repo.CalculatGradeAverage()
	return average
}

func (s Service) HighestStudent() entity.Student {
	student := s.repo.HighestStudent()
	return student
}

func (s Service) LowesStudent() entity.Student {
	student := s.repo.LowesStudent()
	return student
}