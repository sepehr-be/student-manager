package studenthandler

import (
	"net/http"
	"strconv"
	"student-app/dto"
	studentvalidator "student-app/validator/studentValidator"

	"github.com/labstack/echo/v4"
)

func (h Handler) AddCourse(c echo.Context) error {
	var course dto.NewCoursRequest
	bErr := c.Bind(&course)
	if bErr != nil{
		return c.JSON(http.StatusUnprocessableEntity,bErr.Error())
	}
	stringID := c.Param("id")

	id,err := strconv.Atoi(stringID)
	if err != nil{
		return c.JSON(http.StatusBadRequest,"invalid student id")
	}

	vErr := studentvalidator.CourseValidator(course.Course)
	if vErr != nil{
		return c.JSON(http.StatusUnprocessableEntity,vErr.Error())
	}

	aErr := h.studentSvc.AddCourse(course.Course,id)
	if aErr != nil {
		return c.JSON(http.StatusBadRequest,aErr.Error())
	}

	return c.JSON(http.StatusOK,"cours added")
}