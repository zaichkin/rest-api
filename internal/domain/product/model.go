package product

type Product struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Size        string `json:"size"`
	Category    int    `json:"category"`
	Gender      int    `json:"gender"`
	Brand       int    `json:"brand"`
	Description string `json:"description"`
	Img         string `json:"image"`
}
