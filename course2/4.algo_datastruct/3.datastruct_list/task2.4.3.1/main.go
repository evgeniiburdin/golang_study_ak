package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// Commit represents a commit with a message, UUID, and date
type Commit struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
	Date    string `json:"date"`
}

// Node represents a node in a double linked list
type Node struct {
	data *Commit
	prev *Node
	next *Node
}

// DoubleLinkedList represents a double linked list
type DoubleLinkedList struct {
	head *Node // начальный элемент в списке
	tail *Node // последний элемент в списке
	curr *Node // текущий элемент меняется при использовании методов next, prev
	len  int   // количество элементов в списке
}

// LinkedLister is an interface for linked list operations
type LinkedLister interface {
	LoadData(path string) error
	Init(c []Commit)
	Len() int
	SetCurrent(n int) error
	Current() *Node
	Next() *Node
	Prev() *Node
	Insert(n int, c Commit) error
	Push(n int, c Commit) error
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	GetByIndex(n int) (*Node, error)
	Pop() *Node
	Shift() *Node
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

// LoadData loads data from a JSON file at the given path into the list.
func (d *DoubleLinkedList) LoadData(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var commits []Commit
	err = json.Unmarshal(bytes, &commits)
	if err != nil {
		return err
	}
	QuickSort(commits)
	d.Init(commits)
	return nil
}

// Init инициализация двунаправленного связного списка
func (d *DoubleLinkedList) Init(commits []Commit) {
	if len(commits) == 0 {
		return
	}

	d.head = &Node{data: &commits[0]}
	d.tail = d.head
	d.curr = d.head
	d.len = 1

	for i := 1; i < len(commits); i++ {
		newNode := &Node{data: &commits[i], prev: d.tail}
		d.tail.next = newNode
		d.tail = newNode
		d.len++
	}
}

// Len получение длины списка
func (d *DoubleLinkedList) Len() int {
	return d.len
}

// SetCurrent устанавливает текущий элемент на позицию n
func (d *DoubleLinkedList) SetCurrent(n int) error {
	if d.head == nil {
		return errors.New("head is nil")
	}
	if n < 0 || n >= d.len {
		return errors.New("index out of range")
	}
	d.curr = d.head
	for i := 0; i < n; i++ {
		d.curr = d.curr.next
	}
	return nil
}

// Current возвращает текущий элемент
func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

// Next возвращает следующий элемент и сдвигает текущий
func (d *DoubleLinkedList) Next() *Node {
	d.curr = d.curr.next
	return d.curr
}

// Prev возвращает предыдущий элемент и сдвигает текущий
func (d *DoubleLinkedList) Prev() *Node {
	d.curr = d.curr.prev
	return d.curr
}

// Insert вставляет элемент c после позиции n
func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n < 0 || n > d.len {
		return errors.New("index out of bounds")
	}
	newNode := &Node{data: &c}
	if n == 0 {
		if d.head == nil {
			d.head = newNode
			d.tail = newNode
		} else {
			newNode.next = d.head
			d.head.prev = newNode
			d.head = newNode
		}
	} else if n == d.len {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	} else {
		current := d.head
		for i := 0; i < n; i++ {
			current = current.next
		}
		newNode.next = current
		newNode.prev = current.prev
		current.prev.next = newNode
		current.prev = newNode
	}
	d.len++
	return nil
}

// Push добавляет элемент в конец списка
func (d *DoubleLinkedList) Push(n int, c Commit) error {
	newNode := &Node{data: &c}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		d.curr = newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}
	d.len++
	return nil
}

// Delete удаляет элемент на позиции n
func (d *DoubleLinkedList) Delete(n int) error {
	err := d.SetCurrent(n)
	if err != nil {
		return err
	}
	err = d.DeleteCurrent()
	if err != nil {
		return err
	}
	return nil
}

// DeleteCurrent удаляет текущий элемент
func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.curr == nil {
		return errors.New("current is nil")
	}
	if d.curr == d.head {
		d.head = d.head.next
		if d.head != nil {
			d.head.prev = nil
		}
	} else if d.curr == d.tail {
		d.tail = d.tail.prev
		if d.tail != nil {
			d.tail.next = nil
		}
	} else {
		d.curr.prev.next = d.curr.next
		d.curr.next.prev = d.curr.prev
	}
	d.len--
	return nil
}

// Index возвращает индекс текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {
	index := 0
	current := d.head
	for current != nil {
		if current == d.curr {
			return index, nil
		}
		current = current.next
		index++
	}
	return -1, errors.New("current not found")
}

// GetByIndex возвращает элемент по индексу n
func (d *DoubleLinkedList) GetByIndex(n int) (*Node, error) {
	currentCurr := d.curr
	err := d.SetCurrent(n)
	if err != nil {
		return nil, err
	}
	defer func() {
		d.curr = currentCurr
	}()
	return d.curr, nil
}

