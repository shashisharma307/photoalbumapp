package models

import (
	"time"
)

type User struct {
	UserId int `gorm:"primary_key" json:"user_id"`
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Contact int64 `json:"contact"`
	Address string `json:"address"`
	Email string `json:"email"`
	Password string `json:"password"`
	Create time.Time
	Albums []Album `gorm:"foreignkey:UserId" json:"albums"` //one user can have multiple albums, so here in Album table or entity there is a user_id columun. which is referening to foreign key column.
}
