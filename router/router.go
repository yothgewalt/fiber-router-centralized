package router

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

var rootGroupPath string = "/v1/generics"

type Router interface {
	ListenThoughtPort(opts ListenOpts)
}

type routerDependency struct {
	fiberApp   *fiber.App
	definition []RouteGroup
}

func NewRouter(opts RouterOpts) Router {
	opts.FiberApp.Use(opts.Middlewares...)

	return &routerDependency{opts.FiberApp, opts.Definition}
}

func (r *routerDependency) ListenThoughtPort(opts ListenOpts) {
	rootGroup := r.fiberApp.Group(rootGroupPath, func(c *fiber.Ctx) error { return c.Next() })
	for _, rg := range r.definition {
		group := rootGroup.Group(rg.Path, rg.Handler...)

		for _, ro := range rg.Routes {
			switch ro.Method {
			case fiber.MethodConnect:
				group.Connect(ro.Path, ro.Handler...)

			case fiber.MethodDelete:
				group.Delete(ro.Path, ro.Handler...)

			case fiber.MethodGet:
				group.Get(ro.Path, ro.Handler...)

			case fiber.MethodHead:
				group.Head(ro.Path, ro.Handler...)

			case fiber.MethodOptions:
				group.Options(ro.Path, ro.Handler...)

			case fiber.MethodPatch:
				group.Patch(ro.Path, ro.Handler...)

			case fiber.MethodPost:
				group.Post(ro.Path, ro.Handler...)

			case fiber.MethodPut:
				group.Put(ro.Path, ro.Handler...)

			case fiber.MethodTrace:
				group.Trace(ro.Path, ro.Handler...)
			}

			annouce := fmt.Sprintf(
				"Method: %s | Route: %s | Create a route",
				ro.Method, rootGroupPath+rg.Path+ro.Path,
			)
			log.Println(annouce)
		}
	}

	go func() {
		if err := r.fiberApp.Listen(fmt.Sprintf(":%d", opts.Port)); err != nil {
			log.Panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctxWithTimeout, cancel := context.WithTimeout(opts.Context, 5*time.Second)
	defer cancel()

	if err := r.fiberApp.ShutdownWithContext(ctxWithTimeout); err != nil {
		log.Panic(err)
	}
}
