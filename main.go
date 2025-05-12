package main

import (
	"student-app/config"
	"student-app/delivery/server"
	"student-app/delivery/server/studenthandler"
	"student-app/repository"
	studentservice "student-app/service/studentservice"
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