package studenthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetAll(c echo.Context) error {
	students := h.studentSvc.GetAll()

	return c.JSON(http.StatusOK,students)
}