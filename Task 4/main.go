package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Діапазони оцінок
var ranges = []struct {
	name  string
	lower int
	upper int
}{
	{"0-9", 0, 9},
	{"10-19", 10, 19},
	{"20-29", 20, 29},
	{"30-39", 30, 39},
	{"40-49", 40, 49},
	{"50-59", 50, 59},
	{"60-69", 60, 69},
	{"70-79", 70, 79},
	{"80-89", 80, 89},
	{"90-100", 90, 100},
}

func main() {
	// Зчитування оцінок від користувача
	fmt.Println("Введіть оцінки студентів (розділені пробілом):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	scoreStrings := strings.Split(input, " ")

	// Перетворення оцінок зі строкового типу у цілі числа
	var scores []int
	for _, s := range scoreStrings {
		score, err := strconv.Atoi(s)
		if err == nil {
			scores = append(scores, score)
		} else {
			fmt.Printf("Неправильний ввід: %s. Пропускаємо.\n", s)
		}
	}

	// Ініціалізуємо карту для підрахунку оцінок у кожному діапазоні
	histogram := make(map[string]int)

	// Підраховуємо кількість оцінок у кожному діапазоні
	for _, score := range scores {
		for _, r := range ranges {
			if score >= r.lower && score <= r.upper {
				histogram[r.name]++
				break
			}
		}
	}

	// Виводимо гістограму
	fmt.Println("Гістограма розподілу оцінок:")
	for _, r := range ranges {
		fmt.Printf("%s: %d\n", r.name, histogram[r.name])
	}
}
