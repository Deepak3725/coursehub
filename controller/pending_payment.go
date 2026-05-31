package controller

import (
	"database/sql"
	"encoding/json"
	"myweb/model"
	httpResp "myweb/response"
	"myweb/utils/helper"
	"net/http"

	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func InitiatePayment(w http.ResponseWriter,r *http.Request){
	var payment model.Payment
	decoder:=json.NewDecoder(r.Body)
	if err:=decoder.Decode(&payment);err!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,err.Error())
		return
	}

	course:=model.Course{CourseId: payment.CourseId} 

	//we are getting that course matching the id to extract creator's mail...
	if CourseErr:=course.GetCoursebyId();CourseErr!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,CourseErr.Error())
		return
	}


	token:=uuid.New().String() //generates unique string using uuid github package
	payment.Token=token
	if Perr:=payment.CreatePay();Perr!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,Perr.Error())
		return
	}

	//Now we set up sending maiL:
	const baseUrl = "https://coursehub-10gd.onrender.com"//address where my web is running
	yes_link:=fmt.Sprintf("%s/confirmpayment?token=%s&action=yes",baseUrl,payment.Token)//it would be sth like basurl/confirmpayment?token=xxx&action yes...
	no_link:=fmt.Sprintf("%s/confirmpayment?token=%s&action=no",baseUrl,payment.Token)

	body := fmt.Sprintf(
    "A payment for your course has been made.\n\nJRNL Number: %s\n\nClick this link if payment is Valid: %s\n\nor\n\nClick this link if payment is NOT valid: %s",
    payment.Reference_Num, yes_link, no_link,
)
    go func(){
    if Mail_err := helper.Send_Mail(*course.Creator_mail, "Payment Verification", body); Mail_err != nil {
        fmt.Println("Email error:", Mail_err.Error())
        return
    }
}()
	


	httpResp.RespondwithJSON(w,http.StatusAccepted,map[string]string{"Payment Status":"Success"})
}
func ConfirmPayment(w http.ResponseWriter,r *http.Request){
	//get what user clicks..
	token:=r.URL.Query().Get("token")//from request body url return value of token..
	action:=r.URL.Query().Get("action")
//r.Url==>full url and .Query()extract query parameters ?after this and Get will get values matching keys..
	payment:=model.Payment{Token: token}
	if err:=payment.ReadPayment();err!=nil{
		switch err{
		case sql.ErrNoRows:
			httpResp.RespondwithError(w,http.StatusNotFound,err.Error())
			return
		default:
			httpResp.RespondwithError(w,http.StatusBadRequest,err.Error())
			return
		}
	}
    var msgbody string
	if action=="yes"{
		purchase:=model.Purchase{
			Course_id: payment.CourseId,
			User_email: payment.Buyer_mail,
		}

		if PurErr:=purchase.Create();PurErr!=nil{
			httpResp.RespondwithError(w,http.StatusBadRequest,PurErr.Error())
			return
		}
		msgbody=fmt.Sprintf("Your Payment with JRNL Num:%s was Successfull item added to your cart please kindly check",payment.Reference_Num)
	}else{
		msgbody=fmt.Sprintf("Your Payment with JRNL Num:%s was not Successful please make Payment again",payment.Reference_Num)
	}
	
	if MailErr:=helper.Send_Mail(payment.Buyer_mail,"Payment Result",msgbody);MailErr!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,MailErr.Error())
		return
	}

	httpResp.RespondwithJSON(w,http.StatusOK,map[string]string{"Payment":"Success"})

}
func ReadPayment(w http.ResponseWriter,r *http.Request){
	sid:=mux.Vars(r)["sid"]

	payment:=model.Payment{Token: sid}

	if err:=payment.ReadPayment();err!=nil{
		switch err{
		case sql.ErrNoRows:
				httpResp.RespondwithError(w,http.StatusNotFound,err.Error())
				return
		default:
			httpResp.RespondwithError(w,http.StatusBadRequest,err.Error())
			return
		}
	}
	httpResp.RespondwithJSON(w,http.StatusOK,payment)
}

