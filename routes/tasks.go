package routes

import (
	"github.com/Fancy11111/gotmx/components"
	"github.com/Fancy11111/gotmx/persistence"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	store *persistence.Queries
}

func newTaskHandler(store *persistence.Queries) *TaskHandler {
	return &TaskHandler{
		store: store,
	}
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	ps, err := h.store.GetTasks(c.Context())
	if err != nil {
		return err
	}
	return renderView(c, components.Layout(components.Tasks(ps), "Tasks"))
}
