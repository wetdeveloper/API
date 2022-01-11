package crud_api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Num(c echo.Context) error {
	num := c.Param("num")
	return c.String(http.StatusOK, num)
}

func CrudForm(c echo.Context) error {
	return c.Render(http.StatusOK, "crudPage.html", map[string]interface{}{})
	}
