package main

import (
	"fmt"
	"math"
)

func main() {
	// Данные из задачи
	a := 0.05
	b := 0.06
	x := []float64{0.2, 0.95, 0.15, 0.15, 0.26, 0.37, 0.48, 0.56}

	// Массивы для хранения результатов
	yA := make([]float64, len(x))
	yB := make([]float64, len(x))

	// Вычисление значений функции для Задачи A
	for i, xi := range x {
		// Вычисляем числитель и знаменатель
		numerator := math.Acos(xi*xi - b*b)
		denominator := math.Asin(xi*xi - a*a)

		// Проверка на деление на ноль
		if denominator == 0 {
			fmt.Printf("Для x = %.2f: Результат не определен (деление на ноль)\n", xi)
			yA[i] = math.NaN()
		} else {
			// Вычисляем значение функции
			yA[i] = numerator / denominator
			fmt.Printf("Для x = %.2f: yA = %.6f\n", xi, yA[i])
		}
	}

	// Вычисление значений функции для Задачи B
	for i, xi := range x {
		// Вычисляем числитель и знаменатель
		numerator := math.Acos(xi*xi - b*b)
		denominator := math.Asin(xi*xi - a*a)

		// Проверка на деление на ноль
		if denominator == 0 {
			fmt.Printf("Для x = %.2f: Результат не определен (деление на ноль)\n", xi)
			yB[i] = math.NaN()
		} else {
			// Вычисляем значение функции
			yB[i] = numerator / denominator
			fmt.Printf("Для x = %.2f: yB = %.6f\n", xi, yB[i])
		}
	}

	// Вывод результатов
	fmt.Println("Результаты для Задачи A:")
	for i, yi := range yA {
		fmt.Printf("Для x = %.2f: yA = %.6f\n", x[i], yi)
	}

	fmt.Println("Результаты для Задачи B:")
	for i, yi := range yB {
		fmt.Printf("Для x = %.2f: yB = %.6f\n", x[i], yi)
	}
}
