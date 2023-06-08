package main

import (
	"Go_Aulia-Choirin_Aulia-Choirin-Aulia-Choirin/22_UnitTesting/Praktikum/config"

	"Go_Aulia-Choirin_Aulia-Choirin-Aulia-Choirin/22_UnitTesting/Praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()
	// create a new echo instance
	e := routes.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}