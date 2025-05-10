package studenthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) CalculatGradeAverage(c echo.Context) error {
	average := h.studentSvc.CalculatGradeAverage()
	return c.JSON(http.StatusOK,average)
}