// Pop удаляет и возвращает последний элемент списка
func (d *DoubleLinkedList) Pop() *Node {
	if d.len == 0 {
		return nil
	}
	last := d.tail
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	}
	d.len--
	return last
}

// Shift удаляет и возвращает первый элемент списка
func (d *DoubleLinkedList) Shift() *Node {
	if d.head == nil {
		return nil
	}
	first := d.head
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	}
	d.len--
	return first
}

// SearchUUID ищет коммит по UUID
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	current := d.head
	for current != nil {
		if current.data.UUID == uuID {
			return current
		}
		current = current.next
	}
	return nil
}

// Search поиск по сообщению
func (d *DoubleLinkedList) Search(message string) *Node {
	current := d.head
	for current != nil {
		if current.data.Message == message {
			return current
		}
		current = current.next
	}
	return nil
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	if d.head == nil {
		return d
	}

	current := d.head
	var prev *Node = nil
	d.tail = d.head

	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	d.head = prev

	return d
}

// QuickSort реализация алгоритма QuickSort для сортировки коммитов по дате
func QuickSort(commits []Commit) {
	if len(commits) <= 1 {
		return
	}

	// Выбираем опорный элемент (в данном случае - середину)
	pivot := commits[len(commits)/2].Date

	// Разделяем список на элементы меньше и больше опорного элемента
	i, j := 0, len(commits)-1
	for i <= j {
		for commits[i].Date < pivot {
			i++
		}
		for commits[j].Date > pivot {
			j--
		}
		if i <= j {
			commits[i], commits[j] = commits[j], commits[i]
			i++
			j--
		}
	}

	// Рекурсивно сортируем две половины
	if j > 0 {
		QuickSort(commits[:j+1])
	}
	if i < len(commits)-1 {
		QuickSort(commits[i:])
	}
}

// GenerateData генерирует случайные коммиты
func GenerateData(numCommits int) []Commit {
	var commits []Commit
	gofakeit.Seed(0) // Инициализация случайного генератора
	for i := 0; i < numCommits; i++ {
		commit := Commit{
			Message: gofakeit.Sentence(5), // Генерация случайного предложения из 5 слов
			UUID:    gofakeit.UUID(),      // Генерация случайного UUID
			Date:    gofakeit.DateRange(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)).Format("2006-01-02"),
			// Генерация случайной даты между 2020 и 2022 годами
		}
		commits = append(commits, commit)
	}
	return commits
}

func main() {
	// Создаем и инициализируем двунаправленный связный список
	list := &DoubleLinkedList{}
	commits := GenerateData(10) // Генерируем 10 случайных коммитов
	list.Init(commits)

	// Выводим длину списка
	fmt.Println("Initial length of the list:", list.Len())

	// Вставляем новый коммит после позиции 3
	newCommit := Commit{
		Message: "New commit",
		UUID:    gofakeit.UUID(),
		Date:    time.Now().Format("2006-01-02"),
	}
	err := list.Insert(3, newCommit)
	if err != nil {
		fmt.Println("Error inserting new commit:", err)
	}

	// Выводим длину списка после вставки
	fmt.Println("Length of the list after insertion:", list.Len())

	// Удаляем коммит на позиции 2
	err = list.Delete(2)
	if err != nil {
		fmt.Println("Error deleting commit:", err)
	}

	// Выводим длину списка после удаления
	fmt.Println("Length of the list after deletion:", list.Len())

	// Поиск коммита по UUID
	uuidToSearch := commits[5].UUID
	node := list.SearchUUID(uuidToSearch)
	if node != nil {
		fmt.Printf("Commit found with UUID %s: %s\n", uuidToSearch, node.data.Message)
	} else {
		fmt.Printf("Commit with UUID %s not found\n", uuidToSearch)
	}

	// Поиск коммита по сообщению
	messageToSearch := "New commit"
	node = list.Search(messageToSearch)
	if node != nil {
		fmt.Printf("Commit found with message '%s': %s\n", messageToSearch, node.data.UUID)
	} else {
		fmt.Printf("Commit with message '%s' not found\n", messageToSearch)
	}

	// Выводим текущий элемент
	fmt.Printf("Current element: %s\n", list.Current().data.Message)

	// Перемещаемся на следующий элемент и выводим его
	nextNode := list.Next()
	if nextNode != nil {
		fmt.Printf("Next element: %s\n", nextNode.data.Message)
	}

	// Перемещаемся на предыдущий элемент и выводим его
	prevNode := list.Prev()
	if prevNode != nil {
		fmt.Printf("Previous element: %s\n", prevNode.data.Message)
	}

	// Сортируем список
	list.Reverse()

	// Выводим элементы списка в обратном порядке
	fmt.Println("List elements after reversing:")
	current := list.head
	for current != nil {
		fmt.Printf("%s -> ", current.data.Message)
		current = current.next
	}
	fmt.Println("nil")
}

// Остальной код вашего пакета, включая структуры и функции, остается неизменным.
