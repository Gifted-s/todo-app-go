package main

import (
	"net/http"
	"./router"
)


func main(){ 
	router.Router()
	http.ListenAndServe(":3000", router.Router())
}