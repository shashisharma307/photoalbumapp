package routes

import (
	"image"
	"image/png"
	"photoalbum/config"
	"photoalbum/dto"
	"photoalbum/models"
	"photoalbum/repository"
	"photoalbum/utils"

	//"photoalbum/utils"
	"bytes"
	"encoding/base64"
	"github.com/disintegration/imaging"
	"github.com/gorilla/mux"
	"image/jpeg"
	//"log"
	"net/http"
	"strconv"
	"strings"
)


func AddImageToAlbumHandler(w http.ResponseWriter, r *http.Request) {

	var imageentity models.Image
	var buf bytes.Buffer
	var img image.Image

	vars := mux.Vars(r)
	key := vars["albumId"]
	s, err := strconv.Atoi(key)

	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("fileupload")
	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()



	if strings.Contains(strings.ToLower(handler.Filename), ".jpg") || strings.Contains(strings.ToLower(handler.Filename), ".jpeg"){
		img, err = jpeg.Decode(file)
		if err != nil {
			dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = imaging.Encode(&buf, img, imaging.JPEG)
		if err != nil {
			dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if strings.Contains(strings.ToLower(handler.Filename), ".png"){
		img, err = png.Decode(file)
		if err != nil {
			dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = imaging.Encode(&buf, img, imaging.PNG)
		if err != nil {
			dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	image := base64.StdEncoding.EncodeToString(buf.Bytes())
	imageentity.ImageName = r.Form.Get("file_name")
	imageentity.Imagefile = image
	imageentity.AlbumId = s

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	imagerepo := repository.GetImageRespository(db)
	a, err := imagerepo.Save(imageentity)

	if err !=nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else{
		config.ProducerMessage(struct {
			AlbumId int
			ImageId int
			Message string
		}{
			AlbumId: s,
			ImageId: a.ImageID,
			Message: "Create Image in Album successfully",
		},
		)
		dto.RespondWithJSON(w, http.StatusCreated, a)
		return
	}
	dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	return
}

func GetImagesInAlbumHandler (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["albumId"]
	s, err := strconv.Atoi(key)
	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	db, err:= config.GetConnection()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	imagerepo := repository.GetImageRespository(db)
	images, err:= imagerepo.GetAllImagesInAlbum(s)
	if err !=nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	imagedtos := utils.ToAlbumDTO(images)
	dto.RespondWithJSON(w, http.StatusOK, imagedtos)
	return
}

func GetImageInAlbumHandler (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["albumId"]
	albumid, err := strconv.Atoi(aid)
	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	imgId := vars["imageId"]
	imageId, err := strconv.Atoi(imgId)
	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	db, err:= config.GetConnection()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	imagerepo := repository.GetImageRespository(db)
	images, err:= imagerepo.GetImageInAlbum(albumid, imageId)
	if err !=nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	imagedtos := utils.ToImageDTO(images)
	dto.RespondWithJSON(w, http.StatusOK, imagedtos)
	return
}


func DeleteImageInAlbumHandler (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["albumId"]
	albumid, err := strconv.Atoi(aid)
	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	imgId := vars["imageId"]
	imageId, err := strconv.Atoi(imgId)
	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	db, err:= config.GetConnection()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	imagerepo := repository.GetImageRespository(db)
	_, err = imagerepo.DeleteImageInAlbum(albumid, imageId)
	if err !=nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}else {
		config.ProducerMessage(struct {
			AlbumId int
			ImageId int
			Message string
		}{
			AlbumId: albumid,
			ImageId: imageId,
			Message: "Delete Image in Album Successfully",
		},
		)
		dto.RespondWithJSON(w, http.StatusAccepted, "Request Accepted succesffully")
		return
	}
	dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	return
}
