package calculator

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
   )
   
   func Calc(expression string) (float64, error) {
	tokens, err := tokenize(expression)
	if err != nil {
	 return 0, err
	}
	rpn, err := toRPN(tokens)
	if err != nil {
	 return 0, err
	}
	return evaluateRPN(rpn)
   }
   
   func tokenize(expression string) ([]string, error) {
	var tokens []string
	var current strings.Builder
   
	for _, ch := range expression {
	 switch {
	 case unicode.IsDigit(ch) || ch == '.':
	  current.WriteRune(ch)
	 case ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '(' || ch == ')':
	  if current.Len() > 0 {
	   tokens = append(tokens, current.String())
	   current.Reset()
	  }
	  tokens = append(tokens, string(ch))
	 case unicode.IsSpace(ch):
	  if current.Len() > 0 {
	   tokens = append(tokens, current.String())
	   current.Reset()
	  }
	 default:
	  return nil, errors.New("invalid character in expression")
	 }
	}
   
	if current.Len() > 0 {
	 tokens = append(tokens, current.String())
	}
   
	return tokens, nil
   }
   
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
	default:
	 return 0, errors.New("unknown operator")
	}
 
	stack = append(stack, result)
   }
  }

  if len(stack) != 1 {
	return 0, errors.New("invalid expression")
   }
  
   return stack[0], nil
  }
  
  // isNumber проверяет, является ли строка числом.
  func isNumber(token string) bool {
   _, err := strconv.ParseFloat(token, 64)
   return err == nil
  }
  
  
   