package worker

type CreateWorkerDTO struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Workspace   int    `json:"workspace"`
}

type UpdateWorkerDTO struct {
	Id          string  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Workspace   *int    `json:"workspace,omitempty"`
}

type DeleteWorkerDTO struct {
	Id string `json:"id"`
}

type GetWorkerDTO struct {
	Id string `json:"id"`
}
