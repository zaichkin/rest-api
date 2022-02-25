package category

type CreateCategoryDTO struct {
	Parent int    `json:"parent,omitempty"`
	Name   string `json:"name"`
}

type UpdateCategoryDTO struct {
	Id     string  `json:"id,omitempty"`
	Parent *int    `json:"parent,omitempty"`
	Name   *string `json:"name,omitempty"`
}

type DeleteCategoryDTO struct {
	Id string `json:"id"`
}

type GetCategoryDTO struct {
	Id string `json:"id"`
}
