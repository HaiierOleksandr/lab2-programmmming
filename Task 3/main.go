package main

import (
	"fmt"
	"strconv"
)

func reverseNumber(n int) (int, error) {
	// Перетворимо число в рядок
	strNum := strconv.Itoa(n)
	reversedStr := ""

	// Проходимо по числу в зворотному порядку
	for i := len(strNum) - 1; i >= 0; i-- {
		reversedStr += string(strNum[i])
	}

	// Перетворимо перевернутий рядок назад у число
	reversedNum, err := strconv.Atoi(reversedStr)
	if err != nil {
		return 0, err
	}

	return reversedNum, nil
}

func main() {
	var number int
	fmt.Print("Введіть число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Помилка вводу:", err)
		return
	}

	reversed, err := reverseNumber(number)
	if err != nil {
		fmt.Println("Помилка обробки:", err)
		return
	}

	fmt.Printf("Перевернуте число: %d\n", reversed)
}
