package main

import (
	"fmt"
)

func main() {
	// Запитуємо користувача ввести кількість елементів
	var n int
	fmt.Print("Введіть кількість елементів: ")
	fmt.Scan(&n)

	// Створюємо зріз для зберігання елементів
	arr := make([]int, n)
	fmt.Println("Введіть елементи:")

	// Зчитуємо елементи від користувача
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Створюємо карту для підрахунку повторень
	counts := make(map[int]int)

	// Перебираємо зріз і рахуємо частоту кожного елемента
	for _, num := range arr {
		counts[num]++
	}

	// Виводимо повторювані елементи, якщо вони є
	fmt.Println("Повторювані елементи:")
	found := false
	for num, count := range counts {
		if count > 1 {
			fmt.Println(num)
			found = true
		}
	}
	if !found {
		fmt.Println("Повторюваних елементів немає")
	}
}
