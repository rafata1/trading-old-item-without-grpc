package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"trading.olditem.app/core/repo"
	_type "trading.olditem.app/core/type"
)

func addCorsHeader(res *http.ResponseWriter) {
	(*res).Header().Add("Access-Control-Allow-Origin", "*")
	(*res).Header().Add("Access-Control-Allow-Headers", "Authorization,Accept,Origin,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Content-Range,Range")
	(*res).Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE,PATCH")
	(*res).Header().Add("Content-Type", "application/json, charset=utf-8")
	(*res).Header().Add("Access-Control-Allow-Credentials", "true")
	(*res).Header().Add("Access-Control-Max-Age", "3600")
}




func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to trading-old-item server")
}

func decodeJson(request *http.Request, v interface{}) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		panic(err)
	}
}

func encodeJson(v interface{}, jData *[]byte) {
	*jData, _ = json.Marshal(v)
}

func Login(writer http.ResponseWriter, request *http.Request) {

	//handle preflight
	addCorsHeader(&writer)
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusOK)
		return
	}

	var logReq _type.LoginRequest
	decodeJson(request, &logReq)

	statusCode, detail, result := repo.Login(logReq.Email, logReq.Password)

	if statusCode != 200 {
		logRes := _type.LoginResponse{StatusCode: statusCode, Detail: detail}
		var jData []byte
		encodeJson(logRes, &jData)

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

func Signup(writer http.ResponseWriter, request *http.Request) {

	//handle preflight
	addCorsHeader(&writer)
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

func AddPost(writer http.ResponseWriter, request *http.Request) {

	//handle preflight
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Headers", "Authorization,Accept,Origin,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Content-Range,Range")
	writer.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE,PATCH")
	writer.Header().Add("Content-Type", "application/json;charset=UTF-8")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Max-Age", "3600")

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
	writer.WriteHeader(http.StatusOK)
	writer.Write(jData)
}

func EditPost(writer http.ResponseWriter, request *http.Request)  {
	addCorsHeader(&writer)
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusOK)
		return
	}

	var req _type.EditPostRequest
	decodeJson(request, &req)
	statusCode, detail := repo.EditPost(req.ID, req.OwnerID, req.Name, req.Brand, req.Type, req.Amount, req.Description, req.ImageURL)
	res := _type.EditPostResponse{ StatusCode: statusCode, Detail: detail }

	var jData []byte
	encodeJson(res, &jData)
	writer.WriteHeader(http.StatusOK)
	writer.Write(jData)
}

func DeactivatePost(writer http.ResponseWriter, request *http.Request)  {

	addCorsHeader(&writer)
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusOK)
		return
	}

	var req _type.DeactivatePostRequest
	decodeJson(request, &req)
	statusCode, detail := repo.DeactivatePost(req.PostID)
	res := _type.DeactivatePostRes{ statusCode, detail }

	var jData []byte
	encodeJson(res, &jData)
	writer.WriteHeader(http.StatusOK)
	writer.Write(jData)
}

func GetAllPost(writer http.ResponseWriter, request *http.Request) {

	addCorsHeader(&writer)
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

func GetAllUser(writer http.ResponseWriter, request *http.Request)  {

	addCorsHeader(&writer)
	statusCode, detail, result := repo.GetAllUser()

	if statusCode != 200 {
		res := _type.GetAllUserRes{StatusCode: statusCode, Detail: detail}
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

func GetUserByID(writer http.ResponseWriter, request *http.Request) {
	addCorsHeader(&writer)
	params := mux.Vars(request)
	user_id, _ := strconv.Atoi(params["user_id"])
	statusCode, detail, result := repo.GetUserByID(user_id)

	if statusCode != 200 {
		res := _type.GetUserByIDRes{StatusCode: statusCode, Detail: detail}
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
	addCorsHeader(&writer)
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

	addCorsHeader(&writer)
	params := mux.Vars(request)
	owner_id, _ := strconv.Atoi(params["owner_id"])
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

func CreateTransaction(writer http.ResponseWriter, request *http.Request) {

	addCorsHeader(&writer)
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusOK)
		return
	}

	var createTransactionReq _type.CreateTransactionRequest
	decodeJson(request, &createTransactionReq)

	statusCode, detail := repo.CreateTransaction(createTransactionReq.FromPostID, createTransactionReq.ToPostID)

	res := &_type.CreateTransactionRes{
		statusCode,detail,
	}

	var jData []byte
	encodeJson(res, &jData)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jData)

}

func Get_B_BY_A(writer http.ResponseWriter, request *http.Request)  {
	addCorsHeader(&writer)
	params := mux.Vars(request)
	a_ID, _ := strconv.Atoi(params["product_id"])
	statusCode, detail, result := repo.Get_B_By_A(a_ID)

	if statusCode != 200 {
		res := _type.Get_B_BY_A_Res{StatusCode: statusCode, Detail: detail}
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

func GetUserHistory(writer http.ResponseWriter, request *http.Request)  {
	addCorsHeader(&writer)
	params := mux.Vars(request)
	userID, _ := strconv.Atoi(params["user_id"])
	statusCode, detail, result := repo.GetUserHistory(userID)

	if statusCode != 200 {
		res := _type.GetHistoryRes{StatusCode: statusCode, Detail: detail}
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

func CompleteTransaction(writer http.ResponseWriter, request *http.Request)  {

	addCorsHeader(&writer)
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusOK)
		return
	}

	var req _type.CompleteTransRequest
	decodeJson(request, &req)

	statusCode, detail := repo.CompleteTransaction(req.TransactionID)

	res := &_type.CompleteTransRes{
		statusCode,detail,
	}

	var jData []byte
	encodeJson(res, &jData)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jData)
}

func SearchByKeyword(writer http.ResponseWriter, request *http.Request)  {
	addCorsHeader(&writer)
	params := mux.Vars(request)
	keyword, _ := params["keyword"]
	statusCode, detail, result := repo.SearchHardCore(keyword)

	if statusCode != 200 {
		res := _type.SearchRes{StatusCode: statusCode, Detail: detail}
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
