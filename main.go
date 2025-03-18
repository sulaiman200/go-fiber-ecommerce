package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sulaiman200/go-fiber-ecommerce/database"
	"github.com/sulaiman200/go-fiber-ecommerce/models"
)

func main() {
	app := fiber.New()

	//connect to the database
	database.ConnectDatabase()

	//Run Migrations
	database.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Cart{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("E-commerce API is running>>>")
	})

	log.Fatal((app.Listen(":3000")))
}
