package controllers

import (
	m "consoleshop/models"
	"log"
	"net/http"

	"consoleshop/database"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetConsolesWithQuantity(c echo.Context) error {
	var consoleWithQuantitysWithQuantity []m.ConsoleWithQuantity
	database.DBconnection.Preload("Console").Preload("Console.Manufacturer").Find(&consoleWithQuantitysWithQuantity)
	return c.JSON(http.StatusOK, consoleWithQuantitysWithQuantity)
}

func GetConsoleWithQuantity(c echo.Context) error {
	id := c.Param("id")
	var consoleWithQuantity m.ConsoleWithQuantity
	database.DBconnection.Preload("Console").Preload("Console.Manufacturer").Find(&consoleWithQuantity, "ID = ?", id)
	return c.JSON(http.StatusOK, consoleWithQuantity)
}

func AddConsoleWithQuantity(c echo.Context) error {
	consoleWithQuantity := m.ConsoleWithQuantity{}
	err := c.Bind(&consoleWithQuantity)
	if err != nil {
		log.Printf("Failed: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	database.DBconnection.Create(&consoleWithQuantity)
	return c.JSON(http.StatusOK, "Added new consoleWithQuantity.")
}

func DeleteConsoleWithQuantity(c echo.Context) error {
	id := c.Param("id")
	database.DBconnection.Delete(&m.ConsoleWithQuantity{}, id)
	return c.JSON(http.StatusOK, "Deleted consoleWithQuantity with the id: "+id)
}

func UpdateConsoleWithQuantity(c echo.Context) error {
	id := c.Param("id")
	var consoleWithQuantityToUpdate m.ConsoleWithQuantity
	database.DBconnection.Find(&consoleWithQuantityToUpdate, "ID = ?", id)

	consoleWithQuantityFromBody := m.ConsoleWithQuantity{}
	err := c.Bind(&consoleWithQuantityFromBody)
	if err != nil {
		log.Printf("Failed: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if consoleWithQuantityFromBody.ConsoleID != 0 {
		consoleWithQuantityToUpdate.ConsoleID = consoleWithQuantityFromBody.ConsoleID
	}

	if consoleWithQuantityFromBody.Quantity != 0 {
		consoleWithQuantityToUpdate.Quantity = consoleWithQuantityFromBody.Quantity
	}

	if consoleWithQuantityFromBody.ShippingCartID != 0 {
		consoleWithQuantityToUpdate.ShippingCartID = consoleWithQuantityFromBody.ShippingCartID
	}

	database.DBconnection.Save(&consoleWithQuantityToUpdate)
	return c.JSON(http.StatusOK, "Updated consoleWithQuantity with the id: "+id)
}
