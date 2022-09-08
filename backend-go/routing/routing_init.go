package routing

import (
	c "consoleshop/controllers"
	a "consoleshop/controllers/authentication"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	consolesWithIDRoute := "/consoles/:id"
	usersWithIDRoute := "/users/:id"
	consolesWithQuantityIDRoute := "/consoleswithquantity/:id"
	cartWithIDRoute := "/carts/:id"
	e := echo.New()

	// Set CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Console API
	e.GET("/consoles", c.GetConsoles)
	e.GET(consolesWithIDRoute, c.GetConsole)
	e.POST("/consoles", c.AddConsole)
	e.DELETE(consolesWithIDRoute, c.DeleteConsole)
	e.PUT(consolesWithIDRoute, c.UpdateConsole)

	// Manufacturer API
	e.GET("/manufacturers", c.GetManufactures)
	e.GET("/manufacturers/:id", c.GetConsolesFromManufacturer)
	e.POST("/manufacturers", c.AddManufacturer)
	e.DELETE("/manufactures/:id", c.DeleteManufacturer)
	e.PUT("/manufactures/:id", c.UpdateManufacturer)

	// User API
	e.GET("/users", c.GetUsers)
	e.GET(usersWithIDRoute, c.GetUser)
	e.POST("/users", c.AddUser)
	e.DELETE(usersWithIDRoute, c.DeleteUser)
	e.PUT(usersWithIDRoute, c.UpdateUser)

	// Console With Quantity API
	e.GET("/consoleswithquantity", c.GetConsolesWithQuantity)
	e.GET(consolesWithQuantityIDRoute, c.GetConsoleWithQuantity)
	e.POST("/consoleswithquantity", c.AddConsoleWithQuantity)
	e.DELETE(consolesWithQuantityIDRoute, c.DeleteConsoleWithQuantity)
	e.PUT(consolesWithQuantityIDRoute, c.UpdateConsoleWithQuantity)

	// Shopping Cart API
	e.GET("/carts", c.GetCarts)
	e.GET(cartWithIDRoute, c.GetCart)
	e.POST("/cartsUser", c.GetCartForUser)
	e.POST("/carts", c.AddCart)
	e.DELETE(cartWithIDRoute, c.DeleteCart)
	e.PUT(cartWithIDRoute, c.UpdateCart)

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
