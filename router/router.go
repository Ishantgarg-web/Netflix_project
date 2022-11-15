package router

import (
	"github.com/Ishant-tata/NetFlix_Project_MongoDB/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/all", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/all", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
