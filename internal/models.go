package internal

type CalculatorRequest struct {
	Expression string `json:"expression"`
}

type CalculatorSuccessResponse struct {
	Result float64 `json:"result"`
}

type CalculatorFailureResponse struct {
	Error string `json:"error"`
}
