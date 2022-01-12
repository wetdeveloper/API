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

//Create,update,Delete
func Cud(c echo.Context) error {
			operation := c.FormValue("operation")
			username := c.FormValue("username")
			password := c.FormValue("password")
			if operation == "C" {
				if connection.InsertUser(mydb, username, password) {
					return c.String(http.StatusOK, "Created")
				}
				return c.String(http.StatusOK, "There is an error")
			} else if operation == "U" {
				if connection.UpdateUser(mydb, username, password) {
					return c.String(http.StatusOK, "Updated")
				}
				return c.String(http.StatusOK, "There is an error")
			} else if operation == "D" {
				if connection.DeleteUser(mydb, username) {
					return c.String(http.StatusOK, "Deleted")
				}
				return c.String(http.StatusOK, "There is an error")
			}
			return c.String(http.StatusOK, "Not valid operation")

		}
