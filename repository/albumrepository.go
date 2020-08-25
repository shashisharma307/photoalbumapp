package repository

import (
	"photoalbum/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

type AlbumRepositoryError struct {
	error
}

type AlbumRepository struct {
	DB *gorm.DB
}

func GetAlbumRespository(db *gorm.DB) AlbumRepository{
	return AlbumRepository {DB: db}
}

func (u *AlbumRepository) GetAlbumsByUserId(userid interface{}) ([]models.Album, error){
	var user models.User
	var albums []models.Album

	d :=u.DB.Debug().First(&user, userid)
	if d.Error !=nil{
		return  albums, &AlbumRepositoryError{fmt.Errorf(d.Error.Error())}
	}

	u.DB.Debug().Where("user_id = ?", userid).Find(&albums)

	if len(albums) > 0 {
		return albums, nil
	}else {
		err := fmt.Errorf("No Record Found")
		return nil, &AlbumRepositoryError{err}
	}

	err := fmt.Errorf("Server error")
	return nil, &AlbumRepositoryError{err}
}

func (u *AlbumRepository) GetAlbumByUserId(albumid int, userid int) (models.Album, error){
	var user models.User
	album := models.Album{}
	var images []models.Image

	d :=u.DB.Debug().First(&user, userid)
	if d.Error !=nil{
		return  album, &AlbumRepositoryError{fmt.Errorf(d.Error.Error())}
	}

	u.DB.Debug().First(&album, albumid)

	if album.AlbumID != albumid {
		err := fmt.Errorf("No Record Found")
		return album, &AlbumRepositoryError{err}
	}else {
		u.DB.Debug().Where("album_id = ?", albumid).Find(&images)
		album.Images = images
		return album, nil
	}

	err := fmt.Errorf("Server error")
	return album, &AlbumRepositoryError{err}
}


func (u *AlbumRepository) Save(album models.Album) (*models.Album,error){
	var user models.User
	var ab *models.Album

	d :=u.DB.Debug().First(&user, album.UserId)
	if d.Error !=nil{
		return  ab, &AlbumRepositoryError{fmt.Errorf(d.Error.Error())}
	}

	d = u.DB.Save(&album)
	if d.Error !=nil{
		return  ab, &AlbumRepositoryError{fmt.Errorf(d.Error.Error())}
	}else{
		fmt.Println(d.Value)
		ab = d.Value.(*models.Album)
		return ab, nil
	}
	return  ab, &AlbumRepositoryError{fmt.Errorf("can not create album")}
}

func (u *AlbumRepository) DeleteAlbum(albumid int, userid int) (bool, error){
	var user models.User
	album := models.Album{}


	u.DB.Debug().First(&user, userid)
	if user.UserId != userid{
		return  false, &AlbumRepositoryError{fmt.Errorf("user id not found")}
	}

	db:= u.DB.Debug().Where("album_id = ?", albumid).Delete(album)

	if db.Error!=nil{
		return  false, &AlbumRepositoryError{fmt.Errorf("user id not found")}
	}
	return true, nil



}
