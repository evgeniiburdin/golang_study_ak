package main

import (
	"fmt"
)

// Browser представляет собой структуру, которая содержит историю посещенных URL-адресов
type Browser struct {
	history []string
}

// Visit добавляет новый URL-адрес в историю посещений
func (b *Browser) Visit(url string) {
	b.history = append(b.history, url)
	fmt.Printf("Посещение %s\n", url)
}

// Back возвращает предыдущий URL-адрес из истории и удаляет его
func (b *Browser) Back() string {
	if len(b.history) == 0 {
		return "Нет больше истории для возврата"
	}
	lastIndex := len(b.history) - 1
	lastURL := b.history[lastIndex]
	b.history = b.history[:lastIndex]
	fmt.Printf("Возврат к %s\n", lastURL)
	return lastURL
}

// ShowHistory выводит историю браузера
func (b *Browser) ShowHistory() {
	fmt.Println("История браузера:")
	for i := len(b.history) - 1; i >= 0; i-- {
		fmt.Println(b.history[i])
	}
}

func main() {
	browser := &Browser{}

	// Посещение URL-адресов
	browser.Visit("www.github.com")
	browser.Visit("www.google.com")

	// Возврат к предыдущим URL-адресам
	browser.Back()
	browser.Back()
	browser.Back()

	// Показать историю браузера
	browser.ShowHistory()
}
