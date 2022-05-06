package main

import (
	"github.com/MoraGames/didPlus/go/internal/controllers"
	"github.com/gorilla/mux"
)

func routerInit() *mux.Router {
	//Set DB in the controllers package
	controllers.Db = db
	controllers.Cfg = cfg

	//Create new router base
	router := mux.NewRouter()

	//Create new subrouter for APIs
	//routerApi := router.PathPrefix(api + os.Getenv("API_VERSION")).Subrouter()

	//Set Basic Handlers
	//router.HandleFunc("/", controllers.ServeIndex)
	//router.PathPrefix("/statics/").Handler(http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))
	//router.HandleFunc(signin, controllers.ServeSignIn)
	//router.HandleFunc(signup, controllers.ServeSignUp)

	//Set User Handlers
	//routerApi.HandleFunc(signup, controllers.SignUp)
	//routerApi.HandleFunc(signin, controllers.SignIn)

	return router
}
