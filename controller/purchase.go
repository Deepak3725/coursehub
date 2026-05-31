package controller

import (
	"encoding/json"
	"myweb/model"
	httpResp "myweb/response"
	"net/http"

	"github.com/gorilla/mux"
)

func AddPurchase(w http.ResponseWriter, r *http.Request) {
	var purchase model.Purchase

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&purchase); err != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	if CrErr := purchase.Create(); CrErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, CrErr.Error())
		return
	}
	httpResp.RespondwithJSON(w, http.StatusOK, map[string]string{"Status": "Purchase Successful"})

}

func ReadAllPurchase(w http.ResponseWriter, r *http.Request) {
	var user_email = mux.Vars(r)["sid"]
	var purchase model.PurchasedCourse

	purchases, err := purchase.ReadPurchasedAll(user_email)

	if err != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, err.Error())
		return
	}

	httpResp.RespondwithJSON(w, http.StatusOK, purchases)

}
