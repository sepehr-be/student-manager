package studenthandler

import "github.com/labstack/echo/v4"


func (h Handler) SetStudentRouts(app *echo.Echo) {

	app.GET("/get-all", h.GetAll)
	app.POST("/new-student", h.NewStudent)
	app.POST("/add-course/:id", h.AddCourse)
	app.GET("/calculat-average", h.CalculatGradeAverage)
	app.GET("/highest-student", h.HighestStudent)
	app.GET("/lowest-student", h.LowestStudent)
	

}