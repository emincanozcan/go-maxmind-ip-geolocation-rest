package handlers

import (
	location_list "github.com/emincanozcan/go-maxmind-ip-geolocation-rest/services/location-list"
	"github.com/gofiber/fiber/v2"
)

func LocationListHandler(c *fiber.Ctx) error {
	b := location_list.GetLocationsAsJsonBytes()
	return c.Send(b)
}
