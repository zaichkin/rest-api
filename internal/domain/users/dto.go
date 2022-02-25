package users

type SingUpDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SingInDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
