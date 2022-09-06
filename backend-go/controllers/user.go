package controllers

import (
	m "consoleshop/models"
	"log"
	"net/http"

	"consoleshop/database"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"encoding/json"
)

func GetUsers(c echo.Context) error {
	body := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&body)
	if checkIfAdmin(body) {
		var users []m.User
		database.DBconnection.Find(&users)
		return c.JSON(http.StatusOK, users)
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}

func GetUser(c echo.Context) error {
	body := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&body)
	if checkIfAdmin(body) {
		id := c.Param("id")
		var user m.User
		database.DBconnection.Find(&user, "ID = ?", id)
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}

func AddUser(c echo.Context) error {
	body := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&body)
	if checkIfAdmin(body) {
		user := m.User{}
		err := c.Bind(&user)
		if err != nil {
			log.Printf("Failed: %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		database.DBconnection.Create(&user)
		return c.JSON(http.StatusOK, "Added new user.")
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}

func DeleteUser(c echo.Context) error {
	body := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&body)
	if checkIfAdmin(body) {
		id := c.Param("id")
		database.DBconnection.Delete(&m.User{}, id)
		return c.JSON(http.StatusOK, "Deleted user with the id: "+id)
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}

func UpdateUser(c echo.Context) error {
	body := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&body)
	if checkIfAdmin(body) {
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

		if userFromBody.Email != "" {
			userToUpdate.Email = userFromBody.Email
		}

		if userFromBody.AccessToken != "" {
			userToUpdate.AccessToken = userFromBody.AccessToken
		}

		if userFromBody.LoginToken != "" {
			userToUpdate.LoginToken = userFromBody.LoginToken
		}

		database.DBconnection.Save(&userToUpdate)
		return c.JSON(http.StatusOK, "Updated user with the id: "+id)
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}

func LogoutUser(c echo.Context) error {
	body := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&body)
	email := body["email"].(string)
	if checkIfAuthenticated(body) {
		var userToLogout m.User
		database.DBconnection.Find(&userToLogout, "email = ?", email)
		userToLogout.AccessToken = ""
		userToLogout.LoginToken = ""
		userToLogout.IsLoggedIn = false
		database.DBconnection.Save(&userToLogout)
		return c.JSON(http.StatusOK, "User with the email logout")
	} else {
		return c.JSON(http.StatusForbidden, "Not allowed.")
	}
}
