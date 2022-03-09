package controllers

import (
	"github.com/MoraGames/didPlus/go/internal/utils"
	"net/http"
	"os"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "statics/index.html")
}

func ServeSignIn(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "statics/login.html")
}

func ServeSignUp(w http.ResponseWriter, r *http.Request) {
	serveFile(w, "statics/register.html")
}

func serveFile(w http.ResponseWriter, fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		utils.ResponseError(w, 500, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
}
