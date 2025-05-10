package studenthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func (h Handler) HighestStudent(c echo.Context) error {

	student := h.studentSvc.HighestStudent()

	return c.JSON(http.StatusOK,student)

}