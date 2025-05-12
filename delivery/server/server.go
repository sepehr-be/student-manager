package server

import (
	"log"
	"net/http"
	"student-app/config"
	studenthandler "student-app/delivery/server/studenthandler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
	studenthandler studenthandler.Handler
	Router *echo.Echo
}

func New(config config.Config,studenthandler studenthandler.Handler) Server {
	return Server{
		config: config, Router: echo.New(),
		studenthandler: studenthandler,
	}
}

func (s Server) Serve() {
	app := s.Router

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "everythink is good")
	})
	s.studenthandler.SetStudentRouts(s.Router)

	log.Fatal(s.Router.Start(":" + s.config.Port))
}
