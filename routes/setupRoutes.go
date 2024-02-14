package routes

import (
	"log"
	"strconv"

	"github.com/Fancy11111/gotmx/components"
	"github.com/Fancy11111/gotmx/persistence"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func renderView(c *fiber.Ctx, comp templ.Component) error {
	handler := adaptor.HTTPHandler(templ.Handler(comp))
	return handler(c)
}

type Project struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type NewProject struct {
	Name string `json:"name,omitempty"`
}

func SetupRoutes(app *fiber.App, store *persistence.Queries) {
	app.Static("/static", "./static")

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		log.Print("receiving ping, sending pong")
		return ctx.SendString("pong")
	})

	app.Get("/hello/:name?", func(c *fiber.Ctx) error {
		comp := components.Hello(c.Params("name", "World"))
		return renderView(c, comp)
	})

	app.Get("/projects", func(c *fiber.Ctx) error {
		ps, err := store.GetProjects(c.Context())
		if err != nil {
			return err
		}
		return renderView(c, components.Layout(components.Projects(ps), "Projects"))
	})

	app.Post("/projects", func(c *fiber.Ctx) error {
		p := new(NewProject)
		if err := c.BodyParser(p); err != nil {
			log.Printf("error parsing: %v\n", err)
			return err
		}

		log.Printf("Parsed request: %v\n", p)

		newProject, err := store.InsertProject(c.Context(), p.Name)
		if err != nil {
			log.Printf("error inserting: %v\n", err)
			return err
		}
		return renderView(c, components.Project(newProject))
	})

	app.Delete("/projects/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			log.Printf("Error parsing id %d: %v", id, err)
			return err
		}
		err = store.DeleteProject(c.Context(), id)
		if err != nil {
			log.Printf("Error deleting project with id %d: %v", id, err)
			return err
		}
		c.Status(200)
		return nil
	})

	app.Get("/projects/:id/edit", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			log.Printf("Error parsing id %d: %v", id, err)
			return err
		}

		project, err := store.GetProjectById(c.Context(), id)
		if err != nil {
			log.Printf("Error loading project with id %d: %v", id, err)
			return err
		}

		return renderView(c, components.ProjectEdit(project))
	})

	app.Get("/projects/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			log.Printf("Error parsing id %d: %v", id, err)
			return err
		}

		project, err := store.GetProjectById(c.Context(), id)
		if err != nil {
			log.Printf("Error loading project with id %d: %v", id, err)
			return err
		}

		return renderView(c, components.Project(project))
	})

	app.Put("/projects/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			log.Printf("Error parsing id %d: %v", id, err)
			return err
		}

		p := new(NewProject)
		if err := c.BodyParser(p); err != nil {
			log.Printf("error parsing: %v\n", err)
			return err
		}

		log.Printf("Parsed request: %v\n", p)

		newProject, err := store.UpdateProject(c.Context(), persistence.UpdateProjectParams{
			ID:   id,
			Name: p.Name,
		})
		if err != nil {
			log.Printf("error inserting: %v\n", err)
			return err
		}

		return renderView(c, components.Project(newProject))
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		ps, err := store.GetTasks(c.Context())
		if err != nil {
			return err
		}
		return renderView(c, components.Layout(components.Tasks(ps), "Tasks"))
	})
}
