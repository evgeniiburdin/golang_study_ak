package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestLoadData(t *testing.T) {
	list := &DoubleLinkedList{}

	// Создаем временный JSON файл для тестирования
	tempFile, err := os.CreateTemp("", "testdata*.json")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Генерируем данные и записываем в файл
	commits := GenerateData(5)
	data, err := json.Marshal(commits)
	assert.NoError(t, err)

	_, err = tempFile.Write(data)
	assert.NoError(t, err)
	tempFile.Close()

	err = list.LoadData(tempFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, 5, list.Len())
}

func TestInit(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)
	assert.Equal(t, 5, list.Len())
}

func TestLen(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)
	assert.Equal(t, 5, list.Len())
}

func TestSetCurrent(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	err := list.SetCurrent(3)
	assert.NoError(t, err)
	assert.Equal(t, commits[3].Message, list.Current().data.Message)
}

func TestCurrent(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	assert.Equal(t, commits[0].Message, list.Current().data.Message)
}

func TestNext(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	list.SetCurrent(3)
	nextNode := list.Next()
	assert.Equal(t, commits[4].Message, nextNode.data.Message)
}

func TestPrev(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	list.SetCurrent(3)
	prevNode := list.Prev()
	assert.Equal(t, commits[2].Message, prevNode.data.Message)
}

func TestInsert(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	newCommit := Commit{
		Message: "New commit",
		UUID:    gofakeit.UUID(),
		Date:    time.Now().Format("2006-01-02"),
	}

	err := list.Insert(2, newCommit)
	assert.NoError(t, err)
	assert.Equal(t, 6, list.Len())

	node, err := list.GetByIndex(2)
	assert.NoError(t, err)
	assert.Equal(t, "New commit", node.data.Message)
}

func TestPush(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	newCommit := Commit{
		Message: "New commit",
		UUID:    gofakeit.UUID(),
		Date:    time.Now().Format("2006-01-02"),
	}

	err := list.Push(5, newCommit)
	assert.NoError(t, err)
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "New commit", list.tail.data.Message)
}

func TestDelete(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	err := list.Delete(2)
	assert.NoError(t, err)
	assert.Equal(t, 4, list.Len())
}

func TestDeleteCurrent(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	list.SetCurrent(2)
	err := list.DeleteCurrent()
	assert.NoError(t, err)
	assert.Equal(t, 4, list.Len())
}

func TestIndex(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	list.SetCurrent(3)
	index, err := list.Index()
	assert.NoError(t, err)
	assert.Equal(t, 3, index)
}

func TestGetByIndex(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	node, err := list.GetByIndex(2)
	assert.NoError(t, err)
	assert.Equal(t, commits[2].Message, node.data.Message)
}

func TestPop(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	node := list.Pop()
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, commits[4].Message, node.data.Message)
}

func TestShift(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	node := list.Shift()
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, commits[0].Message, node.data.Message)
}

func TestSearchUUID(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	uuid := commits[3].UUID
	node := list.SearchUUID(uuid)
	assert.NotNil(t, node)
	assert.Equal(t, uuid, node.data.UUID)
}

func TestSearch(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	message := commits[3].Message
	node := list.Search(message)
	assert.NotNil(t, node)
	assert.Equal(t, message, node.data.Message)
}

func TestReverse(t *testing.T) {
	list := &DoubleLinkedList{}
	commits := GenerateData(5)
	list.Init(commits)

	reversedList := list.Reverse()
	assert.Equal(t, list.Len(), reversedList.Len())

	originalHead := list.head.data.Message
	originalTail := list.tail.data.Message

	assert.Equal(t, originalHead, reversedList.tail.data.Message)
	assert.Equal(t, originalTail, reversedList.head.data.Message)
}

func TestQuickSort(t *testing.T) {
	commits := GenerateData(5)
	QuickSort(commits)

	for i := 1; i < len(commits); i++ {
		assert.True(t, commits[i-1].Date <= commits[i].Date)
	}
}

func TestGenerateData(t *testing.T) {
	commits := GenerateData(10)
	assert.Equal(t, 10, len(commits))
	for _, commit := range commits {
		assert.NotEmpty(t, commit.Message)
		assert.NotEmpty(t, commit.UUID)
		assert.NotEmpty(t, commit.Date)
	}
}
