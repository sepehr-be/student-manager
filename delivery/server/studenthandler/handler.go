package studenthandler

import (
	"net/http"
	"strconv"
	"student-app/dto"
	studentservice "student-app/service/studentservice"
	"student-app/validator/studentvalidator"

	"github.com/labstack/echo/v4"
)

type Handler struct{
	studentSvc studentservice.Service
}

func New(studentSvc studentservice.Service) Handler {
	return Handler{
		studentSvc: studentSvc,
	}
}


func (h Handler) NewStudent(c echo.Context) error {
	var studen dto.NewStudentRequest
	bErr := c.Bind(&studen)
	if bErr != nil {
		return c.JSON(http.StatusBadRequest, bErr.Error())
	}

	vErr := studentvalidator.StudentValidator(studen)
	if vErr != nil{
		return c.JSON(http.StatusUnprocessableEntity,vErr.Error())
	}

	id, nErr := h.studentSvc.NewStudent(studen)
	if nErr != nil {
		return c.JSON(http.StatusBadRequest, nErr.Error())
	}

	return c.JSON(http.StatusOK,id)
}


func (h Handler) GetAll(c echo.Context) error {
	students := h.studentSvc.GetAll()

	return c.JSON(http.StatusOK,students)
}


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


func (h Handler) CalculatGradeAverage(c echo.Context) error {
	average := h.studentSvc.CalculatGradeAverage()
	return c.JSON(http.StatusOK,average)
}


func (h Handler) HighestStudent(c echo.Context) error {

	student := h.studentSvc.HighestStudent()

	return c.JSON(http.StatusOK,student)

}


func (h Handler) LowestStudent(c echo.Context) error {

	student := h.studentSvc.LowesStudent()

	return c.JSON(http.StatusOK,student)

}

