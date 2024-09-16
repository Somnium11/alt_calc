package calculator

import (
	"errors"
	"strconv"
	"strings"
)

// Calculate выполняет операции со строками.
func Calculate(input string) (string, error) {
	// Разбиваем строку на операнды и оператор
	tokens := strings.Fields(input)
	if len(tokens) != 3 {
		return "", errors.New("неверный формат ввода, ожидается: 'строка оператор строка или число'")
	}

	// Операнды
	str1 := tokens[0]
	operator := tokens[1]
	str2 := tokens[2]

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
			n = len(str1) // Чтобы избежать выхода за пределы строки
		}
		return str1[:n], nil
	default:
		return "", errors.New("неподдерживаемый оператор: " + operator)
	}
}
