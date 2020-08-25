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

func UsersGETHandler(w http.ResponseWriter, r *http.Request){
	db, err:= config.GetConnection()

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	userrepository := repository.GetUserRespository(db)
	users, err:= userrepository.GetAll();

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		dto.RespondWithJSON(w, http.StatusOK, users)
	}
}

func UserGETHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	s, err := strconv.Atoi(key)

	if err != nil {dto.RespondWithError(w, http.StatusInternalServerError, err.Error())}

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	userrepository := repository.GetUserRespository(db)
	users, err:= userrepository.GetByID(s)



	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		u:= utils.ToUserDTO(users)
		dto.RespondWithJSON(w, http.StatusOK, u)
		return
	}
}


func CreateNewUser(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)

	if err !=nil{
		log.Println("unable to read body")
		dto.RespondWithError(w,http.StatusInternalServerError, err.Error())
	}

	var userrequest dto.UserRequest
	err = json.Unmarshal(reqBody, &userrequest)

	if err!=nil{
		log.Println("unable to unmarshal body")
		dto.RespondWithError(w,http.StatusInternalServerError, err.Error())
	}

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	userrepository := repository.GetUserRespository(db)
	user := utils.ToUserEntity(userrequest)

	userresp, err := userrepository.Save(user)

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		userdto := utils.ToUserDTO(userresp)
		dto.RespondWithJSON(w, http.StatusCreated, userdto)
	}

}


func UserDeleteHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	s, err := strconv.Atoi(key)

	if err != nil {dto.RespondWithError(w, http.StatusInternalServerError, err.Error())}

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	userrepository := repository.GetUserRespository(db)
	_, err = userrepository.Delete(s)

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		dto.RespondWithJSON(w, http.StatusOK, "user successfully deleted")
	}
}

