package crud-api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func num(c echo.Context) error {
	num := c.Param("num")
	return c.String(http.StatusOK, num)
}
