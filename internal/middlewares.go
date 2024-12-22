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

func PostMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			log.Printf("Handled invalid request to %s. Method %s is not allowed\n", r.URL, r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			response := CalculatorFailureResponse{Error: "Method not allowed"}
			json.NewEncoder(w).Encode(response)
			return
		}
		next.ServeHTTP(w, r)
	})
}
