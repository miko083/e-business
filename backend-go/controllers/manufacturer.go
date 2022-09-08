package controllers

import (
	"bytes"
	m "consoleshop/models"
	"log"
	"net/http"

	"consoleshop/database"

	"github.com/labstack/echo/v4"

	"io/ioutil"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetManufactures(c echo.Context) error {
	var manufacturers []m.Manufacturer
	database.DBconnection.Find(&manufacturers)
	return c.JSON(http.StatusOK, manufacturers)
}

func GetConsolesFromManufacturer(c echo.Context) error {
	id := c.Param("id")
	var consoles []m.Console
	database.DBconnection.Preload("Manufacturer").Find(&consoles, "manufacturer_id = ?", id)
	return c.JSON(http.StatusOK, consoles)
}

func AddManufacturer(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAdmin(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		manufacturer := m.Manufacturer{}
		err := c.Bind(&manufacturer)
		if err != nil {
			log.Printf("Failed: %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		database.DBconnection.Create(&manufacturer)
		return c.JSON(http.StatusOK, "Added new manufacturer.")
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}

func DeleteManufacturer(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAdmin(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		id := c.Param("id")
		database.DBconnection.Delete(&m.Manufacturer{}, id)
		return c.JSON(http.StatusOK, "Deleted manufacturer with the id: "+id)
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}

func UpdateManufacturer(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAdmin(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		id := c.Param("id")
		var manufacturerToUpdate m.Manufacturer
		database.DBconnection.Find(&manufacturerToUpdate, "ID = ?", id)

		manufacturerFromBody := m.Manufacturer{}
		err := c.Bind(&manufacturerFromBody)
		if err != nil {
			log.Printf("Failed: %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		if manufacturerFromBody.Name != "" {
			manufacturerToUpdate.Name = manufacturerFromBody.Name
		}

		if manufacturerFromBody.OriginCountry != "" {
			manufacturerToUpdate.OriginCountry = manufacturerFromBody.OriginCountry
		}

		database.DBconnection.Save(&manufacturerToUpdate)
		return c.JSON(http.StatusOK, "Updated manufacturer with the id: "+id)
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}
