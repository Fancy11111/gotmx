package routes

type Project struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type NewProject struct {
	Name string `json:"name,omitempty"`
}
