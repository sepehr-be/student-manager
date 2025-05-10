package studenthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) LowestStudent(c echo.Context) error {

	student := h.studentSvc.LowesStudent()

	return c.JSON(http.StatusOK,student)

}