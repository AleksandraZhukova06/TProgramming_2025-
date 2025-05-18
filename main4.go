package main

import (
	"fmt"
)

// TV представляет структуру телевизора
type TV struct {
	channel int
}

// NewTV - конструктор, создающий новый телевизор с заданным каналом
func NewTV(channel int) *TV {
	return &TV{channel: channel}
}

// GetChannel возвращает текущий номер канала
func (tv *TV) GetChannel() int {
	return tv.channel
}

// SetChannel устанавливает новый номер канала
func (tv *TV) SetChannel(channel int) {
	tv.channel = channel
}

// DrawASCII выводит ASCII графикой изображение телевизора
func (tv *TV) DrawASCII() {
	fmt.Println("Телевизор:")
	fmt.Println(" __________")
	fmt.Println("|          |")
	fmt.Println("|          |")
	fmt.Println("|          |")
	fmt.Println("|          |")
	fmt.Println("|          |")
	fmt.Println("|          |")
	fmt.Println("|          |")
	fmt.Println("|__________|")
	fmt.Printf(" Канал: %d\n", tv.channel)
}

func main() {
	// Создаем новый телевизор с каналом 5
	tv := NewTV(5)

	// Выводим текущий канал
	fmt.Printf("Текущий канал: %d\n", tv.GetChannel())

	// Устанавливаем новый канал
	tv.SetChannel(10)

	// Выводим новый канал
	fmt.Printf("Новый канал: %d\n", tv.GetChannel())

	// Рисуем телевизор в ASCII графике
	tv.DrawASCII()
}
