package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"trading.olditem.app/core/repo"
	"trading.olditem.app/core/type"
)


func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "POST,OPTIONS")
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to trading-old-item server")
}

func decodeJson(request *http.Request, v interface{} ) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		panic(err)
	}
}

func encodeJson(v interface{}, jData *[]byte)  {
	*jData, _ = json.Marshal(v)
}

func Login(writer http.ResponseWriter, request *http.Request) {


	//handle preflight
	addCorsHeader(writer)
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusOK)
		return
	}


	var logReq _type.LoginRequest
	decodeJson(request, &logReq)

	statusCode, detail := repo.Login(logReq.Email, logReq.Password)

	logRes := _type.LoginResponse{StatusCode: statusCode, Detail: detail}
	var jData []byte
	encodeJson(logRes, &jData)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jData)
}

func Signup(writer http.ResponseWriter, request *http.Request)  {


	//handle preflight
	addCorsHeader(writer)
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusOK)
		return
	}

	var signupReq _type.SignupRequest
	decodeJson(request, &signupReq)

	statusCode, detail := repo.SignUp(signupReq.Email, signupReq.Username, signupReq.Password, signupReq.PhoneNumber,
		signupReq.Gender, signupReq.DOB)

	signupRes := _type.SignupResponse{StatusCode: statusCode, Detail: detail}
	var jData []byte
	encodeJson(signupRes, &jData)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jData)

}

func AddPost(writer http.ResponseWriter, request *http.Request)  {

	//handle preflight
	addCorsHeader(writer)
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusOK)
		return
	}

	var addPostReq _type.AddPostRequest
	decodeJson(request, &addPostReq)
	statusCode, detail := repo.AddPost(addPostReq.OwnerID, addPostReq.Name, addPostReq.Brand, addPostReq.Type,
		addPostReq.Amount, addPostReq.Description, addPostReq.ImageURL)

	addPostRes := _type.AddPostResponse{StatusCode: statusCode, Detail: detail}
	var jData []byte
	encodeJson(addPostRes, &jData)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jData)

}

func GetAllPost(writer http.ResponseWriter, request *http.Request)  {
	statusCode, detail, result := repo.GetAllPost()

	if statusCode != 200 {
		res := _type.GetAllPostRes{StatusCode: statusCode, Detail: detail}
		var jData []byte
		encodeJson(res, &jData)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jData)
		return
	}

	var jData []byte
	encodeJson(result, &jData)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jData)

}

func GetPostByID(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	post_id, _ := strconv.Atoi(params["post_id"])
	statusCode, detail, result := repo.GetPostByID(post_id)


	if statusCode != 200 {
		res := _type.GetPostByIDRes{StatusCode: statusCode, Detail: detail}
		var jData []byte
		encodeJson(res, &jData)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jData)
		return
	}

	var jData []byte
	encodeJson(result, &jData)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jData)

}

func GetPostsOfUser(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	owner_id,_ := strconv.Atoi(params["owner_id"])
	statusCode, detail, result := repo.GetPostsOfUser(owner_id)

	if statusCode != 200 {
		res := _type.GetUserPostRes{StatusCode: statusCode, Detail: detail}
		var jData []byte
		encodeJson(res, &jData)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jData)
		return
	}

	var jData []byte
	encodeJson(result, &jData)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jData)

}
