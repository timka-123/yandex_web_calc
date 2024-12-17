package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"yandex_web_calc/internal"
)

func TestInvalidExpressionCalculatorHandler(t *testing.T) {
	expressions := []string{"2+2)", "fhewnycrkuve44tne", "43567i2356875723", "2+2+"}

	for _, expression := range expressions {
		requestData := internal.CalculatorRequest{Expression: expression}
		body, _ := json.Marshal(requestData)
		req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(internal.CalculatorHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
		}

		var response internal.CalculatorFailureResponse
		json.NewDecoder(rr.Body).Decode(&response)
		if response.Error == "" {
			t.Errorf("Expected an error message but got none")
		}
	}
}

func TestValidExpressionCalculatorHandler(t *testing.T) {
	requestData := internal.CalculatorRequest{Expression: "2+2"}
	body, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(internal.CalculatorHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response internal.CalculatorSuccessResponse
	json.NewDecoder(rr.Body).Decode(&response)
	if response.Result != 4 {
		t.Errorf("Expected result 4 but got %v", response.Result)
	}
}
