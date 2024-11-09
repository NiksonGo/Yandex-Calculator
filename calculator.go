package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Calc вычисляет арифметическое выражение и возвращает результат или ошибку, если выражение некорректно
func Calc(expression string) (float64, error) {
	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	rpn, err := toRPN(tokens)
	if err != nil {
		return 0, err
	}

	result, err := evaluateRPN(rpn)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// tokenize разбивает строку на токены
func tokenize(expression string) ([]string, error) {
	var tokens []string
	var numStr strings.Builder

	for _, ch := range expression {
		if unicode.IsDigit(ch) || ch == '.' {
			numStr.WriteRune(ch)
		} else {
			if numStr.Len() > 0 {
				tokens = append(tokens, numStr.String())
				numStr.Reset()
			}
			if ch == ' ' {
				continue
			} else if strings.ContainsRune("+-*/()", ch) {
				tokens = append(tokens, string(ch))
			} else {
				return nil, errors.New("invalid character in expression")
			}
		}
	}

	if numStr.Len() > 0 {
		tokens = append(tokens, numStr.String())
	}

	return tokens, nil
}

// toRPN преобразует токены в обратную польскую нотацию
func toRPN(tokens []string) ([]string, error) {
	var output []string
	var operators []string

	precedence := map[string]int{
		"+": 1, "-": 1,
		"*": 2, "/": 2,
	}

	for _, token := range tokens {
		switch {
		case isNumber(token):
			output = append(output, token)
		case token == "(":
			operators = append(operators, token)
		case token == ")":
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			operators = operators[:len(operators)-1]
		default:
			for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[token] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, errors.New("mismatched parentheses")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

// evaluateRPN вычисляет значение выражения, записанного в обратной польской нотации
func evaluateRPN(rpn []string) (float64, error) {
	var stack []float64

	for _, token := range rpn {
		if isNumber(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				result = a / b
			}

			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}

// isNumber проверяет, является ли строка числом
func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func main() {
	expression := "3 + 5 * ( 2 - 8 )"
	result, err := Calc(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
