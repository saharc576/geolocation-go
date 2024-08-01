package api

import "github.com/gofiber/fiber/v2"

// Define constants for HTTP methods
const (
	MethodGET  = "GET"
	MethodPOST = "POST"
)

// Endpoint defines an HTTP endpoint with method, path, and handler.
//
// Method specifies the HTTP method (e.g., GET, POST).
// Path specifies the URL path for the endpoint.
// Handler is the function that handles requests to the endpoint.
type Endpoint interface {
	Method() string
	Path() string
	Handler() fiber.Handler
}
