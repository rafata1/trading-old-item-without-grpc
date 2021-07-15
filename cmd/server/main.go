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

	r.HandleFunc("/", handler2.Greeting)
	r.HandleFunc("/api/login", handler2.Login)
	r.HandleFunc("/api/signup", handler2.Signup)

	port := os.Getenv("PORT")
	myServer := http.Server{Addr: ":"+port, Handler: r }
	fmt.Println("Listening on port " + port)
	myServer.ListenAndServe()
}
