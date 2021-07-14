package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	handler2 "trading.olditem.app/core/handler"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/login", handler2.Login).Methods(http.MethodPost)
	myServer := http.Server{Addr: ":8000", Handler: r }
	fmt.Println("Listening on port 8000")
	myServer.ListenAndServe()
}
