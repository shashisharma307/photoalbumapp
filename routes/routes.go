package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/", HomeGetHandler).Methods("GET")
	r.HandleFunc("/user", UsersGETHandler).Methods("GET")
	r.HandleFunc("/user/{id}", UserGETHandler).Methods("GET")
	r.HandleFunc("/user", CreateNewUser).Methods("POST")
	r.HandleFunc("/user/{id}", UserDeleteHandler).Methods("DELETE")

	r.HandleFunc("/album", CreateNewAlbum).Methods("POST")
	r.HandleFunc("/album/{userId}", AlbumsGETHandler).Methods("GET")
	r.HandleFunc("/album/{userId}/{albumId}", AlbumGETHandler).Methods("GET") //done
	r.HandleFunc("/album/{userId}/{albumId}", AlbumsDeleteHandler).Methods("DELETE") //done
	r.HandleFunc("/album/image/{albumId}", AddImageToAlbumHandler).Methods("POST") //done
	r.HandleFunc("/albm/image/{albumId}", GetImagesInAlbumHandler).Methods("GET") //done
	r.HandleFunc("/album/image/{albumId}/{imageId}", GetImageInAlbumHandler).Methods("GET")
	r.HandleFunc("/album/image/{albumId}/{imageId}", DeleteImageInAlbumHandler).Methods("DELETE")








	return r
}
