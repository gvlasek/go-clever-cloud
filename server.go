package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Model struct {
	Name string
}

func market(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	template, _ := template.ParseFiles("market.gtpl")
	template.Execute(w, Model{})
}

func api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response, _ := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=10")
	io.Copy(w, response.Body)
}

func main() {
	log.Print("Starting server...")
	router := httprouter.New()
	router.GET("/market", market)
	router.GET("/api", api)
	router.GET("/", market)
	log.Fatal(http.ListenAndServe(":8080", router))
}
