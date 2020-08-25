package utils

import (
	"photoalbum/dto"
	"photoalbum/models"
	"time"
)

func ToUserEntity(r dto.UserRequest) models.User {
	return models.User{Fname: r.Fname, Lname: r.Lname,Email: r.Email,Contact: r.Contact,Address: r.Address, Password: r.Password, Create: time.Now(),}
}


func ToUserDTO(r models.User) dto.UserDTO {
	return dto.UserDTO{Fname: r.Fname, Lname: r.Lname,Email: r.Email,Contact: r.Contact,Address: r.Address}
}


func ToAlbumEntity(r dto.AlbumRequest) models.Album {
	return models.Album{AlbumName: r.AlbumName, CreatedAt: time.Now(), UpdatedAt: time.Now(), UserId: r.UserId, Description: r.AlbumDescription}
}

func ToAlbumDTO(r models.Album)  dto.AlbumDTO {
		var albumdto dto.AlbumDTO
		albumdto.AlbumName = r.AlbumName
		albumdto.Description = r.Description
		albumdto.CreatedAt = r.CreatedAt
		albumdto.UpdatedAt = r.UpdatedAt
		albumdto.UserId = r.UserId
		albumdto.AlbumID = r.AlbumID

		var imagedtos []dto.ImageDTO
		for _, v := range r.Images{
			var imagedto dto.ImageDTO
			imagedto.ImageID = v.ImageID
			imagedto.UpdatedAt = v.UpdatedAt
			imagedto.CreatedAt = v.CreatedAt
			imagedto.Imagefile = v.Imagefile
			imagedto.ImageName = v.ImageName
			imagedto.AlbumId = v.AlbumId
			imagedtos = append(imagedtos, imagedto)
		}
		albumdto.Images = imagedtos
		return albumdto
}


func ToImageDTOs(r []models.Image) []dto.ImageDTO {
	dtos := make([]dto.ImageDTO, 0)

	for _, v := range r{
		var imageDto dto.ImageDTO
		//b := make([]byte, len(v.AlbumThumbnail))
		//albumdto.AlbumThumbnail = base64.StdEncoding.EncodeToString(b)

		imageDto.ImageID = v.ImageID
		imageDto.ImageName = v.ImageName
		imageDto.UpdatedAt = v.UpdatedAt
		imageDto.CreatedAt = v.CreatedAt
		imageDto.Imagefile = v.Imagefile
		imageDto.AlbumId = v.AlbumId

		dtos = append(dtos, imageDto)
	}
	return dtos
}

func ToImageDTO(r models.Image) dto.ImageDTO {

		var imageDto dto.ImageDTO

		imageDto.ImageID = r.ImageID
		imageDto.ImageName = r.ImageName
		imageDto.UpdatedAt = r.UpdatedAt
		imageDto.CreatedAt = r.CreatedAt
		imageDto.Imagefile = r.Imagefile
		imageDto.AlbumId = r.AlbumId
		return imageDto
}



