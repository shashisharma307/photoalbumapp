package dto

type UserDTO struct {
	UserId int `json:"user_id"`
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Contact int64 `json:"contact"`
	Address string `json:"address"`
	Email string `json:"email"`
}
