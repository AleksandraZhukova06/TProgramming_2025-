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

    // Вычисление значений функции
    for _, xi := range x {
        // Вычисляем числитель и знаменатель
        numerator := math.Acos(xi*xi - b*b)
        denominator := math.Asin(xi*xi - a*a)

        // Проверка на деление на ноль
        if denominator == 0 {
            fmt.Printf("Для x = %.2f: Результат не определен (деление на ноль)\n", xi)
        } else {
            // Вычисляем значение функции
            yi := numerator / denominator
            fmt.Printf("Для x = %.2f: y = %.6f\n", xi, yi)
        }
    }
}