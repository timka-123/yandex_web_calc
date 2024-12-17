package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
		r.Header.Set("Content-Type", "application/json")
		var requestData CalculatorRequest

		// I wanted to make this code more beautiful but my beautiful version does not working:(
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		err := json.Unmarshal(bodyBytes, &requestData)

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
		next.ServeHTTP(w, r)
	})
}
