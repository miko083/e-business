package controllers

import (
	"bytes"
	m "consoleshop/models"
	"io/ioutil"
	"log"
	"net/http"

	"consoleshop/database"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetConsoles(c echo.Context) error {
	var consoles []m.Console
	database.DBconnection.Preload("Manufacturer").Find(&consoles)
	return c.JSON(http.StatusOK, consoles)
}

func GetConsole(c echo.Context) error {
	id := c.Param("id")
	var console m.Console
	database.DBconnection.Preload("Manufacturer").Find(&console, "ID = ?", id)
	return c.JSON(http.StatusOK, console)
}

func AddConsole(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAdmin(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		console := m.Console{}
		err := c.Bind(&console)
		if err != nil {
			log.Printf("Failed: %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		database.DBconnection.Create(&console)
		return c.JSON(http.StatusOK, "Added new console.")
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func DeleteConsole(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAdmin(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		id := c.Param("id")
		database.DBconnection.Delete(&m.Console{}, id)
		return c.JSON(http.StatusOK, "Deleted console with the id: "+id)
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func UpdateConsole(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAdmin(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		id := c.Param("id")
		var consoleToUpdate m.Console
		database.DBconnection.Find(&consoleToUpdate, "ID = ?", id)

		consoleFromBody := m.Console{}
		err := c.Bind(&consoleFromBody)
		if err != nil {
			log.Printf("Failed: %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		if consoleFromBody.Name != "" {
			consoleToUpdate.Name = consoleFromBody.Name
		}

		if consoleFromBody.Price != 0 {
			consoleToUpdate.Price = consoleFromBody.Price
		}

		if consoleFromBody.ManufacturerID != 0 {
			consoleToUpdate.ManufacturerID = consoleFromBody.ManufacturerID
		}

		database.DBconnection.Save(&consoleToUpdate)
		return c.JSON(http.StatusOK, "Updated console with the id: "+id)
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}
