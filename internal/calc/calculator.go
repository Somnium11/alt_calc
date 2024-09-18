package calculator

import (
	"errors"
	"strconv"
	"strings"
)

func Calculate(input string) (string, error) {
	operators := []string{"+", "-", "*", "/"}

	var operator string
	var pos int
	for _, op := range operators {
		pos = strings.Index(input, " "+op+" ")
		if pos != -1 {
			operator = op
			break
		}
	}

	if operator == "" {
		return "", errors.New("неверный формат ввода, оператор не найден")
	}

	str1 := input[:pos]
	str2 := input[pos+len(operator)+2:] 

	str1 = strings.TrimSpace(str1)
	str2 = strings.TrimSpace(str2)

	switch operator {
	case "+":
		return str1 + str2, nil
	case "-":
		return strings.ReplaceAll(str1, str2, ""), nil
	case "*":
		n, err := strconv.Atoi(str2)
		if err != nil || n < 0 {
			return "", errors.New("второй операнд должен быть положительным целым числом для операции '*'")
		}
		return strings.Repeat(str1, n), nil
	case "/":
		n, err := strconv.Atoi(str2)
		if err != nil || n < 0 {
			return "", errors.New("второй операнд должен быть положительным целым числом для операции '/'")
		}
		if n > len(str1) {
			n = len(str1)
		}
		return str1[:n], nil
	default:
		return "", errors.New("неподдерживаемый оператор: " + operator)
	}
}