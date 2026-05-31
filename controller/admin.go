package controller

import (
	"database/sql"
	"encoding/json"
	"myweb/model"
	httpResp "myweb/response"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&admin)
	if err != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	sErr := admin.SignUp()

	if sErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, sErr.Error())
		return
	}
	httpResp.RespondwithJSON(w, http.StatusCreated, map[string]string{"Status": "Admin Created"})
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin

	Decoder := json.NewDecoder(r.Body)
	err := Decoder.Decode(&admin)

	if err != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()

	LErr := admin.LogIn()
	if LErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, LErr.Error())
		return
	}
	//session Management
	token_val := uuid.New().String()
	cookie := http.Cookie{
		Name:     "coursehub-cookie",
		Value:    token_val,
		HttpOnly: true,
		Secure:   true,
	}
	session := model.Session{
		Token:     cookie.Value,
		User_mail: admin.Email,
	}
	SessErr := session.Createsession()
	if SessErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, SessErr.Error())
		return
	}

	http.SetCookie(w, &cookie)
	httpResp.RespondwithJSON(w, http.StatusOK, map[string]string{"Status": "Logged-In"})

}

func LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("coursehub-cookie")
	if err!=nil{
		switch err{
		case http.ErrNoCookie:
		    httpResp.RespondwithError(w,http.StatusSeeOther,err.Error())
		    return
		default:
			httpResp.RespondwithError(w,http.StatusInternalServerError,"Internal Server Error")
			return
		}
	}
	http.SetCookie(w, &http.Cookie{
		Value:   cookie.Value,
		Expires: time.Now(),
	})
	session:=model.Session{
		Token: cookie.Value,
	}
	Delerr:=session.DeleteSession()

	if Delerr != nil {
		switch Delerr {
		case sql.ErrNoRows:
			httpResp.RespondwithError(w, http.StatusNotFound, Delerr.Error())
			return
		default:
			httpResp.RespondwithError(w, http.StatusBadRequest, Delerr.Error())
			return
		}
	}
	httpResp.RespondwithJSON(w,http.StatusOK,map[string]string{"Logged-Out":"Success"})
}


func VerifyCookie(w http.ResponseWriter,r *http.Request)bool{
	cookie,err:=r.Cookie("coursehub-cookie")
	if err!=nil{
		switch err{
		case http.ErrNoCookie:
			httpResp.RespondwithError(w,http.StatusSeeOther,err.Error())
		    return false
		default:
			httpResp.RespondwithError(w,http.StatusInternalServerError,"Internal Server Error")
			return false
		}
	}
	session:=model.Session{
		Token: cookie.Value,
	}
	Gerr:=session.GetSession()
	if Gerr!=nil{
		switch Gerr{
		case sql.ErrNoRows:
			httpResp.RespondwithError(w,http.StatusNotFound,Gerr.Error())
			return false
		default:
			httpResp.RespondwithError(w,http.StatusBadRequest,Gerr.Error())
			return false
		}
	}

	if cookie.Value!=session.Token{
		httpResp.RespondwithError(w,http.StatusUnauthorized,"Cookies Does not match")
		return false
	}
	return true

}

//We are getting session...
func GetFromCookie(w http.ResponseWriter,r *http.Request){
	cookie,err:=r.Cookie("coursehub-cookie")//Get cookie 
	if err!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,err.Error())
		return
	}
	Session:=model.Session{Token: cookie.Value}//send cookie value

	if Serr:=Session.GetSession();Serr!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,Serr.Error())
		return
	}
	httpResp.RespondwithJSON(w,http.StatusOK,map[string]string{"Email":Session.User_mail})
}