package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "fbservice"
)

func main() {
    sm:=mux.NewRouter()
	
	getsubRouter:=sm.Methods(http.MethodGet).Subrouter()
	getsubRouter.HandleFunc("/name",fbservice.GetUserName)

    
	getsubRouter.HandleFunc("/email",fbservice.GetUserEmail)
	getsubRouter.HandleFunc("/mydetails",fbservice.GetUserDeatils)

    putsubRouter:=sm.Methods(http.MethodPost).Subrouter()
    putsubRouter.HandleFunc("/postimage",fbservice.PostPicture)


    http.ListenAndServe(":8080", sm)
}