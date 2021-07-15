package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"trading.olditem.app/core/repo"
	"trading.olditem.app/core/type"
)

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
