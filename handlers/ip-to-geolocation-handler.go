package handlers

import (
	"encoding/json"
	ip_to_geolocation "github.com/emincanozcan/go-maxmind-ip-geolocation-rest/services/ip-to-geolocation"
	"github.com/gofiber/fiber/v2"
)

func IpToGeolocationHandler(c *fiber.Ctx) error {
	ipAddr := c.Params("ipAddr")
	data := ip_to_geolocation.Geolocation(ipAddr)
	b, _ := json.Marshal(data)
	return c.Send(b)
}
