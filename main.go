package main

import (
	"student-app/config"
	"student-app/delivery/server"
	studenthandler "student-app/delivery/server/studentHandler"
	"student-app/repository"
	studentservice "student-app/service/studentService"
)


func main() {
	config := new(config.Config)
	config.Port = "3000"

	repo := repository.New()
	studentSvc := studentservice.New(repo)
	studentHandler := studenthandler.New(studentSvc)


	server := server.New(*config,studentHandler)
	server.Serve()


}