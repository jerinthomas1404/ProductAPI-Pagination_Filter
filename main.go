package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jeirnthomas1404/ProductAPI-Pagination_Filter/configs"
	"github.com/jeirnthomas1404/ProductAPI-Pagination_Filter/controllers"
)

func main() {
	client := configs.ConnectDB()
	app := fiber.New()

	app.Use(cors.New()) //To avoid CORS error
	app.Post("/api/products/populate", controllers.PopulateData(client))
	app.Get("/api/products/frontend", controllers.GetData(client))
	app.Get("/api/products/backend", controllers.GetFilteredData(client))
	app.Listen(":8080")
}
