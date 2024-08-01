// Package api provides HTTP handlers for the geolocation service.
//
// This package contains the definition of HTTP endpoints and their handlers,
// which process incoming HTTP requests and return appropriate responses.
package api

import (
	"github.com/gofiber/fiber/v2"
)

// Define a struct that implements the Endpoint interface
type GetEnrichGeolocation struct {
	MethodValue  string
	PathValue    string
	HandlerValue fiber.Handler
}

// Implement the Endpoint interface methods
func (e GetEnrichGeolocation) Method() string {
	return e.MethodValue
}

func (e GetEnrichGeolocation) Path() string {
	return e.PathValue
}

func (e GetEnrichGeolocation) Handler() fiber.Handler {
	return e.HandlerValue
}

// NewGetEnrichGeolocation initializes the GET /enrich/geolocation endpoint.
func NewGetEnrichGeolocation() Endpoint {
	return GetEnrichGeolocation{
		MethodValue:  MethodGET,
		PathValue:    "/enrich/geolocation",
		HandlerValue: enrichGeolocationHandler,
	}
}
