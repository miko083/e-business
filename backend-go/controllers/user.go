package controllers

import (
	m "consoleshop/models"
	"log"
	"net/http"

	"consoleshop/database"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetUsers(c echo.Context) error {
	var users []m.User
	database.DBconnection.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	var user m.User
	database.DBconnection.Find(&user, "ID = ?", id)
	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) error {
	user := m.User{}
	err := c.Bind(&user)
	if err != nil {
		log.Printf("Failed: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	database.DBconnection.Create(&user)
	return c.JSON(http.StatusOK, "Added new user.")
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	database.DBconnection.Delete(&m.User{}, id)
	return c.JSON(http.StatusOK, "Deleted user with the id: "+id)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var userToUpdate m.User
	database.DBconnection.Find(&userToUpdate, "ID = ?", id)

	userFromBody := m.User{}
	err := c.Bind(&userFromBody)
	if err != nil {
		log.Printf("Failed: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if userFromBody.FirstName != "" {
		userToUpdate.FirstName = userFromBody.FirstName
	}

	if userFromBody.LastName != "" {
		userToUpdate.LastName = userFromBody.LastName
	}

	database.DBconnection.Save(&userToUpdate)
	return c.JSON(http.StatusOK, "Updated user with the id: "+id)
}
