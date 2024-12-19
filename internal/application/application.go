package application

import (
	"encoding/json"
	"net/http"
	"github.com/dltzk/go-calc/pkg/calculation"
)

type Expression struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result"`
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {

	var expr Expression

	if err := json.NewDecoder(r.Body).Decode(&expr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := calculation.Calc(expr.Expression)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(Response{
		Result: result,
	}); err != nil {
		panic(err)
	}
}

func RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalculatorHandler)
	return http.ListenAndServe("localhost:8080", nil)
}
