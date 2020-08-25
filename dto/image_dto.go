package dto

import "time"

type ImageDTO struct{
	ImageID int `json:"image_id"`
	ImageName string `json:"image_name"`
	Imagefile string `json:"imagefile"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	AlbumId int `json: album_id`
}
