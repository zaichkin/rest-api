package users

type User struct {
	Uuid     string `json:"uuid"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
