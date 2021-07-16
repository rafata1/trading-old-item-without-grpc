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

	r.HandleFunc("/", handler2.Greeting).Methods("GET")
	r.HandleFunc("/api/get_all_post", handler2.GetAllPost).Methods("GET")
	r.HandleFunc("/api/login", handler2.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/signup", handler2.Signup).Methods("POST","OPTIONS")
	r.HandleFunc("/api/add_post", handler2.AddPost).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/get_user_posts/{owner_id}", handler2.GetPostsOfUser).Methods("GET")
	r.HandleFunc("/api/get_post_by_id/{post_id}", handler2.GetPostByID).Methods("GET")
	port := os.Getenv("PORT")
	myServer := http.Server{Addr: ":"+port, Handler: r }
	fmt.Println("Listening on port " + port)
	myServer.ListenAndServe()
}
