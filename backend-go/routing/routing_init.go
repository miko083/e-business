package routing

import (
	c "consoleshop/controllers"
	a "consoleshop/controllers/authentication"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// Set CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Console API
	e.GET("/consoles", c.GetConsoles)
	e.GET("/consoles/:id", c.GetConsole)
	e.POST("/consoles", c.AddConsole)
	e.DELETE("/consoles/:id", c.DeleteConsole)
	e.PUT("/consoles/:id", c.UpdateConsole)

	// Manufacturer API
	e.GET("/manufacturers", c.GetManufactures)
	e.GET("/manufacturers/:id", c.GetConsolesFromManufacturer)
	e.POST("/manufacturers", c.AddManufacturer)
	e.DELETE("/manufactures/:id", c.DeleteManufacturer)
	e.PUT("/manufactures/:id", c.UpdateManufacturer)

	// User API
	e.GET("/users", c.GetUsers)
	e.GET("/users/:id", c.GetUser)
	e.POST("/users", c.AddUser)
	e.DELETE("/users/:id", c.DeleteUser)
	e.PUT("/users/:id", c.UpdateUser)

	// Console With Quantity API
	e.GET("/consoleswithquantity", c.GetConsolesWithQuantity)
	e.GET("/consoleswithquantity/:id", c.GetConsoleWithQuantity)
	e.POST("/consoleswithquantity", c.AddConsoleWithQuantity)
	e.DELETE("/consoleswithquantity/:id", c.DeleteConsoleWithQuantity)
	e.PUT("/consoleswithquantity/:id", c.UpdateConsoleWithQuantity)

	// Shopping Cart API
	e.GET("/carts", c.GetCarts)
	e.GET("/carts/:id", c.GetCart)
	e.POST("/cartsUser", c.GetCartForUser)
	e.POST("/carts", c.AddCart)
	e.DELETE("/carts/:id", c.DeleteCart)
	e.PUT("/carts/:id", c.UpdateCart)

	// Payments
	e.POST("/payments", c.MakePayment)

	// Login
	e.GET("/auth/google/login", a.AuthGoogleLogin)
	e.GET("/auth/google/callback", a.AuthGoogleCallback)

	e.GET("/auth/github/login", a.AuthGithubLogin)
	e.GET("/auth/github/callback", a.AuthGithubCallback)

	// Logout
	e.POST("/logout", c.LogoutUser)

	return e
}
