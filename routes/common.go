package routes

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func renderView(c *fiber.Ctx, comp templ.Component) error {
	handler := adaptor.HTTPHandler(templ.Handler(comp))
	return handler(c)
}
