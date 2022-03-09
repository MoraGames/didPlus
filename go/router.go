package main

import (
	"github.com/MoraGames/didPlus/go/internal/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func routerInit() * mux.Router {
	//Set db in the controllers package
	controllers.Db = db
	controllers.Cfg = cfg

	// Router base
	router := mux.NewRouter()

	// Router api
	routerApi := router.PathPrefix(api + os.Getenv("API_VERSION")).Subrouter()

	// Rotte base
	router.HandleFunc("/", controllers.ServeIndex)
	router.PathPrefix("/statics/").Handler(http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))
	router.HandleFunc(signin, controllers.ServeSignIn)
	router.HandleFunc(signup, controllers.ServeSignUp)

	// Rotte utente
	routerApi.HandleFunc(signup, controllers.SignUp)
	routerApi.HandleFunc(signin, controllers.SignIn)

	return router
}