package studentservice

import "student-app/entity"


func (s Service) LowesStudent() entity.Student {
	student := s.repo.LowesStudent()
	return student
}