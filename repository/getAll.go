package repository

import "student-app/entity"


func (r *StudenRepository) GetAll() []entity.Student {
	studenList := make([]entity.Student,0,len(r.Data))

	for _,student := range r.Data{
		studenList = append(studenList, student)
	}

	return studenList
}