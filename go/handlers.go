package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)

	html, err := os.ReadFile("index.html")
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

	html, err := os.ReadFile("login.html")
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

	html, err := os.ReadFile("register.html")
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

}
