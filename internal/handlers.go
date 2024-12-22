package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"yandex_web_calc/pkg"
)

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	var requestData CalculatorRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		log.Printf("Handled invalid request to %s. Maybe data was mailformed\n", r.URL)
		w.WriteHeader(http.StatusUnprocessableEntity)
		response := CalculatorFailureResponse{Error: "Invalid request"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if requestData.Expression == "" {
		log.Printf("Handled invalid request to %s. Expression %s was bad\n", r.URL, requestData.Expression)
		response := CalculatorFailureResponse{Error: "Expression is required"}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := pkg.Calc(requestData.Expression)
	if err != nil {
		log.Printf("Handled invalid request to %s. Expression %s was bad\n", r.URL, requestData.Expression)
		response := CalculatorFailureResponse{Error: err.Error()}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := CalculatorSuccessResponse{Result: result}
	json.NewEncoder(w).Encode(response)
}
