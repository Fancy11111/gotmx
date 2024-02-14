package routes

import "github.com/Fancy11111/gotmx/persistence"

type Handlers struct {
	Project *ProjectHandler
	Task    *TaskHandler
}

func SetupHandlers(store *persistence.Queries) Handlers {
	return Handlers{
		Project: newProjectHandler(store),
		Task:    newTaskHandler(store),
	}
}
