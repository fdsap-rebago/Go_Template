package routers

import (
	"Template/struct/errors"
	"Template/struct/publics"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func try() {
	fmt.Println(errors.ErrorModel{})
}

func SetupPublicRoutes(app *fiber.App) {

	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", publics.CheckServiceHealth)
}
