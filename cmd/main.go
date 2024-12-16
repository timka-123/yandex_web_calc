package main

import (
	"net/http"
	"yandex_web_calc/internal"
)

func main() {
	http.HandleFunc("/api/v1/calculate", internal.PanicMiddleware(internal.CalculatorMiddleware(internal.CalculatorHandler)))
	http.ListenAndServe(":8080", nil)
}
