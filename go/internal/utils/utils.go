package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ResponseError (w http.ResponseWriter, statusCode int, err error) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf("{\"statusCode\": %d, \"error\": \"%s\"}", statusCode, err.Error())))
}

func PortValid(port string) error {
	portInt, err := strconv.Atoi(strings.Trim(port, ":"))
	if err != nil {
		return errors.New("Invalid Port: Conversion of port to int not valid")
	}

	if portInt < 1024 || portInt > 49151 {
		return errors.New("Port not valid. [1024-49151]")
	}
	return nil
}

//TODO: Modificare il cookie in modo sensato
func SetCookies(w http.ResponseWriter, email string, ) {
	cookie := &http.Cookie{
		Name:     "didPlus",
		Value:    "Email=" + RandomString(15),
		Secure:   false,
		HttpOnly: false,
		Path:     "/",
		Expires: time.Now().Add(10 * time.Minute),
	}

	fmt.Println("COOKIE", cookie)
	http.SetCookie(w, cookie)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}