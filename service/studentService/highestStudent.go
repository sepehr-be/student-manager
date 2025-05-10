package studentservice

import "student-app/entity"


func (s Service) HighestStudent() entity.Student {
	student := s.repo.HighestStudent()
	return student
}