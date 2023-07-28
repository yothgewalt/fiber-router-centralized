package main

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"yoth.dev/router"
)

func main() {
	router.NewRouter(
		router.RouterOpts{
			FiberApp: fiber.New(
				fiber.Config{
					DisableStartupMessage: true,
					JSONEncoder:           sonic.Marshal,
					JSONDecoder:           sonic.Unmarshal,
				},
			),
			Middlewares: []interface{}{cors.New(), logger.New()},
			Definition: []router.RouteGroup{
				{
					Path: "/healthy",
					Routes: []router.Route{
						{
							Path:   "/now",
							Method: fiber.MethodGet,
							Handler: []func(*fiber.Ctx) error{
								func(c *fiber.Ctx) error {
									return c.Status(fiber.StatusOK).JSON(fiber.Map{
										"message": "it's okay.",
									})
								},
							},
						},
					},
					Handler: []func(*fiber.Ctx) error{
						func(c *fiber.Ctx) error { return c.Next() },
					},
				},
			},
		},
	).ListenThoughtPort(
		router.ListenOpts{
			Context: context.Background(),
			Port:    8080,
		},
	)
}
