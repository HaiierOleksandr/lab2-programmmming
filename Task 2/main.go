package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var todoList []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Менеджер списку справ ===")
		fmt.Println("1. Додати завдання")
		fmt.Println("2. Видалити завдання")
		fmt.Println("3. Відобразити всі завдання")
		fmt.Println("4. Вийти")
		fmt.Print("Оберіть опцію: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			addTask(scanner)
		case "2":
			deleteTask(scanner)
		case "3":
			displayTasks()
		case "4":
			fmt.Println("Вихід...")
			return
		default:
			fmt.Println("Невірна опція. Спробуйте ще раз.")
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Введіть завдання: ")
	scanner.Scan()
	task := scanner.Text()
	todoList = append(todoList, task)
	fmt.Println("Завдання додано.")
}

func deleteTask(scanner *bufio.Scanner) {
	if len(todoList) == 0 {
		fmt.Println("Немає завдань для видалення.")
		return
	}

	fmt.Print("Введіть номер завдання для видалення: ")
	scanner.Scan()
	indexStr := scanner.Text()
	index, err := strconv.Atoi(indexStr)

	if err != nil || index <= 0 || index > len(todoList) {
		fmt.Println("Невірний номер завдання.")
		return
	}

	todoList = append(todoList[:index-1], todoList[index:]...)
	fmt.Println("Завдання видалено.")
}

func displayTasks() {
	if len(todoList) == 0 {
		fmt.Println("Список завдань порожній.")
		return
	}

	fmt.Println("\nСписок завдань:")
	for i, task := range todoList {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}
