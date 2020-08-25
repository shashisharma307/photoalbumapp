package dto

type AlbumRequest struct {
	AlbumName string `json:"album_name"`
	AlbumDescription string `json:"description"`
	UserId int `json:"user_id"`
}
