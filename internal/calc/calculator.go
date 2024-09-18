package calculator

import (
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
		panic("неверный формат ввода, оператор не найден")
	}

	str1 := input[:pos]
	str2 := input[pos+len(operator)+2:]

	str1 = strings.TrimSpace(str1)
	str2 = strings.TrimSpace(str2)

	if !(strings.HasPrefix(str1, "\"") && strings.HasSuffix(str1, "\"")) {
		panic("первый операнд должен быть строкой в кавычках")
	}

	str1 = str1[1 : len(str1)-1]

	switch operator {
	case "+":
		if !(strings.HasPrefix(str2, "\"") && strings.HasSuffix(str2, "\"")) {
			panic("второй операнд для сложения должен быть строкой в кавычках")
		}
		str2 = str2[1 : len(str2)-1]
		result := str1 + str2
		return limitString(result), nil
	case "-":
		if !(strings.HasPrefix(str2, "\"") && strings.HasSuffix(str2, "\"")) {
			panic("второй операнд для вычитания должен быть строкой в кавычках")
		}
		str2 = str2[1 : len(str2)-1]
		result := strings.ReplaceAll(str1, str2, "")
		return limitString(result), nil
	case "*":
		n, err := strconv.Atoi(str2)
		if err != nil || n < 1 {
			panic("второй операнд должен быть положительным целым числом для операции '*'")
		}
		result := strings.Repeat(str1, n)
		return limitString(result), nil
	case "/":
		n, err := strconv.Atoi(str2)
		if err != nil || n < 1 {
			panic("второй операнд должен быть положительным целым числом для операции '/'")
		}
		if n > len(str1) {
			n = len(str1)
		}
		result := str1[:n]
		return limitString(result), nil
	default:
		// Если оператор не поддерживается, возвращаем ошибку
		panic("неподдерживаемый оператор: " + operator)
	}
}

func limitString(str string) string {
	if len(str) > 40 {
		return str[:40] + "..."
	}
	return str
}
