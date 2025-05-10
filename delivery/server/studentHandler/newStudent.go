package studenthandler

import (
	"net/http"
	"student-app/dto"
	studentvalidator "student-app/validator/studentValidator"

	"github.com/labstack/echo/v4"
)

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
