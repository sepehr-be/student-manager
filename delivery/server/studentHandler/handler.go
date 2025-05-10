package studenthandler

import studentservice "student-app/service/studentService"

type Handler struct{
	studentSvc studentservice.Service
}

func New(studentSvc studentservice.Service) Handler {
	return Handler{
		studentSvc: studentSvc,
	}
}