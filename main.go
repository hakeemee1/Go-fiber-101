package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/routes"
)

func main() {
	app := fiber.New()

	routes.InetRoutes(app) // for using in router_inet.go

	app.Listen(":3000")
}
