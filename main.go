package main

import (
	"encoding/json"
	"fmt"
	ip_to_geolocation "github.com/emincanozcan/go-maxmind-ip-geolocation-rest/services/ip-to-geolocation"
	"github.com/emincanozcan/go-maxmind-ip-geolocation-rest/services/location-list"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"time"
)

func main() {
	// cache locations before starting server.
	location_list.GetLocationsAsJsonBytes()
	app := fiber.New()
	app.Use(compress.New())
	app.Get("/locations", LocationHandler)
	app.Get("/ip-to-geolocation/:ipAddr", ipToGeoLocationHandler)
	app.Listen(":8090")
}

func LocationHandler(c *fiber.Ctx) error {
	t := time.Now()
	b := location_list.GetLocationsAsJsonBytes()
	err := c.Send(b)
	if err != nil {
		return err
	}
	fmt.Println(time.Now().Sub(t))
	return nil
}

func ipToGeoLocationHandler(c *fiber.Ctx) error {
	ipAddr := c.Params("ipAddr")
	data := ip_to_geolocation.Geolocation(ipAddr)
	b, _ := json.Marshal(data)
	return c.Send(b)
}
