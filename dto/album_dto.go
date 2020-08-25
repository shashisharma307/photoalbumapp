package dto

import "time"

type AlbumDTO struct{
	AlbumID int `json:"album_id"`
	AlbumName string `json:"album_name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Images []ImageDTO `json:"images" `
	UserId int `json: "user_id"`
}
