package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"trading.olditem.app/core/type"
)

func Login(writer http.ResponseWriter, request *http.Request) {

	var logReq _type.LoginRequest
	body, err := ioutil.ReadAll(request.Body)
	err = json.Unmarshal(	body, &logReq)
	fmt.Println(logReq)
	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	logRes := _type.LoginResponse{StatusCode: 200, Detail: "Nhan duoc roi"}
	jData, _ := json.Marshal(logRes)
	writer.Write(jData)
}
