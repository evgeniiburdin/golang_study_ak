package main

import (
	"fmt"
	"student.vkusvill.ru/evgeniiburdin/go-course/course2/2.oop/3.oop_abstraction/task2.2.3.5/hashmap"
)

func main() {
	m := hashmap.NewHashMap(hashmap.WithHashCRC64())
	since := hashmap.MeassureTime(func() {
		m.Set("key", "value")
		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = hashmap.NewHashMap(hashmap.WithHashCRC32())
	since = hashmap.MeassureTime(func() {
		m.Set("key", "value")
		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)
}
