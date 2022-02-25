package worker

type Worker struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Workspace   int    `json:"workspace"`
}
