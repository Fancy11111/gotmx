package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Fancy11111/gotmx/persistence"
	"github.com/Fancy11111/gotmx/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func setupRoutes(app *fiber.App, handler routes.Handlers) {
	app.Static("/static", "./static")

	projects := app.Group("/projects")

	projects.Use(func(c *fiber.Ctx) error {
		hx := len(c.GetReqHeaders()["Hx-Request"]) > 0

		log.Debug().Bool("hx-request", hx).Msg("Is HX-Request?")
		return c.Next()
	})

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	projects.Get("/", handler.Project.GetProjects)
	projects.Post("/", handler.Project.CreateProject)
	projects.Get("/:id", handler.Project.GetProject)
	projects.Get("/:id/edit", handler.Project.GetProjectForEdit)
	projects.Put("/:id", handler.Project.UpdateProject)
	projects.Delete("/:id", handler.Project.DeleteProject)

	tasks := app.Group("/tasks")
	tasks.Get("/", handler.Task.GetTasks)
}

func StartServer() {
	store, err := persistence.SetupStore()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not setup persistence store")
	}

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Debug().Str("method", string(c.Request().Header.Method())).Str("path", string(c.Request().URI().Path())).Send()
		return c.Next()
	})

	handler := routes.SetupHandlers(store)

	setupRoutes(app, handler)

	go func() {
		log.Info().Int("port", 3000).Msg("Starting webserver")
		app.Listen(":3000")
	}()

	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sig := <-c
	println()
	log.Info().Interface("signal", sig).Msg("Got signal")
	log.Info().Msg("Shutting down")
	app.Shutdown()
}
