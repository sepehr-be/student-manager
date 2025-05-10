package repository


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