package controllers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MoraGames/didPlus/go/internal/config"
	"github.com/MoraGames/didPlus/go/internal/models/user"
	"github.com/MoraGames/didPlus/go/internal/utils"
	"io"
	"log"
	"net/http"
	"strings"
)

var Db *sql.DB
var Cfg *config.Config

func SignUp(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		utils.ResponseError(w, 500, err)
		return
	}

	// Parse request json
	var u user.User
	err = json.Unmarshal(content, &u)
	if err != nil {
		log.Println(err)
		utils.ResponseError(w, 400, err)
		return
	}

	// Validate user
	if err = u.IsValid(); err != nil {
		log.Println(err)
		utils.ResponseError(w, 400, err)
		return
	}

	// Check exist
	ok, err := user.CheckUserExist(u.Email, Db)
	if err != nil {
		log.Println(err)
		utils.ResponseError(w, 500, err)
		return
	}

	if ok {
		utils.ResponseError(w, 409, errors.New("User already exist"))
		return
	}

	// Encrypt password
	u.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))

	// Create user
	u.SetDb(Db)
	newUser, err := u.Create()
	if err != nil {
		log.Println(err)
		utils.ResponseError(w, 500, err)
		return
	}
	newUser.Password = ""

	// Create json
	data, err := json.MarshalIndent(newUser, " ", "\t")
	if err != nil {
		log.Println(err)
		utils.ResponseError(w, 500, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("http://%s%s/signin", Cfg.Server.Host, Cfg.Server.Port))
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(302)
	w.Write(data)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		utils.ResponseError(w, 500, err)
		return
	}

	// Parse request json
	var u user.User
	err = json.Unmarshal(content, &u)
	if err != nil {
		log.Println(err)
		utils.ResponseError(w, 400, err)
		return
	}

	// Validate user
	if err = u.IsValid(); err != nil {
		log.Println(err)
		utils.ResponseError(w, 400, err)
		return
	}

	newUser, err := user.GetUserByEmail(u.Email, Db)
	if err != nil {
		log.Println(err)
		utils.ResponseError(w, 500, nil)
		return
	}

	// If not Exist
	if newUser.Email == "" {
		log.Println(err)
		utils.ResponseError(w, 400, errors.New("Incorrect email or password"))
		return
	}

	// Encrypt password
	u.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))

	fmt.Println("DEBUG:", u.Email)
	fmt.Printf("DEBUG:\"%s\"\n", u.Password)
	fmt.Printf("DEBUG:\"%s\"\n", newUser.Password)

	//TODO: Da eliminare, equalfold è un buco di sicurezza perchè non tiene conto del case-sensitive
	if !strings.EqualFold(u.Password, newUser.Password) {
		utils.ResponseError(w, 400, errors.New("Incorrect email or password"))
		return
	}

	utils.SetCookies(w, u.Email)
	w.Header().Set("Location", fmt.Sprintf("http://%s%s/", Cfg.Server.Host, Cfg.Server.Port))
	w.WriteHeader(302)
	w.Write([]byte(""))
}
