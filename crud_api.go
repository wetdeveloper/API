package crud_api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wetdeveloper/connection"
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
	mydb,_:= connection.Connect()
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
//Read
func Read(c echo.Context) error {
			mydb,_:= connection.Connect()
			return c.Render(http.StatusOK, "userslist.html", connection.ListUsers(mydb))
		}
