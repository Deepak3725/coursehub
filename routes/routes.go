package routes

import (
	"fmt"
	"log"
	"myweb/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/adduser", controller.AddUser).Methods("POST")
	router.HandleFunc("/addcourse", controller.AddCourse).Methods("POST")
	router.HandleFunc("/getsellincourse/{sid}", controller.GetSellingCourse).Methods("GET")
	router.HandleFunc("/updatecourse/{sid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/deletecourse/{sid}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/addpurchase", controller.AddPurchase).Methods("POST")
	router.HandleFunc("/getuser/{sid}", controller.GetUser).Methods("GET")
	router.HandleFunc("/getallcourses", controller.GetAllcourse).Methods("GET")
	router.HandleFunc("/login", controller.LogIn).Methods("POST")
	router.HandleFunc("/signup", controller.SignUp).Methods("POST")
	router.HandleFunc("/logout", controller.LogOut)
	router.HandleFunc("/readAllPurchases/{sid}",controller.ReadAllPurchase).Methods("GET")
	router.HandleFunc("/getsession",controller.GetFromCookie).Methods("GET")
	router.HandleFunc("/initiatepayment",controller.InitiatePayment).Methods("POST")
	router.HandleFunc("/confirmpayment",controller.ConfirmPayment).Methods("GET")
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)
	log.Println("Application running on port " + ":8080")
	err := http.ListenAndServe(":8080", router) // port ,use the mux of router if nil uses default mux
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
