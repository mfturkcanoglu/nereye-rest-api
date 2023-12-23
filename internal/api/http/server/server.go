package server

import (
	"log"
	"net/http"
)

func Init(){
	router := Setup()
	log.Fatalln(http.ListenAndServe(":3000", router))
}