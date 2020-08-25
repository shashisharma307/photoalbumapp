package dto

type UserRequest struct {
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Contact int64 `json:"contact"`
	Address string `json:"address"`
	Email string `json:"email"`
	Password string `json:"password"`
}
