package components

type Root string

const (
	RootProjects = "projects"
	RootTasks    = "tasks"
	RootTags     = "tags"
)

type Page struct {
	Name  string
	Route Root
}

var pages = []Page{
	{
		Name:  "Projects",
		Route: "projects",
	},
	{
		Name:  "Tasks",
		Route: "tasks",
	},
	{
		Name:  "Tags",
		Route: "tags",
	},
}

func SetupTemplates() {
}
