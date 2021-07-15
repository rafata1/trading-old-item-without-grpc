package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	handler2 "trading.olditem.app/core/handler"
	"trading.olditem.app/core/repo"
)

func main() {

	//conect to database
	repo.ConnectToDB()
	r := mux.NewRouter()

	r.HandleFunc("/api/login", handler2.Login).Methods(http.MethodPost)
	r.HandleFunc("/api/signup", handler2.Signup).Methods(http.MethodPost)

	myServer := http.Server{Addr: ":8000", Handler: r }
	fmt.Println("Listening on port 8000")
	myServer.ListenAndServe()
}
