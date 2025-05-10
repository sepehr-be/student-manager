package repository

import "student-app/entity"

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