package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"alternative_calc/internal/calc"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите операцию :")

	for {
		fmt.Print("> ") // Приглашение к вводу
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		result, err := calculator.Calculate(input)
		if err != nil {
			fmt.Println("Ошибка:", err) // Вывод ошибки
		} else {
			fmt.Println("Результат:", result) // Вывод результата
		}
	}
}
