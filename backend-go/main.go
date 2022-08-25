package main

import (
	"consoleshop/database"
	"consoleshop/routing"
)

func main() {
	database.Connect()
	defer database.DBconnection.Close()
	database.Setup()

	e := routing.Init()
	e.Logger.Fatal(e.Start(":8000"))

}
