package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)

	html, err := os.ReadFile("pages\\index.html")
	if err != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"code": 500, "err":true, "msg":"error opening the file: %s"}`, err.Error())))
		return
	}

	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(html)
}

func signinPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)

	html, err := os.ReadFile("pages\\login.html")
	if err != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"code": 500, "err":true, "msg":"error opening the file: %s"}`, err.Error())))
		return
	}

	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(html)
}

func signupPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)

	html, err := os.ReadFile("pages\\register.html")
	if err != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"code": 500, "err":true, "msg":"error opening the file: %s"}`, err.Error())))
		return
	}

	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(html)
}

func signin(w http.ResponseWriter, r *http.Request) {

}

func signup(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		responseError(w, 500, err)
		return
	}

	var user User
	err = json.Unmarshal(content, &user)
	if err != nil {
		log.Println(err)
		responseError(w, 400, err)
		return
	}

	fmt.Println(user)
}
