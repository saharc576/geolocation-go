package api

import (
	"geolocation/pkg/libs/geolocation"
	"log"

	"github.com/gofiber/fiber/v2"
)

// The main handler for the enrich geolocation function
func enrichGeolocationHandler(c *fiber.Ctx) error {
	ip := c.Query("ip")
	if ip == "" {
		return c.Status(fiber.StatusBadRequest).SendString("IP address is required")
	}
	log.Println("Checking for IP:", ip)
	result, err := geolocation.GetGeoLocationIpBased(ip)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(result)
}
