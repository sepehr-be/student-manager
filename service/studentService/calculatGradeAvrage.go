package studentservice




func (s Service) CalculatGradeAverage() float64 {
	average := s.repo.CalculatGradeAverage()
	return average
}