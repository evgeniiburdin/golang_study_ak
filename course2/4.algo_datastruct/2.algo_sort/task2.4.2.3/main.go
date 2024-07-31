package main

// User - Структура пользователя
type User struct {
	ID   int
	Name string
	Age  int
}

func Merge(arr1 []User, arr2 []User) []User {
	result := make([]User, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i].Age <= arr2[j].Age {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	if i < len(arr1) {
		result = append(result, arr1[i:]...)
	}

	if j < len(arr2) {
		result = append(result, arr2[j:]...)
	}

	return result
}
