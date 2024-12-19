package calculation

import (
	"strconv"
)

func Factor(expression string, i *int, l *int, answer *[]string, numbers map[string]struct{}) bool {
	if *i >= *l {
		return false
	}
	_, ok := numbers[string(expression[*i])]
	if string(expression[*i]) == "(" {
		*i++
		if !Expr(expression, i, l, answer, numbers) {
			return false
		}
		if *i >= *l {
			return false
		} else {
			if string(expression[*i]) == ")" {
				*i++
				return true
			} else {return false}
		}
	} else if ok {
		*answer = append(*answer, string(expression[*i]))
		*i++
		return true
	} else {
		return false
	}
}

func TList(expression string, i *int, l *int, answer *[]string, numbers map[string]struct{}) bool {
	if *i >= *l {
		return true
	}
	if string(expression[*i]) == "*" || string(expression[*i]) == "/" {
		symbol := string(expression[*i])
		*i++
		if !Factor(expression, i, l, answer, numbers) {
			return false
		}
		*answer = append(*answer, symbol)
		if !TList(expression, i, l, answer, numbers) {
			return false
		}
		return true
	}
	return true
}

func EList(expression string, i *int, l *int, answer *[]string, numbers map[string]struct{}) bool {
	if *i >= *l {
		return true
	}
	if string(expression[*i]) == "+" || string(expression[*i]) == "-" {
		symbol := string(expression[*i])
		*i++
		if !Term(expression, i, l, answer, numbers) {
			return false
		}
		*answer = append(*answer, symbol)
		if !EList(expression, i, l, answer, numbers) {
			return false
		}
	}
	return true
}

func Term(expression string, i *int, l *int, answer *[]string, numbers map[string]struct{}) bool {
	if !Factor(expression, i, l, answer, numbers) {
		return false
	}
	if !TList(expression, i, l, answer, numbers) {
		return false
	}
	return true
}

func Expr(expression string, i *int, l *int, answer *[]string, numbers map[string]struct{}) bool {
	if !Term(expression, i, l, answer, numbers) {
		return false
	}
	if !EList(expression, i, l, answer, numbers) {
		return false
	}
	return true
}

func Calc(expression string) (float64, error) {
	answer := []string{}
	numbers := map[string]struct{}{"1": {}, "2": {}, "3": {}, "4": {}, "5": {}, "6": {}, "7": {}, "8": {}, "9": {}, "0": {}}
	var i int = 0
	var l int = len(expression)
	result := Expr(expression, &i, &l, &answer, numbers)

	stack := []float64{}

	if result && i == l {

		for j := range answer {
			_, ok := numbers[string(answer[j])]
			if ok {
				value, _ := strconv.ParseFloat(string(answer[j]), 64)
				stack = append(stack, value)
			} else {
				var first float64 = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				var second float64 = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				switch string(answer[j]) {
				case "+":
					stack = append(stack, first+second)
				case "-":
					stack = append(stack, second-first)
				case "*":
					stack = append(stack, first*second)
				case "/":
					if first == 0 {
						return 0, ErrDivisionByZero
					}
					stack = append(stack, second/first)
				}
			}
		}
		return stack[0], nil

	}
	return 0, ErrInvalidExpression
}
