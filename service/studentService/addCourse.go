package studentservice

import "fmt"


func (s Service) AddCourse(c string,stdID int) error {
	err := s.repo.AddCourse(c,stdID)
	if err != nil {
		return fmt.Errorf("error: %v",err)
	}

	return nil
}