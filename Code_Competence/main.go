package main

import (
	"Code_Competence/config"
	"Code_Competence/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
