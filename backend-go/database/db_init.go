package database

import (
	m "consoleshop/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Setup() {
	DBconnection.AutoMigrate(&m.User{}, &m.Manufacturer{}, &m.Console{}, &m.ConsoleWithQuantity{}, &m.ShippingCart{})
	seed(DBconnection)
}

func seed(db *gorm.DB) {
	users := []m.User{
		{FirstName: "John", LastName: "Smith", Email: "test1@gmail.com"},
		{FirstName: "George", LastName: "Washington", Email: "test2@gmail.com"},
	}
	for _, u := range users {
		db.Create(&u)
	}

	var johnSmith m.User
	db.First(&johnSmith, "first_name = ?", "John")

	manufacturers := []m.Manufacturer{
		{Name: "Microsoft", OriginCountry: "USA"},
		{Name: "Sony", OriginCountry: "Japan"},
	}
	for _, m := range manufacturers {
		db.Create(&m)
	}

	var microsoft, sony m.Manufacturer
	db.First(&microsoft, "name = ?", "Microsoft")
	db.First(&sony, "name = ?", "Sony")

	consoles := []m.Console{
		{Name: "Xbox Series X", Price: 2499, Manufacturer: microsoft},
		{Name: "Playstation 5", Price: 2599, Manufacturer: sony},
	}

	for _, c := range consoles {
		db.Create(&c)
	}

	var xboxSeriesX m.Console
	db.First(&xboxSeriesX, "name = ?", "Xbox Series X")

	var playstation5 m.Console
	db.First(&playstation5, "name = ?", "Playstation 5")

	consolesWithQuantity := []m.ConsoleWithQuantity{
		{Console: xboxSeriesX, Quantity: 5},
		{Console: playstation5, Quantity: 2},
	}

	shippingCart := m.ShippingCart{
		User: johnSmith, ConsolesWithQuantity: consolesWithQuantity,
	}

	db.Create(&shippingCart)

}
