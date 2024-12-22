package main

import (
	"log"
	"net/http"
	"yandex_web_calc/internal"
)

func main() {
	http.HandleFunc("/api/v1/calculate", internal.PanicMiddleware(internal.PostMiddleware(internal.CalculatorHandler)))
	log.Println("Starting server on 8080 port")
	http.ListenAndServe(":8080", nil)
}
