package main

import (
	"github.com/emincanozcan/go-maxmind-ip-geolocation-rest/handlers"
	"github.com/emincanozcan/go-maxmind-ip-geolocation-rest/services/location-list"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"log"
)

func main() {
	// cache locations before starting server.
	location_list.GetLocationsAsJsonBytes()

	app := fiber.New()
	app.Use(compress.New())
	app.Get("/locations", handlers.LocationListHandler)
	app.Get("/ip-to-geolocation/:ipAddr", handlers.IpToGeolocationHandler)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
