package application

import (
	"encoding/json"
	"net/http"
	"github.com/dltzk/go-calc/pkg/calculation"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/gorilla/mux"
	"fmt"
)

type Expression struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result"`
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {

	var expr Expression

	if r.Method != "POST" {
		w.WriteHeader(http.StatusInternalServerError)
		response := calculation.Error{Error: calculation.InternalServerError}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&expr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := calculation.Error{Error: calculation.InternalServerError}
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := calculation.Calc(expr.Expression)
	if err != nil {
		if err == calculation.ErrDivisionByZero {
			w.WriteHeader(http.StatusUnprocessableEntity)
			response := calculation.Error{Error: calculation.DivisionByZero}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		response := calculation.Error{Error: calculation.ExpressionIsNotValid}
		json.NewEncoder(w).Encode(response)
		return
	}

	encoder := json.NewEncoder(w)

	err = encoder.Encode(Response{
		Result: result,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := calculation.Error{Error: calculation.InternalServerError}
		json.NewEncoder(w).Encode(response)
		return
	}

}

func setupLogger()  *zap.Logger {
	// Настраиваем конфигурацию логгера
	config := zap.NewProductionConfig()

	// Уровень логирования
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	
	// Настраиваем логгер с конфигурацией
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Ошибка настройки логгера: %v\n", err)
	}
	
		return logger
}

func loggingMiddleware(logger *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("New request",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
			)
			next.ServeHTTP(w, r)
		})
	}
}

func RunServer() error {
	r := mux.NewRouter()
	logger := setupLogger()
	r.Use(loggingMiddleware(logger))

	r.HandleFunc("/api/v1/calculate", CalculatorHandler)
	http.Handle("/api/v1/calculate", r)

	logger.Info("Server started", zap.Int("port", 8080))
	return http.ListenAndServe("localhost:8080", nil)
}
