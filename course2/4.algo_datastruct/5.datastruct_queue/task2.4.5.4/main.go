package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	items []int
}

// Push добавляет элемент в стек
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// Pop удаляет и возвращает последний элемент из стека
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Попытка извлечения элемента из пустого стека")
	}
	lastIndex := len(s.items) - 1
	element := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return element
}

// Interpreter представляет собой интерпретатор с использованием стека
type Interpreter struct {
	stack Stack
}

// Execute выполняет операции на стеке
func (i *Interpreter) Execute(code string) {
	tokens := strings.Fields(code)
	for _, token := range tokens {
		switch token {
		case "+":
			b := i.stack.Pop()
			a := i.stack.Pop()
			i.stack.Push(a + b)
		case "-":
			b := i.stack.Pop()
			a := i.stack.Pop()
			i.stack.Push(a - b)
		case "*":
			b := i.stack.Pop()
			a := i.stack.Pop()
			i.stack.Push(a * b)
		case "/":
			b := i.stack.Pop()
			a := i.stack.Pop()
			i.stack.Push(a / b)
		default:
			// Попробуем преобразовать токен в целое число и добавить его в стек
			value, err := strconv.Atoi(token)
			if err != nil {
				panic(fmt.Sprintf("Неверный токен: %s", token))
			}
			i.stack.Push(value)
		}
	}
}

func main() {
	interpreter := Interpreter{}

	code := "5 3 +"
	interpreter.Execute(code)

	fmt.Println(interpreter.stack.items[0]) // Вывод результата выполнения программы: 8
}
