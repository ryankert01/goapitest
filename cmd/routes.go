package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryankert01/goapitest/handlers"
)

func setupRoutes(app *fiber.App) {
	// client: get ADs
	app.Get("/api/v1/ad", handlers.GetAds)

	// admin: create AD
	app.Post("/api/v1/ad", handlers.CreateAd)
}
