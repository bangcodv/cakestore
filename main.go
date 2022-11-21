package main

import (
	h "cakestore/delivery"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", h.HandlerIndex)
	router.HandleFunc("/cakes", h.HandlerListCakes).Methods("GET")
	router.HandleFunc("/cakes", h.HandlerAddCake).Methods("POST")
	router.HandleFunc("/cakes/{id}", h.HandlerDetailCakes).Methods("GET")
	router.HandleFunc("/cakes/{id}", h.HandlerUpdateCake).Methods("PUT")
	router.HandleFunc("/cakes/{id}", h.HandlerDeleteCake).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}
