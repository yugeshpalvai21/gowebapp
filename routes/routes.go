package routes

import (
	"gowebapp/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	userController := &controllers.UserController{}
	pagesController := &controllers.PagesController{}

	router.HandleFunc("/", pagesController.Home).Methods("GET")
	router.HandleFunc("/users", userController.GetUserByID).Methods("GET")

	return router
}
