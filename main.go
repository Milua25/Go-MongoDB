package main

import (
	"github.com/Golang-Personal-Projects/GolangTutorial/GoMongoDB/database"
	"github.com/Golang-Personal-Projects/GolangTutorial/GoMongoDB/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func setupRoutes(app *fiber.App) {
	// endpoints
	app.Get("/employee/", routes.GetEmployees)
	app.Get("/employee/:id", routes.GetEmployee)
	app.Put("/employee/:id", routes.UpdateEmployee)
	app.Post("/employee/", routes.CreateEmployee)
	app.Delete("/employee/:id", routes.DeleteEmployee)
}

func main() {
	if err := database.Connection(); err != nil {
		log.Fatalln(err)
	}
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
