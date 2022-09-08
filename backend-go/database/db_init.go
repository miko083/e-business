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
	nameQuery := "name = ?"
	users := []m.User{
		{FirstName: "John", LastName: "Smith", Email: "test1@gmail.com"},
		{FirstName: "George", LastName: "Washington", Email: "test2@gmail.com"},
	}
	for _, u := range users {
		db.Create(&u)
	}

	var johnSmith m.User
	db.First(&johnSmith, "email = ?", "test1@gmail.com")

	manufacturers := []m.Manufacturer{
		{Name: "Microsoft", OriginCountry: "USA"},
		{Name: "Sony", OriginCountry: "Japan"},
		{Name: "Nintendo", OriginCountry: "Japan"},
	}
	for _, m := range manufacturers {
		db.Create(&m)
	}

	var microsoft, sony, nintendo m.Manufacturer
	db.First(&microsoft, nameQuery, "Microsoft")
	db.First(&sony, nameQuery, "Sony")
	db.First(&nintendo, nameQuery, "Nintendo")

	consoles := []m.Console{
		{Name: "Xbox Series X", Price: 2499, Manufacturer: microsoft},
		{Name: "Xbox Series S", Price: 1299, Manufacturer: microsoft},
		{Name: "Xbox One S", Price: 1399, Manufacturer: microsoft},
		{Name: "Xbox One X", Price: 2299, Manufacturer: microsoft},

		{Name: "Playstation 5", Price: 2599, Manufacturer: sony},
		{Name: "Playstation 4 Pro", Price: 2199, Manufacturer: sony},
		{Name: "Playstation 4", Price: 1699, Manufacturer: sony},

		{Name: "Nintendo Switch", Price: 1499, Manufacturer: nintendo},
	}

	for _, c := range consoles {
		db.Create(&c)
	}

	var xboxSeriesX m.Console
	db.First(&xboxSeriesX, nameQuery, "Xbox Series X")

	var playstation5 m.Console
	db.First(&playstation5, nameQuery, "Playstation 5")

	consolesWithQuantity := []m.ConsoleWithQuantity{
		{Console: xboxSeriesX, Quantity: 5},
		{Console: playstation5, Quantity: 2},
	}

	shippingCart := m.ShippingCart{
		User: johnSmith, ConsolesWithQuantity: consolesWithQuantity,
	}

	db.Create(&shippingCart)

}
