package calculation

import "errors"

type Error struct {
	Error string`json:"error"`
}

var (
	ExpressionIsNotValid = "Expression is not valid"
	InternalServerError = "Internal server error"
	ErrInvalidExpression = errors.New("Invalid expression")
)