package router

import "github.com/gofiber/fiber/v2"

type Handler = []func(*fiber.Ctx) error

type RouteGroup struct {
	Path    string
	Routes  []Route
	Handler Handler
}

type Route struct {
	Path    string
	Method  string
	Handler Handler
}
