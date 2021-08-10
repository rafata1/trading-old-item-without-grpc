package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
	handler2 "trading.olditem.app/core/handler"
	"trading.olditem.app/core/repo"
)

func main() {

	//conect to database

	for true {
		if repo.ConnectToDB() == false {
			fmt.Println("retrying connect to DB after 60 seconds")
		} else {
			break
		}
		time.Sleep(60 * time.Second)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handler2.Greeting).Methods("GET")

	//user services
	r.HandleFunc("/api/login", handler2.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/signup", handler2.Signup).Methods("POST","OPTIONS")
	r.HandleFunc("/api/get_all_user", handler2.GetAllUser).Methods("GET")
	r.HandleFunc("/api/get_user_by_id/{user_id}", handler2.GetUserByID).Methods("GET")
	r.HandleFunc("/api/get_user_history/{user_id}", handler2.GetUserHistory).Methods("GET")

	//post service
	r.HandleFunc("/api/get_all_post", handler2.GetAllPost).Methods("GET")
	r.HandleFunc("/api/add_post", handler2.AddPost).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/get_user_posts/{owner_id}", handler2.GetPostsOfUser).Methods("GET")
	r.HandleFunc("/api/get_post_by_id/{post_id}", handler2.GetPostByID).Methods("GET")
	r.HandleFunc("/api/search_post/{keyword}", handler2.SearchByKeyword).Methods("GET")
	r.HandleFunc("/api/deactivate_post", handler2.DeactivatePost).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/edit_post", handler2.EditPost).Methods("POST", "OPTIONS")

	//trading service
	r.HandleFunc("/api/create_transaction", handler2.CreateTransaction).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/get_product_want_to_trade_with/{product_id}", handler2.Get_B_BY_A).Methods("GET")
	r.HandleFunc("/api/complete_transaction", handler2.CompleteTransaction).Methods("POST", "OPTIONS")


	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	myServer := http.Server{Addr: ":"+port, Handler: r }
	fmt.Println("Listening on port " + port)
	myServer.ListenAndServe()
}
