package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	handler2 "trading.olditem.app/core/handler"
	"trading.olditem.app/core/repo"
)

func main() {

	//conect to database
	repo.ConnectToDB()
	r := mux.NewRouter()

	r.HandleFunc("/", handler2.Greeting).Methods(http.MethodGet)
	r.HandleFunc("/api/login", handler2.Login).Methods(http.MethodPost)
	r.HandleFunc("/api/signup", handler2.Signup).Methods(http.MethodPost)

	port := os.Getenv("PORT")
	myServer := http.Server{Addr: ":"+port, Handler: r }
	fmt.Println("Listening on port " + port)
	myServer.ListenAndServe()
}
