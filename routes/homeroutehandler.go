package routes

import (
	"photoalbum/dto"
	"net/http"
)

func HomeGetHandler(w http.ResponseWriter, r *http.Request){
	dto.RespondWithJSON(w, http.StatusOK, "Welcome home page")
}