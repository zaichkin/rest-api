package product

type CreateProductDTO struct {
	Title       string `json:"title"`
	Price       int    `json:"price,omitempty"`
	Size        string `json:"size,omitempty"`
	Category    int    `json:"category"`
	Gender      int    `json:"gender"`
	Brand       int    `json:"brand"`
	Description string `json:"description,omitempty"`
	Img         string `json:"image,omitempty"`
}

type UpdateProductDTO struct {
	Id          string  `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Price       *int    `json:"price,omitempty"`
	Size        *string `json:"size,omitempty"`
	Category    *int    `json:"category,omitempty"`
	Gender      *int    `json:"gendery,omitempty"`
	Brand       *int    `json:"brand,omitempty"`
	Description *string `json:"description,omitempty"`
	Img         *string `json:"image,omitempty"`
}

type DeleteProductDTO struct {
	Id string `json:"id"`
}

type GetProductDTO struct {
	Id string `json:"id"`
}
