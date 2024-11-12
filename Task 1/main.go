package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введіть елементи масиву через пробіл:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Розділення рядка на окремі елементи
	elements := strings.Fields(input)
	counter := make(map[string]int)

	// Підрахунок кількості кожного елемента
	for _, element := range elements {
		counter[element]++
	}

	// Виведення результатів
	fmt.Println("Кількість повторень кожного елемента:")
	for element, count := range counter {
		fmt.Printf("%s: %d разів\n", element, count)
	}
}
