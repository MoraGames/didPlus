package main

import (
	"fmt"
	"net/http"
)

func responseError (w http.ResponseWriter, statusCode int, err error) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf("{\"statusCode\": %d, \"error\": \"%s\"}", statusCode, err)))
}