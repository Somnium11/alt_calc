package main

import (
	calculator "alternative_calc/internal/calc"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите операцию :")

	for {
		fmt.Print("> ") // Приглашение к вводу
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		result := calculator.Calculate(input)
		fmt.Println("Результат:", result) // Вывод результата
	}
}
