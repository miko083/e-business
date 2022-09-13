package controllers

import (
	"bytes"
	m "consoleshop/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"consoleshop/database"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

const consolePreloadString = "ConsolesWithQuantity.Console"
const consoleManufacturerPreloadString = "ConsolesWithQuantity.Console.Manufacturer"
const forbiddenMessage = "Not allowed."
const userMailPaymentDoneQuery = "user_email = ? AND payment_done = ?"
const keyForStripe = "sk_test_51LhMCXG0dnjPPiJiXx5HgrnffpgiHILoVWFZzodFTFir36GKgVfbyM7la0w07Z1RaWj9SqZJatBOgNQPcV33RIw700lQVbpxet"
const frontEndLink = "http://localhost:3000/?stripe_token="

func GetCarts(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAdmin(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		var shippingCartsWithQuantity []m.ShippingCart
		database.DBconnection.Preload("User").Preload(consolePreloadString).Preload(consoleManufacturerPreloadString).Find(&shippingCartsWithQuantity)
		return c.JSON(http.StatusOK, shippingCartsWithQuantity)
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func GetCart(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAdmin(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		id := c.Param("id")
		var shippingCart m.ShippingCart
		database.DBconnection.Preload("User").Preload("ConsolesWithQuantity").Preload(consolePreloadString).Preload(consoleManufacturerPreloadString).Find(&shippingCart, "ID = ?", id)
		return c.JSON(http.StatusOK, shippingCart)
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func GetCartForUser(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAuthenticated(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		body := make(map[string]interface{})
		json.NewDecoder(c.Request().Body).Decode(&body)
		email := body["user_email"].(string)
		var shippingCart m.ShippingCart
		database.DBconnection.Preload("ConsolesWithQuantity").Preload(consolePreloadString).Preload(consoleManufacturerPreloadString).Find(&shippingCart, userMailPaymentDoneQuery, email, false)
		return c.JSON(http.StatusOK, shippingCart)
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func AddCart(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAuthenticated(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		shippingCartOld := m.ShippingCart{}
		body := make(map[string]interface{})
		json.NewDecoder(c.Request().Body).Decode(&body)
		email := body["user_email"].(string)
		query := database.DBconnection.Find(&shippingCartOld, userMailPaymentDoneQuery, email, false)
		if query.RowsAffected > 0 {
			database.DBconnection.Delete(&shippingCartOld)
		}

		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		shippingCart := m.ShippingCart{}
		c.Bind(&shippingCart)
		database.DBconnection.Find(&shippingCartOld, userMailPaymentDoneQuery, email, false)
		database.DBconnection.Create(&shippingCart)
		return c.JSON(http.StatusOK, "Added new shipping cart.")
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func DeleteCart(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAuthenticated(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		id := c.Param("id")
		database.DBconnection.Delete(&m.ShippingCart{}, id)
		return c.JSON(http.StatusOK, "Deleted shipping cart with the id: "+id)
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func UpdateCart(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAuthenticated(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
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
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func PreparePayment(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAuthenticated(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		body := make(map[string]interface{})
		json.NewDecoder(c.Request().Body).Decode(&body)
		email := body["user_email"].(string)
		money_to_pay := int64(body["money_to_pay"].(float64))
		var shippingCart m.ShippingCart
		query := database.DBconnection.Preload("ConsolesWithQuantity").Preload(consolePreloadString).Preload(consoleManufacturerPreloadString).Find(&shippingCart, "user_email = ?", email)

		if query.RowsAffected > 0 {
			m := make(map[string]string)
			m["stripe_token"] = preparePaymentIntent(money_to_pay)
			return c.JSON(http.StatusOK, m)
		}
		return c.JSON(http.StatusBadRequest, "No shipping cart.")
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func MakePayment(c echo.Context) error {
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
	if checkIfAuthenticated(bodyBytes) {
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		body := make(map[string]interface{})
		json.NewDecoder(c.Request().Body).Decode(&body)
		email := body["user_email"].(string)
		var shippingCart m.ShippingCart
		query := database.DBconnection.Preload("ConsolesWithQuantity").Preload(consolePreloadString).Preload(consoleManufacturerPreloadString).Find(&shippingCart, "user_email = ?", email)

		if query.RowsAffected > 0 {
			shippingCart.PaymentDone = true
			database.DBconnection.Save(&shippingCart)
			return c.JSON(http.StatusOK, "Done.")
		}
		return c.JSON(http.StatusBadRequest, "No shipping cart.")
	}
	return c.JSON(http.StatusForbidden, forbiddenMessage)
}

func preparePaymentIntent(moneyToPay int64) string {
	stripe.Key = keyForStripe
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(moneyToPay),
		Currency: stripe.String(string(stripe.CurrencyPLN)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		log.Printf("pi.New: %v", err)
	}

	return pi.ClientSecret
}
