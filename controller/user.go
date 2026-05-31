package controller

import (
	"encoding/json"
	"myweb/model"
	httpResp "myweb/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddUser(w http.ResponseWriter,r *http.Request){
	var user model.User

	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&user)//read json data and write into user struct
	if err!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,err.Error())
		return
	}
	r.Body.Close()
	if CreateErr:=user.Create();CreateErr!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,CreateErr.Error())
		return
	}
	httpResp.RespondwithJSON(w,http.StatusOK,map[string]string{"status":"User Added"})


}

func GetUser(w http.ResponseWriter,r *http.Request){
	sid:=mux.Vars(r)["sid"]
	new_id,err:=getUserId(sid)
	if err!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,err.Error())
		return
	}
	user:=model.User{UserId: int(new_id)}

	if RErr:=user.Read();RErr!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,RErr.Error())
		return
	}
	httpResp.RespondwithJSON(w,http.StatusOK,user)

}

func getUserId(old_id string)(int64,error){
	new_id,err:=strconv.Atoi(old_id)
	return int64(new_id),err
}
