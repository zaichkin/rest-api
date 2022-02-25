package brand

type UpdateBrandDTO struct {
	Id          string  `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Img         *string `json:"image,omitempty"`
	Description *string `json:"description,omitempty"`
}

type CreateBrandDTO struct {
	Title       string `json:"title"`
	Img         string `json:"image,omitempty"`
	Description string `json:"description,omitempty"`
}

type DeleteBrandDTO struct {
	Id string `json:"id"`
}

type GetBrandDTO struct {
	Id string `json:"id"`
}
