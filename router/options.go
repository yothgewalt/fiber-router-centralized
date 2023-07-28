package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type RouterOpts struct {
	FiberApp    *fiber.App
	Middlewares []interface{}
	Definition  []RouteGroup
}

type ListenOpts struct {
	Context context.Context
	Port    uint32
}
