package repository

import (
	"fmt"
	"strings"
	"student-app/dto"
	"student-app/entity"
)

type StudenRepository struct {
	Data map[int]entity.Student
}

func New() *StudenRepository {
	return &StudenRepository{
		Data: make(map[int]entity.Student),
	}
}

func (r *StudenRepository) getNextIDFromMap() int {
	maxID := 0
	for id := range r.Data {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}


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

func (r *StudenRepository) GetAll() []entity.Student {
	studenList := make([]entity.Student,0,len(r.Data))

	for _,student := range r.Data{
		studenList = append(studenList, student)
	}

	return studenList
}

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

func (r *StudenRepository) CalculatGradeAverage() float64 {
	var sum float64
	var count float64
	for _,student := range r.Data{
		sum += student.Grade
		count ++
	}
	average := sum / count
	
	return average
}

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