package studenthandler

import "github.com/labstack/echo/v4"


func (h Handler) SetStudentRouts(app *echo.Echo) {
	studentsGrp := app.Group("/students")
	studentsGrp.GET("", h.GetAll)
	studentsGrp.POST("", h.NewStudent)
	studentsGrp.POST("/course/:id", h.AddCourse)
	studentsGrp.GET("/average", h.CalculatGradeAverage)
	studentsGrp.GET("/highest", h.HighestStudent)
	studentsGrp.GET("/lowest", h.LowestStudent)
	

}