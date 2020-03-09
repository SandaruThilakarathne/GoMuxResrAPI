package main

import (
	auth "./auth"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	con "./connections"
)

func main()  {
	//auth.Test()
	con.CreateClient()
	r := mux.NewRouter()
	r.HandleFunc("/test", auth.Test).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
