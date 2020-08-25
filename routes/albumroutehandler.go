package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"photoalbum/config"
	"photoalbum/dto"
	"photoalbum/repository"
	"photoalbum/utils"
	"strconv"
)

func CreateNewAlbum(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)

	if err !=nil{
		log.Println("unable to read body")
		dto.RespondWithError(w,http.StatusInternalServerError, err.Error())
	}

	var albumrequest dto.AlbumRequest
	err = json.Unmarshal(reqBody, &albumrequest)

	if err!=nil{
		log.Println("unable to unmarshal body")
		dto.RespondWithError(w,http.StatusInternalServerError, err.Error())
	}

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	albumrepo := repository.GetAlbumRespository(db)
	album := utils.ToAlbumEntity(albumrequest)

	respalbum, err:= albumrepo.Save(album);

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		albumdto := utils.ToAlbumDTO(*respalbum)
		config.ProducerMessage(struct{
			AlbumId int
			Message string
		}{
			AlbumId: albumdto.AlbumID,
			Message: "Album Create Successfully",
		},
		)
		dto.RespondWithJSON(w, http.StatusCreated, albumdto)
	}
}


func AlbumsGETHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["userid"]

	userid, err := strconv.Atoi(key)

	if err != nil {dto.RespondWithError(w, http.StatusInternalServerError, err.Error())}

	db, err:= config.GetConnection()

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	albumrepo := repository.GetAlbumRespository(db)
	albums, err:= albumrepo.GetAlbumsByUserId(userid);

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		dto.RespondWithJSON(w, http.StatusOK, albums)
	}
}



func AlbumGETHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["userId"]
	userid, err := strconv.Atoi(key)
	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	aid := vars["albumId"]
	albumid, err := strconv.Atoi(aid)

	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	db, err:= config.GetConnection()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	albumrepo := repository.GetAlbumRespository(db)
	album, err:= albumrepo.GetAlbumByUserId(albumid, userid);

	albumdto := utils.ToAlbumDTO(album)
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		dto.RespondWithJSON(w, http.StatusOK, albumdto)
	}
}


func AlbumsDeleteHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["userId"]
	userid, err := strconv.Atoi(key)
	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	aid := vars["albumId"]
	albumid, err := strconv.Atoi(aid)

	if err != nil {
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	db, err:= config.GetConnection()

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	albumrepo := repository.GetAlbumRespository(db)
	_, err = albumrepo.DeleteAlbum(albumid, userid);

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		config.ProducerMessage(struct{
			AlbumId int
			Message string
		}{
				AlbumId: albumid,
				Message: "Delete succesfully",
		},
		)
		dto.RespondWithJSON(w, http.StatusAccepted, "Request accepted successfully")
	}
}
