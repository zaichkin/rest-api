package category

type Category struct {
	Id     int    `json:"id"`
	Parent int    `json:"parent"`
	Name   string `json:"name"`
}
