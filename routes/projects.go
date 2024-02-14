package routes

import (
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/Fancy11111/gotmx/components"
	"github.com/Fancy11111/gotmx/persistence"
	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	store *persistence.Queries
}

func newProjectHandler(store *persistence.Queries) *ProjectHandler {
	return &ProjectHandler{
		store: store,
	}
}

func (h *ProjectHandler) GetProjects(c *fiber.Ctx) error {
	ps, err := h.store.GetProjects(c.Context())
	if err != nil {
		log.Error().Err(err).Msg("Error loading projects")
		return err
	}
	return renderView(c, components.Layout(components.Projects(ps), "Projects"))
}

func (h *ProjectHandler) GetProject(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", c.Params("id")).Msg("Error parsing id")
		return err
	}

	project, err := h.store.GetProjectById(c.Context(), id)
	if err != nil {
		log.Error().Err(err).Int64("id", id).Msg("Error loading project")
		return err
	}

	return renderView(c, components.Project(project))
}

func (h *ProjectHandler) GetProjectForEdit(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", c.Params("id")).Msg("Error parsing id")
		return err
	}

	project, err := h.store.GetProjectById(c.Context(), id)
	if err != nil {
		log.Error().Err(err).Int64("id", id).Msg("Error loading project")
		return err
	}

	return renderView(c, components.ProjectEdit(project))
}

func (h *ProjectHandler) UpdateProject(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", c.Params("id")).Msg("Error parsing id")
		return err
	}

	p := new(NewProject)
	if err := c.BodyParser(p); err != nil {
		log.Error().Err(err).Interface("body", p).Msg("error parsing body")
		return err
	}

	log.Debug().Interface("body", p).Msg("Parsed request")

	newProject, err := h.store.UpdateProject(c.Context(), persistence.UpdateProjectParams{
		ID:   id,
		Name: p.Name,
	})
	if err != nil {
		log.Error().Err(err).Int64("id", id).Msg("error updating project")
		return err
	}

	return renderView(c, components.Project(newProject))
}

func (h *ProjectHandler) DeleteProject(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", c.Params("id")).Msg("Error parsing id")
		return err
	}
	err = h.store.DeleteProject(c.Context(), id)
	if err != nil {
		log.Error().Err(err).Int64("id", id).Msg("error deleting project")
		return err
	}
	c.Status(200)
	return nil
}

func (h *ProjectHandler) CreateProject(c *fiber.Ctx) error {
	p := new(NewProject)
	if err := c.BodyParser(p); err != nil {
		log.Error().Err(err).Interface("body", p).Msg("error parsing body")
		return err
	}

	log.Debug().Interface("body", p).Msg("Parsed request")

	newProject, err := h.store.InsertProject(c.Context(), p.Name)
	if err != nil {
		log.Error().Err(err).Interface("body", p).Msg("error inserting project")
		return err
	}
	return renderView(c, components.Project(newProject))
}
