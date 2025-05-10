package studentservice

import (
	"fmt"
	"student-app/dto"
)

func (s *Service) NewStudent(req dto.NewStudentRequest) (dto.NewStudentResponse, error) {
	ID, err := s.repo.NewStudent(req)
	if err != nil {
		return dto.NewStudentResponse{}, fmt.Errorf("error :%v", err)
	}
	return dto.NewStudentResponse{ID: ID.ID}, nil
}
