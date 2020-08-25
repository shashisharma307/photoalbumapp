package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"photoalbum/models"
)


type ImageRepositoryError struct {
	error
}




type ImageRepository struct {
	DB *gorm.DB
}

func GetImageRespository(db *gorm.DB) ImageRepository{
	return ImageRepository {DB: db}
}

func (u *ImageRepository) Save(album models.Image) (models.Image,error){
	d := u.DB.Save(&album)
	if d.Error !=nil{
		return  album, &ImageRepositoryError{fmt.Errorf(d.Error.Error())}
	}else{
		return album, nil
	}
	return  album, &ImageRepositoryError{fmt.Errorf("can not create album")}
}

func (u *ImageRepository) GetAll(albumid interface{}) ([]models.Image, error){
	var images []models.Image
	u.DB.Debug().Where("album_id = ?", albumid).Find(&images)
	if len(images) > 0 {
		return images, nil
	}

	err := fmt.Errorf("Server error")
	return nil, &ImageRepositoryError{err}
}

func (u *ImageRepository) GetAllImagesInAlbum(albumid interface{})(models.Album, error){
	var album models.Album
	var images []models.Image

	u.DB.Debug().Where("album_id = ?", albumid).Find(&album)

	if album.AlbumID == 0 && !(album.AlbumID == albumid){
		log.Print("Album succesffully fetched")
		err := fmt.Errorf("No record found")
		return album, &ImageRepositoryError{err}
	}
	u.DB.Debug().Where("album_id = ?", albumid).Find(&images)
	if len(images) > 0 {
		album.Images = images
		return album, nil
	}
	err := fmt.Errorf("Server error")
	return album, &ImageRepositoryError{err}
}



func (u *ImageRepository) GetImageInAlbum(albumid int ,imageid int) (models.Image, error){

	var image models.Image
	var album models.Album

	u.DB.Debug().Where("album_id = ?", albumid).Find(&album)

	if album.AlbumID != albumid{
		err := fmt.Errorf("album id not found")
		return image, &ImageRepositoryError{err}
	}

	u.DB.Debug().Where("image_id = ?", imageid).Find(&image)
	if image.ImageID != imageid {
		err := fmt.Errorf("image not found")
		return image, &ImageRepositoryError{err}
	}
	return image, nil
}

func (u *ImageRepository) DeleteImageInAlbum(albumid int ,imageid int) (bool, error){
	var image models.Image
	var album models.Album

	u.DB.Debug().Where("album_id = ?", albumid).Find(&album)

	if album.AlbumID != albumid{
		err := fmt.Errorf("album id not found")
		return false, &ImageRepositoryError{err}
	}


	db := u.DB.Debug().Where("image_id = ?", imageid).Delete(&image)

	if db.Error!=nil {
		err := fmt.Errorf(db.Error.Error())
		return false, &ImageRepositoryError{err}
	}
	return true, nil
}
