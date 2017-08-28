package routes

import (
	"github.com/TheoRev/gocomments/controllers"
	"github.com/gorilla/mux"
)

// SetVoteRouter es la ruta para añadir votos
func SetVoteRouter(router *mux.Router) {
	prefix := "/api/votes"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.VoteRegister).Methods("POST")
}
