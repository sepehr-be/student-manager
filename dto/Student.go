package dto

type NewStudentRequest struct {
	Name   string   `json:"name"`
	Age    uint     `json:"age"`
	Grade  float64  `json:"grade"`
	Course []string `json:"course"`
}

type NewStudentResponse struct {
	ID int `json:"id"`
}

type NewCoursRequest struct {
	Course string `json:"course"`
}

