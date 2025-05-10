package studentservice

import "student-app/entity"


func (s Service) GetAll() []entity.Student {
	studenList := s.repo.GetAll()

	return studenList
}