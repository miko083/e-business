package controllers

import (
	"bytes"
	m "consoleshop/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"consoleshop/database"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetCarts(c echo.Context) error {
	var shippingCartsWithQuantity []m.ShippingCart
	database.DBconnection.Preload("User").Preload("ConsolesWithQuantity.Console").Preload("ConsolesWithQuantity.Console.Manufacturer").Find(&shippingCartsWithQuantity)
	return c.JSON(http.StatusOK, shippingCartsWithQuantity)
}

func GetCart(c echo.Context) error {
	id := c.Param("id")
	var shippingCart m.ShippingCart
	database.DBconnection.Preload("User").Preload("ConsolesWithQuantity").Preload("ConsolesWithQuantity.Console").Preload("ConsolesWithQuantity.Console.Manufacturer").Find(&shippingCart, "ID = ?", id)
	return c.JSON(http.StatusOK, shippingCart)
}

func GetCartForUser(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAuthenticated(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		body := make(map[string]interface{})
		json.NewDecoder(c.Request().Body).Decode(&body)
		email := body["user_email"].(string)
		var shippingCart m.ShippingCart
		database.DBconnection.Preload("ConsolesWithQuantity").Preload("ConsolesWithQuantity.Console").Preload("ConsolesWithQuantity.Console.Manufacturer").Find(&shippingCart, "user_email = ? AND payment_done = ?", email, false)
		return c.JSON(http.StatusOK, shippingCart)
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}

func AddCart(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAuthenticated(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		shippingCart := m.ShippingCart{}
		c.Bind(&shippingCart)
		database.DBconnection.Create(&shippingCart)
		return c.JSON(http.StatusOK, "Added new shipping cart.")
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}

func DeleteCart(c echo.Context) error {
	id := c.Param("id")
	database.DBconnection.Delete(&m.ShippingCart{}, id)
	return c.JSON(http.StatusOK, "Deleted shipping cart with the id: "+id)
}

func UpdateCart(c echo.Context) error {
	id := c.Param("id")
	var shippingCartToUpdate m.ShippingCart
	database.DBconnection.Find(&shippingCartToUpdate, "ID = ?", id)

	shippingCartFromBody := m.ShippingCart{}
	err := c.Bind(&shippingCartFromBody)
	if err != nil {
		log.Printf("Failed: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if shippingCartFromBody.ConsolesWithQuantity != nil {
		shippingCartToUpdate.ConsolesWithQuantity = shippingCartFromBody.ConsolesWithQuantity
	}

	database.DBconnection.Save(&shippingCartToUpdate)
	return c.JSON(http.StatusOK, "Updated shipping cart with the id: "+id)
}

func MakePayment(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAuthenticated(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		body := make(map[string]interface{})
		json.NewDecoder(c.Request().Body).Decode(&body)
		email := body["user_email"].(string)
		fmt.Println(email)
		var shippingCart m.ShippingCart
		query := database.DBconnection.Preload("ConsolesWithQuantity").Preload("ConsolesWithQuantity.Console").Preload("ConsolesWithQuantity.Console.Manufacturer").Find(&shippingCart, "user_email = ?", email)
		if query.RowsAffected > 0 {
			shippingCart.PaymentDone = true
			database.DBconnection.Save(&shippingCart)
			return c.JSON(http.StatusOK, "Done.")
		}
		return c.JSON(http.StatusBadRequest, "No shipping cart.")
	}
	return c.JSON(http.StatusForbidden, "Not allowed.")
}
