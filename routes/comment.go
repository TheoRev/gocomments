package routes

import (
	"github.com/TheoRev/gocomments/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCommentRouter establece la ruta de los comentarios
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods("POST")
	subRouter.HandleFunc("/", controllers.CommentGetAll).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
