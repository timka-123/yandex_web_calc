package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func PanicMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Printf("New request to %s\n", r.URL)
		defer func() {
			if err := recover(); err != nil {
				log.Printf("New error at %s: %s\n", r.URL, err)
				w.WriteHeader(http.StatusInternalServerError)
				response := CalculatorFailureResponse{Error: "Internal server error, sorry"}
				json.NewEncoder(w).Encode(response)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func CalculatorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestData CalculatorRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestData)

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
		r.PostForm.Set("expression", requestData.Expression)

		next.ServeHTTP(w, r)
	})
}
