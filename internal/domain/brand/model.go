package brand

type Brand struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Img         string `json:"image"`
	Description string `json:"description"`
}
