package routes

import (
	"gowebapp/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	userController := &controllers.UserController{}

	router.HandleFunc("/users", userController.GetUserByID).Methods("GET")

	return router
}
